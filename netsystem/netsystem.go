package netsystem

import (
	"encoding/binary"
	"fmt"
	"gateserver/configsystem"
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"gateserver/logsystem"
	"net"
	"sync"
	"time"
)

const (
	netCheckTimeout = time.Second * 5
)

type NetSystem struct {
	svrNodes     []Session
	svrTimer     *time.Ticker
	netComponent networks.Component
	gtListener   *net.TCPListener
	gtClients    map[interface{}]Session
}

var This *NetSystem
var thisOnce sync.Once

func NewNetSystemInstance(index int) *NetSystem {
	thisOnce.Do(func() {
		This = &NetSystem{}
		This.svrNodes = make([]Session, 0)
		This.svrTimer = time.NewTicker(netCheckTimeout)
		This.gtClients = make(map[interface{}]Session)

		wslistenattr := configsystem.This.GetServerAttr(
			configs.ServerIdWs,
			1,
			configs.ServerIdGt,
		)
		if wslistenattr == nil {
			This = nil
			return
		}
		This.svrNodes = append(This.svrNodes, NewWorldServer(1, wslistenattr))

		gslistenattrs := make([]Session, 0)
		for i := 1; ; i++ {
			gslistenattr := configsystem.This.GetServerAttr(
				configs.ServerIdGs,
				i,
				configs.ServerIdGt,
			)
			if gslistenattr == nil {
				break
			}
			gslistenattrs = append(gslistenattrs, NewGameServer(i, gslistenattr))
		}
		if len(gslistenattrs) == 0 {
			This = nil
			return
		}
		This.svrNodes = append(This.svrNodes, gslistenattrs...)

		cslistenattr := configsystem.This.GetServerAttr(
			configs.ServerIdCt,
			1,
			configs.ServerIdGt,
		)
		if cslistenattr == nil {
			This = nil
			return
		}
		This.svrNodes = append(This.svrNodes, NewChatServer(1, cslistenattr))

		gtlistenattr := configsystem.This.GetServerAttr(
			configs.ServerIdGt,
			index,
			configs.ServerIdCl,
		)
		if gtlistenattr == nil {
			This = nil
			return
		}
		This.netComponent = networks.NewComponent(This, This)
		if This.netComponent == nil {
			This = nil
			return
		}

		This.gtListener = This.netComponent.Listen(gtlistenattr.Ip, uint16(gtlistenattr.Port))
		if This.gtListener == nil {
			This = nil
			return
		}

		logsystem.This.Inf("on listen GT%d: [%s:%d] [ok].", index, gtlistenattr.Ip, gtlistenattr.Port)
	})

	return This
}

func (ss *NetSystem) checkSession() {
	fmt.Println("check session")
}

func (ss *NetSystem) OnUnpack(data []byte) int {
	l := len(data)
	if l < 4 {
		return networks.FragmentContinue
	}

	h := int(binary.LittleEndian.Uint32(data))
	if h <= 0 {
		return networks.FragmentFatal
	}

	if h > l {
		return networks.FragmentContinue
	}

	return h
}

func (ss *NetSystem) OnConnected(listener *net.TCPListener, conn networks.Connection) {
	logsystem.This.Inf(
		"on connected, local addr:%s, remote addr:%s",
		conn.GetLocalAddr(),
		conn.GetRemoteAddr(),
	)

	if listener == ss.gtListener {
		client := NewClient(conn)
		if _, ok := ss.gtClients[conn]; ok {
			logsystem.This.Err("on connected, system error:2.")
			conn.Disconnect()
			return
		}

		conn.SetData(client)
		ss.gtClients[conn] = client
		client.OnConnected()
	} else {
		server, ok := conn.GetData().(Session)
		if !ok {
			logsystem.This.Err("on connected, system error:3.")
			conn.Disconnect()
			return
		}
		server.OnConnected()
	}
}

func (ss *NetSystem) OnFatal(err error, conn networks.Connection) {
	logsystem.This.Err(
		"on fatal, local addr:%s, remote addr:%s, errmsg:%s",
		conn.GetLocalAddr(),
		conn.GetRemoteAddr(),
		err.Error(),
	)
}

func (ss *NetSystem) OnClosed(conn networks.Connection) {
	logsystem.This.Inf(
		"on closed, local addr:%s, remote addr:%s",
		conn.GetLocalAddr(),
		conn.GetRemoteAddr(),
	)

	server, ok := conn.GetData().(Session)
	if !ok {
		logsystem.This.Err("on closed, system error:1.")
		conn.Disconnect()
		return
	}
	conn.SetData(nil)
	server.OnClosed()
}

func (ss *NetSystem) OnReceived(data []byte, conn networks.Connection) {
	logsystem.This.Inf(
		"on recv, local addr:%s, remote addr:%s",
		conn.GetLocalAddr(),
		conn.GetRemoteAddr(),
	)

	server, ok := conn.GetData().(Session)
	if !ok {
		logsystem.This.Err("on recv, system error:1.")
		conn.Disconnect()
		return
	}
	server.OnReceived(data)
}

func (ss *NetSystem) Do() bool {
	select {
	case <-ss.svrTimer.C:
		ss.checkSession()
	default:
	}

	return ss.netComponent.Do()
}

func (ss *NetSystem) Close() {
	ss.svrTimer.Stop()

	for _, session := range ss.svrNodes {
		session.Disconnect()
	}

	for _, client := range ss.gtClients {
		client.Disconnect()
	}

	ss.gtListener.Close()
	ss.netComponent.Close()
}
