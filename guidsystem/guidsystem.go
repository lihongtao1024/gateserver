package guidsystem

import (
	"gateserver/internal/guids"
	"sync"
)

var Instance guids.Component
var thisOnce sync.Once

func NewInstance(index int) guids.Component {
	thisOnce.Do(func() {
		Instance = guids.NewComponent(uint8(index))
	})

	return Instance
}
