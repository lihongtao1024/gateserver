package clients

import (
	"gateserver/internal/timers"
	"gateserver/logsystem"
	"gateserver/timersystem"
)

const ConnectedTimeout = 2000

type ClientConnectedState struct {
	tmrTimeout timers.Timer
}

func (state *ClientConnectedState) GetType() int {
	return ClientConnected
}

func (state *ClientConnectedState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientConnectedState.", client.GetLogicName())

	state.tmrTimeout = timersystem.Instance.AddTimer(state, ConnectedTimeout, 1)
	state.tmrTimeout.SetData(client)
}

func (state *ClientConnectedState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientConnectedState.", client.GetLogicName())

	timersystem.Instance.DelTimer(state.tmrTimeout)
}

func (state *ClientConnectedState) OnReceived(o interface{}, data []byte) {
	client := o.(*Client)
	if err := client.VerifyHandShakeReq(data); err != nil {
		logsystem.Instance.Err("[%s] %s.", client.GetLogicName(), err)
		client.Disconnect()
		return
	}

	client.SendHandShakeRsp()
	client.SwitchState(&ClientWorkingState{})
}

func (state *ClientConnectedState) OnTimer() {
	client := state.tmrTimeout.GetData().(*Client)
	logsystem.Instance.Dbg("[%s] timeout ClientConnectedState.", client.GetLogicName())

	client.Disconnect()
}
