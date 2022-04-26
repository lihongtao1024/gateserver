package netsystem

type Session interface {
	OnConnected()
	OnFatal(err error)
	OnClosed()
	OnReceived(data []byte)
}
