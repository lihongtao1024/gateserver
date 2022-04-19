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

var This *configs.Config
var thisOnce sync.Once

func NewConfigSystemInstance() *configs.Config {
	thisOnce.Do(func() {
		This = configs.LoadConfig(configPath)
		if This == nil {
			fmt.Fprintf(os.Stdout, "load config '%s' [fail].\n", configPath)
		}
	})
	return This
}
