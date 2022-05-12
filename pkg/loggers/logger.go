package loggers

import (
	"fmt"
	"gateserver/pkg"
	"log"
	"os"
	"sync"
	"time"
)

const (
	fOpen  = os.O_CREATE | os.O_APPEND | os.O_WRONLY
	mOpen  = os.ModeAppend | os.ModePerm
	fLog   = log.Ldate | log.Lmicroseconds
	preDbg = "[DBG] "
	preInf = "[INF] "
	preWrn = "[WRN] "
	preErr = "[ERR] "
	preSys = "[SYS] "
)

type loggerEvt struct {
	flag    pkg.LogType
	format  string
	content []interface{}
}

type loggerImpl struct {
	logYear    int
	logMonth   int
	logDay     int
	logFlag    pkg.LogType
	logFile    *os.File
	dbgLogger  *log.Logger
	dbgConsole *log.Logger
	infLogger  *log.Logger
	infConsole *log.Logger
	wrnLogger  *log.Logger
	wrnConsole *log.Logger
	errLogger  *log.Logger
	errConsole *log.Logger
	sysLogger  *log.Logger
	sysConsole *log.Logger
	outputName string
	outputDir  string
	logEvents  chan *loggerEvt
	waitGroup  *sync.WaitGroup
	exitOnce   *sync.Once
}

func NewLogger(flag pkg.LogType, outname, outdir string) pkg.Logger {
	comp := &loggerImpl{}
	comp.logFlag = flag
	comp.outputName = outname
	comp.outputDir = outdir
	comp.logEvents = make(chan *loggerEvt, 4096)
	comp.waitGroup = &sync.WaitGroup{}
	comp.exitOnce = &sync.Once{}

	if !loggerInit(comp) {
		return nil
	}

	go loggerHandler(comp)
	return comp
}

func loggerWrite(comp *loggerImpl, event *loggerEvt) {
	local := time.Now().Local()
	if local.Year() != comp.logYear ||
		local.Month() != time.Month(comp.logMonth) ||
		local.Day() != comp.logDay {
		comp.logFile.Close()
		loggerInit(comp)
	}

	content := fmt.Sprintf(event.format, event.content...)
	switch event.flag {
	case pkg.LogTypeDbg:
		comp.dbgLogger.Println(content)
		comp.dbgConsole.Println(content)
	case pkg.LogTypeInf:
		comp.infLogger.Println(content)
		comp.infConsole.Println(content)
	case pkg.LogTypeWrn:
		comp.wrnLogger.Println(content)
		comp.wrnConsole.Println(content)
	case pkg.LogTypeErr:
		comp.errLogger.Println(content)
		comp.errConsole.Println(content)
	default:
		comp.sysLogger.Println(content)
		comp.sysConsole.Println(content)
	}
}

func loggerHandler(comp *loggerImpl) {
	comp.waitGroup.Add(1)

	defer func() {
		comp.logFile.Sync()
		comp.logFile.Close()
		close(comp.logEvents)
		comp.waitGroup.Done()
	}()

	for {
		message := <-comp.logEvents
		if message.flag == pkg.LogTypeExit {
			return
		}

		if comp.logFlag&message.flag == 0 {
			continue
		}

		loggerWrite(comp, message)
	}
}

func loggerInit(comp *loggerImpl) bool {
	local := time.Now().Local()
	midname := fmt.Sprintf("%04d_%02d_%02d", local.Year(), local.Month(), local.Day())
	file, err := os.OpenFile(comp.outputDir+"/"+comp.outputName+midname+".log", fOpen, mOpen)
	if err != nil {
		return false
	}

	comp.logFile = file
	comp.dbgLogger = log.New(file, preDbg, fLog)
	comp.dbgConsole = log.New(os.Stdout, preDbg, fLog)
	comp.infLogger = log.New(file, preInf, fLog)
	comp.infConsole = log.New(os.Stdout, preInf, fLog)
	comp.wrnLogger = log.New(file, preWrn, fLog)
	comp.wrnConsole = log.New(os.Stdout, preWrn, fLog)
	comp.errLogger = log.New(file, preErr, fLog)
	comp.errConsole = log.New(os.Stderr, preErr, fLog)
	comp.sysLogger = log.New(file, preSys, fLog)
	comp.sysConsole = log.New(os.Stdout, preSys, fLog)

	comp.logYear = local.Year()
	comp.logMonth = int(local.Month())
	comp.logDay = local.Day()

	return true
}

func (comp *loggerImpl) Dbg(format string, content ...interface{}) {
	comp.logEvents <- &loggerEvt{pkg.LogTypeDbg, format, content}
}

func (comp *loggerImpl) Inf(format string, content ...interface{}) {
	comp.logEvents <- &loggerEvt{pkg.LogTypeInf, format, content}
}

func (comp *loggerImpl) Wrn(format string, content ...interface{}) {
	comp.logEvents <- &loggerEvt{pkg.LogTypeWrn, format, content}
}

func (comp *loggerImpl) Err(format string, content ...interface{}) {
	comp.logEvents <- &loggerEvt{pkg.LogTypeErr, format, content}
}

func (comp *loggerImpl) Sys(format string, content ...interface{}) {
	comp.logEvents <- &loggerEvt{pkg.LogTypeSys, format, content}
}

func (comp *loggerImpl) Close() {
	comp.exitOnce.Do(func() {
		comp.logEvents <- &loggerEvt{flag: pkg.LogTypeExit}
		comp.waitGroup.Wait()
	})
}
