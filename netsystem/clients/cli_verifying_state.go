package clients

import (
	"gateserver/internal/timers"
	"gateserver/logsystem"
	"gateserver/timersystem"
)

const VerifyingTimeout = 180000

type ClientVerifyingState struct {
	tmrTimeout timers.Timer
}

func (state *ClientVerifyingState) GetType() int {
	return ClientVerifying
}

func (state *ClientVerifyingState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientVerifyingState.", client.GetLogicName())

	state.tmrTimeout = timersystem.Instance.AddTimer(state, VerifyingTimeout, 1)
	state.tmrTimeout.SetData(client)
}

func (state *ClientVerifyingState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientVerifyingState.", client.GetLogicName())

	timersystem.Instance.DelTimer(state.tmrTimeout)
}

func (state *ClientVerifyingState) OnReceived(o interface{}, data []byte) {

}

func (state *ClientVerifyingState) OnTimer() {
	client := state.tmrTimeout.GetData().(*Client)
	logsystem.Instance.Dbg("[%s] timeout ClientVerifyingState.", client.GetLogicName())

	client.Disconnect()
}
