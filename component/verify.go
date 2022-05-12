package component

import (
	"gateserver/pkg"
	"gateserver/pkg/protocols"
)

type Verify interface {
	PostRequest(client Client, proto *protocols.LoginReq)
	ReceiveResponse(client Client, suid pkg.Guid, err pkg.ErrorCode)
	HasRequest(client Client) bool
	CancleRequest(client Client)
	Close()
}
