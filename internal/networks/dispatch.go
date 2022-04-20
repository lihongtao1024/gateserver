package networks

type Dispatch interface {
	OnConnected(listener Listener, conn Connection)
	OnFatal(err error, conn Connection)
	OnClosed(conn Connection)
	OnReceived(data []byte, conn Connection)
}
