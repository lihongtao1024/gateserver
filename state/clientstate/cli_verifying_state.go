package clientstate

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/singleton"
)

const VerifyingTimeout = 180000

type ClientVerifyingState struct {
	tmrTimeout pkg.Timer
}

func (state *ClientVerifyingState) GetType() int {
	return int(component.ClientVerifying)
}

func (state *ClientVerifyingState) OnEnter(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] enter ClientVerifyingState.",
		client.GetLogicName(),
	)

	state.tmrTimeout = singleton.TimerInstance.AddTimer(
		state,
		VerifyingTimeout,
		1,
	)
	state.tmrTimeout.SetData(client)
}

func (state *ClientVerifyingState) OnLeave(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] leave ClientVerifyingState.",
		client.GetLogicName(),
	)

	singleton.TimerInstance.DelTimer(state.tmrTimeout)
}

func (state *ClientVerifyingState) OnReceived(o interface{},
	data []byte) {

}

func (state *ClientVerifyingState) OnTimer() {
	client := state.tmrTimeout.GetData().(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] timeout ClientVerifyingState.",
		client.GetLogicName(),
	)

	client.(component.Session).Disconnect()
}
