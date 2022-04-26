package netsystem

import (
	"gateserver/logsystem"
	"gateserver/protosystem"
)

type ClientConnectedState struct {
}

func (state *ClientConnectedState) GetType() int {
	return ClientConnected
}

func (state *ClientConnectedState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientConnectedState.", client.GetLogicName())
}

func (state *ClientConnectedState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientConnectedState.", client.GetLogicName())
}

func (state *ClientConnectedState) OnReceived(o interface{}, data []byte) {
	client := o.(*Client)
	if err := protosystem.Instance.VerifyClientHandShakeReq(data); err != nil {
		logsystem.Instance.Err("[%s] %s.", client.GetLogicName(), err)
		client.Disconnect()
		return
	}

	data = protosystem.Instance.BuildClientHandShakeRsp()
	client.Send(data)

	client.SwitchState(&ClientWorkingState{})
}
