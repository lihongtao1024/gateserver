package singleton

import (
	"gateserver/component"
	"sync"
)

var NetInstance component.Network
var netOnce sync.Once

func NewNetwork(network component.Network) component.Network {
	netOnce.Do(func() {
		NetInstance = network
	})

	return NetInstance
}
