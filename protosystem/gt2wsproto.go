package protosystem

import "gateserver/internal/protocols"

type GT2WSProto struct {
	protoImpl *protocols.ClientWS
}

func NewGT2WSProto() *GT2WSProto {
	proto := &GT2WSProto{}
	proto.protoImpl = protocols.NewClientWS(proto)
	return proto
}

func (proto *GT2WSProto) DispatchProto(data []byte) bool {
	return proto.protoImpl.DispatchProto(data)
}
