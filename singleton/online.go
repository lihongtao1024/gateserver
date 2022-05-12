package singleton

import (
	"gateserver/component"
	"sync"
)

var OnlineInstance component.Online
var onlineOnce sync.Once

func NewOnline(online component.Online) component.Online {
	onlineOnce.Do(func() {
		OnlineInstance = online
	})

	return online
}
