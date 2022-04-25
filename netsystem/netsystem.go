package netsystem

import (
	"encoding/binary"
	"gateserver/configsystem"
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"gateserver/logsystem"
	"sync"
	"time"
	"unsafe"
)

const (
	netCheckTimeout  = time.Second * 5
	netPackageHeader = int(unsafe.Sizeof(int32(0)))
)

type NetSystem struct {
	svrSessions  []*Server
	svrTimer     *time.Ticker
	netComponent networks.Component
	gtListener   networks.Listener
	gtClients    map[interface{}]Session
}

var TheNet *NetSystem
var thisOnce sync.Once

func NewNetSystemInstance(index int) *NetSystem {
	thisOnce.Do(func() {
		TheNet = &NetSystem{}
		TheNet.svrSessions = make([]*Server, 0)
		TheNet.svrTimer = time.NewTicker(netCheckTimeout)
		TheNet.gtClients = make(map[interface{}]Session)

		wslistenattr := configsystem.TheConfig.GetServerAttr(
			configs.ServerIdWs,
			1,
			configs.ServerIdGt,
		)
		if wslistenattr == nil {
			TheNet = nil
			return
		}
		TheNet.svrSessions = append(
			TheNet.svrSessions,
			NewServer(1, configs.ServerIdWs, wslistenattr),
		)

		gslistenattrs := make([]*Server, 0)
		for i := 1; ; i++ {
			gslistenattr := configsystem.TheConfig.GetServerAttr(
				configs.ServerIdGs,
				i,
				configs.ServerIdGt,
			)
			if gslistenattr == nil {
				break
			}
			gslistenattrs = append(
				gslistenattrs,
				NewServer(i, configs.ServerIdGs, gslistenattr),
			)
		}
		if len(gslistenattrs) == 0 {
			TheNet = nil
			return
		}
		TheNet.svrSessions = append(TheNet.svrSessions, gslistenattrs...)

		cslistenattr := configsystem.TheConfig.GetServerAttr(
			configs.ServerIdCt,
			1,
			configs.ServerIdGt,
		)
		if cslistenattr == nil {
			TheNet = nil
			return
		}
		TheNet.svrSessions = append(
			TheNet.svrSessions,
			NewServer(1, configs.ServerIdCt, cslistenattr),
		)

		gtlistenattr := configsystem.TheConfig.GetServerAttr(
			configs.ServerIdGt,
			index,
			configs.ServerIdCl,
		)
		if gtlistenattr == nil {
			TheNet = nil
			return
		}
		TheNet.netComponent = networks.NewComponent(TheNet, TheNet)
		if TheNet.netComponent == nil {
			TheNet = nil
			return
		}

		TheNet.gtListener = TheNet.netComponent.Listen(
			gtlistenattr.Ip,
			uint16(gtlistenattr.Port),
		)
		if TheNet.gtListener == nil {
			TheNet = nil
			return
		}

		logsystem.TheLog.Inf(
			"listen on [GT%d]: [%s:%d] [ok].",
			index,
			gtlistenattr.Ip,
			gtlistenattr.Port,
		)
	})

	return TheNet
}

func (ss *NetSystem) checkServer() {
	for _, server := range ss.svrSessions {
		if server.IsState(ServerIdle) {
			server.SwitchState(&ServerConnectingState{})
		}
	}
}

func (ss *NetSystem) OnUnpack(data []byte) int {
	l := len(data)
	if l < netPackageHeader {
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

func (ss *NetSystem) OnConnected(listener networks.Listener, conn networks.Connection) {
	if listener == ss.gtListener {
		client := NewClient(conn)
		if _, ok := ss.gtClients[conn]; ok {
			logsystem.TheLog.Err("on connected, system error:1.")
			conn.Disconnect()
			return
		}

		conn.SetData(client)
		ss.gtClients[conn] = client
		client.OnConnected()
	} else {
		server, ok := conn.GetData().(Session)
		if !ok {
			logsystem.TheLog.Err("on connected, system error:2.")
			conn.Disconnect()
			return
		}
		server.OnConnected()
	}
}

func (ss *NetSystem) OnFatal(err error, conn networks.Connection) {
	server, ok := conn.GetData().(Session)
	if !ok {
		logsystem.TheLog.Err("on fatal, system error:1.")
		conn.Disconnect()
		return
	}

	server.OnFatal(err)
}

func (ss *NetSystem) OnClosed(conn networks.Connection) {
	server, ok := conn.GetData().(Session)
	if !ok {
		logsystem.TheLog.Err("on closed, system error:1.")
		conn.Disconnect()
		return
	}

	conn.SetData(nil)
	server.OnClosed()
}

func (ss *NetSystem) OnReceived(data []byte, conn networks.Connection) {
	logsystem.TheLog.Inf(
		"on recv, local addr:%s, remote addr:%s.",
		conn.GetLocalAddr(),
		conn.GetRemoteAddr(),
	)

	server, ok := conn.GetData().(Session)
	if !ok {
		logsystem.TheLog.Err("on recv, system error:1.")
		conn.Disconnect()
		return
	}
	server.OnReceived(data)
}

func (ss *NetSystem) Do() bool {
	select {
	case <-ss.svrTimer.C:
		ss.checkServer()
	default:
	}

	return ss.netComponent.Do()
}

func (ss *NetSystem) Close() {
	ss.svrTimer.Stop()

	for _, session := range ss.svrSessions {
		session.Disconnect()
	}

	for _, client := range ss.gtClients {
		client.Disconnect()
	}

	ss.gtListener.Close()
	ss.netComponent.Close()
}
