package verify

import (
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/pkg/protocols"
	"gateserver/singleton"
	"gateserver/state/clientstate"
)

type verifyImpl interface {
	PostRequest(client component.Client)
	ReceiveResponse(client component.Client, err pkg.ErrorCode)
	Close()
}

type verifyComponent struct {
	verifyImpl
	verifyClients map[string]component.Client
}

func NewVerify() *verifyComponent {
	var comp verifyImpl
	if singleton.CfgInstance.IsVirtual() {
		comp = newVirtualImpl()
	} else {
		comp = newRealImpl()
	}

	if comp == nil {
		return nil
	}

	return &verifyComponent{
		verifyImpl:    comp,
		verifyClients: make(map[string]component.Client),
	}
}

func newVirtualImpl() verifyImpl {
	vi := &virtualImpl{}
	if !vi.loadJson() {
		singleton.LogInstance.Err("init verify system [fail].")
		return nil
	}

	return vi
}

func newRealImpl() verifyImpl {
	return &realImpl{}
}

func (verify *verifyComponent) PostRequest(client component.Client,
	proto *protocols.LoginReq) {
	if verify.HasRequest(client) {
		client.SendLoginAck(pkg.ErrorLoginPend)
		client.(component.Session).Disconnect()
		return
	}

	client.SetSid(proto.Sid)
	client.SetUName(proto.Username)
	client.SetIp(proto.Ip)
	client.SetPasssword(proto.Pwd_content)
	client.SetHardware(proto.Hwid)
	client.SetLongitude(proto.Longitude)
	client.SetLatitude(proto.Latitude)
	client.(component.Session).SwitchState(&clientstate.ClientVerifyingState{})

	verify.verifyImpl.PostRequest(client)
}

func (verify *verifyComponent) ReceiveResponse(client component.Client,
	suid pkg.Guid, err pkg.ErrorCode) {
	if client.GetSuid() != suid {
		return
	}

	verify.verifyImpl.ReceiveResponse(client, err)
	client.SendLoginAck(err)

	if err == pkg.ErrorOk {
		if client.GetUid() == 0 {
			panic("illegal uid = 0")
		}

		oldclient := singleton.OnlineInstance.GetOnline(
			client.GetSid(),
			client.GetUid(),
		)
		if oldclient != nil {
			singleton.LogInstance.Sys(
				"[%] kick repeated login.",
				oldclient.GetLogicName(),
			)
			oldclient.SendKickNtf(pkg.ErrorLoginAgain)
			singleton.OnlineInstance.DeleteOnline(oldclient)
			oldclient.SetUName("")
			oldclient.SetSid(0)
			oldclient.SetUid(0)
			oldclient.(component.Session).Disconnect()
		}

		client.(component.Session).SwitchState(&clientstate.ClientRequestingState{})
	} else {
		client.(component.Session).Disconnect()
	}
}

func (verify *verifyComponent) HasRequest(client component.Client) (result bool) {
	_, result = verify.verifyClients[client.GetUName()]
	return
}

func (verify *verifyComponent) CancleRequest(client component.Client) {
	delete(verify.verifyClients, client.GetUName())
}

func (verify *verifyComponent) Close() {
	verify.verifyImpl.Close()
}
