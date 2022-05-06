package verifysystem

import (
	"gateserver/configsystem"
	"gateserver/internal/protocols"
	"gateserver/logsystem"
	"gateserver/netsystem/clients"
	"sync"
)

type VerifyComponent interface {
	PostRequest(client *clients.Client)
	ReceiveResponse(client *clients.Client, err error)
	Close()
}

type VerifySystem struct {
	verifyComponent VerifyComponent
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

func (verify *VerifySystem) PostRequest(client *clients.Client, o interface{}) {
	proto := o.(*protocols.LoginReq)

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

func (verify *VerifySystem) ReceiveResponse(client *clients.Client, err error) {
	verify.verifyComponent.ReceiveResponse(client, err)
	if err == nil {
		client.SwitchState(&clients.ClientRequestingState{})
	} else {
		client.Disconnect()
	}
}

func (verify *VerifySystem) Close() {
	verify.verifyComponent.Close()
}
