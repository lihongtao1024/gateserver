package protocol

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/pkg/protocols"
	"gateserver/singleton"
)

type gt2wsProto struct {
	*protocols.ClientWS
}

func newGT2WSProto() *gt2wsProto {
	protos := &gt2wsProto{}
	protos.ClientWS = protocols.NewClientWS(protos)
	return protos
}

func (protos *gt2wsProto) OnLoginReq(proto *protocols.LoginReq) {
	client := singleton.ProtoInstance.GetDecodeSession().(component.Client)

	if singleton.NetInstance.IsClientLimit() {
		client.SendLoginAck(pkg.ErrorGTOverload)
		client.(component.Session).Disconnect()
	}

	if proto.Client_type != uint8(singleton.AppInstance.GetClientType()) {
		client.SendLoginAck(pkg.ErrorLoginIllegal)
		client.(component.Session).Disconnect()
	}

	if proto.Sid == 0 {
		client.SendLoginAck(pkg.ErrorLoginIllegal)
		client.(component.Session).Disconnect()
	}

	if proto.Username == "" || len(proto.Pwd_content) == 0 {
		client.SendLoginAck(pkg.ErrorLoginPwd)
		client.(component.Session).Disconnect()
	}

	singleton.VerifyInstance.PostRequest(client, proto)
}
