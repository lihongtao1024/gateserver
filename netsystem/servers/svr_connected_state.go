package servers

import (
	"gateserver/logsystem"
)

type ServerConnectedState struct {
}

func (state *ServerConnectedState) GetType() int {
	return ServerConnected
}

func (state *ServerConnectedState) OnEnter(o interface{}) {
	server := o.(*Server)
	logsystem.Instance.Dbg("[%s] enter ServerConnectedState.", server.GetLogicName())

	server.SendHandShakeReq()
}

func (state *ServerConnectedState) OnLeave(o interface{}) {
	server := o.(*Server)
	logsystem.Instance.Dbg("[%s] leave ServerConnectedState.", server.GetLogicName())
}

func (state *ServerConnectedState) OnReceived(o interface{}, data []byte) {
	server := o.(*Server)
	if err := server.VerifyHandShakeRsp(data); err != nil {
		logsystem.Instance.Err("[%s] %s.", server.GetLogicName(), err)
		server.Disconnect()
		return
	}

	server.SwitchState(&ServerWorkingState{})
}
