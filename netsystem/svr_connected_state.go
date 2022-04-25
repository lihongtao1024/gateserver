package netsystem

import (
	"gateserver/logsystem"
)

type ServerConnectedState struct {
}

func (state *ServerConnectedState) GetType() int {
	return ServerConnected
}

func (state *ServerConnectedState) OnEnter(o interface{}) {
	logsystem.TheLog.Dbg("[%s] enter ServerConnectedState.", o.(Session).GetLogicName())
	//o.(*Server).Send([]byte{})
}

func (state *ServerConnectedState) OnLeave(o interface{}) {
	logsystem.TheLog.Dbg("[%s] leave ServerConnectedState.", o.(Session).GetLogicName())
}

func (state *ServerConnectedState) OnReceived(o interface{}, data []byte) {

}
