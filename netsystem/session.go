package netsystem

type Session interface {
	GetIndex() int
	GetLogicName() string
	OnConnected()
	OnFatal(err error)
	OnClosed()
	OnReceived(data []byte)
	Disconnect()
}
