package messages

import (
	"fmt"
	"gateserver/internal/protocols"
	"gateserver/netsystem/clients"
	"gateserver/protosystem"
)

type GT2WSProto struct {
	*protocols.ClientWS
}

func NewGT2WSProto() *GT2WSProto {
	protos := &GT2WSProto{}
	protos.ClientWS = protocols.NewClientWS(protos)
	return protos
}

func (protos *GT2WSProto) GetMid() uint16 {
	return protos.ClientWS.GetMid()
}

func (protos *GT2WSProto) Dispatch(data []byte) bool {
	return protos.ClientWS.DispatchProto(data)
}

func (protos *GT2WSProto) OnLoginReq(proto *protocols.LoginReq) {
	client := protosystem.Instance.GetDecodeSession().(*clients.Client)
	fmt.Println(proto)
	client.SendHandShakeRsp()
}
