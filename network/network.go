package network

import (
	"encoding/binary"
	"gateserver/client"
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/pkg/networks"
	"gateserver/server"
	"gateserver/singleton"
	"gateserver/state/serverstate"
	"net"
	"unsafe"
)

const (
	netCheckTimeout  = uint32(5000)
	netPackageHeader = int(unsafe.Sizeof(int32(0)))
)

type networkImpl struct {
	pkg.Timer
	pkg.TcpComponent
	gtIndex     int
	wsServer    component.Server
	gsServer    []component.Server
	ctServer    component.Server
	gtClients   map[pkg.TcpConnection]component.Client
	tcpListener pkg.TcpListener
}

func NewNetwork(index int) component.Network {
	network := &networkImpl{}
	network.gtIndex = index

	network.Timer = singleton.TimerInstance.AddTimer(
		network,
		netCheckTimeout,
		pkg.InfiniteTimer,
	)
	network.gtClients = make(map[pkg.TcpConnection]component.Client)

	network.TcpComponent = networks.NewComponent(network, network)
	if network.TcpComponent == nil {
		return nil
	}

	wslistenattr := singleton.CfgInstance.GetListenAttr(
		pkg.ServerIdWs,
		1,
		pkg.ServerIdGt,
	)
	if wslistenattr == nil {
		return nil
	}

	network.wsServer = server.NewServer(
		1,
		pkg.ServerIdWs,
		wslistenattr,
		network.TcpComponent,
	)

	cslistenattr := singleton.CfgInstance.GetListenAttr(
		pkg.ServerIdCt,
		1,
		pkg.ServerIdGt,
	)
	if cslistenattr == nil {
		return nil
	}

	network.ctServer = server.NewServer(
		1,
		pkg.ServerIdCt,
		cslistenattr,
		network.TcpComponent,
	)

	network.gsServer = make([]component.Server, 0)
	for i := 1; ; i++ {
		gslistenattr := singleton.CfgInstance.GetListenAttr(
			pkg.ServerIdGs,
			i,
			pkg.ServerIdGt,
		)
		if gslistenattr == nil {
			break
		}

		network.gsServer = append(
			network.gsServer,
			server.NewServer(
				i,
				pkg.ServerIdGs,
				gslistenattr,
				network.TcpComponent,
			),
		)
	}

	if len(network.gsServer) == 0 {
		return nil
	}

	network.OnTimer()
	return network
}

func (network *networkImpl) deleteClient(client component.Client,
	conn pkg.TcpConnection) {
	delete(network.gtClients, conn)
}

func (network *networkImpl) OnTimer() {
	allworking := true

	if !network.wsServer.(component.Session).IsState(int(component.ServerWorking)) {
		allworking = false
	}

	if network.wsServer.(component.Session).IsState(int(component.ServerIdle)) {
		network.wsServer.(component.Session).SwitchState(&serverstate.ServerConnectingState{})
	}

	if !network.ctServer.(component.Session).IsState(int(component.ServerWorking)) {
		allworking = false
	}

	if network.ctServer.(component.Session).IsState(int(component.ServerIdle)) {
		network.ctServer.(component.Session).SwitchState(&serverstate.ServerConnectingState{})
	}

	for _, server := range network.gsServer {
		if !server.(component.Session).IsState(int(component.ServerWorking)) {
			allworking = false
		}

		if server.(component.Session).IsState(int(component.ServerIdle)) {
			server.(component.Session).SwitchState(&serverstate.ServerConnectingState{})
		}
	}

	if allworking {
		network.OnListen()
	}
}

func (network *networkImpl) OnUnpack(data []byte) int {
	l := len(data)
	if l < netPackageHeader {
		return int(pkg.TcpFragmentContinue)
	}

	h := int(binary.LittleEndian.Uint32(data))
	if h <= 0 {
		return int(pkg.TcpFragmentFatal)
	}

	if h > l {
		return int(pkg.TcpFragmentContinue)
	}

	return h
}

func (network *networkImpl) OnConnected(listener pkg.TcpListener, conn pkg.TcpConnection) {
	if listener != nil && listener == network.tcpListener {
		client := client.NewClient(conn)
		if _, ok := network.gtClients[conn]; ok {
			singleton.LogInstance.Err("on connected, system error:1.")
			conn.Disconnect()
			return
		}

		conn.SetData(client)
		network.gtClients[conn] = client
		client.(component.Session).OnConnected()
	} else {
		server, ok := conn.GetData().(component.Session)
		if !ok {
			singleton.LogInstance.Err("on connected, system error:2.")
			conn.Disconnect()
			return
		}
		server.OnConnected()
	}
}

func (network *networkImpl) OnFatal(err error, conn pkg.TcpConnection) {
	session, ok := conn.GetData().(component.Session)
	if !ok {
		singleton.LogInstance.Err("on fatal, system error:1.")
		conn.Disconnect()
		return
	}

	session.OnFatal(err)
}

func (network *networkImpl) OnClosed(conn pkg.TcpConnection) {
	session, ok := conn.GetData().(component.Session)
	if !ok {
		singleton.LogInstance.Err("on closed, system error:1.")
		conn.Disconnect()
		return
	}

	conn.SetData(nil)

	if client, ok := session.(component.Client); ok {
		network.deleteClient(client, conn)
		singleton.VerifyInstance.CancleRequest(client)
		singleton.OnlineInstance.DeleteRequest(client)
		singleton.OnlineInstance.DeleteOnline(client)
	}

	session.OnClosed()
}

func (network *networkImpl) OnReceived(data []byte, conn pkg.TcpConnection) {
	session, ok := conn.GetData().(component.Session)
	if !ok {
		singleton.LogInstance.Err("on recv, system error:1.")
		conn.Disconnect()
		return
	}
	session.OnReceived(data[unsafe.Sizeof(uint32(0)):])
}

func (network *networkImpl) OnListen() {
	if network.tcpListener != nil {
		return
	}

	listenattr := singleton.CfgInstance.GetListenAttr(
		pkg.ServerIdGt,
		network.gtIndex,
		pkg.ServerIdCl,
	)
	if listenattr == nil {
		singleton.LogInstance.Err(
			"listen on [GT%d] [fail], invalid ip/port config.",
			network.gtIndex,
		)
		return
	}

	network.tcpListener = network.Listen(
		listenattr.Ip,
		uint16(listenattr.Port),
	)
	if network.tcpListener == nil {
		singleton.LogInstance.Err(
			"listen on [GT%d]: [%s:%d] [fail].",
			network.gtIndex,
			listenattr.Ip,
			listenattr.Port,
		)
		return
	}

	singleton.LogInstance.Inf(
		"listen on [GT%d]: [%s:%d] [ok].",
		network.gtIndex,
		listenattr.Ip,
		listenattr.Port,
	)
}

func (network *networkImpl) GetWSServer() component.Server {
	if !network.wsServer.(component.Session).IsState(
		int(component.ServerWorking),
	) {
		return nil
	}

	return network.wsServer
}

func (network *networkImpl) GetGSServer(index int) component.Server {
	if index >= len(network.gsServer) {
		return nil
	}

	gs := network.gsServer[index]
	if !gs.(component.Session).IsState(int(component.ServerWorking)) {
		return nil
	}

	return gs
}

func (network *networkImpl) GetCTServer() component.Server {
	if !network.ctServer.(component.Session).IsState(
		int(component.ServerWorking),
	) {
		return nil
	}

	return network.ctServer
}

func (network *networkImpl) IsClientLimit() bool {
	return len(network.gtClients) >= singleton.CfgInstance.GetServerAttr(
		singleton.AppInstance.GetType(),
		singleton.AppInstance.GetIndex(),
	).Users
}

func (network *networkImpl) Close() {
	singleton.TimerInstance.DelTimer(network.Timer)

	network.wsServer.(component.Session).Disconnect()
	network.ctServer.(component.Session).Disconnect()
	for _, server := range network.gsServer {
		server.(component.Session).Disconnect()
	}

	for _, client := range network.gtClients {
		client.(component.Session).Disconnect()
	}

	((*net.TCPListener)(network.tcpListener)).Close()
	network.TcpComponent.Close()
}
