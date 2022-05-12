package serverstate

import (
	"gateserver/component"
	"gateserver/singleton"
)

type ServerConnectedState struct {
}

func (state *ServerConnectedState) GetType() int {
	return int(component.ServerConnected)
}

func (state *ServerConnectedState) OnEnter(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] enter ServerConnectedState.",
		server.GetLogicName(),
	)
	server.SendServerHandShakeReq()
}

func (state *ServerConnectedState) OnLeave(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] leave ServerConnectedState.",
		server.GetLogicName(),
	)
}

func (state *ServerConnectedState) OnReceived(o interface{}, data []byte) {
	server := o.(component.Server)
	if err := server.VerifyHandShakeRsp(data); err != nil {
		singleton.LogInstance.Err(
			"[%s] %s.",
			server.GetLogicName(),
			err,
		)
		server.(component.Session).Disconnect()
		return
	}

	server.(component.Session).SwitchState(&ServerWorkingState{})
}
