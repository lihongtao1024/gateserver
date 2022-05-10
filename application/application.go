package application

import (
	"fmt"
	"gateserver/appsystem"
	"gateserver/configsystem"
	"gateserver/guidsystem"
	"gateserver/internal/applications"
	"gateserver/internal/guids"
	"gateserver/logsystem"
	"gateserver/netsystem"
	"gateserver/protosystem"
	"gateserver/protosystem/protocols"
	"gateserver/timersystem"
	"gateserver/verifysystem"
	"os"
	"time"
)

const (
	appHeartBeatTimeout = time.Millisecond
)

type Application struct {
	applications.Component
}

func NewApplication() *Application {
	app := &Application{}
	app.Component = applications.NewApplication(app)
	return app
}

func (app *Application) OnInit() bool {
	fmt.Fprintf(os.Stdout, "init %s [wait].\n", app.GetLogicName())
	appsystem.NewInstance(app)

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

	if netsystem.NewInstance(app.GetIndex()) == nil {
		return false
	}
	logsystem.Instance.Sys("init net system [ok].")

	if guidsystem.NewInstance(app.GetIndex()) == nil {
		return false
	}
	app.Component.SetRid(guidsystem.Instance.CreateGuid(guids.GuidGlobal))
	logsystem.Instance.Sys("init guid system [ok].")

	if protosystem.NewInstance(
		app.GetIndex(),
		protocols.NewGT2WSProto(),
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

func (app *Application) OnUninit() {
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

	if guidsystem.Instance != nil {
		logsystem.Instance.Sys("uninit guid system [ok].")
		guidsystem.Instance.Close()
	}

	if logsystem.Instance != nil {
		logsystem.Instance.Sys("uninit log system [ok].")
		logsystem.Instance.Close()
	}

	fmt.Fprintf(os.Stdout, "uninit %s [ok].\n", app.GetLogicName())
}

func (app *Application) OnWorking() {
	busy := false
	timersystem.Instance.Do()
	busy = netsystem.Instance.Do() || busy
	if !busy {
		time.Sleep(appHeartBeatTimeout)
	}
}

func (app *Application) OnClosing() bool {
	return true
}
