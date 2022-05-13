package clientstate

import (
	"gateserver/component"
	"gateserver/singleton"
)

type ClientIdleState struct {
}

func (state *ClientIdleState) GetType() int {
	return int(component.ClientIdle)
}

func (state *ClientIdleState) OnEnter(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] enter ClientIdleState.",
		client.GetLogicName(),
	)
}

func (state *ClientIdleState) OnLeave(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] leave ClientIdleState.",
		client.GetLogicName(),
	)
}

func (state *ClientIdleState) OnReceived(o interface{},
	data []byte) {

}
