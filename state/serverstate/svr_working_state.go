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

func (state *ServerWorkingState) OnReceived(o interface{},
	data []byte) {
	server := o.(component.Server)
	result, isc, proto, clients := singleton.ProtoInstance.ParseServerProto(
		data,
	)
	if !result {
		return
	}

	if !isc {
		singleton.ProtoInstance.SetDecodeSession(server)
		singleton.ProtoInstance.DispatchProto(proto)
		singleton.ProtoInstance.SetDecodeSession(nil)
		return
	}

	result, _, _ = singleton.ProtoInstance.IsServerWatch(proto)
	if !result {
		for _, client := range clients {
			client.SendClientData(proto)
		}
		return
	}

	for _, client := range clients {
		singleton.ProtoInstance.SetDecodeSession(client)
		singleton.ProtoInstance.DispatchProto(proto)
		singleton.ProtoInstance.SetDecodeSession(nil)
	}
}
