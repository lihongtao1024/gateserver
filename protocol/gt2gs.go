package protocol

import "gateserver/pkg/protocols"

type gt2gsProto struct {
	*protocols.ClientGS
}

func newGT2GSProto() *gt2gsProto {
	protos := &gt2gsProto{}
	protos.ClientGS = protocols.NewClientGS(protos)
	return protos
}
