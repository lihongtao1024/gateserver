package verify

import (
	"gateserver/component"
	"gateserver/pkg"
)

type realImpl struct {
}

func (impl *realImpl) PostRequest(client component.Client) {
}

func (impl *realImpl) ReceiveResponse(client component.Client, err pkg.ErrorCode) {
}

func (impl *realImpl) Close() {
}
