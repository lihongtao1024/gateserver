package component

import (
	"gateserver/pkg"
)

type ServerState int

const (
	ServerIdle = ServerState(iota)
	ServerConnecting
	ServerConnected
	ServerWorking
)

type Server interface {
	GetType() pkg.ServerType
	GetIndex() int
	GetLogicName() string
	Connect() bool
	SendServerHandShakeReq() bool
	VerifyHandShakeRsp(data []byte) error
}
