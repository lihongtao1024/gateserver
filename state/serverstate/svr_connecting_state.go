package serverstate

import (
	"gateserver/component"
	"gateserver/singleton"
)

type ServerConnectingState struct {
}

func (state *ServerConnectingState) GetType() int {
	return int(component.ServerConnecting)
}

func (state *ServerConnectingState) OnEnter(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] enter ServerConnectingState.",
		server.GetLogicName(),
	)
	if !server.Connect() {
		singleton.LogInstance.Err("connect to [%s] [fail].")
	}
}

func (state *ServerConnectingState) OnLeave(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] leave ServerConnectingState.",
		server.GetLogicName(),
	)
}

func (state *ServerConnectingState) OnReceived(o interface{}, data []byte) {

}
