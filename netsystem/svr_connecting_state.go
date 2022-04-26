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
	server := o.(*Server)
	logsystem.Instance.Dbg("[%s] enter ServerConnectingState.", server.GetLogicName())
	if !server.Connect() {
		logsystem.Instance.Err("connect to [%s] [fail].")
	}
}

func (state *ServerConnectingState) OnLeave(o interface{}) {
	server := o.(*Server)
	logsystem.Instance.Dbg("[%s] leave ServerConnectingState.", server.GetLogicName())
}

func (state *ServerConnectingState) OnReceived(o interface{}, data []byte) {

}
