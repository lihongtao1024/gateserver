///>本代码由测试工具自动生成,请勿手动修改
package protocols

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

type ChannelMsgReq struct { //>属性通知
	Mid     uint16
	Pid     uint16
	Channel uint8   //>1：当前 2：队伍 3：帮派 4：地图 5：p2p 6：私聊 7：喇叭 8：招募 9：门派 10：世界
	Msg     []uint8 //>消息内容
}

type ChannelMsgNtf struct { //>属性通知
	Mid      uint16
	Pid      uint16
	Channel  uint8   //>1：当前 2：队伍 3：帮派 4：地图 5：p2p 6：私聊 7：喇叭 8：招募 9：门派 10：世界
	Chat_msg ChatMsg //>消息内容
}

type TipsMsgCSNtf struct { //>客户端左上角消息通知
	Mid uint16
	Pid uint16
	Msg []uint8 //>消息内容
}

type NoticeMsgNtf struct { //>Notice消息通知
	Mid    uint16
	Pid    uint16
	Type   uint8   //>类型由客户端自定义,对服务器无意义
	Msg    []uint8 //>消息内容
	Scroll uint8   //>消息滚动次数
}

type EmojiDataNtf struct { //>聊天表情通知
	Mid    uint16
	Pid    uint16
	Emojis EmojiData //>频道开关数组
}

func (proto *ChannelMsgReq) GetMid() uint16 {
	return 104
}

func (proto *ChannelMsgReq) GetPid() uint16 {
	return 1
}

func (proto *ChannelMsgReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(104)) {
		return false
	}

	if !writeProtoInteger(b, uint16(1)) {
		return false
	}

	if !writeProtoInteger(b, proto.Channel) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ChannelMsgReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Channel) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *ChannelMsgNtf) GetMid() uint16 {
	return 104
}

func (proto *ChannelMsgNtf) GetPid() uint16 {
	return 2
}

func (proto *ChannelMsgNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(104)) {
		return false
	}

	if !writeProtoInteger(b, uint16(2)) {
		return false
	}

	if !writeProtoInteger(b, proto.Channel) {
		return false
	}

	if !writeProtoCustom(b, &proto.Chat_msg) {
		return false
	}

	return true
}

func (proto *ChannelMsgNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Channel) {
		return false
	}

	if !readProtoCustom(b, &proto.Chat_msg) {
		return false
	}

	return true
}

func (proto *TipsMsgCSNtf) GetMid() uint16 {
	return 104
}

func (proto *TipsMsgCSNtf) GetPid() uint16 {
	return 3
}

func (proto *TipsMsgCSNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(104)) {
		return false
	}

	if !writeProtoInteger(b, uint16(3)) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *TipsMsgCSNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	return true
}

func (proto *NoticeMsgNtf) GetMid() uint16 {
	return 104
}

func (proto *NoticeMsgNtf) GetPid() uint16 {
	return 4
}

func (proto *NoticeMsgNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(104)) {
		return false
	}

	if !writeProtoInteger(b, uint16(4)) {
		return false
	}

	if !writeProtoInteger(b, proto.Type) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Msg, uint16(65535)) {
		return false
	}

	if !writeProtoInteger(b, proto.Scroll) {
		return false
	}

	return true
}

func (proto *NoticeMsgNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Type) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Msg, uint16(65535)) {
		return false
	}

	if !readProtoInteger(b, &proto.Scroll) {
		return false
	}

	return true
}

func (proto *EmojiDataNtf) GetMid() uint16 {
	return 104
}

func (proto *EmojiDataNtf) GetPid() uint16 {
	return 5
}

func (proto *EmojiDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(104)) {
		return false
	}

	if !writeProtoInteger(b, uint16(5)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Emojis) {
		return false
	}

	return true
}

func (proto *EmojiDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Emojis) {
		return false
	}

	return true
}

type IChannelMsgReq interface {
	OnChannelMsgReq(proto *ChannelMsgReq)
}

type IChannelMsgNtf interface {
	OnChannelMsgNtf(proto *ChannelMsgNtf)
}

type ITipsMsgCSNtf interface {
	OnTipsMsgCSNtf(proto *TipsMsgCSNtf)
}

type INoticeMsgNtf interface {
	OnNoticeMsgNtf(proto *NoticeMsgNtf)
}

type IEmojiDataNtf interface {
	OnEmojiDataNtf(proto *EmojiDataNtf)
}

type ClientCS struct {
	protoDispatch interface{}
}

func NewClientCS[T any](dispatch *T) *ClientCS {
	return &ClientCS{dispatch}
}

func (protos *ClientCS) GetMid() uint16 {
	return 104
}

func (protos *ClientCS) DispatchProto(data []byte) bool {
	b := bytes.NewBuffer(data)

	mid := binary.LittleEndian.Uint16(data)
	if mid != protos.GetMid() {
		return false
	}

	pid := binary.LittleEndian.Uint16(data[unsafe.Sizeof(uint16(0)):])
	switch pid {
	case 1:
		{
			t, ok := protos.protoDispatch.(IChannelMsgReq)
			if !ok {
				return false
			}

			proto := &ChannelMsgReq{}
			if !proto.Read(b) {
				fmt.Println("read ChannelMsgReq fail, system error.")
				return false
			}

			t.OnChannelMsgReq(proto)
		}
	case 2:
		{
			t, ok := protos.protoDispatch.(IChannelMsgNtf)
			if !ok {
				return false
			}

			proto := &ChannelMsgNtf{}
			if !proto.Read(b) {
				fmt.Println("read ChannelMsgNtf fail, system error.")
				return false
			}

			t.OnChannelMsgNtf(proto)
		}
	case 3:
		{
			t, ok := protos.protoDispatch.(ITipsMsgCSNtf)
			if !ok {
				return false
			}

			proto := &TipsMsgCSNtf{}
			if !proto.Read(b) {
				fmt.Println("read TipsMsgCSNtf fail, system error.")
				return false
			}

			t.OnTipsMsgCSNtf(proto)
		}
	case 4:
		{
			t, ok := protos.protoDispatch.(INoticeMsgNtf)
			if !ok {
				return false
			}

			proto := &NoticeMsgNtf{}
			if !proto.Read(b) {
				fmt.Println("read NoticeMsgNtf fail, system error.")
				return false
			}

			t.OnNoticeMsgNtf(proto)
		}
	case 5:
		{
			t, ok := protos.protoDispatch.(IEmojiDataNtf)
			if !ok {
				return false
			}

			proto := &EmojiDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read EmojiDataNtf fail, system error.")
				return false
			}

			t.OnEmojiDataNtf(proto)
		}
	default:
		{
			fmt.Println("illegal protocol, Mid =", mid, "Pid =", pid)
		}
	}

	return true
}
