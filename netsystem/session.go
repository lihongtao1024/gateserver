package netsystem

type Session interface {
	GetIndex() int
	GetLogicName() string
	OnConnected()
	OnClosed()
	OnReceived(data []byte)
	Disconnect()
}
