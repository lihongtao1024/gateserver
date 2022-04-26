package netsystem

import (
	"fmt"
	"gateserver/internal/machines"
	"gateserver/internal/networks"
	"gateserver/logsystem"
)

type Client struct {
	cliUid   uint32
	cliConn  networks.Connection
	cliState machines.Machine
}

func NewClient(conn networks.Connection) *Client {
	client := &Client{cliConn: conn}
	client.cliState = machines.NewMachine(client)
	client.cliState.SwitchState(&ClientIdleState{})
	return client
}

func (client *Client) GetUid() uint32 {
	return 0
}

func (client *Client) GetLogicName() string {
	return fmt.Sprintf("%p Uid:%d", client, client.cliUid)
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
	state.(SessionState).OnReceived(client, data)
}

func (client *Client) IsState(s int) bool {
	return client.cliState.IsState(s)
}

func (client *Client) SwitchState(state SessionState) {
	client.cliState.SwitchState(state)
}

func (client *Client) Send(data []byte) bool {
	return client.cliConn.Send(data)
}

func (client *Client) Disconnect() {
	if !client.cliState.IsState(ServerConnected) &&
		!client.cliState.IsState(ServerWorking) {
		return
	}

	client.cliConn.Disconnect()
}
