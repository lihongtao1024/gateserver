package timersystem

import (
	"fmt"
	"gateserver/internal/timers"
	"os"
	"sync"
)

var Instance timers.Component
var thisOnce sync.Once

func NewInstance() timers.Component {
	thisOnce.Do(func() {
		Instance = timers.NewComponent()
		if Instance == nil {
			fmt.Fprintf(os.Stderr, "init timer system [fail].\n")
		}
	})

	return Instance
}
