package clients

import (
	"gateserver/internal/timers"
	"gateserver/logsystem"
	"gateserver/protosystem"
	"gateserver/timersystem"
)

const WorkingTimeout = 150000

type ClientWorkingState struct {
	tmrTimeout timers.Timer
}

func (state *ClientWorkingState) GetType() int {
	return ClientVerifying
}

func (state *ClientWorkingState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientWorkingState.", client.GetLogicName())

	client.SendRandKey()
	state.tmrTimeout = timersystem.Instance.AddTimer(state, WorkingTimeout, 1)
	state.tmrTimeout.SetData(client)
}

func (state *ClientWorkingState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientWorkingState.", client.GetLogicName())

	timersystem.Instance.DelTimer(state.tmrTimeout)
}

func (state *ClientWorkingState) OnReceived(o interface{}, data []byte) {
	client := o.(*Client)
	protosystem.Instance.SetDecodeSession(client)
	protosystem.Instance.ReadProto(data)
	protosystem.Instance.SetDecodeSession(nil)
}

func (state *ClientWorkingState) OnTimer() {
	client := state.tmrTimeout.GetData().(*Client)
	logsystem.Instance.Dbg("[%s] timeout ClientWorkingState.", client.GetLogicName())

	client.Disconnect()
}
