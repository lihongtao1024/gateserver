package protocol

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/pkg/protocols"
	"gateserver/singleton"
	"gateserver/state/clientstate"
)

const (
	logoutGame = iota + 1
	logoutWorld
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

func (protos *gt2wsProto) OnLoginAck(proto *protocols.LoginAck) {
	client := singleton.ProtoInstance.GetDecodeSession().(component.Client)
	if client.GetSuid() != pkg.Guid(proto.Suid) {
		return
	}

	client.SendClientProto(proto)

	if proto.Errcode != int32(pkg.ErrorOk) {
		client.(component.Session).Disconnect()
		return
	}

	client.SendRealNtf()
	singleton.VerifyInstance.CancleRequest(client)
	singleton.OnlineInstance.DeleteRequest(client)
	singleton.OnlineInstance.AddOnline(client)

	client.(component.Session).SwitchState(&clientstate.ClientLoggedInState{})
}

func (protos *gt2wsProto) OnLogoutAck(proto *protocols.LogoutAck) {
	client := singleton.ProtoInstance.GetDecodeSession().(component.Client)
	client.SendClientProto(proto)

	if proto.Errcode != int32(pkg.ErrorOk) {
		return
	}

	switch proto.Type {
	case logoutGame:
		client.SetGuid(pkg.InvalidGuid)
		client.SetGSid(0)
		client.SetAid(^uint32(0))
		client.(component.Session).SwitchState(&clientstate.ClientLoggedInState{})
	case logoutWorld:
		client.(component.Session).Disconnect()
	default:
	}
}

func (protos *gt2wsProto) OnEnterGSAck(proto *protocols.EnterGSAck) {
	client := singleton.ProtoInstance.GetDecodeSession().(component.Client)
	client.SendClientProto(proto)

	if proto.Errcode != int32(pkg.ErrorOk) {
		return
	}

	client.SetGuid(pkg.Guid(proto.Guid))
	client.SetGSid(uint8(proto.Gsindex))
	client.SetAid(proto.Arrayid)

	client.(component.Session).SwitchState(&clientstate.ClientPlayingState{})
}

func (protos *gt2wsProto) OnKickNtf(proto *protocols.KickNtf) {
	client := singleton.ProtoInstance.GetDecodeSession().(component.Client)
	client.SendClientProto(proto)
	client.(component.Session).Disconnect()
}
