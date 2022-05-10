package verifysystem

import (
	"gateserver/internal/errors"
	"gateserver/netsystem/clients"
)

type realImpl struct {
}

func (impl *realImpl) PostRequest(client *clients.Client) {
}

func (impl *realImpl) ReceiveResponse(client *clients.Client, err errors.ErrorCode) {
}

func (impl *realImpl) Close() {
}
