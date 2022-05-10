package netsystem

import (
	"encoding/binary"
	"gateserver/configsystem"
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"gateserver/internal/timers"
	"gateserver/logsystem"
	"gateserver/netsystem/clients"
	"gateserver/netsystem/servers"
	"gateserver/netsystem/sessions"
	"gateserver/timersystem"
	"gateserver/verifysystem"
	"sync"
	"unsafe"
)

const (
	netCheckTimeout  = uint32(5000)
	netPackageHeader = int(unsafe.Sizeof(int32(0)))
)

type NetSystem struct {
	gtIndex      int
	svrSessions  []*servers.Server
	svrTimer     timers.Timer
	netComponent networks.Component
	gtListener   networks.Listener
	gtClients    map[interface{}]*clients.Client
}

var Instance *NetSystem
var thisOnce sync.Once

func NewInstance(index int) *NetSystem {
	thisOnce.Do(func() {
		Instance = &NetSystem{}
		Instance.gtIndex = index
		Instance.svrSessions = make([]*servers.Server, 0)
		Instance.svrTimer = timersystem.Instance.AddTimer(
			Instance,
			netCheckTimeout,
			timers.InfiniteTimer,
		)
		Instance.gtClients = make(map[interface{}]*clients.Client)

		Instance.netComponent = networks.NewComponent(Instance, Instance)
		if Instance.netComponent == nil {
			Instance = nil
			return
		}

		wslistenattr := configsystem.Instance.GetServerAttr(
			configs.ServerIdWs,
			1,
			configs.ServerIdGt,
		)
		if wslistenattr == nil {
			Instance = nil
			return
		}
		Instance.svrSessions = append(
			Instance.svrSessions,
			servers.NewServer(1, configs.ServerIdWs, wslistenattr, Instance.netComponent),
		)

		gslistenattrs := make([]*servers.Server, 0)
		for i := 1; ; i++ {
			gslistenattr := configsystem.Instance.GetServerAttr(
				configs.ServerIdGs,
				i,
				configs.ServerIdGt,
			)
			if gslistenattr == nil {
				break
			}
			gslistenattrs = append(
				gslistenattrs,
				servers.NewServer(i, configs.ServerIdGs, gslistenattr, Instance.netComponent),
			)
		}
		if len(gslistenattrs) == 0 {
			Instance = nil
			return
		}
		Instance.svrSessions = append(Instance.svrSessions, gslistenattrs...)

		cslistenattr := configsystem.Instance.GetServerAttr(
			configs.ServerIdCt,
			1,
			configs.ServerIdGt,
		)
		if cslistenattr == nil {
			Instance = nil
			return
		}
		Instance.svrSessions = append(
			Instance.svrSessions,
			servers.NewServer(1, configs.ServerIdCt, cslistenattr, Instance.netComponent),
		)

		Instance.OnTimer()
	})

	return Instance
}

func (ss *NetSystem) deleteClient(client *clients.Client, conn networks.Connection) {
	delete(ss.gtClients, conn)
}

func (ss *NetSystem) OnTimer() {
	allworking := true

	for _, server := range ss.svrSessions {
		if !server.IsState(servers.ServerWorking) {
			allworking = false
		}

		if server.IsState(servers.ServerIdle) {
			server.SwitchState(&servers.ServerConnectingState{})
		}
	}

	if allworking {
		ss.OnListen()
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

func (ss *NetSystem) Connect(ip string, port uint16) networks.Connection {
	return ss.netComponent.Connect(ip, port)
}

func (ss *NetSystem) OnConnected(listener networks.Listener, conn networks.Connection) {
	if listener != nil && listener == ss.gtListener {
		client := clients.NewClient(conn)
		if _, ok := ss.gtClients[conn]; ok {
			logsystem.Instance.Err("on connected, system error:1.")
			conn.Disconnect()
			return
		}

		conn.SetData(client)
		ss.gtClients[conn] = client
		client.OnConnected()
	} else {
		server, ok := conn.GetData().(sessions.Session)
		if !ok {
			logsystem.Instance.Err("on connected, system error:2.")
			conn.Disconnect()
			return
		}
		server.OnConnected()
	}
}

func (ss *NetSystem) OnFatal(err error, conn networks.Connection) {
	session, ok := conn.GetData().(sessions.Session)
	if !ok {
		logsystem.Instance.Err("on fatal, system error:1.")
		conn.Disconnect()
		return
	}

	session.OnFatal(err)
}

func (ss *NetSystem) OnClosed(conn networks.Connection) {
	session, ok := conn.GetData().(sessions.Session)
	if !ok {
		logsystem.Instance.Err("on closed, system error:1.")
		conn.Disconnect()
		return
	}

	conn.SetData(nil)

	if client, ok := session.(*clients.Client); ok {
		ss.deleteClient(client, conn)
		verifysystem.Instance.CancleRequest(client)
	}

	session.OnClosed()
}

func (ss *NetSystem) OnReceived(data []byte, conn networks.Connection) {
	session, ok := conn.GetData().(sessions.Session)
	if !ok {
		logsystem.Instance.Err("on recv, system error:1.")
		conn.Disconnect()
		return
	}
	session.OnReceived(data[unsafe.Sizeof(uint32(0)):])
}

func (ss *NetSystem) OnListen() {
	if ss.gtListener != nil {
		return
	}

	listenattr := configsystem.Instance.GetServerAttr(
		configs.ServerIdGt,
		ss.gtIndex,
		configs.ServerIdCl,
	)
	if listenattr == nil {
		logsystem.Instance.Err(
			"listen on [GT%d] [fail], invalid ip/port config.",
			ss.gtIndex,
		)
		return
	}

	ss.gtListener = ss.netComponent.Listen(
		listenattr.Ip,
		uint16(listenattr.Port),
	)
	if ss.gtListener == nil {
		logsystem.Instance.Err(
			"listen on [GT%d]: [%s:%d] [fail].",
			ss.gtIndex,
			listenattr.Ip,
			listenattr.Port,
		)
		return
	}

	logsystem.Instance.Inf(
		"listen on [GT%d]: [%s:%d] [ok].",
		ss.gtIndex,
		listenattr.Ip,
		listenattr.Port,
	)
}

func (ss *NetSystem) Do() bool {
	return ss.netComponent.Do()
}

func (ss *NetSystem) Close() {
	timersystem.Instance.DelTimer(ss.svrTimer)

	for _, session := range ss.svrSessions {
		session.Disconnect()
	}

	for _, client := range ss.gtClients {
		client.Disconnect()
	}

	ss.gtListener.Close()
	ss.netComponent.Close()
}
