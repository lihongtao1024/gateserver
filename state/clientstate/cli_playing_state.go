package clientstate

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/singleton"
)

type ClientPlayingState struct {
}

func (state *ClientPlayingState) GetType() int {
	return int(component.ClientPlaying)
}

func (state *ClientPlayingState) OnEnter(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] enter ClientPlayingState.",
		client.GetLogicName(),
	)
}

func (state *ClientPlayingState) OnLeave(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] leave ClientPlayingState.",
		client.GetLogicName(),
	)
}

func (state *ClientPlayingState) OnReceived(o interface{},
	data []byte) {
	client := o.(component.Client)

	result, mid, pid := singleton.ProtoInstance.IsClientWatch(data)
	if result {
		singleton.LogInstance.Dbg(
			"[%s] received unexpected protocol:[mid=%d, pid=%d].",
			client.GetLogicName(),
			mid,
			pid,
		)
		return
	}

	switch {
	case singleton.ProtoInstance.IsGSProtocol(mid):
		{
			client.SendServerData(pkg.ServerIdGs, data)
		}
	case singleton.ProtoInstance.IsCTProtocol(mid):
		{
			client.SendServerData(pkg.ServerIdCt, data)
		}
	case singleton.ProtoInstance.IsWSProtocol(mid):
		{
			client.SendServerData(pkg.ServerIdWs, data)
		}
	default:
		{
			singleton.LogInstance.Dbg(
				"[%s] received unexpected protocol:[mid=%d, pid=%d].",
				client.GetLogicName(),
				mid,
				pid,
			)
			return
		}
	}
}
