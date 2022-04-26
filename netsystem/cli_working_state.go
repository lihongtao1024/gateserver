package netsystem

import (
	"gateserver/logsystem"
)

type ClientWorkingState struct {
}

func (state *ClientWorkingState) GetType() int {
	return ClientVerifying
}

func (state *ClientWorkingState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientWorkingState.", client.GetLogicName())

}

func (state *ClientWorkingState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientWorkingState.", client.GetLogicName())

}

func (state *ClientWorkingState) OnReceived(o interface{}, data []byte) {

}
