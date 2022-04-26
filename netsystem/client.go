package netsystem

import (
	"gateserver/internal/machines"
	"gateserver/internal/networks"
	"gateserver/logsystem"
)

type Client struct {
	cliConn  networks.Connection
	cliState machines.Machine
}

func NewClient(conn networks.Connection) Session {
	client := &Client{cliConn: conn}
	client.cliState = machines.NewMachine(client)
	return client
}

func (client *Client) GetIndex() int {
	return 0
}

func (client *Client) GetLogicName() string {
	return ""
}

func (client *Client) OnConnected() {
	logsystem.TheLog.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		client.GetLogicName(),
		client.cliConn.GetLocalAddr(),
		client.cliConn.GetRemoteAddr(),
	)
}

func (client *Client) OnFatal(err error) {
	logsystem.TheLog.Err(
		"on fatal [%s]: local addr:%s, remote addr:%s, errmsg:'%s'.",
		client.GetLogicName(),
		client.cliConn.GetLocalAddr(),
		client.cliConn.GetRemoteAddr(),
		err.Error(),
	)
}

func (client *Client) OnClosed() {
	logsystem.TheLog.Inf(
		"on closed [%s]: local addr:%s, remote addr:%s.",
		client.GetLogicName(),
		client.cliConn.GetLocalAddr(),
		client.cliConn.GetRemoteAddr(),
	)
	client.cliConn = nil
}

func (client *Client) OnReceived(data []byte) {

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
