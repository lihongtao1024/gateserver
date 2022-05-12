package pkg

const InfiniteTimer = ^uint32(0)

type Timer interface {
	GetInterval() uint32
	GetCount() uint32
	SetData(interface{})
	GetData() interface{}
}

type TimerDispatcher interface {
	OnTimer()
}

type TimerComponent interface {
	AddTimer(dispatcher TimerDispatcher, interval uint32, count uint32) Timer
	ModifyTimer(t Timer, dispatcher TimerDispatcher, interval uint32, count uint32) bool
	DelTimer(t Timer)
	Do()
	Close()
}
