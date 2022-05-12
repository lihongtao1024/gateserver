package pkg

type Guid uint64
type GuidType int

const InvalidGuid = Guid(0)

const (
	GuidNull = GuidType(iota)
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
	GuidMax = GuidType(256)
)

type GuidBuilder interface {
	CreateGuid(i GuidType) Guid
	Close()
}
