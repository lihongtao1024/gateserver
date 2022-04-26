package netsystem

import (
	"gateserver/internal/configs"
	"gateserver/internal/machines"
	"gateserver/internal/networks"
	"gateserver/logsystem"
	"strconv"
)

type Server struct {
	svrIndex int
	svrType  int
	svrAttr  *configs.ListenAttr
	svrConn  networks.Connection
	svrState machines.Machine
}

func NewServer(i, t int, attr *configs.ListenAttr) *Server {
	svr := &Server{
		svrIndex: i,
		svrType:  t,
		svrAttr:  attr,
	}
	svr.svrState = machines.NewMachine(svr)
	svr.svrState.SwitchState(&ServerIdleState{})
	return svr
}

func (svr *Server) GetType() int {
	return svr.svrType
}

func (svr *Server) GetIndex() int {
	return svr.svrIndex
}

func (svr *Server) GetLogicName() string {
	return configs.GetServerName(svr.svrType) + strconv.Itoa(svr.svrIndex)
}

func (svr *Server) OnConnected() {
	logsystem.TheLog.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
	)

	svr.SwitchState(&ServerConnectedState{})
}

func (svr *Server) OnFatal(err error) {
	if svr.svrConn.IsDialFatal() {
		logsystem.TheLog.Err(
			"on fatal [%s]: errmsg:'%s'.",
			svr.GetLogicName(),
			err.Error(),
		)

		svr.svrConn = nil
		svr.SwitchState(&ServerIdleState{})
		return
	}

	logsystem.TheLog.Err(
		"on fatal [%s]: local addr:%s, remote addr:%s, errmsg:'%s'.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
		err.Error(),
	)
}

func (svr *Server) OnClosed() {
	logsystem.TheLog.Inf(
		"on closed [%s]: local addr:%s, remote addr:%s.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
	)

	svr.svrConn = nil
	svr.SwitchState(&ServerIdleState{})
}

func (svr *Server) OnReceived(data []byte) {
	state := svr.svrState.GetState()
	state.(SessionState).OnReceived(svr, data)
}

func (svr *Server) Connect() bool {
	svr.svrConn = TheNet.netComponent.Connect(
		svr.svrAttr.Ip,
		uint16(svr.svrAttr.Port),
	)
	if svr.svrConn == nil {
		logsystem.TheLog.Err(
			"connect to [%s]: [%s:%d] [fail].",
			svr.GetLogicName(),
			svr.svrAttr.Ip,
			svr.svrAttr.Port,
		)
		return false
	}

	svr.svrConn.SetData(svr)
	logsystem.TheLog.Inf(
		"connect to [%s]: [%s:%d] [wait].",
		svr.GetLogicName(),
		svr.svrAttr.Ip,
		svr.svrAttr.Port,
	)
	return true
}

func (svr *Server) IsState(s int) bool {
	return svr.svrState.IsState(s)
}

func (svr *Server) SwitchState(state SessionState) {
	svr.svrState.SwitchState(state)
}

func (svr *Server) Send(data []byte) bool {
	return svr.svrConn.Send(data)
}

func (svr *Server) Disconnect() {
	if !svr.svrState.IsState(ServerConnected) &&
		!svr.svrState.IsState(ServerWorking) {
		return
	}

	svr.svrConn.Disconnect()
}
