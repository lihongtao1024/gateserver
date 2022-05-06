package verifysystem

import "gateserver/netsystem/clients"

type realImpl struct {
}

func (impl *realImpl) PostRequest(client *clients.Client) {
}

func (impl *realImpl) ReceiveResponse(client *clients.Client, err error) {
}

func (impl *realImpl) Close() {
}
