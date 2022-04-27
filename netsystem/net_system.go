package netsystem

import (
	"encoding/binary"
	"gateserver/configsystem"
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"gateserver/internal/timers"
	"gateserver/logsystem"
	"gateserver/timersystem"
	"sync"
	"unsafe"
)

const (
	netCheckTimeout  = uint32(5000)
	netPackageHeader = int(unsafe.Sizeof(int32(0)))
)

type NetSystem struct {
	svrSessions  []*Server
	svrTimer     timers.Timer
	netComponent networks.Component
	gtListener   networks.Listener
	gtClients    map[interface{}]*Client
}

var Instance *NetSystem
var thisOnce sync.Once

func NewNetSystemInstance(index int) *NetSystem {
	thisOnce.Do(func() {
		Instance = &NetSystem{}
		Instance.svrSessions = make([]*Server, 0)
		Instance.svrTimer = timersystem.Instance.AddTimer(
			Instance,
			netCheckTimeout,
			timers.InfiniteTimer,
		)
		Instance.gtClients = make(map[interface{}]*Client)

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
			NewServer(1, configs.ServerIdWs, wslistenattr),
		)

		gslistenattrs := make([]*Server, 0)
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
				NewServer(i, configs.ServerIdGs, gslistenattr),
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
			NewServer(1, configs.ServerIdCt, cslistenattr),
		)

		gtlistenattr := configsystem.Instance.GetServerAttr(
			configs.ServerIdGt,
			index,
			configs.ServerIdCl,
		)
		if gtlistenattr == nil {
			Instance = nil
			return
		}
		Instance.netComponent = networks.NewComponent(Instance, Instance)
		if Instance.netComponent == nil {
			Instance = nil
			return
		}

		Instance.gtListener = Instance.netComponent.Listen(
			gtlistenattr.Ip,
			uint16(gtlistenattr.Port),
		)
		if Instance.gtListener == nil {
			Instance = nil
			return
		}

		Instance.OnTimer()

		logsystem.Instance.Inf(
			"listen on [GT%d]: [%s:%d] [ok].",
			index,
			gtlistenattr.Ip,
			gtlistenattr.Port,
		)
	})

	return Instance
}

func (ss *NetSystem) OnTimer() {
	for _, server := range ss.svrSessions {
		if server.IsState(ServerIdle) {
			//server.SwitchState(&ServerConnectingState{})
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
			logsystem.Instance.Err("on connected, system error:1.")
			conn.Disconnect()
			return
		}

		conn.SetData(client)
		ss.gtClients[conn] = client
		client.OnConnected()
	} else {
		server, ok := conn.GetData().(Session)
		if !ok {
			logsystem.Instance.Err("on connected, system error:2.")
			conn.Disconnect()
			return
		}
		server.OnConnected()
	}
}

func (ss *NetSystem) OnFatal(err error, conn networks.Connection) {
	server, ok := conn.GetData().(Session)
	if !ok {
		logsystem.Instance.Err("on fatal, system error:1.")
		conn.Disconnect()
		return
	}

	server.OnFatal(err)
}

func (ss *NetSystem) OnClosed(conn networks.Connection) {
	server, ok := conn.GetData().(Session)
	if !ok {
		logsystem.Instance.Err("on closed, system error:1.")
		conn.Disconnect()
		return
	}

	conn.SetData(nil)
	server.OnClosed()
}

func (ss *NetSystem) OnReceived(data []byte, conn networks.Connection) {
	server, ok := conn.GetData().(Session)
	if !ok {
		logsystem.Instance.Err("on recv, system error:1.")
		conn.Disconnect()
		return
	}
	server.OnReceived(data[unsafe.Sizeof(uint32(0)):])
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
