package protocol

import "gateserver/pkg/protocols"

type gt2ctProto struct {
	*protocols.ClientCS
}

func newGT2CTProto() *gt2ctProto {
	protos := &gt2ctProto{}
	protos.ClientCS = protocols.NewClientCS(protos)
	return protos
}
