package networks

import "net"

type Dispatch interface {
	OnConnected(listener *net.TCPListener, conn Connection)
	OnFatal(err error, conn Connection)
	OnClosed(conn Connection)
	OnReceived(data []byte, conn Connection)
}
