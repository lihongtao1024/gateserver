package singleton

import (
	"gateserver/pkg"
	"sync"
)

var LogInstance pkg.Logger
var logOnce sync.Once

func NewLogger(log pkg.Logger) pkg.Logger {
	logOnce.Do(func() {
		LogInstance = log
	})

	return LogInstance
}
