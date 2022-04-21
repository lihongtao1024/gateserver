package netsystem

import (
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"gateserver/logsystem"
	"strconv"
)

type WorldServer struct {
	svrIndex int
	svrAttr  *configs.ListenAttr
	svrConn  networks.Connection
}

func NewWorldServer(index int, attr *configs.ListenAttr) Server {
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
	logsystem.This.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
	)
}

func (svr *WorldServer) OnFatal(err error) {
	if svr.svrConn.IsDialFatal() {
		logsystem.This.Err(
			"on fatal [%s]: errmsg:'%s'.",
			svr.GetLogicName(),
			err.Error(),
		)
		svr.svrConn = nil
		return
	}

	logsystem.This.Err(
		"on fatal [%s]: local addr:%s, remote addr:%s, errmsg:'%s'.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
		err.Error(),
	)
}

func (svr *WorldServer) OnClosed() {
	logsystem.This.Inf(
		"on closed [%s]: local addr:%s, remote addr:%s.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
	)
	svr.svrConn = nil
}

func (svr *WorldServer) OnReceived(data []byte) {

}

func (svr *WorldServer) Connect() bool {
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

func (svr *WorldServer) IsConnected() bool {
	return svr.svrConn != nil
}

func (svr *WorldServer) Disconnect() {
	if svr.svrConn == nil {
		return
	}

	svr.svrConn.Disconnect()
}
