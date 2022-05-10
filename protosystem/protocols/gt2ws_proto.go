package protocols

import (
	"gateserver/internal/protocols"
	"gateserver/netsystem/clients"
	"gateserver/protosystem"
	"gateserver/verifysystem"
)

type GT2WSProto struct {
	*protocols.ClientWS
}

func NewGT2WSProto() *GT2WSProto {
	protos := &GT2WSProto{}
	protos.ClientWS = protocols.NewClientWS(protos)
	return protos
}

func (protos *GT2WSProto) OnLoginReq(proto *protocols.LoginReq) {
	client := protosystem.Instance.GetDecodeSession().(*clients.Client)
	verifysystem.Instance.PostRequest(client, proto)
}
