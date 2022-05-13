package clientstate

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/singleton"
)

const LoggedInTimeout = 60000

type ClientLoggedInState struct {
	tmrTimeout pkg.Timer
}

func (state *ClientLoggedInState) GetType() int {
	return int(component.ClientLoggedIn)
}

func (state *ClientLoggedInState) OnEnter(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] enter ClientLoggedInState.",
		client.GetLogicName(),
	)

	state.tmrTimeout = singleton.TimerInstance.AddTimer(
		state,
		LoggedInTimeout,
		1,
	)
	state.tmrTimeout.SetData(client)
}

func (state *ClientLoggedInState) OnLeave(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] leave ClientLoggedInState.",
		client.GetLogicName(),
	)
	singleton.TimerInstance.DelTimer(state.tmrTimeout)
}

func (state *ClientLoggedInState) OnReceived(o interface{},
	data []byte) {
	client := o.(component.Client)

	result, mid, pid := singleton.ProtoInstance.IsClientWatch(data)
	if result {
		singleton.LogInstance.Dbg(
			"[%s] received unexpected protocol:[mid=%d, pid=%d].",
			client.GetLogicName(),
			mid,
			pid,
		)
		return
	}

	if !singleton.ProtoInstance.IsWSProtocol(mid) {
		singleton.LogInstance.Dbg(
			"[%s] received unexpected protocol:[mid=%d, pid=%d].",
			client.GetLogicName(),
			mid,
			pid,
		)
		return
	}

	client.SendServerData(pkg.ServerIdWs, data)
}

func (state *ClientLoggedInState) OnTimer() {
	client := state.tmrTimeout.GetData().(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] timeout ClientLoggedInState.",
		client.GetLogicName(),
	)

	client.SendKickNtf(pkg.ErrorTimeoutKick)
	client.(component.Session).Disconnect()
}
