package netsystem

import (
	"gateserver/logsystem"
)

type ServerConnectingState struct {
}

func (state *ServerConnectingState) GetType() int {
	return ServerConnecting
}

func (state *ServerConnectingState) OnEnter(o interface{}) {
	logsystem.Instance.Dbg("[%s] enter ServerConnectingState.", o.(Session).GetLogicName())
	if !o.(*Server).Connect() {
		logsystem.Instance.Err("connect to [%s] [fail].")
	}
}

func (state *ServerConnectingState) OnLeave(o interface{}) {
	logsystem.Instance.Dbg("[%s] leave ServerConnectingState.", o.(Session).GetLogicName())
}

func (state *ServerConnectingState) OnReceived(o interface{}, data []byte) {

}
