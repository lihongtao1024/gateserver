package sessions

type Session interface {
	OnConnected()
	OnFatal(err error)
	OnClosed()
	OnReceived(data []byte)
	Disconnect()
}

type SessionState interface {
	GetType() int
	OnEnter(o interface{})
	OnLeave(o interface{})
	OnReceived(o interface{}, data []byte)
}
