package netsystem

import (
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"gateserver/logsystem"
	"strconv"
)

type GameServer struct {
	svrIndex int
	svrAttr  *configs.ListenAttr
	svrConn  networks.Connection
}

func NewGameServer(index int, attr *configs.ListenAttr) Server {
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
	logsystem.This.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
	)
}

func (svr *GameServer) OnFatal(err error) {
	logsystem.This.Err(
		"on fatal [%s]: local addr:%s, remote addr:%s, errmsg:'%s'.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
		err.Error(),
	)
}

func (svr *GameServer) OnClosed() {
	logsystem.This.Inf(
		"on closed [%s]: local addr:%s, remote addr:%s.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
	)
	svr.svrConn = nil
}

func (svr *GameServer) OnReceived(data []byte) {

}

func (svr *GameServer) Connect() bool {
	svr.svrConn = This.netComponent.Connect(svr.svrAttr.Ip, uint16(svr.svrAttr.Port))
	if svr.svrConn == nil {
		logsystem.This.Err(
			"connect to [%s]: [%s:%d] [fail].",
			svr.GetLogicName(),
			svr.svrAttr.Ip,
			svr.svrAttr.Port,
		)
		return false
	}

	svr.svrConn.SetData(svr)
	logsystem.This.Inf(
		"connect to [%s]: [%s:%d] [wait].",
		svr.GetLogicName(),
		svr.svrAttr.Ip,
		svr.svrAttr.Port,
	)
	return true
}

func (svr *GameServer) IsConnected() bool {
	return svr.svrConn != nil
}

func (svr *GameServer) Disconnect() {
	if svr.svrConn == nil {
		return
	}

	svr.svrConn.Disconnect()
}
