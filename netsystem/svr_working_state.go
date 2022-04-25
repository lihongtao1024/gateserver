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
	logsystem.TheLog.Dbg("[%s] enter ServerWorkingState.", o.(Session).GetLogicName())
}

func (state *ServerWorkingState) OnLeave(o interface{}) {
	logsystem.TheLog.Dbg("[%s] leave ServerWorkingState.", o.(Session).GetLogicName())
}

func (state *ServerWorkingState) OnReceived(o interface{}, data []byte) {

}