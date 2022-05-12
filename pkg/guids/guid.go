package guids

import (
	"gateserver/pkg"
	"time"
)

/*
 64 bit Guid:
     type         sid         time         inc
 |63......56|55........48|47........16|15........0
     8bit        8 bit        32 bit      16 bit
*/

type guidbuilderImpl struct {
	guidArray [pkg.GuidMax]pkg.Guid
}

func NewGuidBuilder(index int) pkg.GuidBuilder {
	comp := &guidbuilderImpl{}
	times := uint32(time.Now().Local().Unix())

	for i := 0; i < len(comp.guidArray); i++ {
		guid := &comp.guidArray[i]
		*guid = pkg.Guid(i)<<56 | pkg.Guid(index)<<48 | pkg.Guid(times)<<16
	}

	return comp
}

func (comp *guidbuilderImpl) CreateGuid(i pkg.GuidType) pkg.Guid {
	if i <= pkg.GuidNull || i >= pkg.GuidMax {
		panic("unexpected Guid type")
	}

	comp.guidArray[i]++
	return comp.guidArray[i]
}

func (comp *guidbuilderImpl) Close() {

}
