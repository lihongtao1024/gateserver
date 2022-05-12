package singleton

import (
	"gateserver/pkg"
	"sync"
)

var GuidInstance pkg.GuidBuilder
var guidOnce sync.Once

func NewGuidBuilder(builder pkg.GuidBuilder) pkg.GuidBuilder {
	guidOnce.Do(func() {
		GuidInstance = builder
	})

	return GuidInstance
}
