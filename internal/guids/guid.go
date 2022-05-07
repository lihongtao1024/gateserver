package guids

import "time"

/*
 64 bit Guid:
     type         sid         time         inc
 |63......56|55........48|47........16|15........0
     8bit        8 bit        32 bit      16 bit
*/
type Guid = uint64

const InvalidGuid = Guid(0)

const (
	GuidNull = iota
	GuidGlobal
	GuidMap
	GuidPlayer
	GuidNpc
	GuidItem
	GuidPet
	GuidGuard
	GuidMonster
	GuidTeam
	GuidTimer
	GuidTrigger
	GuidShop
	GuidAuction
	GuidGuild
	GuidMessage
	GuidMail
	GuidRobot
	GuidPlayerFighter
	GuidMonsterFighter
	GuidPetFighter
	GuidGuardFighter
	GuidRobotFighter
	GuidBeast
	GuidBeastFighter
	GuidLadder
	GuidLadderFighter
	GuidDummyFighter
	GuidMax = 256
)

type Component interface {
	CreateGuid(i int) Guid
	Close()
}

type guidComponent struct {
	guidArray [GuidMax]Guid
}

func NewComponent(index uint8) Component {
	comp := &guidComponent{}
	times := uint32(time.Now().Local().Unix())

	for i := 0; i < len(comp.guidArray); i++ {
		guid := &comp.guidArray[i]
		*guid = Guid(i)<<56 | Guid(index)<<48 | Guid(times)<<16
	}

	return comp
}

func (comp *guidComponent) CreateGuid(i int) Guid {
	if i <= GuidNull || i >= GuidMax {
		panic("unexpected Guid type")
	}

	comp.guidArray[i]++
	return comp.guidArray[i]
}

func (comp *guidComponent) Close() {

}
