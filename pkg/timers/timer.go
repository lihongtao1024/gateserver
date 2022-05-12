package timers

import "gateserver/pkg"

type loopLink struct {
	tmrImpl *timerImpl
	pPrev   *loopLink
	pNext   *loopLink
}

type timerHeader = loopLink

type timerImpl struct {
	tmrInterval   uint32
	tmrTimeout    uint32
	tmrCount      uint32
	tmrLink       loopLink
	tmrData       interface{}
	tmrDispatcher pkg.TimerDispatcher
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
