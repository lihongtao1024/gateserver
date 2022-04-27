package timers

import (
	"time"
)

const (
	InfiniteTimer = ^uint32(0)
)
const (
	BucketB1 = uint32(8)
	BucketB2 = uint32(6)
	BucketS2 = uint32(8)
	BucketS3 = uint32(14)
	BucketS4 = uint32(20)
	BucketS5 = uint32(26)
	BucketA1 = uint32(1) << BucketB1
	BucketA2 = uint32(1) << BucketB2
)

type Component interface {
	AddTimer(callback Callback, interval uint32, count uint32) Timer
	ModifyTimer(t Timer, callback Callback, interval uint32, count uint32) bool
	DelTimer(t Timer)
	Do()
	Close()
}

type timerComponent struct {
	lastTick    uint32
	tmrBuckets1 [BucketA1]timerHeader
	tmrBuckets2 [BucketA2]timerHeader
	tmrBuckets3 [BucketA2]timerHeader
	tmrBuckets4 [BucketA2]timerHeader
	tmrBuckets5 [BucketA2]timerHeader
}

func NewComponent() Component {
	comp := &timerComponent{
		lastTick: uint32(time.Now().UnixMilli()),
	}

	comp.initBucket(comp.tmrBuckets1[:])
	comp.initBucket(comp.tmrBuckets2[:])
	comp.initBucket(comp.tmrBuckets3[:])
	comp.initBucket(comp.tmrBuckets4[:])
	comp.initBucket(comp.tmrBuckets5[:])
	return comp
}

func (comp *timerComponent) initBucket(buckets []timerHeader) {
	l := len(buckets)
	for i := 0; i < l; i++ {
		bucket := &buckets[i]
		bucket.pPrev = bucket
		bucket.pNext = bucket
	}
}

func (comp *timerComponent) addTimerImpl(timer *timerImpl) {
	var bucket *timerHeader

	interval := timer.tmrTimeout - comp.lastTick
	switch {
	case interval < uint32(1)<<BucketS2:
		{
			bucket = &comp.tmrBuckets1[timer.tmrTimeout&(uint32(1)<<BucketB1-1)]
		}
	case interval < uint32(1)<<BucketS3:
		{
			bucket = &comp.tmrBuckets2[(timer.tmrTimeout>>BucketS2)&(uint32(1)<<BucketB2-1)]
		}
	case interval < uint32(1)<<BucketS4:
		{
			bucket = &comp.tmrBuckets3[(timer.tmrTimeout>>BucketS3)&(uint32(1)<<BucketB2-1)]
		}
	case interval < uint32(1)<<BucketS5:
		{
			bucket = &comp.tmrBuckets4[(timer.tmrTimeout>>BucketS4)&(uint32(1)<<BucketB2-1)]
		}
	default:
		{
			bucket = &comp.tmrBuckets5[(timer.tmrTimeout>>BucketS5)&(uint32(1)<<BucketB2-1)]
		}
	}

	bucket.pPrev.pNext = &timer.tmrLink
	timer.tmrLink.pNext = bucket
	timer.tmrLink.pPrev = bucket.pPrev
	bucket.pPrev = &timer.tmrLink
}

func (comp *timerComponent) onTimerImpl(timer *timerImpl) {
	timer.tmrLink.pPrev = &timer.tmrLink
	timer.tmrLink.pNext = &timer.tmrLink

	if timer.tmrCount == InfiniteTimer {
		timer.tmrCallback.OnTimer()
	} else {
		timer.tmrCount--
		timer.tmrCallback.OnTimer()
	}

	if timer.tmrCount == 0 {
		return
	}

	if timer.tmrLink.pPrev != &timer.tmrLink ||
		timer.tmrLink.pNext != &timer.tmrLink {
		return
	}

	timer.tmrTimeout = comp.lastTick + timer.tmrInterval
	comp.addTimerImpl(timer)
}

func (comp *timerComponent) wheelImpl(buckets []timerHeader, index uint32,
	fn func(*timerImpl)) bool {
	th := &buckets[index]

	var hr timerHeader
	hr.pPrev = th.pPrev
	hr.pPrev.pNext = &hr
	hr.pNext = th.pNext
	hr.pNext.pPrev = &hr

	th.pPrev = th
	th.pNext = th

	for tmr := hr.pNext; tmr != &hr; {
		tmp := tmr.pNext
		fn(tmr.tmrImpl)
		tmr = tmp
	}

	return index == 0
}

func (comp *timerComponent) AddTimer(callback Callback, interval uint32,
	count uint32) Timer {
	if interval == 0 {
		interval = 1
	}

	if count == 0 {
		return nil
	}

	timer := &timerImpl{
		tmrInterval: interval,
		tmrTimeout:  interval + comp.lastTick,
		tmrCount:    count,
		tmrCallback: callback,
	}
	timer.tmrLink.tmrImpl = timer

	comp.addTimerImpl(timer)
	return timer
}

func (comp *timerComponent) ModifyTimer(t Timer, callback Callback, interval uint32,
	count uint32) bool {
	timer := t.(*timerImpl)
	comp.DelTimer(t)

	timer.tmrInterval = interval
	timer.tmrTimeout = interval + comp.lastTick
	timer.tmrCount = count
	timer.tmrCallback = callback
	comp.addTimerImpl(timer)
	return true
}

func (comp *timerComponent) DelTimer(t Timer) {
	timer := t.(*timerImpl)
	timer.tmrCount = 0
	timer.tmrLink.pNext.pPrev = timer.tmrLink.pPrev
	timer.tmrLink.pPrev.pNext = timer.tmrLink.pNext
	timer.tmrLink.pNext = &timer.tmrLink
	timer.tmrLink.pPrev = &timer.tmrLink
}

func (comp *timerComponent) Do() {
	tick := uint32(time.Now().UnixMilli())

	for ; int(tick-comp.lastTick) > 0; comp.lastTick++ {
		index := comp.lastTick & (uint32(1)<<BucketS2 - 1)
		if index == 0 &&
			comp.wheelImpl(
				comp.tmrBuckets2[:],
				(comp.lastTick>>BucketS2)&(uint32(1)<<BucketB2-1),
				comp.addTimerImpl,
			) &&
			comp.wheelImpl(
				comp.tmrBuckets3[:],
				(comp.lastTick>>BucketS3)&(uint32(1)<<BucketB2-1),
				comp.addTimerImpl,
			) &&
			comp.wheelImpl(
				comp.tmrBuckets4[:],
				(comp.lastTick>>BucketS4)&(uint32(1)<<BucketB2-1),
				comp.addTimerImpl,
			) {
			comp.wheelImpl(
				comp.tmrBuckets5[:],
				(comp.lastTick>>BucketS5)&(uint32(1)<<BucketB2-1),
				comp.addTimerImpl,
			)
		}

		comp.wheelImpl(comp.tmrBuckets1[:], index, comp.onTimerImpl)
	}
}

func (comp *timerComponent) Close() {

}
