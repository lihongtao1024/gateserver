package serverstate

import (
	"gateserver/component"
	"gateserver/singleton"
)

type ServerWorkingState struct {
}

func (state *ServerWorkingState) GetType() int {
	return int(component.ServerWorking)
}

func (state *ServerWorkingState) OnEnter(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] enter ServerWorkingState.",
		server.GetLogicName(),
	)
}

func (state *ServerWorkingState) OnLeave(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] leave ServerWorkingState.",
		server.GetLogicName(),
	)
}

func (state *ServerWorkingState) OnReceived(o interface{}, data []byte) {

}
