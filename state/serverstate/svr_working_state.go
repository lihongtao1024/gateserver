package serverstate

import (
	"gateserver/component"
	"gateserver/singleton"
)

type ServerWorkingState struct {
}

func (state *ServerWorkingState) GetType() int {
	return int(component.ServerWorking)
}

func (state *ServerWorkingState) OnEnter(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] enter ServerWorkingState.",
		server.GetLogicName(),
	)
}

func (state *ServerWorkingState) OnLeave(o interface{}) {
	server := o.(component.Server)
	singleton.LogInstance.Dbg(
		"[%s] leave ServerWorkingState.",
		server.GetLogicName(),
	)
}

func (state *ServerWorkingState) OnReceived(o interface{}, data []byte) {
	server := o.(component.Server)
	result, toclient, data, clients := singleton.ProtoInstance.ParseServerProto(data)
	if !result {
		return
	}

	if !toclient {
		singleton.ProtoInstance.SetDecodeSession(server)
		singleton.ProtoInstance.DispatchProto(data)
		singleton.ProtoInstance.SetDecodeSession(nil)
		return
	}

	result, _, _ = singleton.ProtoInstance.IsServerWatch(data)
	if !result {
		for _, client := range clients {
			client.(component.Session).Send(data)
		}
		return
	}

	for _, client := range clients {
		singleton.ProtoInstance.SetDecodeSession(client)
		singleton.ProtoInstance.DispatchProto(data)
		singleton.ProtoInstance.SetDecodeSession(nil)
	}
}
