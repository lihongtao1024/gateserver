package netsystem

type Server interface {
	GetIndex() int
	GetLogicName() string
	IsConnected() bool
	OnConnected()
	OnFatal(err error)
	OnClosed()
	OnReceived(data []byte)
	Connect() bool
	Disconnect()
}
