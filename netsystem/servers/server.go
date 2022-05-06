package servers

import (
	"gateserver/internal/configs"
	"gateserver/internal/machines"
	"gateserver/internal/networks"
	"gateserver/logsystem"
	"gateserver/netsystem/sessions"
	"gateserver/protosystem"
	"strconv"
)

const (
	ServerIdle       = 0
	ServerConnecting = 1
	ServerConnected  = 2
	ServerWorking    = 3
)

type Server struct {
	svrIndex int
	svrType  int
	svrAttr  *configs.ListenAttr
	svrConn  networks.Connection
	svrComp  networks.Component
	svrState machines.Machine
}

func NewServer(i, t int, attr *configs.ListenAttr, comp networks.Component) *Server {
	svr := &Server{
		svrIndex: i,
		svrType:  t,
		svrAttr:  attr,
		svrComp:  comp,
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
	logsystem.Instance.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
	)

	svr.SwitchState(&ServerConnectedState{})
}

func (svr *Server) OnFatal(err error) {
	if svr.svrConn.IsDialFatal() {
		logsystem.Instance.Err(
			"on fatal [%s]: errmsg:'%s'.",
			svr.GetLogicName(),
			err.Error(),
		)

		svr.svrConn = nil
		svr.SwitchState(&ServerIdleState{})
		return
	}

	logsystem.Instance.Err(
		"on fatal [%s]: local addr:%s, remote addr:%s, errmsg:'%s'.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
		err.Error(),
	)
}

func (svr *Server) OnClosed() {
	logsystem.Instance.Inf(
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
	state.(sessions.SessionState).OnReceived(svr, data)
}

func (svr *Server) Connect() bool {
	svr.svrConn = svr.svrComp.Connect(
		svr.svrAttr.Ip,
		uint16(svr.svrAttr.Port),
	)
	if svr.svrConn == nil {
		logsystem.Instance.Err(
			"connect to [%s]: [%s:%d] [fail].",
			svr.GetLogicName(),
			svr.svrAttr.Ip,
			svr.svrAttr.Port,
		)
		return false
	}

	svr.svrConn.SetData(svr)
	logsystem.Instance.Inf(
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

func (svr *Server) SwitchState(state sessions.SessionState) {
	svr.svrState.SwitchState(state)
}

func (svr *Server) SendHandShakeReq() bool {
	data := protosystem.Instance.BuildServerHandShakeReq()
	return svr.Send(data)
}

func (svr *Server) VerifyHandShakeRsp(data []byte) error {
	return protosystem.Instance.VerifyServerHandShakeRsp(uint16(svr.GetIndex()), data)
}

func (svr *Server) Send(data []byte) bool {
	return svr.svrConn.Send(data)
}

func (svr *Server) Disconnect() {
	if svr.svrConn == nil {
		return
	}

	svr.svrConn.Disconnect()
}
