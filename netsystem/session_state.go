package netsystem

const (
	ServerIdle       = 0
	ServerConnecting = 1
	ServerConnected  = 2
	ServerWorking    = 3
)

const (
	ClientIdle       = 0
	ClientConnected  = 1
	ClientWorking    = 2
	ClientVerifying  = 3
	ClientRequesting = 4
	ClientLoggedIn   = 5
	ClientPlaying    = 6
)

type SessionState interface {
	GetType() int
	OnEnter(o interface{})
	OnLeave(o interface{})
	OnReceived(o interface{}, data []byte)
}
