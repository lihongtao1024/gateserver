package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/pkg/protocols"
	"gateserver/singleton"
	"unsafe"
)

type protoDispatcher interface {
	GetMid() uint16
	DispatchProto(data []byte) bool
}

type protocolImpl struct {
	protoDispatcher []protoDispatcher
	decodeSession   interface{}
	protoBuf        [4194304]byte
}

func NewProtocol() component.Protocol {
	protos := &protocolImpl{protoDispatcher: make([]protoDispatcher, 0)}
	protos.protoDispatcher = append(
		protos.protoDispatcher,
		newGT2WSProto(),
	)
	return protos
}

func (protos *protocolImpl) SetDecodeSession(o interface{}) {
	protos.decodeSession = o
}

func (protos *protocolImpl) GetDecodeSession() interface{} {
	return protos.decodeSession
}

func (protos *protocolImpl) DispatchProto(data []byte) (result bool) {
	found := false
	mid := binary.LittleEndian.Uint16(data)

	for _, dispatcher := range protos.protoDispatcher {
		if dispatcher.GetMid() == mid {
			found = true
			result = dispatcher.DispatchProto(data)
		}
	}

	if !found {
		result = false
		pid := binary.LittleEndian.Uint16(data[unsafe.Sizeof(uint16(0)):])
		singleton.LogInstance.Dbg(
			"handle illegal proto: [mid:%d, pid:%d].",
			mid,
			pid,
		)
	}

	return
}

func (protos *protocolImpl) BuildServerHandShakeReq() []byte {
	bs := protos.protoBuf[unsafe.Sizeof(uint32(0)):]
	hr := *(**protocols.ServerHandShakeReq)(unsafe.Pointer(&bs))
	hr.Type = uint16(singleton.AppInstance.GetType())
	hr.Index = uint16(singleton.AppInstance.GetIndex())
	hr.Zone = uint32(singleton.CfgInstance.GetZoneId())
	hr.Version = protocols.NetVersion
	hr.Lenght = 0

	l := uint32(unsafe.Sizeof(*hr)) + uint32(unsafe.Sizeof(uint32(0)))
	binary.LittleEndian.PutUint32(protos.protoBuf[:], l)
	return protos.protoBuf[:l]
}

func (proto *protocolImpl) VerifyServerHandShakeRsp(index uint16,
	data []byte) error {
	hr := *(**protocols.ServerHandShakeRsp)(unsafe.Pointer(&data))
	if len(data) != int(unsafe.Sizeof(*hr)) {
		return errors.New(
			"verify ServerHandShakeRsp fail, illegal packet length",
		)
	}

	if hr.Index != uint32(index) {
		return fmt.Errorf(
			"verify ServerHandShakeRsp fail, illegal index[%d - %d]",
			hr.Index,
			index,
		)
	}

	zid := uint32(singleton.CfgInstance.GetZoneId())
	if hr.Zone != zid {
		return fmt.Errorf(
			"verify ServerHandShakeRsp fail, illegal zone[%d - %d]",
			hr.Zone,
			zid,
		)
	}

	if hr.Version != protocols.NetVersion {
		return fmt.Errorf(
			"verify ServerHandShakeRsp fail, illegal version[%d - %d]",
			hr.Version,
			protocols.NetVersion,
		)
	}
	return nil
}

func (protos *protocolImpl) BuildClientHandShakeRsp() []byte {
	bs := protos.protoBuf[unsafe.Sizeof(uint32(0)):]
	hr := *(**protocols.ClientHandShakeRsp)(unsafe.Pointer(&bs))
	hr.Return = 0
	hr.Load = 0
	hr.Delay = 10

	l := uint32(unsafe.Sizeof(*hr)) + uint32(unsafe.Sizeof(uint32(0)))
	binary.LittleEndian.PutUint32(protos.protoBuf[:], l)
	return protos.protoBuf[:l]
}

func (protos *protocolImpl) VerifyClientHandShakeReq(data []byte) error {
	hr := *(**protocols.ClientHandShakeReq)(unsafe.Pointer(&data))
	if len(data) != int(unsafe.Sizeof(*hr)) {
		return errors.New(
			"verify ClientHandShakeReq fail, illegal packet length",
		)
	}

	zid := uint32(singleton.CfgInstance.GetZoneId())
	if hr.Zone != zid {
		return fmt.Errorf(
			"verify ClientHandShakeReq fail, illegal zone[%d - %d]",
			hr.Zone,
			zid,
		)
	}

	if hr.Version == 0 {
		return fmt.Errorf(
			"verify ClientHandShakeReq fail, illegal version[%d - %d]",
			hr.Version,
			protocols.NetVersion,
		)
	}

	return nil
}

func (protos *protocolImpl) BuildClientProto(proto pkg.WriterProto) (
	result bool, data []byte) {
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

func (protos *protocolImpl) BuildServerProto(client component.Client,
	proto pkg.WriterProto) (result bool, data []byte) {
	bl := unsafe.Sizeof(uint32(0))
	bs := protos.protoBuf[bl:]
	hr := *(**protocols.C2SProtoHeader)(unsafe.Pointer(&bs))
	hl := unsafe.Sizeof(*hr)
	hr.Sid = uint32(client.GetSid())
	hr.Uid = client.GetUid()

	bs = protos.protoBuf[bl+hl:]
	br := bytes.NewBuffer(bs)
	br.Reset()

	if !proto.Write(br) {
		result = false
		data = nil
		return
	}

	l := uint32(br.Len()) + uint32(bl) + uint32(hl)
	binary.LittleEndian.PutUint32(protos.protoBuf[:], l)
	return true, protos.protoBuf[:l]
}

func (protos *protocolImpl) Close() {

}
