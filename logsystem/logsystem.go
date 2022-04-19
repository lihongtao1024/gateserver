package logsystem

import (
	"fmt"
	"gateserver/configsystem"
	"gateserver/internal/loggers"
	"os"
	"sync"
)

var This loggers.Component
var thisOnce sync.Once

func NewLogSystemInstance(name string) loggers.Component {
	thisOnce.Do(func() {
		config := configsystem.This.GetLogAttr()
		This = loggers.NewLogger(config.Flag, name, config.Output)
		if This == nil {
			fmt.Fprintf(os.Stderr, "init log component [fail].\n")
		}
	})

	return This
}
