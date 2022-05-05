package clients

import (
	"encoding/binary"
	"fmt"
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
	cliUid     uint32
	cliConn    networks.Connection
	cliState   machines.Machine
	cliRandKey []byte
}

func NewClient(conn networks.Connection) *Client {
	client := &Client{cliConn: conn}
	client.cliState = machines.NewMachine(client)
	client.cliState.SwitchState(&ClientIdleState{})
	client.cliRandKey = make([]byte, unsafe.Sizeof(uint64(0)))
	binary.LittleEndian.PutUint64(client.cliRandKey, rand.Uint64())

	return client
}

func (client *Client) GetUid() uint32 {
	return 0
}

func (client *Client) GetLogicName() string {
	return fmt.Sprintf("%p Uid:%d", client, client.cliUid)
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

	result, data := protosystem.Instance.BuildRawProto(proto)
	if !result {
		return false
	}

	return client.Send(data)
}

func (client *Client) Send(data []byte) bool {
	return client.cliConn.Send(data)
}

func (client *Client) Disconnect() {
	if client.cliState.IsState(ClientIdle) {
		return
	}

	client.cliConn.Disconnect()
}
