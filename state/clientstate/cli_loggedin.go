package clientstate

import (
	"gateserver/component"
	"gateserver/singleton"
)

type ClientLoggedInState struct {
}

func (state *ClientLoggedInState) GetType() int {
	return int(component.ClientLoggedIn)
}

func (state *ClientLoggedInState) OnEnter(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] enter ClientLoggedInState.",
		client.GetLogicName(),
	)

	singleton.VerifyInstance.CancleRequest(client)
	singleton.OnlineInstance.DeleteRequest(client)
	singleton.OnlineInstance.AddOnline(client)
}

func (state *ClientLoggedInState) OnLeave(o interface{}) {
	client := o.(component.Client)
	singleton.LogInstance.Dbg(
		"[%s] leave ClientLoggedInState.",
		client.GetLogicName(),
	)
}

func (state *ClientLoggedInState) OnReceived(o interface{}, data []byte) {
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

}
