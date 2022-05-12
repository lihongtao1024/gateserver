package singleton

import (
	"gateserver/pkg"
	"sync"
)

var CfgInstance pkg.Config
var cfgOnce sync.Once

func NewConfig(cfg pkg.Config, path string) pkg.Config {
	cfgOnce.Do(func() {
		if !cfg.LoadConfig(path) {
			return
		}

		CfgInstance = cfg
	})

	return CfgInstance
}
