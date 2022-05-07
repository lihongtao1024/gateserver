package applications

import (
	"gateserver/internal/configs"
	"gateserver/internal/guids"
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

type Component interface {
	GetType() int
	GetLogicName() string
	GetIndex() int
	GetEnvir() string
	GetRid() guids.Guid
	SetRid(guid guids.Guid)
	Start(ty int, idx int, env string)
	Stop()
	IsClosed() bool
}

type Dispatcher interface {
	OnInit() bool
	OnUninit()
	OnClosing() bool
	OnWorking()
}

type appComponent struct {
	appStatus int32
	appType   int
	appIndex  int
	appRid    guids.Guid
	appEnver  string
	startOnce *sync.Once
	stopOnce  *sync.Once
	waitGroup *sync.WaitGroup
	Dispatcher
}

func NewApplication(dispatcher Dispatcher) Component {
	app := &appComponent{
		appStatus:  appIdle,
		waitGroup:  &sync.WaitGroup{},
		startOnce:  &sync.Once{},
		stopOnce:   &sync.Once{},
		Dispatcher: dispatcher,
	}
	return app
}

func (app *appComponent) GetType() int {
	return configs.ServerIdGt
}

func (app *appComponent) GetLogicName() string {
	return configs.GetServerName(configs.ServerIdGt) +
		strconv.Itoa(app.appIndex)
}

func (app *appComponent) GetIndex() int {
	return app.appIndex
}

func (app *appComponent) GetEnvir() string {
	return app.appEnver
}

func (app *appComponent) GetRid() guids.Guid {
	return app.appRid
}

func (app *appComponent) SetRid(guid guids.Guid) {
	app.appRid = guid
}

func (app *appComponent) Start(ty int, idx int, env string) {
	app.startOnce.Do(func() {
		app.waitGroup.Add(1)

		go func() {
			defer app.waitGroup.Done()

			app.appType = ty
			app.appIndex = idx
			app.appEnver = env
			app.setStatus(appRunning)
			app.doApp()
		}()
	})
}

func (app *appComponent) Stop() {
	app.stopOnce.Do(func() {
		if !app.isStatus(appRunning) {
			return
		}

		app.setStatus(appClosing)
		app.waitGroup.Wait()
	})
}

func (app *appComponent) setStatus(status int32) {
	atomic.StoreInt32(&app.appStatus, status)
}

func (app *appComponent) isStatus(status int32) bool {
	return atomic.LoadInt32(&app.appStatus) == status
}

func (app *appComponent) doApp() {
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

func (app *appComponent) IsClosed() bool {
	return app.isStatus(appClosed)
}
