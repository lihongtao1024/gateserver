package pkg

type TcpComponent interface {
	Listen(ip string, port uint16) TcpListener
	Connect(ip string, port uint16) TcpConnection
	Do() bool
	Close()
}
