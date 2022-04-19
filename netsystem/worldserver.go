package netsystem

import (
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"strconv"
)

type WorldServer struct {
	svrIndex int
	svrAttr  *configs.ListenAttr
	svrConn  networks.Connection
}

func NewWorldServer(index int, attr *configs.ListenAttr) Session {
	svr := &WorldServer{
		svrIndex: index,
		svrAttr:  attr,
	}
	return svr
}

func (svr *WorldServer) GetIndex() int {
	return svr.svrIndex
}

func (svr *WorldServer) GetLogicName() string {
	return configs.GetServerName(configs.ServerIdWs) + strconv.Itoa(svr.svrIndex)
}

func (svr *WorldServer) OnConnected() {

}

func (svr *WorldServer) OnClosed() {

}

func (svr *WorldServer) OnReceived(data []byte) {

}

func (svr *WorldServer) Disconnect() {
	if svr.svrConn == nil {
		return
	}

	svr.svrConn.Disconnect()
}
