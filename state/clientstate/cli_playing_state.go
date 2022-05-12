package clientstate

import (
	"gateserver/component"
	"gateserver/singleton"
)

type ClientPlayingState struct {
}

func (state *ClientPlayingState) GetType() int {
	return int(component.ClientPlaying)
}

func (state *ClientPlayingState) OnEnter(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] enter ClientPlayingState.",
		client.GetLogicName(),
	)
}

func (state *ClientPlayingState) OnLeave(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] leave ClientPlayingState.",
		client.GetLogicName(),
	)
}

func (state *ClientPlayingState) OnReceived(o interface{}, data []byte) {

}
