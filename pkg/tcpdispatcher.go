package pkg

type TcpDispatcher interface {
	OnConnected(listener TcpListener, conn TcpConnection)
	OnFatal(err error, conn TcpConnection)
	OnClosed(conn TcpConnection)
	OnReceived(data []byte, conn TcpConnection)
}
