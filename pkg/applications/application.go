package applications

import (
	"gateserver/pkg"
	"gateserver/pkg/configs"
	"strconv"
	"sync"
	"sync/atomic"
)

const (
	appIdle    = 0
	appRunning = 1
	appClosing = -1
	appClosed  = -2
)

type applicationImpl struct {
	appStatus int32
	appIndex  int
	appRid    pkg.Guid
	appType   pkg.ServerType
	appEnver  string
	startOnce *sync.Once
	stopOnce  *sync.Once
	waitGroup *sync.WaitGroup
	pkg.ApplicationDispatcher
}

func NewApplication(dispatcher pkg.ApplicationDispatcher) pkg.Application {
	app := &applicationImpl{
		appStatus:             appIdle,
		waitGroup:             &sync.WaitGroup{},
		startOnce:             &sync.Once{},
		stopOnce:              &sync.Once{},
		ApplicationDispatcher: dispatcher,
	}
	return app
}

func (app *applicationImpl) GetType() pkg.ServerType {
	return pkg.ServerIdGt
}

func (app *applicationImpl) GetLogicName() string {
	return configs.GetServerName(pkg.ServerIdGt) +
		strconv.Itoa(app.appIndex)
}

func (app *applicationImpl) GetIndex() int {
	return app.appIndex
}

func (app *applicationImpl) GetEnvir() string {
	return app.appEnver
}

func (app *applicationImpl) GetRid() pkg.Guid {
	return app.appRid
}

func (app *applicationImpl) SetRid(guid pkg.Guid) {
	app.appRid = guid
}

func (app *applicationImpl) Start(typ1 pkg.ServerType, idx int, env string) {
	app.startOnce.Do(func() {
		app.waitGroup.Add(1)

		go func() {
			defer app.waitGroup.Done()

			app.appType = typ1
			app.appIndex = idx
			app.appEnver = env
			app.setStatus(appRunning)
			app.doApp()
		}()
	})
}

func (app *applicationImpl) Stop() {
	app.stopOnce.Do(func() {
		if !app.isStatus(appRunning) {
			return
		}

		app.setStatus(appClosing)
		app.waitGroup.Wait()
	})
}

func (app *applicationImpl) setStatus(status int32) {
	atomic.StoreInt32(&app.appStatus, status)
}

func (app *applicationImpl) isStatus(status int32) bool {
	return atomic.LoadInt32(&app.appStatus) == status
}

func (app *applicationImpl) doApp() {
	if app.OnInit() {
		for !app.IsClosed() {
			if app.isStatus(appClosing) {
				if app.OnClosing() {
					app.OnUninit()
					app.setStatus(appClosed)
				}
			}
			app.OnWorking()
		}
	} else {
		app.OnUninit()
		app.setStatus(appClosed)
	}
}

func (app *applicationImpl) IsClosed() bool {
	return app.isStatus(appClosed)
}
