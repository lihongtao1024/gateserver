package pkg

type LogType int

const (
	LogTypeNull = LogType(iota)
	LogTypeDbg  = LogType(1 << (iota - 1))
	LogTypeInf
	LogTypeWrn
	LogTypeErr
	LogTypeSys
	LogTypeExit = LogType(-1)
)

type Logger interface {
	Dbg(format string, content ...interface{})
	Inf(format string, content ...interface{})
	Wrn(format string, content ...interface{})
	Err(format string, content ...interface{})
	Sys(format string, content ...interface{})
	Close()
}
