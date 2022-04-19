package netsystem

import (
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"strconv"
)

type ChatServer struct {
	svrIndex int
	svrAttr  *configs.ListenAttr
	svrConn  networks.Connection
}

func NewChatServer(index int, attr *configs.ListenAttr) Session {
	svr := &ChatServer{
		svrIndex: index,
		svrAttr:  attr,
	}
	return svr
}

func (svr *ChatServer) GetIndex() int {
	return svr.svrIndex
}

func (svr *ChatServer) GetLogicName() string {
	return configs.GetServerName(configs.ServerIdCt) + strconv.Itoa(svr.svrIndex)
}

func (svr *ChatServer) OnConnected() {

}

func (svr *ChatServer) OnClosed() {

}

func (svr *ChatServer) OnReceived(data []byte) {

}

func (svr *ChatServer) Disconnect() {
	if svr.svrConn == nil {
		return
	}

	svr.svrConn.Disconnect()
}
