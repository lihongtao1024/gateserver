package appsystem

import (
	"gateserver/internal/applications"
	"sync"
)

var Instance applications.Component
var thisOnce sync.Once

func NewInstance(comp applications.Component) {
	thisOnce.Do(func() {
		Instance = comp
	})
}
