package clientstate

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/singleton"
)

const RequestingTimeout = 15000

type ClientRequestingState struct {
	tmrTimeout pkg.Timer
}

func (state *ClientRequestingState) GetType() int {
	return int(component.ClientRequesting)
}

func (state *ClientRequestingState) OnEnter(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] enter ClientRequestingState.",
		client.GetLogicName(),
	)

	client.SendLoginReq()
	state.tmrTimeout = singleton.TimerInstance.AddTimer(
		state,
		RequestingTimeout,
		1,
	)
	state.tmrTimeout.SetData(client)
}

func (state *ClientRequestingState) OnLeave(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] leave ClientRequestingState.",
		client.GetLogicName(),
	)

	singleton.TimerInstance.DelTimer(state.tmrTimeout)
}

func (state *ClientRequestingState) OnReceived(o interface{}, data []byte) {

}

func (state *ClientRequestingState) OnTimer() {
	client := state.tmrTimeout.GetData().(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] timeout ClientRequestingState.",
		client.GetLogicName(),
	)

	client.(component.Session).Disconnect()
}
