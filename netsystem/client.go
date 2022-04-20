package netsystem

import (
	"gateserver/internal/networks"
	"gateserver/logsystem"
)

type Client struct {
	cliConn networks.Connection
}

func NewClient(conn networks.Connection) Session {
	client := &Client{conn}
	return client
}

func (client *Client) GetIndex() int {
	return 0
}

func (client *Client) GetLogicName() string {
	return ""
}

func (client *Client) OnConnected() {
	logsystem.This.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		client.GetLogicName(),
		client.cliConn.GetLocalAddr(),
		client.cliConn.GetRemoteAddr(),
	)
}

func (client *Client) OnFatal(err error) {
	logsystem.This.Err(
		"on fatal [%s]: local addr:%s, remote addr:%s, errmsg:'%s'.",
		client.GetLogicName(),
		client.cliConn.GetLocalAddr(),
		client.cliConn.GetRemoteAddr(),
		err.Error(),
	)
}

func (client *Client) OnClosed() {
	logsystem.This.Inf(
		"on closed [%s]: local addr:%s, remote addr:%s.",
		client.GetLogicName(),
		client.cliConn.GetLocalAddr(),
		client.cliConn.GetRemoteAddr(),
	)
	client.cliConn = nil
}

func (client *Client) OnReceived(data []byte) {

}

func (client *Client) Disconnect() {
	if client.cliConn == nil {
		return
	}

	client.cliConn.Disconnect()
}
