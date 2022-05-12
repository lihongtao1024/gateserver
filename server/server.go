package server

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/pkg/configs"
	"gateserver/pkg/machines"
	"gateserver/singleton"
	"gateserver/state/serverstate"
	"strconv"
)

type serverImpl struct {
	svrIndex int
	svrType  pkg.ServerType
	svrAttr  *pkg.ListenAttr
	pkg.Machine
	pkg.TcpConnection
	pkg.TcpComponent
}

func NewServer(index int, typ1 pkg.ServerType, attr *pkg.ListenAttr,
	comp pkg.TcpComponent) component.Server {
	server := &serverImpl{
		svrIndex:     index,
		svrType:      typ1,
		svrAttr:      attr,
		TcpComponent: comp,
	}
	server.Machine = machines.NewMachine(server)
	server.SwitchState(&serverstate.ServerIdleState{})
	return server
}

func (server *serverImpl) GetType() pkg.ServerType {
	return server.svrType
}

func (server *serverImpl) GetIndex() int {
	return server.svrIndex
}

func (server *serverImpl) GetLogicName() string {
	return configs.GetServerName(server.svrType) +
		strconv.Itoa(server.svrIndex)
}

func (server *serverImpl) OnConnected() {
	singleton.LogInstance.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		server.GetLogicName(),
		server.GetLocalAddr(),
		server.GetRemoteAddr(),
	)

	server.SwitchState(&serverstate.ServerConnectedState{})
}

func (server *serverImpl) OnFatal(err error) {
	if server.IsDialFatal() {
		singleton.LogInstance.Err(
			"on fatal [%s]: errmsg:'%s'.",
			server.GetLogicName(),
			err.Error(),
		)

		server.TcpConnection = nil
		server.SwitchState(&serverstate.ServerIdleState{})
		return
	}

	singleton.LogInstance.Err(
		"on fatal [%s]: local addr:%s, remote addr:%s, errmsg:'%s'.",
		server.GetLogicName(),
		server.GetLocalAddr(),
		server.GetRemoteAddr(),
		err.Error(),
	)
}

func (server *serverImpl) OnClosed() {
	singleton.LogInstance.Inf(
		"on closed [%s]: local addr:%s, remote addr:%s.",
		server.GetLogicName(),
		server.GetLocalAddr(),
		server.GetRemoteAddr(),
	)

	server.TcpConnection = nil
	server.SwitchState(&serverstate.ServerIdleState{})
}

func (server *serverImpl) OnReceived(data []byte) {
	state := server.GetState()
	state.(component.SessionState).OnReceived(server, data)
}

func (server *serverImpl) Connect() bool {
	server.TcpConnection = server.TcpComponent.Connect(
		server.svrAttr.Ip,
		uint16(server.svrAttr.Port),
	)
	if server.TcpConnection == nil {
		singleton.LogInstance.Err(
			"connect to [%s]: [%s:%d] [fail].",
			server.GetLogicName(),
			server.svrAttr.Ip,
			server.svrAttr.Port,
		)
		return false
	}

	server.SetData(server)
	singleton.LogInstance.Inf(
		"connect to [%s]: [%s:%d] [wait].",
		server.GetLogicName(),
		server.svrAttr.Ip,
		server.svrAttr.Port,
	)
	return true
}

func (server *serverImpl) SendServerHandShakeReq() bool {
	return server.Send(singleton.ProtoInstance.BuildServerHandShakeReq())
}

func (server *serverImpl) VerifyHandShakeRsp(data []byte) error {
	return singleton.ProtoInstance.VerifyServerHandShakeRsp(
		uint16(singleton.AppInstance.GetIndex()),
		data,
	)
}

func (server *serverImpl) Disconnect() {
	if server.TcpConnection == nil {
		return
	}

	server.TcpConnection.Disconnect()
}
