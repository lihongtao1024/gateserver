package component

import "gateserver/pkg"

type Protocol interface {
	SetDecodeSession(o interface{})
	GetDecodeSession() interface{}
	DispatchProto(data []byte) bool
	BuildServerHandShakeReq() []byte
	VerifyServerHandShakeRsp(index uint16, data []byte) error
	BuildClientHandShakeRsp() []byte
	VerifyClientHandShakeReq(data []byte) error
	BuildClientProto(proto pkg.WriterProto) (result bool, data []byte)
	BuildServerProto(client Client, proto pkg.WriterProto) (
		result bool, data []byte)
	Close()
}
