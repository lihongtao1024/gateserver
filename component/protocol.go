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
	IsServerWatch(data []byte) (result bool, mid uint16, pid uint16)
	IsClientWatch(data []byte) (result bool, mid uint16, pid uint16)
	BuildClientProto(proto pkg.WriterProto) (result bool, data []byte)
	BuildServerProto(client Client, proto pkg.WriterProto) (result bool, data []byte)
	ParseServerProto(data []byte) (result bool, client bool, proto []byte, clients []Client)
	Close()
}
