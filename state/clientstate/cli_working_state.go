package clientstate

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/singleton"
)

const WorkingTimeout = 150000

type ClientWorkingState struct {
	tmrTimeout pkg.Timer
}

func (state *ClientWorkingState) GetType() int {
	return int(component.ClientWorking)
}

func (state *ClientWorkingState) OnEnter(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] enter ClientWorkingState.",
		client.GetLogicName(),
	)

	client.SendRandKey()
	state.tmrTimeout = singleton.TimerInstance.AddTimer(
		state,
		WorkingTimeout,
		1,
	)
	state.tmrTimeout.SetData(client)
}

func (state *ClientWorkingState) OnLeave(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] leave ClientWorkingState.",
		client.GetLogicName(),
	)

	singleton.TimerInstance.DelTimer(state.tmrTimeout)
}

func (state *ClientWorkingState) OnReceived(o interface{},
	data []byte) {
	client := o.(component.Client)

	result, mid, pid := singleton.ProtoInstance.IsClientWatch(data)
	if !result {
		singleton.LogInstance.Dbg(
			"[%s] received unexpected protocol:[mid=%d, pid=%d].",
			client.GetLogicName(),
			mid,
			pid,
		)
		return
	}

	singleton.ProtoInstance.SetDecodeSession(client)
	singleton.ProtoInstance.DispatchProto(data)
	singleton.ProtoInstance.SetDecodeSession(nil)
}

func (state *ClientWorkingState) OnTimer() {
	client := state.tmrTimeout.GetData().(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] timeout ClientWorkingState.",
		client.GetLogicName(),
	)

	client.(component.Session).Disconnect()
}
