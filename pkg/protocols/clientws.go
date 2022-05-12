///>本代码由测试工具自动生成,请勿手动修改
package protocols

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

type RandKeyNtf struct { //>随机码
	Mid          uint16
	Pid          uint16
	Code_content []uint8 //>随机码
}

type LoginReq struct { //>客户端登陆请求
	Mid                 uint16
	Pid                 uint16
	Sid                 uint8   //>子区ID
	Username            string  //>账号名
	Ip                  uint32  //>登陆ip地址
	Pwd_content         []uint8 //>密码密文
	Hwid                string  //>硬件码
	Client_type         uint8   //>客户端类型
	Client_version      string  //>客户端版本号
	Client_protocol_md5 string  //>客户端协议MD5
	Longitude           string  //>经度
	Latitude            string  //>纬度
	Session             uint64  //>会话ID(由网关填写)
	Suid                uint64  //>会话ID(由网关填写)
}

type LoginAck struct { //>客户端登陆应答
	Mid     uint16
	Pid     uint16
	Uid     uint32 //>账号ID
	Guid    uint64 //>正在GS中(战斗)的玩家GUID
	Suid    uint64 //>会话ID
	Rid     uint64 //>运行时ID(服务器重启后变化)
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type ReloginReq struct { //>客户端重连请求
	Mid                 uint16
	Pid                 uint16
	Sid                 uint8  //>子区ID
	Uid                 uint32 //>账号ID
	Ip                  uint32 //>登陆ip地址
	Hwid                string //>硬件码
	Client_type         uint8  //>客户端类型
	Client_version      string //>客户端版本号
	Client_protocol_md5 string //>客户端协议MD5
	Longitude           string //>经度
	Latitude            string //>纬度
	Session             uint64 //>会话ID(由网关填写)
	Suid                uint64 //>会话ID
}

type ReloginAck struct { //>客户端重连应答
	Mid     uint16
	Pid     uint16
	Guid    uint64 //>在线玩家GUID
	Rid     uint64 //>运行时ID(服务器重启后变化)
	Gsindex uint32 //>GS索引
	Arrayid uint32 //>GS上玩家对象数组下标
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type LogoutReq struct { //>客户端退出请求
	Mid  uint16
	Pid  uint16
	Type uint8 //>1:回到选角界面 2:退出游戏
}

type LogoutAck struct { //>客户端退出应答
	Mid     uint16
	Pid     uint16
	Type    uint8  //>1:回到选角界面 2:退出游戏
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type KickNtf struct { //>通知客户端踢人
	Mid     uint16
	Pid     uint16
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type PlayerListReq struct { //>玩家简略列表请求
	Mid uint16
	Pid uint16
}

type PlayerListAck struct { //>玩家简略列表回应
	Mid        uint16
	Pid        uint16
	Lastplayer uint64        //>最后一次登陆玩家
	Briefs     []PlayerBrief //>玩家简略列表
	Errcode    int32         //>错误码
	Errmsg     string        //>错误描述
}

type CreatePlayerReq struct { //>创建玩家请求
	Mid         uint16
	Pid         uint16
	Playerbrief PlayerBrief //>玩家信息
}

type CreatePlayerAck struct { //>创建玩家回应
	Mid         uint16
	Pid         uint16
	Playerbrief PlayerBrief //>玩家信息
	Errcode     int32       //>错误码
	Errmsg      string      //>错误描述
}

type DestroyPlayerReq struct { //>销毁玩家请求
	Mid  uint16
	Pid  uint16
	Guid uint64 //>玩家GUID
}

type DestroyPlayerAck struct { //>销毁玩家回应
	Mid     uint16
	Pid     uint16
	Guid    uint64 //>玩家GUID
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type RestorePlayerReq struct { //>销毁玩家请求
	Mid  uint16
	Pid  uint16
	Guid uint64 //>玩家GUID
}

type RestorePlayerAck struct { //>销毁玩家回应
	Mid     uint16
	Pid     uint16
	Guid    uint64 //>玩家GUID
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type EnterGSReq struct { //>玩家进入GS请求
	Mid  uint16
	Pid  uint16
	Guid uint64 //>玩家GUID
}

type EnterGSAck struct { //>玩家进入GS回应
	Mid     uint16
	Pid     uint16
	Guid    uint64 //>玩家GUID
	Gsindex uint32 //>GS索引
	Arrayid uint32 //>GS上玩家对象数组下标
	Errcode int32  //>错误码
	Errmsg  string //>错误描述
}

type TrdLoginReq struct { //>玩家密码验证请求
	Mid                 uint16
	Pid                 uint16
	Sid                 uint8   //>子区ID
	Username            string  //>账号名
	Platform            string  //>平台名称
	Token               []uint8 //>token
	Hwid                string  //>硬件码
	Ip                  uint32  //>登陆ip地址
	Client_type         uint8   //>客户端类型
	Client_version      string  //>客户端版本号
	Client_protocol_md5 string  //>客户端协议MD5
	Longitude           string  //>经度
	Latitude            string  //>纬度
	Session             uint64  //>会话ID(由网关填写)
	Suid                uint64  //>会话ID(由网关填写)
}

type GetWSTimestampReq struct { //>服务器时间戳请求
	Mid uint16
	Pid uint16
}

type GetWSTimestampAck struct { //>服务器时间戳应答
	Mid  uint16
	Pid  uint16
	Now  uint32 //>当前服务器本地时间
	Zone int32  //>当前服务器本地时间减去标准UTC时间的差值
}

type RealnameInfoNtf struct { //>实名认证信息通知
	Mid             uint16
	Pid             uint16
	Realname_status uint32 //>实名认证状态
	Auth            string //>授权信息
	Realname_token  string //>token
}

func (proto *RandKeyNtf) GetMid() uint16 {
	return 101
}

func (proto *RandKeyNtf) GetPid() uint16 {
	return 1
}

func (proto *RandKeyNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(1)) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Code_content, uint8(255)) {
		return false
	}

	return true
}

func (proto *RandKeyNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Code_content, uint8(255)) {
		return false
	}

	return true
}

func (proto *LoginReq) GetMid() uint16 {
	return 101
}

func (proto *LoginReq) GetPid() uint16 {
	return 2
}

func (proto *LoginReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(2)) {
		return false
	}

	if !writeProtoInteger(b, proto.Sid) {
		return false
	}

	if !writeProtoString(b, proto.Username, 255) {
		return false
	}

	if !writeProtoInteger(b, proto.Ip) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Pwd_content, uint8(255)) {
		return false
	}

	if !writeProtoString(b, proto.Hwid, 256) {
		return false
	}

	if !writeProtoInteger(b, proto.Client_type) {
		return false
	}

	if !writeProtoString(b, proto.Client_version, 256) {
		return false
	}

	if !writeProtoString(b, proto.Client_protocol_md5, 256) {
		return false
	}

	if !writeProtoString(b, proto.Longitude, 64) {
		return false
	}

	if !writeProtoString(b, proto.Latitude, 64) {
		return false
	}

	if !writeProtoInteger(b, proto.Session) {
		return false
	}

	if !writeProtoInteger(b, proto.Suid) {
		return false
	}

	return true
}

func (proto *LoginReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Sid) {
		return false
	}

	if !readProtoString(b, &proto.Username, 255) {
		return false
	}

	if !readProtoInteger(b, &proto.Ip) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Pwd_content, uint8(255)) {
		return false
	}

	if !readProtoString(b, &proto.Hwid, 256) {
		return false
	}

	if !readProtoInteger(b, &proto.Client_type) {
		return false
	}

	if !readProtoString(b, &proto.Client_version, 256) {
		return false
	}

	if !readProtoString(b, &proto.Client_protocol_md5, 256) {
		return false
	}

	if !readProtoString(b, &proto.Longitude, 64) {
		return false
	}

	if !readProtoString(b, &proto.Latitude, 64) {
		return false
	}

	if !readProtoInteger(b, &proto.Session) {
		return false
	}

	if !readProtoInteger(b, &proto.Suid) {
		return false
	}

	return true
}

func (proto *LoginAck) GetMid() uint16 {
	return 101
}

func (proto *LoginAck) GetPid() uint16 {
	return 3
}

func (proto *LoginAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(3)) {
		return false
	}

	if !writeProtoInteger(b, proto.Uid) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Suid) {
		return false
	}

	if !writeProtoInteger(b, proto.Rid) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *LoginAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Uid) {
		return false
	}

	if !readProtoInteger(b, &proto.Guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Suid) {
		return false
	}

	if !readProtoInteger(b, &proto.Rid) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *ReloginReq) GetMid() uint16 {
	return 101
}

func (proto *ReloginReq) GetPid() uint16 {
	return 4
}

func (proto *ReloginReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(4)) {
		return false
	}

	if !writeProtoInteger(b, proto.Sid) {
		return false
	}

	if !writeProtoInteger(b, proto.Uid) {
		return false
	}

	if !writeProtoInteger(b, proto.Ip) {
		return false
	}

	if !writeProtoString(b, proto.Hwid, 128) {
		return false
	}

	if !writeProtoInteger(b, proto.Client_type) {
		return false
	}

	if !writeProtoString(b, proto.Client_version, 256) {
		return false
	}

	if !writeProtoString(b, proto.Client_protocol_md5, 256) {
		return false
	}

	if !writeProtoString(b, proto.Longitude, 64) {
		return false
	}

	if !writeProtoString(b, proto.Latitude, 64) {
		return false
	}

	if !writeProtoInteger(b, proto.Session) {
		return false
	}

	if !writeProtoInteger(b, proto.Suid) {
		return false
	}

	return true
}

func (proto *ReloginReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Sid) {
		return false
	}

	if !readProtoInteger(b, &proto.Uid) {
		return false
	}

	if !readProtoInteger(b, &proto.Ip) {
		return false
	}

	if !readProtoString(b, &proto.Hwid, 128) {
		return false
	}

	if !readProtoInteger(b, &proto.Client_type) {
		return false
	}

	if !readProtoString(b, &proto.Client_version, 256) {
		return false
	}

	if !readProtoString(b, &proto.Client_protocol_md5, 256) {
		return false
	}

	if !readProtoString(b, &proto.Longitude, 64) {
		return false
	}

	if !readProtoString(b, &proto.Latitude, 64) {
		return false
	}

	if !readProtoInteger(b, &proto.Session) {
		return false
	}

	if !readProtoInteger(b, &proto.Suid) {
		return false
	}

	return true
}

func (proto *ReloginAck) GetMid() uint16 {
	return 101
}

func (proto *ReloginAck) GetPid() uint16 {
	return 5
}

func (proto *ReloginAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(5)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Rid) {
		return false
	}

	if !writeProtoInteger(b, proto.Gsindex) {
		return false
	}

	if !writeProtoInteger(b, proto.Arrayid) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *ReloginAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Rid) {
		return false
	}

	if !readProtoInteger(b, &proto.Gsindex) {
		return false
	}

	if !readProtoInteger(b, &proto.Arrayid) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *LogoutReq) GetMid() uint16 {
	return 101
}

func (proto *LogoutReq) GetPid() uint16 {
	return 6
}

func (proto *LogoutReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(6)) {
		return false
	}

	if !writeProtoInteger(b, proto.Type) {
		return false
	}

	return true
}

func (proto *LogoutReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Type) {
		return false
	}

	return true
}

func (proto *LogoutAck) GetMid() uint16 {
	return 101
}

func (proto *LogoutAck) GetPid() uint16 {
	return 7
}

func (proto *LogoutAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(7)) {
		return false
	}

	if !writeProtoInteger(b, proto.Type) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *LogoutAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Type) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *KickNtf) GetMid() uint16 {
	return 101
}

func (proto *KickNtf) GetPid() uint16 {
	return 8
}

func (proto *KickNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(8)) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *KickNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *PlayerListReq) GetMid() uint16 {
	return 101
}

func (proto *PlayerListReq) GetPid() uint16 {
	return 9
}

func (proto *PlayerListReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(9)) {
		return false
	}

	return true
}

func (proto *PlayerListReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	return true
}

func (proto *PlayerListAck) GetMid() uint16 {
	return 101
}

func (proto *PlayerListAck) GetPid() uint16 {
	return 10
}

func (proto *PlayerListAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(10)) {
		return false
	}

	if !writeProtoInteger(b, proto.Lastplayer) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Briefs, uint8(255)) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *PlayerListAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Lastplayer) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Briefs, uint8(255)) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *CreatePlayerReq) GetMid() uint16 {
	return 101
}

func (proto *CreatePlayerReq) GetPid() uint16 {
	return 11
}

func (proto *CreatePlayerReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(11)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Playerbrief) {
		return false
	}

	return true
}

func (proto *CreatePlayerReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Playerbrief) {
		return false
	}

	return true
}

func (proto *CreatePlayerAck) GetMid() uint16 {
	return 101
}

func (proto *CreatePlayerAck) GetPid() uint16 {
	return 12
}

func (proto *CreatePlayerAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(12)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Playerbrief) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *CreatePlayerAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Playerbrief) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *DestroyPlayerReq) GetMid() uint16 {
	return 101
}

func (proto *DestroyPlayerReq) GetPid() uint16 {
	return 13
}

func (proto *DestroyPlayerReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(13)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *DestroyPlayerReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Guid) {
		return false
	}

	return true
}

func (proto *DestroyPlayerAck) GetMid() uint16 {
	return 101
}

func (proto *DestroyPlayerAck) GetPid() uint16 {
	return 14
}

func (proto *DestroyPlayerAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(14)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *DestroyPlayerAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *RestorePlayerReq) GetMid() uint16 {
	return 101
}

func (proto *RestorePlayerReq) GetPid() uint16 {
	return 15
}

func (proto *RestorePlayerReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(15)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *RestorePlayerReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Guid) {
		return false
	}

	return true
}

func (proto *RestorePlayerAck) GetMid() uint16 {
	return 101
}

func (proto *RestorePlayerAck) GetPid() uint16 {
	return 16
}

func (proto *RestorePlayerAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(16)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *RestorePlayerAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *EnterGSReq) GetMid() uint16 {
	return 101
}

func (proto *EnterGSReq) GetPid() uint16 {
	return 17
}

func (proto *EnterGSReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(17)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *EnterGSReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Guid) {
		return false
	}

	return true
}

func (proto *EnterGSAck) GetMid() uint16 {
	return 101
}

func (proto *EnterGSAck) GetPid() uint16 {
	return 18
}

func (proto *EnterGSAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(18)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Gsindex) {
		return false
	}

	if !writeProtoInteger(b, proto.Arrayid) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *EnterGSAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Gsindex) {
		return false
	}

	if !readProtoInteger(b, &proto.Arrayid) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	return true
}

func (proto *TrdLoginReq) GetMid() uint16 {
	return 101
}

func (proto *TrdLoginReq) GetPid() uint16 {
	return 19
}

func (proto *TrdLoginReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(19)) {
		return false
	}

	if !writeProtoInteger(b, proto.Sid) {
		return false
	}

	if !writeProtoString(b, proto.Username, 255) {
		return false
	}

	if !writeProtoString(b, proto.Platform, 255) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Token, uint16(65535)) {
		return false
	}

	if !writeProtoString(b, proto.Hwid, 256) {
		return false
	}

	if !writeProtoInteger(b, proto.Ip) {
		return false
	}

	if !writeProtoInteger(b, proto.Client_type) {
		return false
	}

	if !writeProtoString(b, proto.Client_version, 256) {
		return false
	}

	if !writeProtoString(b, proto.Client_protocol_md5, 256) {
		return false
	}

	if !writeProtoString(b, proto.Longitude, 64) {
		return false
	}

	if !writeProtoString(b, proto.Latitude, 64) {
		return false
	}

	if !writeProtoInteger(b, proto.Session) {
		return false
	}

	if !writeProtoInteger(b, proto.Suid) {
		return false
	}

	return true
}

func (proto *TrdLoginReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Sid) {
		return false
	}

	if !readProtoString(b, &proto.Username, 255) {
		return false
	}

	if !readProtoString(b, &proto.Platform, 255) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Token, uint16(65535)) {
		return false
	}

	if !readProtoString(b, &proto.Hwid, 256) {
		return false
	}

	if !readProtoInteger(b, &proto.Ip) {
		return false
	}

	if !readProtoInteger(b, &proto.Client_type) {
		return false
	}

	if !readProtoString(b, &proto.Client_version, 256) {
		return false
	}

	if !readProtoString(b, &proto.Client_protocol_md5, 256) {
		return false
	}

	if !readProtoString(b, &proto.Longitude, 64) {
		return false
	}

	if !readProtoString(b, &proto.Latitude, 64) {
		return false
	}

	if !readProtoInteger(b, &proto.Session) {
		return false
	}

	if !readProtoInteger(b, &proto.Suid) {
		return false
	}

	return true
}

func (proto *GetWSTimestampReq) GetMid() uint16 {
	return 101
}

func (proto *GetWSTimestampReq) GetPid() uint16 {
	return 20
}

func (proto *GetWSTimestampReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(20)) {
		return false
	}

	return true
}

func (proto *GetWSTimestampReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	return true
}

func (proto *GetWSTimestampAck) GetMid() uint16 {
	return 101
}

func (proto *GetWSTimestampAck) GetPid() uint16 {
	return 21
}

func (proto *GetWSTimestampAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(21)) {
		return false
	}

	if !writeProtoInteger(b, proto.Now) {
		return false
	}

	if !writeProtoInteger(b, proto.Zone) {
		return false
	}

	return true
}

func (proto *GetWSTimestampAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Now) {
		return false
	}

	if !readProtoInteger(b, &proto.Zone) {
		return false
	}

	return true
}

func (proto *RealnameInfoNtf) GetMid() uint16 {
	return 101
}

func (proto *RealnameInfoNtf) GetPid() uint16 {
	return 22
}

func (proto *RealnameInfoNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(101)) {
		return false
	}

	if !writeProtoInteger(b, uint16(22)) {
		return false
	}

	if !writeProtoInteger(b, proto.Realname_status) {
		return false
	}

	if !writeProtoString(b, proto.Auth, 64) {
		return false
	}

	if !writeProtoString(b, proto.Realname_token, 64) {
		return false
	}

	return true
}

func (proto *RealnameInfoNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Realname_status) {
		return false
	}

	if !readProtoString(b, &proto.Auth, 64) {
		return false
	}

	if !readProtoString(b, &proto.Realname_token, 64) {
		return false
	}

	return true
}

type IRandKeyNtf interface {
	OnRandKeyNtf(proto *RandKeyNtf)
}

type ILoginReq interface {
	OnLoginReq(proto *LoginReq)
}

type ILoginAck interface {
	OnLoginAck(proto *LoginAck)
}

type IReloginReq interface {
	OnReloginReq(proto *ReloginReq)
}

type IReloginAck interface {
	OnReloginAck(proto *ReloginAck)
}

type ILogoutReq interface {
	OnLogoutReq(proto *LogoutReq)
}

type ILogoutAck interface {
	OnLogoutAck(proto *LogoutAck)
}

type IKickNtf interface {
	OnKickNtf(proto *KickNtf)
}

type IPlayerListReq interface {
	OnPlayerListReq(proto *PlayerListReq)
}

type IPlayerListAck interface {
	OnPlayerListAck(proto *PlayerListAck)
}

type ICreatePlayerReq interface {
	OnCreatePlayerReq(proto *CreatePlayerReq)
}

type ICreatePlayerAck interface {
	OnCreatePlayerAck(proto *CreatePlayerAck)
}

type IDestroyPlayerReq interface {
	OnDestroyPlayerReq(proto *DestroyPlayerReq)
}

type IDestroyPlayerAck interface {
	OnDestroyPlayerAck(proto *DestroyPlayerAck)
}

type IRestorePlayerReq interface {
	OnRestorePlayerReq(proto *RestorePlayerReq)
}

type IRestorePlayerAck interface {
	OnRestorePlayerAck(proto *RestorePlayerAck)
}

type IEnterGSReq interface {
	OnEnterGSReq(proto *EnterGSReq)
}

type IEnterGSAck interface {
	OnEnterGSAck(proto *EnterGSAck)
}

type ITrdLoginReq interface {
	OnTrdLoginReq(proto *TrdLoginReq)
}

type IGetWSTimestampReq interface {
	OnGetWSTimestampReq(proto *GetWSTimestampReq)
}

type IGetWSTimestampAck interface {
	OnGetWSTimestampAck(proto *GetWSTimestampAck)
}

type IRealnameInfoNtf interface {
	OnRealnameInfoNtf(proto *RealnameInfoNtf)
}

type ClientWS struct {
	protoDispatch interface{}
}

func NewClientWS[T any](dispatch *T) *ClientWS {
	return &ClientWS{dispatch}
}

func (protos *ClientWS) GetMid() uint16 {
	return 101
}

func (protos *ClientWS) DispatchProto(data []byte) bool {
	b := bytes.NewBuffer(data)

	mid := binary.LittleEndian.Uint16(data)
	if mid != protos.GetMid() {
		return false
	}

	pid := binary.LittleEndian.Uint16(data[unsafe.Sizeof(uint16(0)):])
	switch pid {
	case 1:
		{
			t, ok := protos.protoDispatch.(IRandKeyNtf)
			if !ok {
				return false
			}

			proto := &RandKeyNtf{}
			if !proto.Read(b) {
				fmt.Println("read RandKeyNtf fail, system error.")
				return false
			}

			t.OnRandKeyNtf(proto)
		}
	case 2:
		{
			t, ok := protos.protoDispatch.(ILoginReq)
			if !ok {
				return false
			}

			proto := &LoginReq{}
			if !proto.Read(b) {
				fmt.Println("read LoginReq fail, system error.")
				return false
			}

			t.OnLoginReq(proto)
		}
	case 3:
		{
			t, ok := protos.protoDispatch.(ILoginAck)
			if !ok {
				return false
			}

			proto := &LoginAck{}
			if !proto.Read(b) {
				fmt.Println("read LoginAck fail, system error.")
				return false
			}

			t.OnLoginAck(proto)
		}
	case 4:
		{
			t, ok := protos.protoDispatch.(IReloginReq)
			if !ok {
				return false
			}

			proto := &ReloginReq{}
			if !proto.Read(b) {
				fmt.Println("read ReloginReq fail, system error.")
				return false
			}

			t.OnReloginReq(proto)
		}
	case 5:
		{
			t, ok := protos.protoDispatch.(IReloginAck)
			if !ok {
				return false
			}

			proto := &ReloginAck{}
			if !proto.Read(b) {
				fmt.Println("read ReloginAck fail, system error.")
				return false
			}

			t.OnReloginAck(proto)
		}
	case 6:
		{
			t, ok := protos.protoDispatch.(ILogoutReq)
			if !ok {
				return false
			}

			proto := &LogoutReq{}
			if !proto.Read(b) {
				fmt.Println("read LogoutReq fail, system error.")
				return false
			}

			t.OnLogoutReq(proto)
		}
	case 7:
		{
			t, ok := protos.protoDispatch.(ILogoutAck)
			if !ok {
				return false
			}

			proto := &LogoutAck{}
			if !proto.Read(b) {
				fmt.Println("read LogoutAck fail, system error.")
				return false
			}

			t.OnLogoutAck(proto)
		}
	case 8:
		{
			t, ok := protos.protoDispatch.(IKickNtf)
			if !ok {
				return false
			}

			proto := &KickNtf{}
			if !proto.Read(b) {
				fmt.Println("read KickNtf fail, system error.")
				return false
			}

			t.OnKickNtf(proto)
		}
	case 9:
		{
			t, ok := protos.protoDispatch.(IPlayerListReq)
			if !ok {
				return false
			}

			proto := &PlayerListReq{}
			if !proto.Read(b) {
				fmt.Println("read PlayerListReq fail, system error.")
				return false
			}

			t.OnPlayerListReq(proto)
		}
	case 10:
		{
			t, ok := protos.protoDispatch.(IPlayerListAck)
			if !ok {
				return false
			}

			proto := &PlayerListAck{}
			if !proto.Read(b) {
				fmt.Println("read PlayerListAck fail, system error.")
				return false
			}

			t.OnPlayerListAck(proto)
		}
	case 11:
		{
			t, ok := protos.protoDispatch.(ICreatePlayerReq)
			if !ok {
				return false
			}

			proto := &CreatePlayerReq{}
			if !proto.Read(b) {
				fmt.Println("read CreatePlayerReq fail, system error.")
				return false
			}

			t.OnCreatePlayerReq(proto)
		}
	case 12:
		{
			t, ok := protos.protoDispatch.(ICreatePlayerAck)
			if !ok {
				return false
			}

			proto := &CreatePlayerAck{}
			if !proto.Read(b) {
				fmt.Println("read CreatePlayerAck fail, system error.")
				return false
			}

			t.OnCreatePlayerAck(proto)
		}
	case 13:
		{
			t, ok := protos.protoDispatch.(IDestroyPlayerReq)
			if !ok {
				return false
			}

			proto := &DestroyPlayerReq{}
			if !proto.Read(b) {
				fmt.Println("read DestroyPlayerReq fail, system error.")
				return false
			}

			t.OnDestroyPlayerReq(proto)
		}
	case 14:
		{
			t, ok := protos.protoDispatch.(IDestroyPlayerAck)
			if !ok {
				return false
			}

			proto := &DestroyPlayerAck{}
			if !proto.Read(b) {
				fmt.Println("read DestroyPlayerAck fail, system error.")
				return false
			}

			t.OnDestroyPlayerAck(proto)
		}
	case 15:
		{
			t, ok := protos.protoDispatch.(IRestorePlayerReq)
			if !ok {
				return false
			}

			proto := &RestorePlayerReq{}
			if !proto.Read(b) {
				fmt.Println("read RestorePlayerReq fail, system error.")
				return false
			}

			t.OnRestorePlayerReq(proto)
		}
	case 16:
		{
			t, ok := protos.protoDispatch.(IRestorePlayerAck)
			if !ok {
				return false
			}

			proto := &RestorePlayerAck{}
			if !proto.Read(b) {
				fmt.Println("read RestorePlayerAck fail, system error.")
				return false
			}

			t.OnRestorePlayerAck(proto)
		}
	case 17:
		{
			t, ok := protos.protoDispatch.(IEnterGSReq)
			if !ok {
				return false
			}

			proto := &EnterGSReq{}
			if !proto.Read(b) {
				fmt.Println("read EnterGSReq fail, system error.")
				return false
			}

			t.OnEnterGSReq(proto)
		}
	case 18:
		{
			t, ok := protos.protoDispatch.(IEnterGSAck)
			if !ok {
				return false
			}

			proto := &EnterGSAck{}
			if !proto.Read(b) {
				fmt.Println("read EnterGSAck fail, system error.")
				return false
			}

			t.OnEnterGSAck(proto)
		}
	case 19:
		{
			t, ok := protos.protoDispatch.(ITrdLoginReq)
			if !ok {
				return false
			}

			proto := &TrdLoginReq{}
			if !proto.Read(b) {
				fmt.Println("read TrdLoginReq fail, system error.")
				return false
			}

			t.OnTrdLoginReq(proto)
		}
	case 20:
		{
			t, ok := protos.protoDispatch.(IGetWSTimestampReq)
			if !ok {
				return false
			}

			proto := &GetWSTimestampReq{}
			if !proto.Read(b) {
				fmt.Println("read GetWSTimestampReq fail, system error.")
				return false
			}

			t.OnGetWSTimestampReq(proto)
		}
	case 21:
		{
			t, ok := protos.protoDispatch.(IGetWSTimestampAck)
			if !ok {
				return false
			}

			proto := &GetWSTimestampAck{}
			if !proto.Read(b) {
				fmt.Println("read GetWSTimestampAck fail, system error.")
				return false
			}

			t.OnGetWSTimestampAck(proto)
		}
	case 22:
		{
			t, ok := protos.protoDispatch.(IRealnameInfoNtf)
			if !ok {
				return false
			}

			proto := &RealnameInfoNtf{}
			if !proto.Read(b) {
				fmt.Println("read RealnameInfoNtf fail, system error.")
				return false
			}

			t.OnRealnameInfoNtf(proto)
		}
	default:
		{
			fmt.Println("illegal protocol, Mid =", mid, "Pid =", pid)
		}
	}

	return true
}
