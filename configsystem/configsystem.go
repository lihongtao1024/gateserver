package configsystem

import (
	"fmt"
	"gateserver/internal/configs"
	"os"
	"sync"
)

const (
	configPath = "./svrinfo.xml"
)

var TheConfig *configs.Config
var thisOnce sync.Once

func NewConfigSystemInstance() *configs.Config {
	thisOnce.Do(func() {
		TheConfig = configs.LoadConfig(configPath)
		if TheConfig == nil {
			fmt.Fprintf(os.Stdout, "load config '%s' [fail].\n", configPath)
		}
	})
	return TheConfig
}
