package netsystem

import (
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"strconv"
)

type GameServer struct {
	svrIndex int
	svrAttr  *configs.ListenAttr
	svrConn  networks.Connection
}

func NewGameServer(index int, attr *configs.ListenAttr) Session {
	svr := &GameServer{
		svrIndex: index,
		svrAttr:  attr,
	}
	return svr
}

func (svr *GameServer) GetIndex() int {
	return svr.svrIndex
}

func (svr *GameServer) GetLogicName() string {
	return configs.GetServerName(configs.ServerIdGs) + strconv.Itoa(svr.svrIndex)
}

func (svr *GameServer) OnConnected() {

}

func (svr *GameServer) OnClosed() {

}

func (svr *GameServer) OnReceived(data []byte) {

}

func (svr *GameServer) Disconnect() {
	if svr.svrConn == nil {
		return
	}

	svr.svrConn.Disconnect()
}
