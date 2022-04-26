package protosystem

import (
	"encoding/binary"
	"errors"
	"fmt"
	"gateserver/configsystem"
	"gateserver/internal/configs"
	"gateserver/internal/protocols"
	"sync"
	"unsafe"
)

type ProtoSystem struct {
	svrIndex   int
	gt2wsProto *GT2WSProto
	sendBuffer [65536]byte
	recvBuffer [65536]byte
}

var TheProto *ProtoSystem
var thisOnce sync.Once

func NewProtoSystemInstance(index int) *ProtoSystem {
	thisOnce.Do(func() {
		TheProto = &ProtoSystem{
			svrIndex:   index,
			gt2wsProto: NewGT2WSProto(),
		}
	})

	return TheProto
}

func (proto *ProtoSystem) BuildServerHandShakeReq() []byte {
	bs := proto.sendBuffer[unsafe.Sizeof(uint32(0)):]
	hr := *(**protocols.ServerHandShakeReq)(unsafe.Pointer(&bs))
	hr.Type = configs.ServerIdGt
	hr.Index = uint16(proto.svrIndex)
	hr.Zone = uint32(configsystem.TheConfig.GetZoneId())
	hr.Version = protocols.NetVersion
	hr.Lenght = 0

	l := uint32(unsafe.Sizeof(*hr)) + uint32(unsafe.Sizeof(uint32(0)))
	binary.LittleEndian.PutUint32(proto.sendBuffer[:], l)
	return proto.sendBuffer[:l]
}

func (proto *ProtoSystem) VerifyServerHandShakeRsp(index uint16, data []byte) error {
	hr := *(**protocols.ServerHandShakeRsp)(unsafe.Pointer(&data))
	if len(data) != int(unsafe.Sizeof(*hr)) {
		return errors.New("verify ServerHandShakeRsp fail, illegal packet length")
	}

	if hr.Index != uint32(index) {
		return fmt.Errorf("verify ServerHandShakeRsp fail, illegal index[%d - %d]",
			hr.Index, index)
	}

	zid := uint32(configsystem.TheConfig.GetZoneId())
	if hr.Zone != zid {
		return fmt.Errorf("verify ServerHandShakeRsp fail, illegal zone[%d - %d]",
			hr.Zone, zid)
	}

	if hr.Version != protocols.NetVersion {
		return fmt.Errorf("verify ServerHandShakeRsp fail, illegal version[%d - %d]",
			hr.Version, protocols.NetVersion)
	}
	return nil
}

func (proto *ProtoSystem) Close() {

}
