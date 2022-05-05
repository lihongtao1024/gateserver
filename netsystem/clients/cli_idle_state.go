package clients

import (
	"gateserver/logsystem"
)

type ClientIdleState struct {
}

func (state *ClientIdleState) GetType() int {
	return ClientIdle
}

func (state *ClientIdleState) OnEnter(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] enter ClientIdleState.", client.GetLogicName())
}

func (state *ClientIdleState) OnLeave(o interface{}) {
	client := o.(*Client)
	logsystem.Instance.Dbg("[%s] leave ClientIdleState.", client.GetLogicName())
}

func (state *ClientIdleState) OnReceived(o interface{}, data []byte) {

}
