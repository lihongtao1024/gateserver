package component

import (
	"gateserver/pkg"
)

type ClientType int
type CientState int

const (
	ClientTypeMT1 = ClientType(iota)
	ClientTypeMT2
)

const (
	ClientIdle = CientState(iota)
	ClientConnected
	ClientWorking
	ClientVerifying
	ClientRequesting
	ClientLoggedIn
	ClientPlaying
)

type Client interface {
	SetUName(name string)
	GetUName() string
	SetUid(uid uint32)
	GetUid() uint32
	SetHeartbeat(heartbeat bool)
	GetHeartbeat() bool
	SetAntiIndulge(antiindulge bool)
	GetAntiIndulge() bool
	SetThirdUser(thirduser bool)
	GetThirdUser() bool
	SetSid(sid uint8)
	GetSid() uint8
	SetGSid(gsid uint8)
	GetGSid() uint8
	SetZid(zid uint32)
	GetZid() uint32
	SetAid(aid uint32)
	GetAid() uint32
	SetIp(ip uint32)
	GetIp() uint32
	GetSuid() pkg.Guid
	SetGuid(guid pkg.Guid)
	GetGuid() pkg.Guid
	SetLongitude(longitude string)
	GetLongitude() string
	SetLatitude(latitude string)
	GetLatitude() string
	SetPasssword(pwd []byte)
	GetPassword() []byte
	SetHardware(hw string)
	GetHardware() string
	SetPlatform(platform string)
	GetPlatform() string
	GetLogicName() string
	GetRandKey() []byte
	VerifyClientHandShakeReq(data []byte) error
	SendClientHandShakeRsp() bool
	SendLoginReq() bool
	SendLoginAck(errcode pkg.ErrorCode, rid ...pkg.Guid) bool
	SendKickNtf(errcode pkg.ErrorCode) bool
}
