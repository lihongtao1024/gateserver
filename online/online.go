package online

import (
	"gateserver/component"
)

type onlineKey uint64

type onlineSystem struct {
	onlineClients  map[onlineKey]component.Client
	requestClients map[onlineKey]component.Client
}

func NewOnline() *onlineSystem {
	return &onlineSystem{
		make(map[onlineKey]component.Client),
		make(map[onlineKey]component.Client),
	}
}

func newOnlineKey1(client component.Client) onlineKey {
	return newOnlineKey2(client.GetSid(), client.GetUid())
}

func newOnlineKey2(sid uint8, uid uint32) onlineKey {
	return onlineKey(sid)<<32 | onlineKey(uid)
}

func (system *onlineSystem) AddOnline(client component.Client) {
	system.onlineClients[newOnlineKey1(client)] = client
}

func (system *onlineSystem) GetOnline(sid uint8, uid uint32) component.Client {
	client, ok := system.onlineClients[newOnlineKey2(sid, uid)]
	if !ok {
		return nil
	}
	return client
}

func (system *onlineSystem) DeleteOnline(client component.Client) {
	delete(system.onlineClients, newOnlineKey1(client))
}

func (system *onlineSystem) AddRequest(client component.Client) {
	system.requestClients[newOnlineKey1(client)] = client
}

func (system *onlineSystem) GetRequest(sid uint8, uid uint32) component.Client {
	client, ok := system.requestClients[newOnlineKey2(sid, uid)]
	if !ok {
		return nil
	}
	return client
}

func (system *onlineSystem) DeleteRequest(client component.Client) {
	delete(system.requestClients, newOnlineKey1(client))
}

func (system *onlineSystem) Close() {

}
