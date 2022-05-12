package component

import "gateserver/pkg"

type SessionState interface {
	GetType() int
	OnEnter(o interface{})
	OnLeave(o interface{})
	OnReceived(o interface{}, data []byte)
}

type Session interface {
	IsState(s int) bool
	SwitchState(state pkg.State)
	OnConnected()
	OnFatal(err error)
	OnClosed()
	OnReceived(data []byte)
	Send(data []byte) bool
	Disconnect()
}
