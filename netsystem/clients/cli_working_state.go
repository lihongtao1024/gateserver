package clients

import (
	"gateserver/logsystem"
	"gateserver/protosystem"
)

type ClientWorkingState struct {
}

func (state *ClientWorkingState) GetType() int {
	return ClientVerifying
}

func (state *ClientWorkingState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientWorkingState.", client.GetLogicName())

	client.SendRandKey()
}

func (state *ClientWorkingState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientWorkingState.", client.GetLogicName())
}

func (state *ClientWorkingState) OnReceived(o interface{}, data []byte) {
	client := o.(*Client)
	protosystem.Instance.SetDecodeSession(client)
	protosystem.Instance.ReadProto(data)
	protosystem.Instance.SetDecodeSession(nil)
}
