package timers

type Timer interface {
	GetInterval() uint32
	GetCount() uint32
}

type Callback interface {
	OnTimer(Timer)
}

type loopLink struct {
	tmrImpl *timerImpl
	pPrev   *loopLink
	pNext   *loopLink
}

type timerHeader = loopLink

type timerImpl struct {
	tmrInterval uint32
	tmrTimeout  uint32
	tmrCount    uint32
	tmrCallback Callback
	tmrLink     loopLink
}

func (timer *timerImpl) GetInterval() uint32 {
	return timer.tmrInterval
}

func (timer *timerImpl) GetCount() uint32 {
	return timer.tmrCount
}
