package clients

import (
	"gateserver/logsystem"
)

type ClientVerifyingState struct {
}

func (state *ClientVerifyingState) GetType() int {
	return ClientVerifying
}

func (state *ClientVerifyingState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientVerifyingState.", client.GetLogicName())
}

func (state *ClientVerifyingState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientVerifyingState.", client.GetLogicName())
}

func (state *ClientVerifyingState) OnReceived(o interface{}, data []byte) {

}
