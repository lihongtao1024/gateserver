package netsystem

import (
	"fmt"
	"gateserver/internal/networks"
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
	fmt.Println("client OnConnected")
}

func (client *Client) OnClosed() {
	fmt.Println("client OnClosed")
}

func (client *Client) OnReceived(data []byte) {
	fmt.Println("client OnReceived")
}

func (client *Client) Disconnect() {
	if client.cliConn == nil {
		return
	}

	client.cliConn.Disconnect()
}
