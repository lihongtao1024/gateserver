package netsystem

const (
	ServerIdle       = 0
	ServerConnecting = 1
	ServerConnected  = 2
	ServerWorking    = 3
)

type SessionState interface {
	GetType() int
	OnEnter(o interface{})
	OnLeave(o interface{})
	OnReceived(o interface{}, data []byte)
}
