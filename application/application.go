package application

import (
	"fmt"
	"gateserver/network"
	"gateserver/online"
	"gateserver/pkg"
	"gateserver/pkg/applications"
	"gateserver/pkg/configs"
	"gateserver/pkg/guids"
	"gateserver/pkg/loggers"
	"gateserver/pkg/timers"
	"gateserver/protocol"
	"gateserver/singleton"
	"gateserver/verify"
	"os"
	"time"
)

const appConfig = "./svrinfo.xml"
const appHeartBeatTimeout = time.Millisecond

type applicationImpl struct {
	pkg.Application
}

func NewApplication() *applicationImpl {
	app := &applicationImpl{}
	app.Application = applications.NewApplication(app)
	return app
}

func (app *applicationImpl) buildLogParam() (flag pkg.LogType, outname, outdir string) {
	attr := singleton.CfgInstance.GetLogAttr()
	flag = pkg.LogTypeSys
	if attr.Dbg {
		flag |= pkg.LogTypeDbg
	}

	if attr.Inf {
		flag |= pkg.LogTypeInf
	}

	if attr.Wrn {
		flag |= pkg.LogTypeWrn
	}

	if attr.Err {
		flag |= pkg.LogTypeErr
	}

	outname = app.GetLogicName()
	outdir = attr.Output
	return
}

func (app *applicationImpl) OnInit() bool {
	fmt.Fprintf(os.Stdout, "init %s [wait].\n", app.GetLogicName())

	singleton.NewApplication(app)
	fmt.Fprintf(os.Stdout, "init singleton<Application> [ok].\n")

	if singleton.NewConfig(configs.NewConfig(), appConfig) == nil {
		fmt.Fprintf(os.Stderr, "init singleton<Config> [fail].\n")
		return false
	}
	fmt.Fprintf(os.Stdout, "init singleton<Config> [ok].\n")

	if singleton.NewLogger(loggers.NewLogger(app.buildLogParam())) == nil {
		fmt.Fprintf(os.Stderr, "init singleton<Logger> [fail].\n")
		return false
	}
	singleton.LogInstance.Sys("init singleton<Logger> [ok].")

	if singleton.NewTimer(timers.NewComponent()) == nil {
		singleton.LogInstance.Err("init singleton<Timer> [fail].")
		return false
	}
	singleton.LogInstance.Sys("init singleton<Timer> [ok].")

	if singleton.NewGuidBuilder(guids.NewGuidBuilder(app.GetIndex())) == nil {
		singleton.LogInstance.Err("init singleton<Guid> [fail].")
		return false
	}
	singleton.LogInstance.Sys("init singleton<Guid> [ok].")

	if singleton.NewNetwork(network.NewNetwork(app.GetIndex())) == nil {
		singleton.LogInstance.Err("init singleton<Network> [fail].")
		return false
	}
	singleton.LogInstance.Sys("init singleton<Network> [ok].")

	if singleton.NewProtocol(protocol.NewProtocol()) == nil {
		singleton.LogInstance.Err("init singleton<Protocol> [fail].")
		return false
	}
	singleton.LogInstance.Sys("init singleton<Protocol> [ok].")

	if singleton.NewVerify(verify.NewVerify()) == nil {
		singleton.LogInstance.Err("init singleton<Verify> [fail].")
		return false
	}
	singleton.LogInstance.Sys("init singleton<Verify> [ok].")

	if singleton.NewOnline(online.NewOnline()) == nil {
		singleton.LogInstance.Err("init singleton<Online> [fail].")
		return false
	}
	singleton.LogInstance.Sys("init singleton<Online> [ok].")

	singleton.LogInstance.Sys("init %s [ok].", app.GetLogicName())
	return true
}

func (app *applicationImpl) OnUninit() {
	if singleton.OnlineInstance != nil {
		singleton.LogInstance.Sys("uninit singleton<Online> [ok].")
		singleton.OnlineInstance.Close()
	}

	if singleton.VerifyInstance != nil {
		singleton.LogInstance.Sys("uninit singleton<Verify> [ok].")
		singleton.VerifyInstance.Close()
	}

	if singleton.ProtoInstance != nil {
		singleton.LogInstance.Sys("uninit singleton<Protocol> [ok].")
		singleton.ProtoInstance.Close()
	}

	if singleton.NetInstance != nil {
		singleton.LogInstance.Sys("uninit singleton<Network> [ok].")
		singleton.NetInstance.Close()
	}

	if singleton.GuidInstance != nil {
		singleton.LogInstance.Sys("uninit singleton<Guid> [ok].")
		singleton.GuidInstance.Close()
	}

	if singleton.TimerInstance != nil {
		singleton.LogInstance.Sys("uninit singleton<Timer> [ok].")
		singleton.TimerInstance.Close()
	}

	if singleton.LogInstance != nil {
		singleton.LogInstance.Sys("uninit singleton<Config> [ok].")
		singleton.LogInstance.Close()
	}

	if singleton.CfgInstance != nil {
		fmt.Fprintf(os.Stdout, "uninit singleton<Config> [ok].\n")
		singleton.CfgInstance.Close()
	}

	fmt.Fprintf(os.Stdout, "uninit %s [ok].\n", app.GetLogicName())
}

func (app *applicationImpl) OnWorking() {
	busy := false
	singleton.TimerInstance.Do()
	busy = singleton.NetInstance.Do() || busy
	if !busy {
		time.Sleep(appHeartBeatTimeout)
	}
}

func (app *applicationImpl) OnClosing() bool {
	return true
}
