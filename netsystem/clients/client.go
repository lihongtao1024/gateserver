package clients

import (
	"encoding/binary"
	"fmt"
	"gateserver/internal/errors"
	"gateserver/internal/machines"
	"gateserver/internal/networks"
	"gateserver/internal/protocols"
	"gateserver/logsystem"
	"gateserver/netsystem/sessions"
	"gateserver/protosystem"
	"math/rand"
	"unsafe"
)

const (
	ClientIdle       = 0
	ClientConnected  = 1
	ClientWorking    = 2
	ClientVerifying  = 3
	ClientRequesting = 4
	ClientLoggedIn   = 5
	ClientPlaying    = 6
)

type Client struct {
	cliHeartBeat   bool
	cliAntiIndulge bool
	cliThirdUser   bool
	cliSid         uint8
	cliGSid        uint8
	cliUid         uint32
	cliZid         uint32
	cliAid         uint32
	cliIp          uint32
	cliSUid        uint64
	cliUname       string
	cliHardware    string
	cliLongitude   string
	cliLatitude    string
	cliPlatform    string
	cliConn        networks.Connection
	cliState       machines.Machine
	cliRandKey     []byte
	cliPassword    []byte
}

func NewClient(conn networks.Connection) *Client {
	client := &Client{cliConn: conn}
	client.cliState = machines.NewMachine(client)
	client.cliState.SwitchState(&ClientIdleState{})
	client.cliRandKey = make([]byte, unsafe.Sizeof(uint64(0)))
	binary.LittleEndian.PutUint64(client.cliRandKey, rand.Uint64())

	return client
}

func (client *Client) SetUName(name string) {
	client.cliUname = name
}

func (client *Client) GetUName() string {
	return client.cliUname
}

func (client *Client) SetUid(uid uint32) {
	client.cliUid = uid
}

func (client *Client) GetUid() uint32 {
	return client.cliUid
}

func (client *Client) SetHeartbeat(heartbeat bool) {
	client.cliHeartBeat = heartbeat
}

func (client *Client) GetHeartbeat() bool {
	return client.cliHeartBeat
}

func (client *Client) SetAntiIndulge(antiindulge bool) {
	client.cliAntiIndulge = antiindulge
}

func (client *Client) GetAntiIndulge() bool {
	return client.cliAntiIndulge
}

func (client *Client) SetThirdUser(thirduser bool) {
	client.cliThirdUser = thirduser
}

func (client *Client) GetThirdUser() bool {
	return client.cliThirdUser
}

func (client *Client) SetSid(sid uint8) {
	client.cliSid = sid
}

func (client *Client) GetSid() uint8 {
	return client.cliSid
}

func (client *Client) SetGSid(gsid uint8) {
	client.cliGSid = gsid
}

func (client *Client) GetGSid() uint8 {
	return client.cliGSid
}

func (client *Client) SetZid(zid uint32) {
	client.cliZid = zid
}

func (client *Client) GetZid() uint32 {
	return client.cliZid
}

func (client *Client) SetAid(aid uint32) {
	client.cliAid = aid
}

func (client *Client) GetAid() uint32 {
	return client.cliAid
}

func (client *Client) SetIp(ip uint32) {
	client.cliIp = ip
}

func (client *Client) GetIp() uint32 {
	return client.cliIp
}

func (client *Client) GetSUid() uint64 {
	return client.cliSUid
}

func (client *Client) SetLongitude(longitude string) {
	client.cliLongitude = longitude
}

func (client *Client) GetLongitude() string {
	return client.cliLongitude
}

func (client *Client) SetLatitude(latitude string) {
	client.cliLatitude = latitude
}

func (client *Client) GetLatitude() string {
	return client.cliLatitude
}

func (client *Client) SetPasssword(pwd []byte) {
	client.cliPassword = make([]byte, len(pwd))
	copy(client.cliPassword, pwd)
}

func (client *Client) GetPassword() []byte {
	return client.cliPassword
}

func (client *Client) SetHardware(hw string) {
	client.cliHardware = hw
}

func (client *Client) GetHardware() string {
	return client.cliHardware
}

func (client *Client) SetPlatform(platform string) {
	client.cliPlatform = platform
}

func (client *Client) GetPlatform() string {
	return client.cliPlatform
}

func (client *Client) GetLogicName() string {
	return fmt.Sprintf("%p Uname:%s,Uid:%d", client, client.cliUname, client.cliUid)
}

func (client *Client) GetRandKey() []byte {
	return client.cliRandKey
}

func (client *Client) OnConnected() {
	logsystem.Instance.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		client.GetLogicName(),
		client.cliConn.GetLocalAddr(),
		client.cliConn.GetRemoteAddr(),
	)

	client.cliState.SwitchState(&ClientConnectedState{})
}

func (client *Client) OnFatal(err error) {
	logsystem.Instance.Err(
		"on fatal [%s]: local addr:%s, remote addr:%s, errmsg:'%s'.",
		client.GetLogicName(),
		client.cliConn.GetLocalAddr(),
		client.cliConn.GetRemoteAddr(),
		err.Error(),
	)
}

func (client *Client) OnClosed() {
	logsystem.Instance.Inf(
		"on closed [%s]: local addr:%s, remote addr:%s.",
		client.GetLogicName(),
		client.cliConn.GetLocalAddr(),
		client.cliConn.GetRemoteAddr(),
	)

	client.cliConn = nil
	client.cliState.SwitchState(&ClientIdleState{})
}

func (client *Client) OnReceived(data []byte) {
	state := client.cliState.GetState()
	state.(sessions.SessionState).OnReceived(client, data)
}

func (client *Client) IsState(s int) bool {
	return client.cliState.IsState(s)
}

func (client *Client) SwitchState(state sessions.SessionState) {
	client.cliState.SwitchState(state)
}

func (client *Client) VerifyHandShakeReq(data []byte) error {
	return protosystem.Instance.VerifyClientHandShakeReq(data)
}

func (client *Client) SendHandShakeRsp() bool {
	data := protosystem.Instance.BuildClientHandShakeRsp()
	return client.Send(data)
}

func (client *Client) SendRandKey() bool {
	proto := &protocols.RandKeyNtf{}
	proto.Code_content = client.GetRandKey()

	result, data := protosystem.Instance.BuildClientProto(proto)
	if !result {
		return false
	}

	return client.Send(data)
}

func (client *Client) SendLoginAck(err errors.SysError) bool {
	proto := &protocols.LoginAck{}
	proto.Uid = client.cliUid
	proto.Guid = 0
	proto.Suid = 0
	proto.Rid = 0

	code := err.Code()
	if code != errors.ErrorOk {
		proto.Errcode = code
		proto.Errmsg = err.Error()
	}

	result, data := protosystem.Instance.BuildClientProto(proto)
	if !result {
		return false
	}

	return client.Send(data)
}

func (client *Client) Send(data []byte) bool {
	return client.cliConn.Send(data)
}

func (client *Client) Disconnect() {
	if client.cliConn == nil {
		return
	}

	client.cliConn.Disconnect()
}
