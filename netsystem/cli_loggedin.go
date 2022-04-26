package netsystem

import (
	"gateserver/logsystem"
)

type ClientLoggedInState struct {
}

func (state *ClientLoggedInState) GetType() int {
	return ClientLoggedIn
}

func (state *ClientLoggedInState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientLoggedInState.", client.GetLogicName())

}

func (state *ClientLoggedInState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientLoggedInState.", client.GetLogicName())

}

func (state *ClientLoggedInState) OnReceived(o interface{}, data []byte) {

}
