package netsystem

type Session interface {
	GetIndex() int
	GetLogicName() string
	IsState(s int) bool
	SwitchState(state SessionState)
	OnConnected()
	OnFatal(err error)
	OnClosed()
	OnReceived(data []byte)
	Send(data []byte) bool
	Disconnect()
}
