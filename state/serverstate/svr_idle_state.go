package serverstate

import (
	"gateserver/component"
	"gateserver/singleton"
)

type ServerIdleState struct {
}

func (state *ServerIdleState) GetType() int {
	return int(component.ServerIdle)
}

func (state *ServerIdleState) OnEnter(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] enter ServerIdleState.",
		server.GetLogicName(),
	)
}

func (state *ServerIdleState) OnLeave(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] leave ServerIdleState.",
		server.GetLogicName(),
	)
}

func (state *ServerIdleState) OnReceived(o interface{}, data []byte) {

}
