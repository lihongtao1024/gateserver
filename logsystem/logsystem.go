package logsystem

import (
	"fmt"
	"gateserver/configsystem"
	"gateserver/internal/loggers"
	"os"
	"sync"
)

var Instance loggers.Component
var thisOnce sync.Once

func NewLogSystemInstance(name string) loggers.Component {
	thisOnce.Do(func() {
		config := configsystem.Instance.GetLogAttr()
		Instance = loggers.NewLogger(config.Flag, name, config.Output)
		if Instance == nil {
			fmt.Fprintf(os.Stderr, "init log component [fail].\n")
		}
	})

	return Instance
}
