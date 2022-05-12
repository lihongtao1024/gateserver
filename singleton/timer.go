package singleton

import (
	"gateserver/pkg"
	"sync"
)

var TimerInstance pkg.TimerComponent
var timerOnce sync.Once

func NewTimer(comp pkg.TimerComponent) pkg.TimerComponent {
	timerOnce.Do(func() {
		TimerInstance = comp
	})

	return TimerInstance
}
