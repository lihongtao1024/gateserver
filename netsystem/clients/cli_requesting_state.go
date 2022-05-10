package clients

import (
	"gateserver/internal/timers"
	"gateserver/logsystem"
	"gateserver/timersystem"
)

const RequestingTimeout = 15000

type ClientRequestingState struct {
	tmrTimeout timers.Timer
}

func (state *ClientRequestingState) GetType() int {
	return ClientPlaying
}

func (state *ClientRequestingState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientRequestingState.", client.GetLogicName())

	state.tmrTimeout = timersystem.Instance.AddTimer(state, RequestingTimeout, 1)
	state.tmrTimeout.SetData(client)
}

func (state *ClientRequestingState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientRequestingState.", client.GetLogicName())

	timersystem.Instance.DelTimer(state.tmrTimeout)
}

func (state *ClientRequestingState) OnReceived(o interface{}, data []byte) {

}

func (state *ClientRequestingState) OnTimer() {
	client := state.tmrTimeout.GetData().(*Client)
	logsystem.Instance.Dbg("[%s] timeout ClientRequestingState.", client.GetLogicName())

	client.Disconnect()
}
