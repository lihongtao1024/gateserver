package singleton

import (
	"gateserver/component"
	"sync"
)

var ProtoInstance component.Protocol
var protoOnce sync.Once

func NewProtocol(proto component.Protocol) component.Protocol {
	protoOnce.Do(func() {
		ProtoInstance = proto
	})

	return ProtoInstance
}
