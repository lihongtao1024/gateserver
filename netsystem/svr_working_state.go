package netsystem

import (
	"gateserver/logsystem"
)

type ServerWorkingState struct {
}

func (state *ServerWorkingState) GetType() int {
	return ServerWorking
}

func (state *ServerWorkingState) OnEnter(o interface{}) {
	server := o.(*Server)
	logsystem.Instance.Dbg("[%s] enter ServerWorkingState.", server.GetLogicName())
}

func (state *ServerWorkingState) OnLeave(o interface{}) {
	server := o.(*Server)
	logsystem.Instance.Dbg("[%s] leave ServerWorkingState.", server.GetLogicName())
}

func (state *ServerWorkingState) OnReceived(o interface{}, data []byte) {

}
