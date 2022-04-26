package applications

import (
	"fmt"
	"gateserver/configsystem"
	"gateserver/internal/configs"
	"gateserver/logsystem"
	"gateserver/netsystem"
	"gateserver/protosystem"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	appIdle    = 0
	appRunning = 1
	appClosing = -1
	appClosed  = -2
)

const (
	appHeartBeatTimeout = time.Millisecond
)

type Application struct {
	appStatus int32
	appIndex  int
	appEnver  string
	startOnce *sync.Once
	stopOnce  *sync.Once
	waitGroup *sync.WaitGroup
}

var TheApp *Application
var thisOnce sync.Once

func NewApplicationInstance() *Application {
	thisOnce.Do(func() {
		TheApp = &Application{
			appStatus: appIdle,
			waitGroup: &sync.WaitGroup{},
			startOnce: &sync.Once{},
			stopOnce:  &sync.Once{},
		}
	})

	return TheApp
}

func (app *Application) GetType() int {
	return configs.ServerIdGt
}

func (app *Application) GetLogicName() string {
	return configs.GetServerName(configs.ServerIdGt) + strconv.Itoa(app.appIndex)
}

func (app *Application) Start(index int, envir string) {
	app.startOnce.Do(func() {
		app.waitGroup.Add(1)

		go func() {
			app.appIndex = index
			app.appEnver = envir
			app.setStatus(appRunning)
			app.doApp()
			app.waitGroup.Done()
		}()
	})
}

func (app *Application) Stop() {
	app.stopOnce.Do(func() {
		if !app.isStatus(appRunning) {
			return
		}

		app.setStatus(appClosing)
		app.waitGroup.Wait()
	})
}

func (app *Application) setStatus(status int32) {
	atomic.StoreInt32(&app.appStatus, status)
}

func (app *Application) isStatus(status int32) bool {
	return atomic.LoadInt32(&app.appStatus) == status
}

func (app *Application) doApp() {
	if app.appInit() {
		busy := false
		for !app.IsClosed() {
			if app.isStatus(appClosing) {
				app.setStatus(appClosed)
			}
			busy = netsystem.TheNet.Do() || busy
			if !busy {
				time.Sleep(appHeartBeatTimeout)
			}
		}
	} else {
		app.setStatus(appClosed)
	}

	app.appUninit()
}

func (app *Application) appInit() bool {
	fmt.Fprintf(os.Stdout, "init %s [wait].\n", app.GetLogicName())

	if configsystem.NewConfigSystemInstance() == nil {
		return false
	}

	if logsystem.NewLogSystemInstance(app.GetLogicName()) == nil {
		return false
	}
	logsystem.TheLog.Sys("init log system [ok].")

	if netsystem.NewNetSystemInstance(app.appIndex) == nil {
		return false
	}
	logsystem.TheLog.Sys("init net system [ok].")

	if protosystem.NewProtoSystemInstance(app.appIndex) == nil {
		return false
	}
	logsystem.TheLog.Sys("init proto system [ok].")

	logsystem.TheLog.Sys("init %s [ok].", app.GetLogicName())
	return true
}

func (app *Application) appUninit() {
	if logsystem.TheLog != nil {
		logsystem.TheLog.Sys("uninit %s [wait].", app.GetLogicName())
	} else {
		fmt.Fprintf(os.Stdout, "uninit %s [wait].\n", app.GetLogicName())
	}

	if netsystem.TheNet != nil {
		logsystem.TheLog.Sys("uninit net system [ok].")
		netsystem.TheNet.Close()
	}

	if protosystem.TheProto != nil {
		logsystem.TheLog.Sys("uninit proto system [ok].")
		protosystem.TheProto.Close()
	}

	if logsystem.TheLog != nil {
		logsystem.TheLog.Sys("uninit log system [ok].")
		logsystem.TheLog.Close()
	}

	fmt.Fprintf(os.Stdout, "uninit %s [ok].\n", app.GetLogicName())
}

func (app *Application) IsClosed() bool {
	return app.isStatus(appClosed)
}
