package logsystem

import (
	"fmt"
	"gateserver/configsystem"
	"gateserver/internal/loggers"
	"os"
	"sync"
)

var TheLog loggers.Component
var thisOnce sync.Once

func NewLogSystemInstance(name string) loggers.Component {
	thisOnce.Do(func() {
		config := configsystem.TheConfig.GetLogAttr()
		TheLog = loggers.NewLogger(config.Flag, name, config.Output)
		if TheLog == nil {
			fmt.Fprintf(os.Stderr, "init log component [fail].\n")
		}
	})

	return TheLog
}
