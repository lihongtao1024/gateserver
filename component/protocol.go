package component

import "gateserver/pkg"

type Protocol interface {
	SetDecodeSession(o interface{})
	GetDecodeSession() interface{}
	BuildServerHandShakeReq() []byte
	VerifyServerHandShakeRsp(index uint16, data []byte) error
	BuildClientHandShakeRsp() []byte
	VerifyClientHandShakeReq(data []byte) error
	IsWSProtocol(mid uint16) bool
	IsCTProtocol(mid uint16) bool
	IsGSProtocol(mid uint16) bool
	DispatchProto(data []byte) bool
	IsServerWatch(data []byte) (result bool, mid uint16, pid uint16)
	IsClientWatch(data []byte) (result bool, mid uint16, pid uint16)
	BuildClientProto(proto pkg.WriterProto) (result bool, data []byte)
	BuildClientData(data []byte) (result bool, proto []byte)
	BuildServerProto(client Client, proto pkg.WriterProto) (result bool, data []byte)
	BuildServerData(client Client, data []byte) (result bool, proto []byte)
	ParseServerProto(data []byte) (result bool, client bool, proto []byte, clients []Client)
	Close()
}
