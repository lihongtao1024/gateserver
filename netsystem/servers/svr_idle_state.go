package servers

import (
	"gateserver/logsystem"
)

type ServerIdleState struct {
}

func (state *ServerIdleState) GetType() int {
	return ServerIdle
}

func (state *ServerIdleState) OnEnter(o interface{}) {
	server := o.(*Server)
	logsystem.Instance.Dbg("[%s] enter ServerIdleState.", server.GetLogicName())
}

func (state *ServerIdleState) OnLeave(o interface{}) {
	server := o.(*Server)
	logsystem.Instance.Dbg("[%s] leave ServerIdleState.", server.GetLogicName())
}

func (state *ServerIdleState) OnReceived(o interface{}, data []byte) {

}
