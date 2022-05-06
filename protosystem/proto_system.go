package protosystem

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"gateserver/configsystem"
	"gateserver/internal/configs"
	"gateserver/internal/protocols"
	"gateserver/logsystem"
	"sync"
	"unsafe"
)

type ProtoDispatcher interface {
	GetMid() uint16
	Dispatch(data []byte) bool
}

type ProtoSystem struct {
	svrIndex        int
	protoDispatcher []ProtoDispatcher
	decodeSession   interface{}
	protoBuf        [4194304]byte
}

var Instance *ProtoSystem
var thisOnce sync.Once

func NewInstance(index int, dispatcher ...ProtoDispatcher) *ProtoSystem {
	thisOnce.Do(func() {
		Instance = &ProtoSystem{
			svrIndex:        index,
			protoDispatcher: append(make([]ProtoDispatcher, 0), dispatcher...),
		}
	})

	return Instance
}

func (protos *ProtoSystem) SetDecodeSession(o interface{}) {
	protos.decodeSession = o
}

func (protos *ProtoSystem) GetDecodeSession() interface{} {
	return protos.decodeSession
}

func (protos *ProtoSystem) ReadProto(data []byte) (result bool) {
	found := false
	mid := binary.LittleEndian.Uint16(data)

	for _, dispatcher := range protos.protoDispatcher {
		if dispatcher.GetMid() == mid {
			found = true
			result = dispatcher.Dispatch(data)
		}
	}

	if !found {
		result = false
		pid := binary.LittleEndian.Uint16(data[unsafe.Sizeof(uint16(0)):])
		logsystem.Instance.Dbg("handle illegal proto: [mid:%d, pid:%d].", mid, pid)
	}

	return
}

func (protos *ProtoSystem) BuildServerHandShakeReq() []byte {
	bs := protos.protoBuf[unsafe.Sizeof(uint32(0)):]
	hr := *(**protocols.ServerHandShakeReq)(unsafe.Pointer(&bs))
	hr.Type = configs.ServerIdGt
	hr.Index = uint16(protos.svrIndex)
	hr.Zone = uint32(configsystem.Instance.GetZoneId())
	hr.Version = protocols.NetVersion
	hr.Lenght = 0

	l := uint32(unsafe.Sizeof(*hr)) + uint32(unsafe.Sizeof(uint32(0)))
	binary.LittleEndian.PutUint32(protos.protoBuf[:], l)
	return protos.protoBuf[:l]
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

	zid := uint32(configsystem.Instance.GetZoneId())
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

func (protos *ProtoSystem) BuildClientHandShakeRsp() []byte {
	bs := protos.protoBuf[unsafe.Sizeof(uint32(0)):]
	hr := *(**protocols.ClientHandShakeRsp)(unsafe.Pointer(&bs))
	hr.Return = 0
	hr.Load = 0
	hr.Delay = 10

	l := uint32(unsafe.Sizeof(*hr)) + uint32(unsafe.Sizeof(uint32(0)))
	binary.LittleEndian.PutUint32(protos.protoBuf[:], l)
	return protos.protoBuf[:l]
}

func (protos *ProtoSystem) VerifyClientHandShakeReq(data []byte) error {
	hr := *(**protocols.ClientHandShakeReq)(unsafe.Pointer(&data))
	if len(data) != int(unsafe.Sizeof(*hr)) {
		return errors.New("verify ClientHandShakeReq fail, illegal packet length")
	}

	zid := uint32(configsystem.Instance.GetZoneId())
	if hr.Zone != zid {
		return fmt.Errorf("verify ClientHandShakeReq fail, illegal zone[%d - %d]",
			hr.Zone, zid)
	}

	if hr.Version == 0 {
		return fmt.Errorf("verify ClientHandShakeReq fail, illegal version[%d - %d]",
			hr.Version, protocols.NetVersion)
	}

	return nil
}

func (protos *ProtoSystem) BuildClientProto(proto protocols.Writer) (result bool, data []byte) {
	bs := protos.protoBuf[unsafe.Sizeof(uint32(0)):]
	br := bytes.NewBuffer(bs)
	br.Reset()

	if !proto.Write(br) {
		result = false
		data = nil
		return
	}

	l := uint32(br.Len()) + uint32(unsafe.Sizeof(uint32(0)))
	binary.LittleEndian.PutUint32(protos.protoBuf[:], l)
	return true, protos.protoBuf[:l]
}

func (protos *ProtoSystem) Close() {

}
