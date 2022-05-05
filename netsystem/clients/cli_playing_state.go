package clients

import (
	"gateserver/logsystem"
)

type ClientPlayingState struct {
}

func (state *ClientPlayingState) GetType() int {
	return ClientPlaying
}

func (state *ClientPlayingState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientPlayingState.", client.GetLogicName())

}

func (state *ClientPlayingState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientPlayingState.", client.GetLogicName())
}

func (state *ClientPlayingState) OnReceived(o interface{}, data []byte) {

}
