package singleton

import (
	"gateserver/component"
	"gateserver/pkg"
	"sync"
)

var AppInstance *applicationImpl
var appOnce sync.Once

type applicationImpl struct {
	pkg.Application
}

func NewApplication(app pkg.Application) *applicationImpl {
	appOnce.Do(func() {
		AppInstance = &applicationImpl{app}
	})

	return AppInstance
}

func (app *applicationImpl) GetClientType() component.ClientType {
	return component.ClientTypeMT2
}
