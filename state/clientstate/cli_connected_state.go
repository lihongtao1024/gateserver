package clientstate

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/singleton"
)

const ConnectedTimeout = 2000

type ClientConnectedState struct {
	tmrTimeout pkg.Timer
}

func (state *ClientConnectedState) GetType() int {
	return int(component.ClientConnected)
}

func (state *ClientConnectedState) OnEnter(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] enter ClientConnectedState.",
		client.GetLogicName(),
	)

	state.tmrTimeout = singleton.TimerInstance.AddTimer(
		state,
		ConnectedTimeout,
		1,
	)
	state.tmrTimeout.SetData(client)
}

func (state *ClientConnectedState) OnLeave(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] leave ClientConnectedState.",
		client.GetLogicName(),
	)

	singleton.TimerInstance.DelTimer(state.tmrTimeout)
}

func (state *ClientConnectedState) OnReceived(o interface{}, data []byte) {
	client := o.(component.Client)
	if err := client.VerifyClientHandShakeReq(data); err != nil {
		singleton.LogInstance.Err(
			"[%s] %s.",
			client.GetLogicName(),
			err,
		)
		client.(component.Session).Disconnect()
		return
	}

	client.SendClientHandShakeRsp()
	client.(component.Session).SwitchState(&ClientWorkingState{})
}

func (state *ClientConnectedState) OnTimer() {
	client := state.tmrTimeout.GetData().(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] timeout ClientConnectedState.",
		client.GetLogicName(),
	)

	client.(component.Session).Disconnect()
}
