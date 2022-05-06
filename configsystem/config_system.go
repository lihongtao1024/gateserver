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

var Instance *configs.Config
var thisOnce sync.Once

func NewInstance() *configs.Config {
	thisOnce.Do(func() {
		Instance = configs.LoadConfig(configPath)
		if Instance == nil {
			fmt.Fprintf(os.Stdout, "load config '%s' [fail].\n", configPath)
		}
	})
	return Instance
}
