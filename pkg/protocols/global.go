///>本代码由测试工具自动生成,请勿手动修改
package protocols

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

type NewGSNtf struct { //>通知各服务器动态添加新GS
	Mid   uint16
	Pid   uint16
	Index uint32 //>新GS的索引id
}

type TerminateNtf struct { //>网络断开通知
	Mid     uint16
	Pid     uint16
	Session uint64 //>会话ID(由网关填写,用于多网关账号互顶区分)
}

type ErrorNtf struct { //>错误通知
	Mid     uint16
	Pid     uint16
	Errcode int32  //>0=成功, 其他表示错误码
	Errmsg  string //>错误码不为0时表示 错误消息
}

type ServerStopNtf struct { //>服务器停止通知
	Mid uint16
	Pid uint16
}

type ServerConfigNtf struct { //>服务器配置通知
	Mid     uint16
	Pid     uint16
	Svr_id  uint64     //>服务器id
	Limit   uint32     //>服务器人数限制
	Listen  []IPConfig //>监听列表
	Connect []IPConfig //>链接链表
}

type DBWrapperPkg struct { //>GS通过WS转发给DB的包装协议
	Mid     uint16
	Pid     uint16
	Wrapper []uint8 //>包装数据
}

type GSWrapperPkg struct { //>DB通过WS转发给GS的包装协议
	Mid     uint16
	Pid     uint16
	Index   uint32  //>gs索引
	Wrapper []uint8 //>包装数据
}

type LoadAuctionObjectDataReq struct { //>加载拍卖行数据请求
	Mid   uint16
	Pid   uint16
	Index uint32 //>gs索引
}

type LoadAuctionObjectDataAck struct { //>加载拍卖行数据回应
	Mid        uint16
	Pid        uint16
	Index      uint32              //>gs索引
	Total_num  uint32              //>总记录数
	Remain_num uint32              //>剩余记录数
	Datas      []AuctionObjectData //>拍卖行道具列表
	Errcode    int32               //>0=成功, 其他表示错误码
	Errmsg     string              //>错误码不为0时表示 错误消息
}

type AddAuctionObjectDataNtf struct { //>拍卖行道具添加通知
	Mid  uint16
	Pid  uint16
	Data AuctionObjectData //>拍卖行道具
}

type DelAuctionObjectDataNtf struct { //>拍卖行道具删除通知
	Mid  uint16
	Pid  uint16
	Guid uint64 //>拍卖行道具GUID
}

type UpdateAuctionObjectDataNtf struct { //>拍卖行道具更新通知
	Mid  uint16
	Pid  uint16
	Data AuctionObjectData //>拍卖行道具
}

type LoadAuctionCookieDataReq struct { //>加载拍卖行个人数据请求
	Mid   uint16
	Pid   uint16
	Index uint32 //>gs索引
}

type LoadAuctionCookieDataAck struct { //>加载拍卖行个人数据回应
	Mid        uint16
	Pid        uint16
	Index      uint32              //>gs索引
	Total_num  uint32              //>总记录数
	Remain_num uint32              //>剩余记录数
	Datas      []AuctionCookieData //>拍卖行个人数据列表
	Errcode    int32               //>0=成功, 其他表示错误码
	Errmsg     string              //>错误码不为0时表示 错误消息
}

type DuplicateAuctionCookieDataNtf struct { //>拍卖行个人数据更新(新建)通知
	Mid  uint16
	Pid  uint16
	Data AuctionCookieData //>拍卖行个人数据
}

type LoadGuildDataReq struct { //>加载帮派数据请求
	Mid   uint16
	Pid   uint16
	Index uint32 //>gs索引
}

type LoadGuildDataAck struct { //>帮派数据应答
	Mid     uint16
	Pid     uint16
	Index   uint32        //>gs索引
	Datas   []DBGuildData //>帮派数据
	Errcode int32         //>0=成功, 其他表示错误码
	Errmsg  string        //>错误码不为0时表示 错误消息
}

type LoadGuildMemberDataAck struct { //>帮派成员数据应答
	Mid     uint16
	Pid     uint16
	Index   uint32              //>gs索引
	Datas   []DBGuildMemberData //>帮派成员数据
	Errcode int32               //>0=成功, 其他表示错误码
	Errmsg  string              //>错误码不为0时表示 错误消息
}

type LoadGuildApplicantDataAck struct { //>帮派申请数据应答
	Mid     uint16
	Pid     uint16
	Index   uint32                 //>gs索引
	Datas   []DBGuildApplicantData //>帮派申请数据
	Errcode int32                  //>0=成功, 其他表示错误码
	Errmsg  string                 //>错误码不为0时表示 错误消息
}

type AddGuildDataNtf struct { //>添加帮派通知
	Mid   uint16
	Pid   uint16
	Guild DBGuildData //>帮派数据
}

type DelGuildDataNtf struct { //>删除帮派通知
	Mid  uint16
	Pid  uint16
	Guid uint64 //>帮派GUID
}

type UpdateGuildDataNtf struct { //>更新帮派通知
	Mid   uint16
	Pid   uint16
	Guild DBGuildData //>帮派数据
}

type AddGuildMemberDataNtf struct { //>添加帮派成员通知
	Mid    uint16
	Pid    uint16
	Member DBGuildMemberData //>帮派数据
}

type DelGuildMemberDataNtf struct { //>删除帮派成员通知
	Mid         uint16
	Pid         uint16
	Player_guid uint64 //>成员GUID
}

type UpdateGuildMemberDataNtf struct { //>更新帮派成员通知
	Mid    uint16
	Pid    uint16
	Member DBGuildMemberData //>帮派成员数据
}

type AddGuildApplicantDataNtf struct { //>添加帮派申请通知
	Mid       uint16
	Pid       uint16
	Applicant DBGuildApplicantData //>帮派申请数据
}

type DelGuildApplicantDataNtf struct { //>删除帮派申请通知
	Mid         uint16
	Pid         uint16
	Player_guid uint64 //>玩家GUID
	Guild_guid  uint64 //>帮派GUID
}

type UpdateGuildApplicantDataNtf struct { //>更新帮派申请通知
	Mid       uint16
	Pid       uint16
	Applicant DBGuildApplicantData //>帮派申请数据
}

type LoadContactDataReq struct { //>加载联系人数据请求
	Mid   uint16
	Pid   uint16
	Index uint32 //>gs索引
}

type LoadContactInfoDataAck struct { //>联系人信息数据应答
	Mid     uint16
	Pid     uint16
	Index   uint32            //>gs索引
	Datas   []ContactInfoData //>联系人信息数据
	Errcode int32             //>0=成功, 其他表示错误码
	Errmsg  string            //>错误码不为0时表示 错误消息
}

type LoadContactDataAck struct { //>加载联系人数据应答
	Mid     uint16
	Pid     uint16
	Index   uint32          //>gs索引
	Datas   []DBContactData //>联系人数据
	Errcode int32           //>0=成功, 其他表示错误码
	Errmsg  string          //>错误码不为0时表示 错误消息
}

type AddContactDataNtf struct { //>添加联系人数据通知
	Mid  uint16
	Pid  uint16
	Data DBContactData //>联系人数据
}

type DelContactDataNtf struct { //>删除联系人数据通知
	Mid         uint16
	Pid         uint16
	Player_guid uint64 //>玩家GUID
	Target_guid uint64 //>联系人GUID
}

type UpdateContactDataNtf struct { //>更新联系人数据通知
	Mid  uint16
	Pid  uint16
	Data DBContactData //>联系人数据
}

type LoadDBVarDataReq struct { //>加载DB变量数据请求
	Mid   uint16
	Pid   uint16
	Index uint32 //>gs索引
}

type LoadDBVarDataAck struct { //>加载DB变量数据应答
	Mid     uint16
	Pid     uint16
	Index   uint32      //>gs索引
	Datas   []DBVarData //>DB变量数据
	Errcode int32       //>0=成功, 其他表示错误码
	Errmsg  string      //>错误码不为0时表示 错误消息
}

type AddDBVarDataNtf struct { //>添加DB变量数据通知
	Mid  uint16
	Pid  uint16
	Data DBVarData //>DB变量数据
}

type DelDBVarDataNtf struct { //>删除联系人数据通知
	Mid        uint16
	Pid        uint16
	Key        []uint8 //>key值
	Value_type uint8   //>值类型
}

type LoadMailDataReq struct { //>加载邮件数据请求
	Mid       uint16
	Pid       uint16
	Index     uint32 //>gs索引
	Recv_guid uint64 //>收件人GUID
}

type LoadMailDataAck struct { //>加载邮件数据应答
	Mid       uint16
	Pid       uint16
	Index     uint32       //>gs索引
	Recv_guid uint64       //>收件人GUID
	Datas     []DBMailData //>邮件数据
	Errcode   int32        //>0=成功, 其他表示错误码
	Errmsg    string       //>错误码不为0时表示 错误消息
}

type AddMailDataNtf struct { //>添加邮件数据通知
	Mid  uint16
	Pid  uint16
	Data DBMailData //>邮件数据
}

type DelMailDataNtf struct { //>删除邮件数据通知
	Mid       uint16
	Pid       uint16
	Mail_guid uint64 //>邮件GUID
	Recv_guid uint64 //>收件人GUID
}

type UpdateMailDataNtf struct { //>更新邮件数据通知
	Mid  uint16
	Pid  uint16
	Data DBMailData //>邮件数据
}

type LoadRanklistDataReq struct { //>加载排行榜数据请求
	Mid   uint16
	Pid   uint16
	Index uint32 //>gs索引
}

type LoadRanklistDataAck struct { //>加载排行榜数据应答
	Mid     uint16
	Pid     uint16
	Index   uint32           //>gs索引
	Errcode int32            //>0=成功, 其他表示错误码
	Errmsg  string           //>错误码不为0时表示 错误消息
	Data    []DBRanklistData //>排行榜数据
}

type AddRanklistDataNtf struct { //>添加排行榜数据通知
	Mid  uint16
	Pid  uint16
	Data DBRanklistData //>排行榜数据
}

type DelRanklistDataNtf struct { //>删除排行榜数据通知
	Mid         uint16
	Pid         uint16
	Object_guid uint64 //>对象GUID
	Rank_type   int32  //>排行榜类型
}

type UpdateRanklistDataNtf struct { //>更新排行数据通知
	Mid  uint16
	Pid  uint16
	Data DBRanklistData //>邮件数据
}

type BillInMockReq struct { //>模拟充值通知请求
	Mid      uint16
	Pid      uint16
	Bill_id  uint32 //>订单ID
	User_id  uint32 //>用户ID
	Added_yb int32  //>用户充值元宝数
	Award_yb int32  //>用户奖励元宝数
	User     string //>用户名
	Desc     string //>充值描述
}

type AddPasturePetDataNtf struct { //>添加牧场宠物通知
	Mid  uint16
	Pid  uint16
	Data PasturePetData //>牧场宠物数据
}

type LoadPasturePetDataReq struct { //>加载牧场宠物请求
	Mid   uint16
	Pid   uint16
	Index uint32 //>gs索引
}

type LoadPasturePetDataAck struct { //>加载牧场宠物应答
	Mid     uint16
	Pid     uint16
	Index   uint32           //>gs索引
	Errcode int32            //>0=成功, 其他表示错误码
	Errmsg  string           //>错误码不为0时表示 错误消息
	Data    []PasturePetData //>牧场宠物
}

type FregmentPkg struct { //>数据包分片协议
	Mid       uint16
	Pid       uint16
	Frag_mark uint8   //>分片标志 0=开始 1=中间 2=结束
	Frag_data []uint8 //>包装数据
}

type DelPasturePetDataNtf struct { //>删除牧场宠物通知
	Mid      uint16
	Pid      uint16
	Pet_guid uint64 //>宠物GUID
}

type UpdatePasturePetDataNtf struct { //>更新牧场宠物通知
	Mid  uint16
	Pid  uint16
	Data PasturePetData //>牧场宠物数据
}

type LoadContactInfoReq struct { //>加载联系人信息请求
	Mid   uint16
	Pid   uint16
	Index uint32 //>gs索引
}

type LoadLadderDataReq struct { //>加载天梯数据请求
	Mid   uint16
	Pid   uint16
	Index uint32 //>gs索引
}

type LoadLadderDataAck struct { //>加载天梯数据应答
	Mid     uint16
	Pid     uint16
	Index   uint32       //>gs索引
	Errcode int32        //>0=成功, 其他表示错误码
	Errmsg  string       //>错误码不为0时表示 错误消息
	Data    []LadderData //>天梯数据
}

type AddLadderDataNtf struct { //>添加天梯数据通知
	Mid  uint16
	Pid  uint16
	Data LadderData //>天梯数据
}

type DelLadderDataNtf struct { //>删除天梯数据通知
	Mid         uint16
	Pid         uint16
	Player_guid uint64 //>玩家GUID
}

type UpdateLadderDataNtf struct { //>更新天梯数据通知
	Mid  uint16
	Pid  uint16
	Data LadderData //>天梯数据
}

type UnloadMailDataReq struct { //>卸载邮件数据请求
	Mid       uint16
	Pid       uint16
	Index     uint32 //>gs索引
	Recv_guid uint64 //>收件人GUID
}

type LoadOfflinePlayerReq struct { //>加载离线玩家
	Mid         uint16
	Pid         uint16
	Index       uint32 //>gs索引
	Seq         uint32 //>序号
	Target_guid uint64 //>目标GUID
	Object_type uint32 //>对象类型
	Object_guid uint64 //>对象GUID
}

type LoadOfflinePlayerAck struct { //>加载离线玩家应答
	Mid          uint16
	Pid          uint16
	Index        uint32     //>gs索引
	Seq          uint32     //>序号
	Target_guid  uint64     //>目标GUID
	Object_type  uint32     //>对象类型
	Object_guid  uint64     //>对象GUID
	Errcode      int32      //>0=成功, 其他表示错误码
	Errmsg       string     //>错误码不为0时表示 错误消息
	Offline_data PlayerData //>离线角色数据
}

type SyncViolateNtf struct { //>同步违禁字通知
	Mid      uint16
	Pid      uint16
	Violates []ViolateData //>违禁字数据
}

type SyncForbidTalkNtf struct { //>同步禁言数据通知
	Mid     uint16
	Pid     uint16
	Forbids []ForbidTalkData //>禁止聊天数据
}

type AddForbidTalkNtf struct { //>添加禁言数据通知
	Mid  uint16
	Pid  uint16
	Data ForbidTalkData //>禁止聊天数据
}

type DelForbidTalkNtf struct { //>删除禁言数据通知
	Mid  uint16
	Pid  uint16
	Guid uint64 //>角色GUID
}

type LogLevelNtf struct { //>日志级别变更通知
	Mid    uint16
	Pid    uint16
	Level  uint32 //>日志级别 0=LogSys 1=LogErr 2=LogWrn 3=LogInf 4=LogDbg
	Enable uint32 //>0关闭 1开启
}

type BillQueryReq struct { //>充值记录请求
	Mid         uint16
	Pid         uint16
	Index       uint32 //>gs索引
	User_id     uint32 //>用户ID
	Start_time  uint32 //>起始时间
	End_time    uint32 //>终止时间
	Callback    string //>回调函数
	Player_guid uint64 //>角色GUID
}

type BillQueryAck struct { //>充值记录回包
	Mid         uint16
	Pid         uint16
	Index       uint32          //>gs索引
	Datas       []BillQueryData //>充值记录
	Errcode     int32           //>0=成功, 其他表示错误码
	Errmsg      string          //>错误码不为0时表示 错误消息
	Callback    string          //>回调函数
	Player_guid uint64          //>角色GUID
}

func (proto *NewGSNtf) GetMid() uint16 {
	return 105
}

func (proto *NewGSNtf) GetPid() uint16 {
	return 1
}

func (proto *NewGSNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(1)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *NewGSNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *TerminateNtf) GetMid() uint16 {
	return 105
}

func (proto *TerminateNtf) GetPid() uint16 {
	return 2
}

func (proto *TerminateNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(2)) {
		return false
	}

	if !writeProtoInteger(b, proto.Session) {
		return false
	}

	return true
}

func (proto *TerminateNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Session) {
		return false
	}

	return true
}

func (proto *ErrorNtf) GetMid() uint16 {
	return 105
}

func (proto *ErrorNtf) GetPid() uint16 {
	return 3
}

func (proto *ErrorNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(3)) {
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

func (proto *ErrorNtf) Read(b *bytes.Buffer) bool {
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

func (proto *ServerStopNtf) GetMid() uint16 {
	return 105
}

func (proto *ServerStopNtf) GetPid() uint16 {
	return 4
}

func (proto *ServerStopNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(4)) {
		return false
	}

	return true
}

func (proto *ServerStopNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	return true
}

func (proto *ServerConfigNtf) GetMid() uint16 {
	return 105
}

func (proto *ServerConfigNtf) GetPid() uint16 {
	return 5
}

func (proto *ServerConfigNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(5)) {
		return false
	}

	if !writeProtoInteger(b, proto.Svr_id) {
		return false
	}

	if !writeProtoInteger(b, proto.Limit) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Listen, uint8(255)) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Connect, uint8(255)) {
		return false
	}

	return true
}

func (proto *ServerConfigNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Svr_id) {
		return false
	}

	if !readProtoInteger(b, &proto.Limit) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Listen, uint8(255)) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Connect, uint8(255)) {
		return false
	}

	return true
}

func (proto *DBWrapperPkg) GetMid() uint16 {
	return 105
}

func (proto *DBWrapperPkg) GetPid() uint16 {
	return 6
}

func (proto *DBWrapperPkg) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(6)) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Wrapper, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *DBWrapperPkg) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Wrapper, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *GSWrapperPkg) GetMid() uint16 {
	return 105
}

func (proto *GSWrapperPkg) GetPid() uint16 {
	return 7
}

func (proto *GSWrapperPkg) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(7)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Wrapper, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *GSWrapperPkg) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Wrapper, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *LoadAuctionObjectDataReq) GetMid() uint16 {
	return 105
}

func (proto *LoadAuctionObjectDataReq) GetPid() uint16 {
	return 8
}

func (proto *LoadAuctionObjectDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(8)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *LoadAuctionObjectDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *LoadAuctionObjectDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadAuctionObjectDataAck) GetPid() uint16 {
	return 9
}

func (proto *LoadAuctionObjectDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(9)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Total_num) {
		return false
	}

	if !writeProtoInteger(b, proto.Remain_num) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint32(4294967295)) {
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

func (proto *LoadAuctionObjectDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Total_num) {
		return false
	}

	if !readProtoInteger(b, &proto.Remain_num) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint32(4294967295)) {
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

func (proto *AddAuctionObjectDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddAuctionObjectDataNtf) GetPid() uint16 {
	return 10
}

func (proto *AddAuctionObjectDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(10)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *AddAuctionObjectDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *DelAuctionObjectDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelAuctionObjectDataNtf) GetPid() uint16 {
	return 11
}

func (proto *DelAuctionObjectDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(11)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *DelAuctionObjectDataNtf) Read(b *bytes.Buffer) bool {
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

func (proto *UpdateAuctionObjectDataNtf) GetMid() uint16 {
	return 105
}

func (proto *UpdateAuctionObjectDataNtf) GetPid() uint16 {
	return 12
}

func (proto *UpdateAuctionObjectDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(12)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *UpdateAuctionObjectDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *LoadAuctionCookieDataReq) GetMid() uint16 {
	return 105
}

func (proto *LoadAuctionCookieDataReq) GetPid() uint16 {
	return 13
}

func (proto *LoadAuctionCookieDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(13)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *LoadAuctionCookieDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *LoadAuctionCookieDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadAuctionCookieDataAck) GetPid() uint16 {
	return 14
}

func (proto *LoadAuctionCookieDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(14)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Total_num) {
		return false
	}

	if !writeProtoInteger(b, proto.Remain_num) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint32(4294967295)) {
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

func (proto *LoadAuctionCookieDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Total_num) {
		return false
	}

	if !readProtoInteger(b, &proto.Remain_num) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint32(4294967295)) {
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

func (proto *DuplicateAuctionCookieDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DuplicateAuctionCookieDataNtf) GetPid() uint16 {
	return 15
}

func (proto *DuplicateAuctionCookieDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(15)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *DuplicateAuctionCookieDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *LoadGuildDataReq) GetMid() uint16 {
	return 105
}

func (proto *LoadGuildDataReq) GetPid() uint16 {
	return 16
}

func (proto *LoadGuildDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(16)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *LoadGuildDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *LoadGuildDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadGuildDataAck) GetPid() uint16 {
	return 17
}

func (proto *LoadGuildDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(17)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint16(65535)) {
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

func (proto *LoadGuildDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint16(65535)) {
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

func (proto *LoadGuildMemberDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadGuildMemberDataAck) GetPid() uint16 {
	return 18
}

func (proto *LoadGuildMemberDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(18)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint32(4294967295)) {
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

func (proto *LoadGuildMemberDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint32(4294967295)) {
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

func (proto *LoadGuildApplicantDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadGuildApplicantDataAck) GetPid() uint16 {
	return 19
}

func (proto *LoadGuildApplicantDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(19)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint32(4294967295)) {
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

func (proto *LoadGuildApplicantDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint32(4294967295)) {
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

func (proto *AddGuildDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddGuildDataNtf) GetPid() uint16 {
	return 20
}

func (proto *AddGuildDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(20)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Guild) {
		return false
	}

	return true
}

func (proto *AddGuildDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Guild) {
		return false
	}

	return true
}

func (proto *DelGuildDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelGuildDataNtf) GetPid() uint16 {
	return 21
}

func (proto *DelGuildDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(21)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *DelGuildDataNtf) Read(b *bytes.Buffer) bool {
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

func (proto *UpdateGuildDataNtf) GetMid() uint16 {
	return 105
}

func (proto *UpdateGuildDataNtf) GetPid() uint16 {
	return 22
}

func (proto *UpdateGuildDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(22)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Guild) {
		return false
	}

	return true
}

func (proto *UpdateGuildDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Guild) {
		return false
	}

	return true
}

func (proto *AddGuildMemberDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddGuildMemberDataNtf) GetPid() uint16 {
	return 23
}

func (proto *AddGuildMemberDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(23)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Member) {
		return false
	}

	return true
}

func (proto *AddGuildMemberDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Member) {
		return false
	}

	return true
}

func (proto *DelGuildMemberDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelGuildMemberDataNtf) GetPid() uint16 {
	return 24
}

func (proto *DelGuildMemberDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(24)) {
		return false
	}

	if !writeProtoInteger(b, proto.Player_guid) {
		return false
	}

	return true
}

func (proto *DelGuildMemberDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Player_guid) {
		return false
	}

	return true
}

func (proto *UpdateGuildMemberDataNtf) GetMid() uint16 {
	return 105
}

func (proto *UpdateGuildMemberDataNtf) GetPid() uint16 {
	return 25
}

func (proto *UpdateGuildMemberDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(25)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Member) {
		return false
	}

	return true
}

func (proto *UpdateGuildMemberDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Member) {
		return false
	}

	return true
}

func (proto *AddGuildApplicantDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddGuildApplicantDataNtf) GetPid() uint16 {
	return 26
}

func (proto *AddGuildApplicantDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(26)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Applicant) {
		return false
	}

	return true
}

func (proto *AddGuildApplicantDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Applicant) {
		return false
	}

	return true
}

func (proto *DelGuildApplicantDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelGuildApplicantDataNtf) GetPid() uint16 {
	return 27
}

func (proto *DelGuildApplicantDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(27)) {
		return false
	}

	if !writeProtoInteger(b, proto.Player_guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Guild_guid) {
		return false
	}

	return true
}

func (proto *DelGuildApplicantDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Player_guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Guild_guid) {
		return false
	}

	return true
}

func (proto *UpdateGuildApplicantDataNtf) GetMid() uint16 {
	return 105
}

func (proto *UpdateGuildApplicantDataNtf) GetPid() uint16 {
	return 28
}

func (proto *UpdateGuildApplicantDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(28)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Applicant) {
		return false
	}

	return true
}

func (proto *UpdateGuildApplicantDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Applicant) {
		return false
	}

	return true
}

func (proto *LoadContactDataReq) GetMid() uint16 {
	return 105
}

func (proto *LoadContactDataReq) GetPid() uint16 {
	return 29
}

func (proto *LoadContactDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(29)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *LoadContactDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *LoadContactInfoDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadContactInfoDataAck) GetPid() uint16 {
	return 30
}

func (proto *LoadContactInfoDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(30)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint32(4294967295)) {
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

func (proto *LoadContactInfoDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint32(4294967295)) {
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

func (proto *LoadContactDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadContactDataAck) GetPid() uint16 {
	return 31
}

func (proto *LoadContactDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(31)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint32(4294967295)) {
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

func (proto *LoadContactDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint32(4294967295)) {
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

func (proto *AddContactDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddContactDataNtf) GetPid() uint16 {
	return 32
}

func (proto *AddContactDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(32)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *AddContactDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *DelContactDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelContactDataNtf) GetPid() uint16 {
	return 33
}

func (proto *DelContactDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(33)) {
		return false
	}

	if !writeProtoInteger(b, proto.Player_guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Target_guid) {
		return false
	}

	return true
}

func (proto *DelContactDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Player_guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Target_guid) {
		return false
	}

	return true
}

func (proto *UpdateContactDataNtf) GetMid() uint16 {
	return 105
}

func (proto *UpdateContactDataNtf) GetPid() uint16 {
	return 34
}

func (proto *UpdateContactDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(34)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *UpdateContactDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *LoadDBVarDataReq) GetMid() uint16 {
	return 105
}

func (proto *LoadDBVarDataReq) GetPid() uint16 {
	return 35
}

func (proto *LoadDBVarDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(35)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *LoadDBVarDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *LoadDBVarDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadDBVarDataAck) GetPid() uint16 {
	return 36
}

func (proto *LoadDBVarDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(36)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint32(4294967295)) {
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

func (proto *LoadDBVarDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint32(4294967295)) {
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

func (proto *AddDBVarDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddDBVarDataNtf) GetPid() uint16 {
	return 37
}

func (proto *AddDBVarDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(37)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *AddDBVarDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *DelDBVarDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelDBVarDataNtf) GetPid() uint16 {
	return 38
}

func (proto *DelDBVarDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(38)) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Key, uint8(255)) {
		return false
	}

	if !writeProtoInteger(b, proto.Value_type) {
		return false
	}

	return true
}

func (proto *DelDBVarDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Key, uint8(255)) {
		return false
	}

	if !readProtoInteger(b, &proto.Value_type) {
		return false
	}

	return true
}

func (proto *LoadMailDataReq) GetMid() uint16 {
	return 105
}

func (proto *LoadMailDataReq) GetPid() uint16 {
	return 39
}

func (proto *LoadMailDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(39)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Recv_guid) {
		return false
	}

	return true
}

func (proto *LoadMailDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Recv_guid) {
		return false
	}

	return true
}

func (proto *LoadMailDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadMailDataAck) GetPid() uint16 {
	return 40
}

func (proto *LoadMailDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(40)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Recv_guid) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint16(65535)) {
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

func (proto *LoadMailDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Recv_guid) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint16(65535)) {
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

func (proto *AddMailDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddMailDataNtf) GetPid() uint16 {
	return 41
}

func (proto *AddMailDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(41)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *AddMailDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *DelMailDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelMailDataNtf) GetPid() uint16 {
	return 42
}

func (proto *DelMailDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(42)) {
		return false
	}

	if !writeProtoInteger(b, proto.Mail_guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Recv_guid) {
		return false
	}

	return true
}

func (proto *DelMailDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Mail_guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Recv_guid) {
		return false
	}

	return true
}

func (proto *UpdateMailDataNtf) GetMid() uint16 {
	return 105
}

func (proto *UpdateMailDataNtf) GetPid() uint16 {
	return 43
}

func (proto *UpdateMailDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(43)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *UpdateMailDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *LoadRanklistDataReq) GetMid() uint16 {
	return 105
}

func (proto *LoadRanklistDataReq) GetPid() uint16 {
	return 44
}

func (proto *LoadRanklistDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(44)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *LoadRanklistDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *LoadRanklistDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadRanklistDataAck) GetPid() uint16 {
	return 45
}

func (proto *LoadRanklistDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(45)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Data, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *LoadRanklistDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Data, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *AddRanklistDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddRanklistDataNtf) GetPid() uint16 {
	return 46
}

func (proto *AddRanklistDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(46)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *AddRanklistDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *DelRanklistDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelRanklistDataNtf) GetPid() uint16 {
	return 47
}

func (proto *DelRanklistDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(47)) {
		return false
	}

	if !writeProtoInteger(b, proto.Object_guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Rank_type) {
		return false
	}

	return true
}

func (proto *DelRanklistDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Object_guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Rank_type) {
		return false
	}

	return true
}

func (proto *UpdateRanklistDataNtf) GetMid() uint16 {
	return 105
}

func (proto *UpdateRanklistDataNtf) GetPid() uint16 {
	return 48
}

func (proto *UpdateRanklistDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(48)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *UpdateRanklistDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *BillInMockReq) GetMid() uint16 {
	return 105
}

func (proto *BillInMockReq) GetPid() uint16 {
	return 49
}

func (proto *BillInMockReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(49)) {
		return false
	}

	if !writeProtoInteger(b, proto.Bill_id) {
		return false
	}

	if !writeProtoInteger(b, proto.User_id) {
		return false
	}

	if !writeProtoInteger(b, proto.Added_yb) {
		return false
	}

	if !writeProtoInteger(b, proto.Award_yb) {
		return false
	}

	if !writeProtoString(b, proto.User, 20) {
		return false
	}

	if !writeProtoString(b, proto.Desc, 255) {
		return false
	}

	return true
}

func (proto *BillInMockReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Bill_id) {
		return false
	}

	if !readProtoInteger(b, &proto.User_id) {
		return false
	}

	if !readProtoInteger(b, &proto.Added_yb) {
		return false
	}

	if !readProtoInteger(b, &proto.Award_yb) {
		return false
	}

	if !readProtoString(b, &proto.User, 20) {
		return false
	}

	if !readProtoString(b, &proto.Desc, 255) {
		return false
	}

	return true
}

func (proto *AddPasturePetDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddPasturePetDataNtf) GetPid() uint16 {
	return 50
}

func (proto *AddPasturePetDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(50)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *AddPasturePetDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *LoadPasturePetDataReq) GetMid() uint16 {
	return 105
}

func (proto *LoadPasturePetDataReq) GetPid() uint16 {
	return 51
}

func (proto *LoadPasturePetDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(51)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *LoadPasturePetDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *LoadPasturePetDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadPasturePetDataAck) GetPid() uint16 {
	return 52
}

func (proto *LoadPasturePetDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(52)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Data, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *LoadPasturePetDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Data, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *FregmentPkg) GetMid() uint16 {
	return 105
}

func (proto *FregmentPkg) GetPid() uint16 {
	return 53
}

func (proto *FregmentPkg) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(53)) {
		return false
	}

	if !writeProtoInteger(b, proto.Frag_mark) {
		return false
	}

	if !writeProtoIntegerArray(b, proto.Frag_data, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *FregmentPkg) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Frag_mark) {
		return false
	}

	if !readProtoIntegerArray(b, &proto.Frag_data, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *DelPasturePetDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelPasturePetDataNtf) GetPid() uint16 {
	return 54
}

func (proto *DelPasturePetDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(54)) {
		return false
	}

	if !writeProtoInteger(b, proto.Pet_guid) {
		return false
	}

	return true
}

func (proto *DelPasturePetDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pet_guid) {
		return false
	}

	return true
}

func (proto *UpdatePasturePetDataNtf) GetMid() uint16 {
	return 105
}

func (proto *UpdatePasturePetDataNtf) GetPid() uint16 {
	return 55
}

func (proto *UpdatePasturePetDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(55)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *UpdatePasturePetDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *LoadContactInfoReq) GetMid() uint16 {
	return 105
}

func (proto *LoadContactInfoReq) GetPid() uint16 {
	return 56
}

func (proto *LoadContactInfoReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(56)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *LoadContactInfoReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *LoadLadderDataReq) GetMid() uint16 {
	return 105
}

func (proto *LoadLadderDataReq) GetPid() uint16 {
	return 57
}

func (proto *LoadLadderDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(57)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	return true
}

func (proto *LoadLadderDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	return true
}

func (proto *LoadLadderDataAck) GetMid() uint16 {
	return 105
}

func (proto *LoadLadderDataAck) GetPid() uint16 {
	return 58
}

func (proto *LoadLadderDataAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(58)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Data, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *LoadLadderDataAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Data, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *AddLadderDataNtf) GetMid() uint16 {
	return 105
}

func (proto *AddLadderDataNtf) GetPid() uint16 {
	return 59
}

func (proto *AddLadderDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(59)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *AddLadderDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *DelLadderDataNtf) GetMid() uint16 {
	return 105
}

func (proto *DelLadderDataNtf) GetPid() uint16 {
	return 60
}

func (proto *DelLadderDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(60)) {
		return false
	}

	if !writeProtoInteger(b, proto.Player_guid) {
		return false
	}

	return true
}

func (proto *DelLadderDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Player_guid) {
		return false
	}

	return true
}

func (proto *UpdateLadderDataNtf) GetMid() uint16 {
	return 105
}

func (proto *UpdateLadderDataNtf) GetPid() uint16 {
	return 61
}

func (proto *UpdateLadderDataNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(61)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *UpdateLadderDataNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *UnloadMailDataReq) GetMid() uint16 {
	return 105
}

func (proto *UnloadMailDataReq) GetPid() uint16 {
	return 62
}

func (proto *UnloadMailDataReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(62)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Recv_guid) {
		return false
	}

	return true
}

func (proto *UnloadMailDataReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Recv_guid) {
		return false
	}

	return true
}

func (proto *LoadOfflinePlayerReq) GetMid() uint16 {
	return 105
}

func (proto *LoadOfflinePlayerReq) GetPid() uint16 {
	return 63
}

func (proto *LoadOfflinePlayerReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(63)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Seq) {
		return false
	}

	if !writeProtoInteger(b, proto.Target_guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Object_type) {
		return false
	}

	if !writeProtoInteger(b, proto.Object_guid) {
		return false
	}

	return true
}

func (proto *LoadOfflinePlayerReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Seq) {
		return false
	}

	if !readProtoInteger(b, &proto.Target_guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Object_type) {
		return false
	}

	if !readProtoInteger(b, &proto.Object_guid) {
		return false
	}

	return true
}

func (proto *LoadOfflinePlayerAck) GetMid() uint16 {
	return 105
}

func (proto *LoadOfflinePlayerAck) GetPid() uint16 {
	return 64
}

func (proto *LoadOfflinePlayerAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(64)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.Seq) {
		return false
	}

	if !writeProtoInteger(b, proto.Target_guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Object_type) {
		return false
	}

	if !writeProtoInteger(b, proto.Object_guid) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	if !writeProtoCustom(b, &proto.Offline_data) {
		return false
	}

	return true
}

func (proto *LoadOfflinePlayerAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.Seq) {
		return false
	}

	if !readProtoInteger(b, &proto.Target_guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Object_type) {
		return false
	}

	if !readProtoInteger(b, &proto.Object_guid) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	if !readProtoCustom(b, &proto.Offline_data) {
		return false
	}

	return true
}

func (proto *SyncViolateNtf) GetMid() uint16 {
	return 105
}

func (proto *SyncViolateNtf) GetPid() uint16 {
	return 65
}

func (proto *SyncViolateNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(65)) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Violates, uint16(65535)) {
		return false
	}

	return true
}

func (proto *SyncViolateNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Violates, uint16(65535)) {
		return false
	}

	return true
}

func (proto *SyncForbidTalkNtf) GetMid() uint16 {
	return 105
}

func (proto *SyncForbidTalkNtf) GetPid() uint16 {
	return 66
}

func (proto *SyncForbidTalkNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(66)) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Forbids, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *SyncForbidTalkNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Forbids, uint32(4294967295)) {
		return false
	}

	return true
}

func (proto *AddForbidTalkNtf) GetMid() uint16 {
	return 105
}

func (proto *AddForbidTalkNtf) GetPid() uint16 {
	return 67
}

func (proto *AddForbidTalkNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(67)) {
		return false
	}

	if !writeProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *AddForbidTalkNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoCustom(b, &proto.Data) {
		return false
	}

	return true
}

func (proto *DelForbidTalkNtf) GetMid() uint16 {
	return 105
}

func (proto *DelForbidTalkNtf) GetPid() uint16 {
	return 68
}

func (proto *DelForbidTalkNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(68)) {
		return false
	}

	if !writeProtoInteger(b, proto.Guid) {
		return false
	}

	return true
}

func (proto *DelForbidTalkNtf) Read(b *bytes.Buffer) bool {
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

func (proto *LogLevelNtf) GetMid() uint16 {
	return 105
}

func (proto *LogLevelNtf) GetPid() uint16 {
	return 69
}

func (proto *LogLevelNtf) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(69)) {
		return false
	}

	if !writeProtoInteger(b, proto.Level) {
		return false
	}

	if !writeProtoInteger(b, proto.Enable) {
		return false
	}

	return true
}

func (proto *LogLevelNtf) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Level) {
		return false
	}

	if !readProtoInteger(b, &proto.Enable) {
		return false
	}

	return true
}

func (proto *BillQueryReq) GetMid() uint16 {
	return 105
}

func (proto *BillQueryReq) GetPid() uint16 {
	return 70
}

func (proto *BillQueryReq) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(70)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoInteger(b, proto.User_id) {
		return false
	}

	if !writeProtoInteger(b, proto.Start_time) {
		return false
	}

	if !writeProtoInteger(b, proto.End_time) {
		return false
	}

	if !writeProtoString(b, proto.Callback, 255) {
		return false
	}

	if !writeProtoInteger(b, proto.Player_guid) {
		return false
	}

	return true
}

func (proto *BillQueryReq) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoInteger(b, &proto.User_id) {
		return false
	}

	if !readProtoInteger(b, &proto.Start_time) {
		return false
	}

	if !readProtoInteger(b, &proto.End_time) {
		return false
	}

	if !readProtoString(b, &proto.Callback, 255) {
		return false
	}

	if !readProtoInteger(b, &proto.Player_guid) {
		return false
	}

	return true
}

func (proto *BillQueryAck) GetMid() uint16 {
	return 105
}

func (proto *BillQueryAck) GetPid() uint16 {
	return 71
}

func (proto *BillQueryAck) Write(b *bytes.Buffer) bool {
	if !writeProtoInteger(b, uint16(105)) {
		return false
	}

	if !writeProtoInteger(b, uint16(71)) {
		return false
	}

	if !writeProtoInteger(b, proto.Index) {
		return false
	}

	if !writeProtoCustomArray(b, proto.Datas, uint32(4294967295)) {
		return false
	}

	if !writeProtoInteger(b, proto.Errcode) {
		return false
	}

	if !writeProtoString(b, proto.Errmsg, 255) {
		return false
	}

	if !writeProtoString(b, proto.Callback, 255) {
		return false
	}

	if !writeProtoInteger(b, proto.Player_guid) {
		return false
	}

	return true
}

func (proto *BillQueryAck) Read(b *bytes.Buffer) bool {
	if !readProtoInteger(b, &proto.Mid) {
		return false
	}

	if !readProtoInteger(b, &proto.Pid) {
		return false
	}

	if !readProtoInteger(b, &proto.Index) {
		return false
	}

	if !readProtoCustomArray(b, &proto.Datas, uint32(4294967295)) {
		return false
	}

	if !readProtoInteger(b, &proto.Errcode) {
		return false
	}

	if !readProtoString(b, &proto.Errmsg, 255) {
		return false
	}

	if !readProtoString(b, &proto.Callback, 255) {
		return false
	}

	if !readProtoInteger(b, &proto.Player_guid) {
		return false
	}

	return true
}

type INewGSNtf interface {
	OnNewGSNtf(proto *NewGSNtf)
}

type ITerminateNtf interface {
	OnTerminateNtf(proto *TerminateNtf)
}

type IErrorNtf interface {
	OnErrorNtf(proto *ErrorNtf)
}

type IServerStopNtf interface {
	OnServerStopNtf(proto *ServerStopNtf)
}

type IServerConfigNtf interface {
	OnServerConfigNtf(proto *ServerConfigNtf)
}

type IDBWrapperPkg interface {
	OnDBWrapperPkg(proto *DBWrapperPkg)
}

type IGSWrapperPkg interface {
	OnGSWrapperPkg(proto *GSWrapperPkg)
}

type ILoadAuctionObjectDataReq interface {
	OnLoadAuctionObjectDataReq(proto *LoadAuctionObjectDataReq)
}

type ILoadAuctionObjectDataAck interface {
	OnLoadAuctionObjectDataAck(proto *LoadAuctionObjectDataAck)
}

type IAddAuctionObjectDataNtf interface {
	OnAddAuctionObjectDataNtf(proto *AddAuctionObjectDataNtf)
}

type IDelAuctionObjectDataNtf interface {
	OnDelAuctionObjectDataNtf(proto *DelAuctionObjectDataNtf)
}

type IUpdateAuctionObjectDataNtf interface {
	OnUpdateAuctionObjectDataNtf(proto *UpdateAuctionObjectDataNtf)
}

type ILoadAuctionCookieDataReq interface {
	OnLoadAuctionCookieDataReq(proto *LoadAuctionCookieDataReq)
}

type ILoadAuctionCookieDataAck interface {
	OnLoadAuctionCookieDataAck(proto *LoadAuctionCookieDataAck)
}

type IDuplicateAuctionCookieDataNtf interface {
	OnDuplicateAuctionCookieDataNtf(proto *DuplicateAuctionCookieDataNtf)
}

type ILoadGuildDataReq interface {
	OnLoadGuildDataReq(proto *LoadGuildDataReq)
}

type ILoadGuildDataAck interface {
	OnLoadGuildDataAck(proto *LoadGuildDataAck)
}

type ILoadGuildMemberDataAck interface {
	OnLoadGuildMemberDataAck(proto *LoadGuildMemberDataAck)
}

type ILoadGuildApplicantDataAck interface {
	OnLoadGuildApplicantDataAck(proto *LoadGuildApplicantDataAck)
}

type IAddGuildDataNtf interface {
	OnAddGuildDataNtf(proto *AddGuildDataNtf)
}

type IDelGuildDataNtf interface {
	OnDelGuildDataNtf(proto *DelGuildDataNtf)
}

type IUpdateGuildDataNtf interface {
	OnUpdateGuildDataNtf(proto *UpdateGuildDataNtf)
}

type IAddGuildMemberDataNtf interface {
	OnAddGuildMemberDataNtf(proto *AddGuildMemberDataNtf)
}

type IDelGuildMemberDataNtf interface {
	OnDelGuildMemberDataNtf(proto *DelGuildMemberDataNtf)
}

type IUpdateGuildMemberDataNtf interface {
	OnUpdateGuildMemberDataNtf(proto *UpdateGuildMemberDataNtf)
}

type IAddGuildApplicantDataNtf interface {
	OnAddGuildApplicantDataNtf(proto *AddGuildApplicantDataNtf)
}

type IDelGuildApplicantDataNtf interface {
	OnDelGuildApplicantDataNtf(proto *DelGuildApplicantDataNtf)
}

type IUpdateGuildApplicantDataNtf interface {
	OnUpdateGuildApplicantDataNtf(proto *UpdateGuildApplicantDataNtf)
}

type ILoadContactDataReq interface {
	OnLoadContactDataReq(proto *LoadContactDataReq)
}

type ILoadContactInfoDataAck interface {
	OnLoadContactInfoDataAck(proto *LoadContactInfoDataAck)
}

type ILoadContactDataAck interface {
	OnLoadContactDataAck(proto *LoadContactDataAck)
}

type IAddContactDataNtf interface {
	OnAddContactDataNtf(proto *AddContactDataNtf)
}

type IDelContactDataNtf interface {
	OnDelContactDataNtf(proto *DelContactDataNtf)
}

type IUpdateContactDataNtf interface {
	OnUpdateContactDataNtf(proto *UpdateContactDataNtf)
}

type ILoadDBVarDataReq interface {
	OnLoadDBVarDataReq(proto *LoadDBVarDataReq)
}

type ILoadDBVarDataAck interface {
	OnLoadDBVarDataAck(proto *LoadDBVarDataAck)
}

type IAddDBVarDataNtf interface {
	OnAddDBVarDataNtf(proto *AddDBVarDataNtf)
}

type IDelDBVarDataNtf interface {
	OnDelDBVarDataNtf(proto *DelDBVarDataNtf)
}

type ILoadMailDataReq interface {
	OnLoadMailDataReq(proto *LoadMailDataReq)
}

type ILoadMailDataAck interface {
	OnLoadMailDataAck(proto *LoadMailDataAck)
}

type IAddMailDataNtf interface {
	OnAddMailDataNtf(proto *AddMailDataNtf)
}

type IDelMailDataNtf interface {
	OnDelMailDataNtf(proto *DelMailDataNtf)
}

type IUpdateMailDataNtf interface {
	OnUpdateMailDataNtf(proto *UpdateMailDataNtf)
}

type ILoadRanklistDataReq interface {
	OnLoadRanklistDataReq(proto *LoadRanklistDataReq)
}

type ILoadRanklistDataAck interface {
	OnLoadRanklistDataAck(proto *LoadRanklistDataAck)
}

type IAddRanklistDataNtf interface {
	OnAddRanklistDataNtf(proto *AddRanklistDataNtf)
}

type IDelRanklistDataNtf interface {
	OnDelRanklistDataNtf(proto *DelRanklistDataNtf)
}

type IUpdateRanklistDataNtf interface {
	OnUpdateRanklistDataNtf(proto *UpdateRanklistDataNtf)
}

type IBillInMockReq interface {
	OnBillInMockReq(proto *BillInMockReq)
}

type IAddPasturePetDataNtf interface {
	OnAddPasturePetDataNtf(proto *AddPasturePetDataNtf)
}

type ILoadPasturePetDataReq interface {
	OnLoadPasturePetDataReq(proto *LoadPasturePetDataReq)
}

type ILoadPasturePetDataAck interface {
	OnLoadPasturePetDataAck(proto *LoadPasturePetDataAck)
}

type IFregmentPkg interface {
	OnFregmentPkg(proto *FregmentPkg)
}

type IDelPasturePetDataNtf interface {
	OnDelPasturePetDataNtf(proto *DelPasturePetDataNtf)
}

type IUpdatePasturePetDataNtf interface {
	OnUpdatePasturePetDataNtf(proto *UpdatePasturePetDataNtf)
}

type ILoadContactInfoReq interface {
	OnLoadContactInfoReq(proto *LoadContactInfoReq)
}

type ILoadLadderDataReq interface {
	OnLoadLadderDataReq(proto *LoadLadderDataReq)
}

type ILoadLadderDataAck interface {
	OnLoadLadderDataAck(proto *LoadLadderDataAck)
}

type IAddLadderDataNtf interface {
	OnAddLadderDataNtf(proto *AddLadderDataNtf)
}

type IDelLadderDataNtf interface {
	OnDelLadderDataNtf(proto *DelLadderDataNtf)
}

type IUpdateLadderDataNtf interface {
	OnUpdateLadderDataNtf(proto *UpdateLadderDataNtf)
}

type IUnloadMailDataReq interface {
	OnUnloadMailDataReq(proto *UnloadMailDataReq)
}

type ILoadOfflinePlayerReq interface {
	OnLoadOfflinePlayerReq(proto *LoadOfflinePlayerReq)
}

type ILoadOfflinePlayerAck interface {
	OnLoadOfflinePlayerAck(proto *LoadOfflinePlayerAck)
}

type ISyncViolateNtf interface {
	OnSyncViolateNtf(proto *SyncViolateNtf)
}

type ISyncForbidTalkNtf interface {
	OnSyncForbidTalkNtf(proto *SyncForbidTalkNtf)
}

type IAddForbidTalkNtf interface {
	OnAddForbidTalkNtf(proto *AddForbidTalkNtf)
}

type IDelForbidTalkNtf interface {
	OnDelForbidTalkNtf(proto *DelForbidTalkNtf)
}

type ILogLevelNtf interface {
	OnLogLevelNtf(proto *LogLevelNtf)
}

type IBillQueryReq interface {
	OnBillQueryReq(proto *BillQueryReq)
}

type IBillQueryAck interface {
	OnBillQueryAck(proto *BillQueryAck)
}

type Global struct {
	protoDispatch interface{}
}

func NewGlobal[T any](dispatch *T) *Global {
	return &Global{dispatch}
}

func (protos *Global) GetMid() uint16 {
	return 105
}

func (protos *Global) DispatchProto(data []byte) bool {
	b := bytes.NewBuffer(data)

	mid := binary.LittleEndian.Uint16(data)
	if mid != protos.GetMid() {
		return false
	}

	pid := binary.LittleEndian.Uint16(data[unsafe.Sizeof(uint16(0)):])
	switch pid {
	case 1:
		{
			t, ok := protos.protoDispatch.(INewGSNtf)
			if !ok {
				return false
			}

			proto := &NewGSNtf{}
			if !proto.Read(b) {
				fmt.Println("read NewGSNtf fail, system error.")
				return false
			}

			t.OnNewGSNtf(proto)
		}
	case 2:
		{
			t, ok := protos.protoDispatch.(ITerminateNtf)
			if !ok {
				return false
			}

			proto := &TerminateNtf{}
			if !proto.Read(b) {
				fmt.Println("read TerminateNtf fail, system error.")
				return false
			}

			t.OnTerminateNtf(proto)
		}
	case 3:
		{
			t, ok := protos.protoDispatch.(IErrorNtf)
			if !ok {
				return false
			}

			proto := &ErrorNtf{}
			if !proto.Read(b) {
				fmt.Println("read ErrorNtf fail, system error.")
				return false
			}

			t.OnErrorNtf(proto)
		}
	case 4:
		{
			t, ok := protos.protoDispatch.(IServerStopNtf)
			if !ok {
				return false
			}

			proto := &ServerStopNtf{}
			if !proto.Read(b) {
				fmt.Println("read ServerStopNtf fail, system error.")
				return false
			}

			t.OnServerStopNtf(proto)
		}
	case 5:
		{
			t, ok := protos.protoDispatch.(IServerConfigNtf)
			if !ok {
				return false
			}

			proto := &ServerConfigNtf{}
			if !proto.Read(b) {
				fmt.Println("read ServerConfigNtf fail, system error.")
				return false
			}

			t.OnServerConfigNtf(proto)
		}
	case 6:
		{
			t, ok := protos.protoDispatch.(IDBWrapperPkg)
			if !ok {
				return false
			}

			proto := &DBWrapperPkg{}
			if !proto.Read(b) {
				fmt.Println("read DBWrapperPkg fail, system error.")
				return false
			}

			t.OnDBWrapperPkg(proto)
		}
	case 7:
		{
			t, ok := protos.protoDispatch.(IGSWrapperPkg)
			if !ok {
				return false
			}

			proto := &GSWrapperPkg{}
			if !proto.Read(b) {
				fmt.Println("read GSWrapperPkg fail, system error.")
				return false
			}

			t.OnGSWrapperPkg(proto)
		}
	case 8:
		{
			t, ok := protos.protoDispatch.(ILoadAuctionObjectDataReq)
			if !ok {
				return false
			}

			proto := &LoadAuctionObjectDataReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadAuctionObjectDataReq fail, system error.")
				return false
			}

			t.OnLoadAuctionObjectDataReq(proto)
		}
	case 9:
		{
			t, ok := protos.protoDispatch.(ILoadAuctionObjectDataAck)
			if !ok {
				return false
			}

			proto := &LoadAuctionObjectDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadAuctionObjectDataAck fail, system error.")
				return false
			}

			t.OnLoadAuctionObjectDataAck(proto)
		}
	case 10:
		{
			t, ok := protos.protoDispatch.(IAddAuctionObjectDataNtf)
			if !ok {
				return false
			}

			proto := &AddAuctionObjectDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddAuctionObjectDataNtf fail, system error.")
				return false
			}

			t.OnAddAuctionObjectDataNtf(proto)
		}
	case 11:
		{
			t, ok := protos.protoDispatch.(IDelAuctionObjectDataNtf)
			if !ok {
				return false
			}

			proto := &DelAuctionObjectDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelAuctionObjectDataNtf fail, system error.")
				return false
			}

			t.OnDelAuctionObjectDataNtf(proto)
		}
	case 12:
		{
			t, ok := protos.protoDispatch.(IUpdateAuctionObjectDataNtf)
			if !ok {
				return false
			}

			proto := &UpdateAuctionObjectDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateAuctionObjectDataNtf fail, system error.")
				return false
			}

			t.OnUpdateAuctionObjectDataNtf(proto)
		}
	case 13:
		{
			t, ok := protos.protoDispatch.(ILoadAuctionCookieDataReq)
			if !ok {
				return false
			}

			proto := &LoadAuctionCookieDataReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadAuctionCookieDataReq fail, system error.")
				return false
			}

			t.OnLoadAuctionCookieDataReq(proto)
		}
	case 14:
		{
			t, ok := protos.protoDispatch.(ILoadAuctionCookieDataAck)
			if !ok {
				return false
			}

			proto := &LoadAuctionCookieDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadAuctionCookieDataAck fail, system error.")
				return false
			}

			t.OnLoadAuctionCookieDataAck(proto)
		}
	case 15:
		{
			t, ok := protos.protoDispatch.(IDuplicateAuctionCookieDataNtf)
			if !ok {
				return false
			}

			proto := &DuplicateAuctionCookieDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DuplicateAuctionCookieDataNtf fail, system error.")
				return false
			}

			t.OnDuplicateAuctionCookieDataNtf(proto)
		}
	case 16:
		{
			t, ok := protos.protoDispatch.(ILoadGuildDataReq)
			if !ok {
				return false
			}

			proto := &LoadGuildDataReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadGuildDataReq fail, system error.")
				return false
			}

			t.OnLoadGuildDataReq(proto)
		}
	case 17:
		{
			t, ok := protos.protoDispatch.(ILoadGuildDataAck)
			if !ok {
				return false
			}

			proto := &LoadGuildDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadGuildDataAck fail, system error.")
				return false
			}

			t.OnLoadGuildDataAck(proto)
		}
	case 18:
		{
			t, ok := protos.protoDispatch.(ILoadGuildMemberDataAck)
			if !ok {
				return false
			}

			proto := &LoadGuildMemberDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadGuildMemberDataAck fail, system error.")
				return false
			}

			t.OnLoadGuildMemberDataAck(proto)
		}
	case 19:
		{
			t, ok := protos.protoDispatch.(ILoadGuildApplicantDataAck)
			if !ok {
				return false
			}

			proto := &LoadGuildApplicantDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadGuildApplicantDataAck fail, system error.")
				return false
			}

			t.OnLoadGuildApplicantDataAck(proto)
		}
	case 20:
		{
			t, ok := protos.protoDispatch.(IAddGuildDataNtf)
			if !ok {
				return false
			}

			proto := &AddGuildDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddGuildDataNtf fail, system error.")
				return false
			}

			t.OnAddGuildDataNtf(proto)
		}
	case 21:
		{
			t, ok := protos.protoDispatch.(IDelGuildDataNtf)
			if !ok {
				return false
			}

			proto := &DelGuildDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelGuildDataNtf fail, system error.")
				return false
			}

			t.OnDelGuildDataNtf(proto)
		}
	case 22:
		{
			t, ok := protos.protoDispatch.(IUpdateGuildDataNtf)
			if !ok {
				return false
			}

			proto := &UpdateGuildDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateGuildDataNtf fail, system error.")
				return false
			}

			t.OnUpdateGuildDataNtf(proto)
		}
	case 23:
		{
			t, ok := protos.protoDispatch.(IAddGuildMemberDataNtf)
			if !ok {
				return false
			}

			proto := &AddGuildMemberDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddGuildMemberDataNtf fail, system error.")
				return false
			}

			t.OnAddGuildMemberDataNtf(proto)
		}
	case 24:
		{
			t, ok := protos.protoDispatch.(IDelGuildMemberDataNtf)
			if !ok {
				return false
			}

			proto := &DelGuildMemberDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelGuildMemberDataNtf fail, system error.")
				return false
			}

			t.OnDelGuildMemberDataNtf(proto)
		}
	case 25:
		{
			t, ok := protos.protoDispatch.(IUpdateGuildMemberDataNtf)
			if !ok {
				return false
			}

			proto := &UpdateGuildMemberDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateGuildMemberDataNtf fail, system error.")
				return false
			}

			t.OnUpdateGuildMemberDataNtf(proto)
		}
	case 26:
		{
			t, ok := protos.protoDispatch.(IAddGuildApplicantDataNtf)
			if !ok {
				return false
			}

			proto := &AddGuildApplicantDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddGuildApplicantDataNtf fail, system error.")
				return false
			}

			t.OnAddGuildApplicantDataNtf(proto)
		}
	case 27:
		{
			t, ok := protos.protoDispatch.(IDelGuildApplicantDataNtf)
			if !ok {
				return false
			}

			proto := &DelGuildApplicantDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelGuildApplicantDataNtf fail, system error.")
				return false
			}

			t.OnDelGuildApplicantDataNtf(proto)
		}
	case 28:
		{
			t, ok := protos.protoDispatch.(IUpdateGuildApplicantDataNtf)
			if !ok {
				return false
			}

			proto := &UpdateGuildApplicantDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateGuildApplicantDataNtf fail, system error.")
				return false
			}

			t.OnUpdateGuildApplicantDataNtf(proto)
		}
	case 29:
		{
			t, ok := protos.protoDispatch.(ILoadContactDataReq)
			if !ok {
				return false
			}

			proto := &LoadContactDataReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadContactDataReq fail, system error.")
				return false
			}

			t.OnLoadContactDataReq(proto)
		}
	case 30:
		{
			t, ok := protos.protoDispatch.(ILoadContactInfoDataAck)
			if !ok {
				return false
			}

			proto := &LoadContactInfoDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadContactInfoDataAck fail, system error.")
				return false
			}

			t.OnLoadContactInfoDataAck(proto)
		}
	case 31:
		{
			t, ok := protos.protoDispatch.(ILoadContactDataAck)
			if !ok {
				return false
			}

			proto := &LoadContactDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadContactDataAck fail, system error.")
				return false
			}

			t.OnLoadContactDataAck(proto)
		}
	case 32:
		{
			t, ok := protos.protoDispatch.(IAddContactDataNtf)
			if !ok {
				return false
			}

			proto := &AddContactDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddContactDataNtf fail, system error.")
				return false
			}

			t.OnAddContactDataNtf(proto)
		}
	case 33:
		{
			t, ok := protos.protoDispatch.(IDelContactDataNtf)
			if !ok {
				return false
			}

			proto := &DelContactDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelContactDataNtf fail, system error.")
				return false
			}

			t.OnDelContactDataNtf(proto)
		}
	case 34:
		{
			t, ok := protos.protoDispatch.(IUpdateContactDataNtf)
			if !ok {
				return false
			}

			proto := &UpdateContactDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateContactDataNtf fail, system error.")
				return false
			}

			t.OnUpdateContactDataNtf(proto)
		}
	case 35:
		{
			t, ok := protos.protoDispatch.(ILoadDBVarDataReq)
			if !ok {
				return false
			}

			proto := &LoadDBVarDataReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadDBVarDataReq fail, system error.")
				return false
			}

			t.OnLoadDBVarDataReq(proto)
		}
	case 36:
		{
			t, ok := protos.protoDispatch.(ILoadDBVarDataAck)
			if !ok {
				return false
			}

			proto := &LoadDBVarDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadDBVarDataAck fail, system error.")
				return false
			}

			t.OnLoadDBVarDataAck(proto)
		}
	case 37:
		{
			t, ok := protos.protoDispatch.(IAddDBVarDataNtf)
			if !ok {
				return false
			}

			proto := &AddDBVarDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddDBVarDataNtf fail, system error.")
				return false
			}

			t.OnAddDBVarDataNtf(proto)
		}
	case 38:
		{
			t, ok := protos.protoDispatch.(IDelDBVarDataNtf)
			if !ok {
				return false
			}

			proto := &DelDBVarDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelDBVarDataNtf fail, system error.")
				return false
			}

			t.OnDelDBVarDataNtf(proto)
		}
	case 39:
		{
			t, ok := protos.protoDispatch.(ILoadMailDataReq)
			if !ok {
				return false
			}

			proto := &LoadMailDataReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadMailDataReq fail, system error.")
				return false
			}

			t.OnLoadMailDataReq(proto)
		}
	case 40:
		{
			t, ok := protos.protoDispatch.(ILoadMailDataAck)
			if !ok {
				return false
			}

			proto := &LoadMailDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadMailDataAck fail, system error.")
				return false
			}

			t.OnLoadMailDataAck(proto)
		}
	case 41:
		{
			t, ok := protos.protoDispatch.(IAddMailDataNtf)
			if !ok {
				return false
			}

			proto := &AddMailDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddMailDataNtf fail, system error.")
				return false
			}

			t.OnAddMailDataNtf(proto)
		}
	case 42:
		{
			t, ok := protos.protoDispatch.(IDelMailDataNtf)
			if !ok {
				return false
			}

			proto := &DelMailDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelMailDataNtf fail, system error.")
				return false
			}

			t.OnDelMailDataNtf(proto)
		}
	case 43:
		{
			t, ok := protos.protoDispatch.(IUpdateMailDataNtf)
			if !ok {
				return false
			}

			proto := &UpdateMailDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateMailDataNtf fail, system error.")
				return false
			}

			t.OnUpdateMailDataNtf(proto)
		}
	case 44:
		{
			t, ok := protos.protoDispatch.(ILoadRanklistDataReq)
			if !ok {
				return false
			}

			proto := &LoadRanklistDataReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadRanklistDataReq fail, system error.")
				return false
			}

			t.OnLoadRanklistDataReq(proto)
		}
	case 45:
		{
			t, ok := protos.protoDispatch.(ILoadRanklistDataAck)
			if !ok {
				return false
			}

			proto := &LoadRanklistDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadRanklistDataAck fail, system error.")
				return false
			}

			t.OnLoadRanklistDataAck(proto)
		}
	case 46:
		{
			t, ok := protos.protoDispatch.(IAddRanklistDataNtf)
			if !ok {
				return false
			}

			proto := &AddRanklistDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddRanklistDataNtf fail, system error.")
				return false
			}

			t.OnAddRanklistDataNtf(proto)
		}
	case 47:
		{
			t, ok := protos.protoDispatch.(IDelRanklistDataNtf)
			if !ok {
				return false
			}

			proto := &DelRanklistDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelRanklistDataNtf fail, system error.")
				return false
			}

			t.OnDelRanklistDataNtf(proto)
		}
	case 48:
		{
			t, ok := protos.protoDispatch.(IUpdateRanklistDataNtf)
			if !ok {
				return false
			}

			proto := &UpdateRanklistDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateRanklistDataNtf fail, system error.")
				return false
			}

			t.OnUpdateRanklistDataNtf(proto)
		}
	case 49:
		{
			t, ok := protos.protoDispatch.(IBillInMockReq)
			if !ok {
				return false
			}

			proto := &BillInMockReq{}
			if !proto.Read(b) {
				fmt.Println("read BillInMockReq fail, system error.")
				return false
			}

			t.OnBillInMockReq(proto)
		}
	case 50:
		{
			t, ok := protos.protoDispatch.(IAddPasturePetDataNtf)
			if !ok {
				return false
			}

			proto := &AddPasturePetDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddPasturePetDataNtf fail, system error.")
				return false
			}

			t.OnAddPasturePetDataNtf(proto)
		}
	case 51:
		{
			t, ok := protos.protoDispatch.(ILoadPasturePetDataReq)
			if !ok {
				return false
			}

			proto := &LoadPasturePetDataReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadPasturePetDataReq fail, system error.")
				return false
			}

			t.OnLoadPasturePetDataReq(proto)
		}
	case 52:
		{
			t, ok := protos.protoDispatch.(ILoadPasturePetDataAck)
			if !ok {
				return false
			}

			proto := &LoadPasturePetDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadPasturePetDataAck fail, system error.")
				return false
			}

			t.OnLoadPasturePetDataAck(proto)
		}
	case 53:
		{
			t, ok := protos.protoDispatch.(IFregmentPkg)
			if !ok {
				return false
			}

			proto := &FregmentPkg{}
			if !proto.Read(b) {
				fmt.Println("read FregmentPkg fail, system error.")
				return false
			}

			t.OnFregmentPkg(proto)
		}
	case 54:
		{
			t, ok := protos.protoDispatch.(IDelPasturePetDataNtf)
			if !ok {
				return false
			}

			proto := &DelPasturePetDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelPasturePetDataNtf fail, system error.")
				return false
			}

			t.OnDelPasturePetDataNtf(proto)
		}
	case 55:
		{
			t, ok := protos.protoDispatch.(IUpdatePasturePetDataNtf)
			if !ok {
				return false
			}

			proto := &UpdatePasturePetDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdatePasturePetDataNtf fail, system error.")
				return false
			}

			t.OnUpdatePasturePetDataNtf(proto)
		}
	case 56:
		{
			t, ok := protos.protoDispatch.(ILoadContactInfoReq)
			if !ok {
				return false
			}

			proto := &LoadContactInfoReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadContactInfoReq fail, system error.")
				return false
			}

			t.OnLoadContactInfoReq(proto)
		}
	case 57:
		{
			t, ok := protos.protoDispatch.(ILoadLadderDataReq)
			if !ok {
				return false
			}

			proto := &LoadLadderDataReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadLadderDataReq fail, system error.")
				return false
			}

			t.OnLoadLadderDataReq(proto)
		}
	case 58:
		{
			t, ok := protos.protoDispatch.(ILoadLadderDataAck)
			if !ok {
				return false
			}

			proto := &LoadLadderDataAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadLadderDataAck fail, system error.")
				return false
			}

			t.OnLoadLadderDataAck(proto)
		}
	case 59:
		{
			t, ok := protos.protoDispatch.(IAddLadderDataNtf)
			if !ok {
				return false
			}

			proto := &AddLadderDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddLadderDataNtf fail, system error.")
				return false
			}

			t.OnAddLadderDataNtf(proto)
		}
	case 60:
		{
			t, ok := protos.protoDispatch.(IDelLadderDataNtf)
			if !ok {
				return false
			}

			proto := &DelLadderDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelLadderDataNtf fail, system error.")
				return false
			}

			t.OnDelLadderDataNtf(proto)
		}
	case 61:
		{
			t, ok := protos.protoDispatch.(IUpdateLadderDataNtf)
			if !ok {
				return false
			}

			proto := &UpdateLadderDataNtf{}
			if !proto.Read(b) {
				fmt.Println("read UpdateLadderDataNtf fail, system error.")
				return false
			}

			t.OnUpdateLadderDataNtf(proto)
		}
	case 62:
		{
			t, ok := protos.protoDispatch.(IUnloadMailDataReq)
			if !ok {
				return false
			}

			proto := &UnloadMailDataReq{}
			if !proto.Read(b) {
				fmt.Println("read UnloadMailDataReq fail, system error.")
				return false
			}

			t.OnUnloadMailDataReq(proto)
		}
	case 63:
		{
			t, ok := protos.protoDispatch.(ILoadOfflinePlayerReq)
			if !ok {
				return false
			}

			proto := &LoadOfflinePlayerReq{}
			if !proto.Read(b) {
				fmt.Println("read LoadOfflinePlayerReq fail, system error.")
				return false
			}

			t.OnLoadOfflinePlayerReq(proto)
		}
	case 64:
		{
			t, ok := protos.protoDispatch.(ILoadOfflinePlayerAck)
			if !ok {
				return false
			}

			proto := &LoadOfflinePlayerAck{}
			if !proto.Read(b) {
				fmt.Println("read LoadOfflinePlayerAck fail, system error.")
				return false
			}

			t.OnLoadOfflinePlayerAck(proto)
		}
	case 65:
		{
			t, ok := protos.protoDispatch.(ISyncViolateNtf)
			if !ok {
				return false
			}

			proto := &SyncViolateNtf{}
			if !proto.Read(b) {
				fmt.Println("read SyncViolateNtf fail, system error.")
				return false
			}

			t.OnSyncViolateNtf(proto)
		}
	case 66:
		{
			t, ok := protos.protoDispatch.(ISyncForbidTalkNtf)
			if !ok {
				return false
			}

			proto := &SyncForbidTalkNtf{}
			if !proto.Read(b) {
				fmt.Println("read SyncForbidTalkNtf fail, system error.")
				return false
			}

			t.OnSyncForbidTalkNtf(proto)
		}
	case 67:
		{
			t, ok := protos.protoDispatch.(IAddForbidTalkNtf)
			if !ok {
				return false
			}

			proto := &AddForbidTalkNtf{}
			if !proto.Read(b) {
				fmt.Println("read AddForbidTalkNtf fail, system error.")
				return false
			}

			t.OnAddForbidTalkNtf(proto)
		}
	case 68:
		{
			t, ok := protos.protoDispatch.(IDelForbidTalkNtf)
			if !ok {
				return false
			}

			proto := &DelForbidTalkNtf{}
			if !proto.Read(b) {
				fmt.Println("read DelForbidTalkNtf fail, system error.")
				return false
			}

			t.OnDelForbidTalkNtf(proto)
		}
	case 69:
		{
			t, ok := protos.protoDispatch.(ILogLevelNtf)
			if !ok {
				return false
			}

			proto := &LogLevelNtf{}
			if !proto.Read(b) {
				fmt.Println("read LogLevelNtf fail, system error.")
				return false
			}

			t.OnLogLevelNtf(proto)
		}
	case 70:
		{
			t, ok := protos.protoDispatch.(IBillQueryReq)
			if !ok {
				return false
			}

			proto := &BillQueryReq{}
			if !proto.Read(b) {
				fmt.Println("read BillQueryReq fail, system error.")
				return false
			}

			t.OnBillQueryReq(proto)
		}
	case 71:
		{
			t, ok := protos.protoDispatch.(IBillQueryAck)
			if !ok {
				return false
			}

			proto := &BillQueryAck{}
			if !proto.Read(b) {
				fmt.Println("read BillQueryAck fail, system error.")
				return false
			}

			t.OnBillQueryAck(proto)
		}
	default:
		{
			fmt.Println("illegal protocol, Mid =", mid, "Pid =", pid)
		}
	}

	return true
}
