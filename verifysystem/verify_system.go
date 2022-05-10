package verifysystem

import (
	"gateserver/configsystem"
	"gateserver/internal/errors"
	"gateserver/internal/protocols"
	"gateserver/logsystem"
	"gateserver/netsystem/clients"
	"sync"
)

type VerifyComponent interface {
	PostRequest(client *clients.Client)
	ReceiveResponse(client *clients.Client, errcode errors.ErrorCode)
	Close()
}

type VerifySystem struct {
	verifyComponent VerifyComponent
	verifyClients   map[string]*clients.Client
}

var Instance *VerifySystem
var thisOnce sync.Once

func NewInstance() *VerifySystem {
	thisOnce.Do(func() {
		var comp VerifyComponent
		if configsystem.Instance.IsVirtual() {
			comp = newVirtualImpl()
		} else {
			comp = newRealImpl()
		}

		if comp == nil {
			return
		}

		Instance = &VerifySystem{verifyComponent: comp}
	})

	return Instance
}

func newVirtualImpl() VerifyComponent {
	vi := &virtualImpl{}
	if !vi.loadJson() {
		logsystem.Instance.Err("init verify system [fail].")
		return nil
	}

	return vi
}

func newRealImpl() VerifyComponent {
	return &realImpl{}
}

func (verify *VerifySystem) PostRequest(client *clients.Client, proto *protocols.LoginReq) {
	if verify.HasRequest(client) {
		client.SendLoginAck(errors.NewError(errors.ErrorLoginPend))
		client.Disconnect()
		return
	}

	client.SetSid(proto.Sid)
	client.SetUName(proto.Username)
	client.SetIp(proto.Ip)
	client.SetPasssword(proto.Pwd_content)
	client.SetHardware(proto.Hwid)
	client.SetLongitude(proto.Longitude)
	client.SetLatitude(proto.Latitude)
	client.SwitchState(&clients.ClientVerifyingState{})

	verify.verifyComponent.PostRequest(client)
}

func (verify *VerifySystem) ReceiveResponse(client *clients.Client, errcode errors.ErrorCode) {
	verify.verifyComponent.ReceiveResponse(client, errcode)
	client.SendLoginAck(errors.NewError(errcode))

	if errcode == errors.ErrorOk {
		client.SwitchState(&clients.ClientRequestingState{})
	} else {
		client.Disconnect()
	}
}

func (verify *VerifySystem) HasRequest(client *clients.Client) (result bool) {
	_, result = verify.verifyClients[client.GetUName()]
	return
}

func (verify *VerifySystem) CancleRequest(client *clients.Client) {
	delete(verify.verifyClients, client.GetUName())
}

func (verify *VerifySystem) Close() {
	verify.verifyComponent.Close()
}
