package component

type Network interface {
	GetWSServer() Server
	GetGSServer(index int) Server
	GetCTServer() Server
	IsClientLimit() bool
	Do() bool
	Close()
}
