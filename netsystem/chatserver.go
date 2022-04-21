package netsystem

import (
	"gateserver/internal/configs"
	"gateserver/internal/networks"
	"gateserver/logsystem"
	"strconv"
)

type ChatServer struct {
	svrIndex int
	svrAttr  *configs.ListenAttr
	svrConn  networks.Connection
}

func NewChatServer(index int, attr *configs.ListenAttr) Server {
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
	logsystem.This.Inf(
		"on connected [%s]: local addr:%s, remote addr:%s.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
	)
}

func (svr *ChatServer) OnFatal(err error) {
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

func (svr *ChatServer) OnClosed() {
	logsystem.This.Inf(
		"on closed [%s]: local addr:%s, remote addr:%s.",
		svr.GetLogicName(),
		svr.svrConn.GetLocalAddr(),
		svr.svrConn.GetRemoteAddr(),
	)
	svr.svrConn = nil
}

func (svr *ChatServer) OnReceived(data []byte) {

}

func (svr *ChatServer) Connect() bool {
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

func (svr *ChatServer) IsConnected() bool {
	return svr.svrConn != nil
}

func (svr *ChatServer) Disconnect() {
	if svr.svrConn == nil {
		return
	}

	svr.svrConn.Disconnect()
}
