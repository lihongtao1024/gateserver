package netsystem

import (
	"gateserver/logsystem"
)

type ClientRequestingState struct {
}

func (state *ClientRequestingState) GetType() int {
	return ClientPlaying
}

func (state *ClientRequestingState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientRequestingState.", client.GetLogicName())

}

func (state *ClientRequestingState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientRequestingState.", client.GetLogicName())

}

func (state *ClientRequestingState) OnReceived(o interface{}, data []byte) {

}
