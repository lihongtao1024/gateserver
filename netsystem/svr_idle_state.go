package netsystem

import (
	"gateserver/logsystem"
)

type ServerIdleState struct {
}

func (state *ServerIdleState) GetType() int {
	return ServerIdle
}

func (state *ServerIdleState) OnEnter(o interface{}) {
	logsystem.TheLog.Dbg("[%s] enter ServerIdleState.", o.(Session).GetLogicName())
}

func (state *ServerIdleState) OnLeave(o interface{}) {
	logsystem.TheLog.Dbg("[%s] leave ServerIdleState.", o.(Session).GetLogicName())
}

func (state *ServerIdleState) OnReceived(o interface{}, data []byte) {

}
