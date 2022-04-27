package timers

type Timer interface {
	GetInterval() uint32
	GetCount() uint32
	SetData(interface{})
	GetData() interface{}
}

type Callback interface {
	OnTimer()
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
	tmrData     interface{}
}

func (timer *timerImpl) GetInterval() uint32 {
	return timer.tmrInterval
}

func (timer *timerImpl) GetCount() uint32 {
	return timer.tmrCount
}

func (timer *timerImpl) SetData(o interface{}) {
	timer.tmrData = o
}

func (timer *timerImpl) GetData() interface{} {
	return timer.tmrData
}
