package pkg

type TcpConnection interface {
	SetData(data interface{})
	GetData() interface{}
	GetLocalAddr() string
	GetRemoteAddr() string
	Send(data []byte) bool
	IsDialFatal() bool
	Disconnect()
}
