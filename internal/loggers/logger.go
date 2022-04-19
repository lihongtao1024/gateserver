package loggers

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Component interface {
	Dbg(format string, content ...interface{})
	Inf(format string, content ...interface{})
	Wrn(format string, content ...interface{})
	Err(format string, content ...interface{})
	Sys(format string, content ...interface{})
	Close()
}

const (
	LogTypeNull = 0
	LogTypeDbg  = 1
	LogTypeInf  = 2
	LogTypeWrn  = 4
	LogTypeErr  = 8
	LogTypeSys  = 16
	LogTypeExit = -1
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

type logEvent struct {
	flag    int
	format  string
	content []interface{}
}

type logComponent struct {
	logFlag    int
	logYear    int
	logMonth   int
	logDay     int
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
	logEvents  chan *logEvent
	waitGroup  *sync.WaitGroup
	exitOnce   *sync.Once
}

func loggerWrite(comp *logComponent, event *logEvent) {
	local := time.Now().Local()
	if local.Year() != comp.logYear ||
		local.Month() != time.Month(comp.logMonth) ||
		local.Day() != comp.logDay {
		comp.logFile.Close()
		loggerInit(comp)
	}

	content := fmt.Sprintf(event.format, event.content...)
	switch event.flag {
	case LogTypeDbg:
		comp.dbgLogger.Println(content)
		comp.dbgConsole.Println(content)
	case LogTypeInf:
		comp.infLogger.Println(content)
		comp.infConsole.Println(content)
	case LogTypeWrn:
		comp.wrnLogger.Println(content)
		comp.wrnConsole.Println(content)
	case LogTypeErr:
		comp.errLogger.Println(content)
		comp.errConsole.Println(content)
	default:
		comp.sysLogger.Println(content)
		comp.sysConsole.Println(content)
	}
}

func loggerHandler(comp *logComponent) {
	comp.waitGroup.Add(1)

	defer func() {
		comp.logFile.Close()
		close(comp.logEvents)
		comp.waitGroup.Done()
	}()

	for {
		message := <-comp.logEvents
		if message.flag == LogTypeExit {
			return
		}

		if comp.logFlag&message.flag == 0 {
			continue
		}

		loggerWrite(comp, message)
	}
}

func loggerInit(comp *logComponent) bool {
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

func NewLogger(flag int, outname, outdir string) Component {
	comp := &logComponent{}
	comp.logFlag = flag | LogTypeSys
	comp.outputName = outname
	comp.outputDir = outdir
	comp.logEvents = make(chan *logEvent, 4096)
	comp.waitGroup = &sync.WaitGroup{}
	comp.exitOnce = &sync.Once{}

	if !loggerInit(comp) {
		return nil
	}

	go loggerHandler(comp)
	return comp
}

func (comp *logComponent) Dbg(format string, content ...interface{}) {
	comp.logEvents <- &logEvent{LogTypeDbg, format, content}
}

func (comp *logComponent) Inf(format string, content ...interface{}) {
	comp.logEvents <- &logEvent{LogTypeInf, format, content}
}

func (comp *logComponent) Wrn(format string, content ...interface{}) {
	comp.logEvents <- &logEvent{LogTypeWrn, format, content}
}

func (comp *logComponent) Err(format string, content ...interface{}) {
	comp.logEvents <- &logEvent{LogTypeErr, format, content}
}

func (comp *logComponent) Sys(format string, content ...interface{}) {
	comp.logEvents <- &logEvent{LogTypeSys, format, content}
}

func (comp *logComponent) Close() {
	comp.exitOnce.Do(func() {
		comp.logEvents <- &logEvent{flag: LogTypeExit}
		comp.waitGroup.Wait()
	})
}
