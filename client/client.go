package client

import (
	"encoding/binary"
	"fmt"
	"gateserver/component"
	"gateserver/pkg"
	"gateserver/pkg/errors"
	"gateserver/pkg/machines"
	"gateserver/pkg/protocols"
	"gateserver/singleton"
	"gateserver/state/clientstate"
	"math/rand"
	"strconv"
	"strings"
	"unsafe"
)

type clientImpl struct {
	cliHeartBeat   bool
	cliAntiIndulge bool
	cliThirdUser   bool
	cliSid         uint8
	cliGSid        uint8
	cliUid         uint32
	cliZid         uint32
	cliAid         uint32
	cliIp          uint32
	cliRealState   uint32
	cliGuid        pkg.Guid
	cliSuid        pkg.Guid
	cliUname       string
	cliHardware    string
	cliLongitude   string
	cliLatitude    string
	cliPlatform    string
	cliRealToken   string
	cliAuthToken   string
	cliRandKey     []byte
	cliPassword    []byte
	pkg.TcpConnection
	pkg.Machine
}

func NewClient(conn pkg.TcpConnection) component.Client {
	client := &clientImpl{TcpConnection: conn}
	client.Machine = machines.NewMachine(client)
	client.SwitchState(&clientstate.ClientIdleState{})
	client.cliRandKey = make([]byte, unsafe.Sizeof(uint64(0)))
	client.cliGuid = pkg.InvalidGuid
	client.cliSuid = singleton.GuidInstance.CreateGuid(pkg.GuidGlobal)
	client.cliAid = ^uint32(0)
	binary.LittleEndian.PutUint64(client.cliRandKey, rand.Uint64())

	return client
}

func (client *clientImpl) SetUName(name string) {
	client.cliUname = name
}

func (client *clientImpl) GetUName() string {
	return client.cliUname
}

func (client *clientImpl) SetUid(uid uint32) {
	client.cliUid = uid
}

func (client *clientImpl) GetUid() uint32 {
	return client.cliUid
}

func (client *clientImpl) SetHeartbeat(heartbeat bool) {
	client.cliHeartBeat = heartbeat
}

func (client *clientImpl) GetHeartbeat() bool {
	return client.cliHeartBeat
}

func (client *clientImpl) SetAntiIndulge(antiindulge bool) {
	client.cliAntiIndulge = antiindulge
}

func (client *clientImpl) GetAntiIndulge() bool {
	return client.cliAntiIndulge
}

func (client *clientImpl) SetThirdUser(thirduser bool) {
	client.cliThirdUser = thirduser
}

func (client *clientImpl) GetThirdUser() bool {
	return client.cliThirdUser
}

func (client *clientImpl) SetSid(sid uint8) {
	client.cliSid = sid
}

func (client *clientImpl) GetSid() uint8 {
	return client.cliSid
}

func (client *clientImpl) SetGSid(gsid uint8) {
	client.cliGSid = gsid
}

func (client *clientImpl) GetGSid() uint8 {
	return client.cliGSid
}

func (client *clientImpl) SetZid(zid uint32) {
	client.cliZid = zid
}

func (client *clientImpl) GetZid() uint32 {
	return client.cliZid
}

func (client *clientImpl) SetAid(aid uint32) {
	client.cliAid = aid
}

func (client *clientImpl) GetAid() uint32 {
	return client.cliAid
}

func (client *clientImpl) SetIp(ip uint32) {
	client.cliIp = ip
}

func (client *clientImpl) GetIp() uint32 {
	return client.cliIp
}

func (client *clientImpl) SetRealState(state uint32) {
	client.cliRealState = state
}

func (client *clientImpl) GetRealState() uint32 {
	return client.cliRealState
}

func (client *clientImpl) GetSuid() pkg.Guid {
	return client.cliSuid
}

func (client *clientImpl) SetGuid(guid pkg.Guid) {
	client.cliGuid = guid
}

func (client *clientImpl) GetGuid() pkg.Guid {
	return client.cliGuid
}

func (client *clientImpl) SetLongitude(longitude string) {
	client.cliLongitude = longitude
}

func (client *clientImpl) GetLongitude() string {
	return client.cliLongitude
}

func (client *clientImpl) SetLatitude(latitude string) {
	client.cliLatitude = latitude
}

func (client *clientImpl) GetLatitude() string {
	return client.cliLatitude
}

func (client *clientImpl) SetPasssword(pwd []byte) {
	client.cliPassword = make([]byte, len(pwd))
	copy(client.cliPassword, pwd)
}

func (client *clientImpl) GetPassword() []byte {
	return client.cliPassword
}

func (client *clientImpl) SetHardware(hw string) {
	client.cliHardware = hw
}

func (client *clientImpl) GetHardware() string {
	return client.cliHardware
}

func (client *clientImpl) SetPlatform(platform string) {
	client.cliPlatform = platform
}

func (client *clientImpl) GetPlatform() string {
	return client.cliPlatform
}

func (client *clientImpl) SetRealToken(token string) {
	client.cliRealToken = token
}

func (client *clientImpl) GetRealToken() string {
	return client.cliRealToken
}

func (client *clientImpl) SetAuthToken(token string) {
	client.cliAuthToken = token
}

func (client *clientImpl) GetAuthToken() string {
	return client.cliAuthToken
}

func (client *clientImpl) GetLogicName() string {
	return fmt.Sprintf(
		"%p Uname:%s,Uid:%d",
		client,
		client.cliUname,
		client.cliUid,
	)
}

func (client *clientImpl) GetRandKey() []byte {
	return client.cliRandKey
}

func (client *clientImpl) OnConnected() {
	singleton.LogInstance.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		client.GetLogicName(),
		client.GetLocalAddr(),
		client.GetRemoteAddr(),
	)

	client.SwitchState(&clientstate.ClientConnectedState{})
}

func (client *clientImpl) OnFatal(err error) {
	singleton.LogInstance.Err(
		"on fatal [%s]: local addr:%s, remote addr:%s, errmsg:'%s'.",
		client.GetLogicName(),
		client.GetLocalAddr(),
		client.GetRemoteAddr(),
		err.Error(),
	)
}

func (client *clientImpl) OnClosed() {
	singleton.LogInstance.Inf(
		"on closed [%s]: local addr:%s, remote addr:%s.",
		client.GetLogicName(),
		client.GetLocalAddr(),
		client.GetRemoteAddr(),
	)

	client.TcpConnection = nil
	client.SwitchState(&clientstate.ClientIdleState{})
}

func (client *clientImpl) OnReceived(data []byte) {
	state := client.GetState()
	state.(component.SessionState).OnReceived(client, data)
}

func (client *clientImpl) VerifyClientHandShakeReq(data []byte) error {
	return singleton.ProtoInstance.VerifyClientHandShakeReq(data)
}

func (client *clientImpl) SendClientHandShakeRsp() bool {
	return client.Send(singleton.ProtoInstance.BuildClientHandShakeRsp())
}

func (client *clientImpl) SendRandKey() bool {
	proto := &protocols.RandKeyNtf{}
	proto.Code_content = client.cliRandKey

	return client.SendClientProto(proto)
}

func (client *clientImpl) SendLoginReq() bool {
	ws := singleton.NetInstance.GetWSServer()
	if ws == nil {
		return false
	}

	proto := &protocols.LoginReq{}
	proto.Sid = client.cliSid
	proto.Username = client.cliUname
	proto.Ip = func(addr string) uint32 {
		ip := strings.Split(addr, ":")[0]
		b := strings.Split(ip, ".")
		b0, _ := strconv.Atoi(b[0])
		b1, _ := strconv.Atoi(b[1])
		b2, _ := strconv.Atoi(b[2])
		b3, _ := strconv.Atoi(b[3])
		return uint32(b0)<<24 | uint32(b1)<<16 | uint32(b2)<<8 | uint32(b3)
	}(client.TcpConnection.GetRemoteAddr())
	proto.Pwd_content = client.cliPassword
	proto.Hwid = client.cliHardware
	proto.Client_type = uint8(singleton.AppInstance.GetClientType())
	proto.Longitude = client.cliLongitude
	proto.Latitude = client.cliLatitude
	proto.Session = uint64(uintptr(unsafe.Pointer(&client.TcpConnection)))
	proto.Suid = uint64(client.cliSuid)

	return client.sendServerProto(ws, proto)
}

func (client *clientImpl) SendLoginAck(errcode pkg.ErrorCode,
	rid ...pkg.Guid) bool {
	proto := &protocols.LoginAck{}
	proto.Uid = client.cliUid
	proto.Guid = uint64(client.cliGuid)
	proto.Suid = uint64(client.cliSuid)

	if rid != nil {
		proto.Rid = uint64(rid[0])
	}

	if errcode != pkg.ErrorOk {
		err := errors.NewError(errcode)
		proto.Errcode = err.Code()
		proto.Errmsg = err.Error()
	}

	return client.SendClientProto(proto)
}

func (client *clientImpl) SendKickNtf(errcode pkg.ErrorCode) bool {
	proto := &protocols.KickNtf{}
	if errcode != pkg.ErrorOk {
		err := errors.NewError(errcode)
		proto.Errcode = err.Code()
		proto.Errmsg = err.Error()
	}

	return client.SendClientProto(proto)
}

func (client *clientImpl) SendRealNtf() bool {
	proto := &protocols.RealnameInfoNtf{}
	proto.Realname_status = client.cliRealState
	proto.Auth = client.cliAuthToken
	proto.Realname_token = client.cliRealToken

	return client.SendClientProto(proto)
}

func (client *clientImpl) SendClientProto(proto pkg.WriterProto) bool {
	result, data := singleton.ProtoInstance.BuildClientProto(proto)
	if !result {
		return false
	}

	return client.Send(data)
}

func (client *clientImpl) SendClientData(data []byte) bool {
	result, data := singleton.ProtoInstance.BuildClientData(data)
	if !result {
		return false
	}

	return client.Send(data)
}

func (client *clientImpl) sendServerProto(server component.Server,
	proto pkg.WriterProto) bool {
	if server == nil {
		return false
	}

	result, data := singleton.ProtoInstance.BuildServerProto(client, proto)
	if !result {
		return false
	}

	return server.(component.Session).Send(data)
}

func (client *clientImpl) SendServerData(typ1 pkg.ServerType,
	data []byte) bool {
	var server component.Server

	switch typ1 {
	case pkg.ServerIdWs:
		server = singleton.NetInstance.GetWSServer()
	case pkg.ServerIdCt:
		server = singleton.NetInstance.GetCTServer()
	case pkg.ServerIdGs:
		server = singleton.NetInstance.GetGSServer(int(client.cliGSid))
	default:
	}

	if server == nil {
		return false
	}

	result, data := singleton.ProtoInstance.BuildServerData(client, data)
	if !result {
		return false
	}

	return server.(component.Session).Send(data)
}

func (client *clientImpl) Disconnect() {
	if client.TcpConnection == nil {
		return
	}

	client.TcpConnection.Disconnect()
}
