package applications

import (
	"fmt"
	"gateserver/configsystem"
	"gateserver/internal/configs"
	"gateserver/logsystem"
	"gateserver/netsystem"
	"gateserver/protosystem"
	"gateserver/protosystem/messages"
	"gateserver/timersystem"
	"gateserver/verifysystem"
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

var Instance *Application
var thisOnce sync.Once

func NewApplicationInstance() *Application {
	thisOnce.Do(func() {
		Instance = &Application{
			appStatus: appIdle,
			waitGroup: &sync.WaitGroup{},
			startOnce: &sync.Once{},
			stopOnce:  &sync.Once{},
		}
	})

	return Instance
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
				app.appUninit()
				app.setStatus(appClosed)
			}
			timersystem.Instance.Do()
			busy = netsystem.Instance.Do() || busy
			if !busy {
				time.Sleep(appHeartBeatTimeout)
			}
		}
	} else {
		app.appUninit()
		app.setStatus(appClosed)
	}
}

func (app *Application) appInit() bool {
	fmt.Fprintf(os.Stdout, "init %s [wait].\n", app.GetLogicName())

	if configsystem.NewInstance() == nil {
		return false
	}

	if logsystem.NewInstance(app.GetLogicName()) == nil {
		return false
	}
	logsystem.Instance.Sys("init log system [ok].")

	if timersystem.NewInstance() == nil {
		return false
	}
	logsystem.Instance.Sys("init timer system [ok].")

	if netsystem.NewInstance(app.appIndex) == nil {
		return false
	}
	logsystem.Instance.Sys("init net system [ok].")

	if protosystem.NewInstance(
		app.appIndex,
		messages.NewGT2WSProto(),
	) == nil {
		return false
	}
	logsystem.Instance.Sys("init proto system [ok].")

	if verifysystem.NewInstance() == nil {
		return false
	}
	logsystem.Instance.Sys("init verify system [ok].")

	logsystem.Instance.Sys("init %s [ok].", app.GetLogicName())
	return true
}

func (app *Application) appUninit() {
	if logsystem.Instance != nil {
		logsystem.Instance.Sys("uninit %s [wait].", app.GetLogicName())
	} else {
		fmt.Fprintf(os.Stdout, "uninit %s [wait].\n", app.GetLogicName())
	}

	if verifysystem.Instance != nil {
		logsystem.Instance.Sys("uninit verify system [ok].")
		verifysystem.Instance.Close()
	}

	if netsystem.Instance != nil {
		logsystem.Instance.Sys("uninit net system [ok].")
		netsystem.Instance.Close()
	}

	if protosystem.Instance != nil {
		logsystem.Instance.Sys("uninit proto system [ok].")
		protosystem.Instance.Close()
	}

	if timersystem.Instance != nil {
		logsystem.Instance.Sys("uninit timer system [ok].")
		timersystem.Instance.Close()
	}

	if logsystem.Instance != nil {
		logsystem.Instance.Sys("uninit log system [ok].")
		logsystem.Instance.Close()
	}

	fmt.Fprintf(os.Stdout, "uninit %s [ok].\n", app.GetLogicName())
}

func (app *Application) IsClosed() bool {
	return app.isStatus(appClosed)
}
