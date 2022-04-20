package networks

import (
	"bytes"
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"
)

type Connection interface {
	SetData(data interface{})
	GetData() interface{}
	GetLocalAddr() string
	GetRemoteAddr() string
	Send(data []byte, len int) bool
	IsDialFatal() bool
	IsDisconnected() bool
	Disconnect()
}

const (
	connSend  = 1
	connFatal = 2
	connClose = 3
)

const (
	connConnected = 0
	connClosing   = 1
	connClosed    = 2
)

type connEvent interface {
	getType() int
}

type connSendEvent struct {
	data []byte
}

func (conn *connSendEvent) getType() int {
	return connSend
}

type connFatalEvent struct {
	err error
}

func (conn *connFatalEvent) getType() int {
	return connFatal
}

type connCloseEvent struct {
}

func (conn *connCloseEvent) getType() int {
	return connClose
}

type tcpConnection struct {
	netStatus  int32
	netComp    *tcpComponent
	netConn    *net.TCPConn
	fatalOnce  *sync.Once
	closeOnce  *sync.Once
	eventChan  chan connEvent
	customData interface{}
}

func NewConnection(comp *tcpComponent, conn *net.TCPConn) Connection {
	connection := &tcpConnection{
		netStatus: connConnected,
		netComp:   comp,
		netConn:   conn,
		fatalOnce: &sync.Once{},
		closeOnce: &sync.Once{},
		eventChan: make(chan connEvent, 4096),
	}
	return connection
}

func (conn *tcpConnection) onAccpeted(listener *net.TCPListener) {
	conn.netComp.postConnected(listener, conn)
	conn.asyncDo(conn.readBytes)
	conn.asyncDo(conn.writeBytes)
}

func (conn *tcpConnection) onConnected(connection *net.TCPConn) {
	conn.netConn = connection
	conn.netComp.postConnected(nil, conn)
	conn.asyncDo(conn.readBytes)
	conn.asyncDo(conn.writeBytes)
}

func (conn *tcpConnection) readBytes() {
	buf := make([]byte, 4096)
	recvbuf := bytes.NewBuffer([]byte{})

	for {
		if conn.isClosed() {
			return
		}

		len, err := conn.netConn.Read(buf)
		if err != nil {
			if err != io.EOF {
				conn.postFatal(err)
			}
			conn.Disconnect()
			break
		}

		if len == 0 {
			conn.Disconnect()
			break
		}

		recvbuf.Write(buf[:len])

		for recvbuf.Len() > 0 {
			fragment := recvbuf.Bytes()
			length := conn.netComp.comFragment.OnUnpack(fragment)
			if length == FragmentContinue {
				break
			} else if length == FragmentFatal {
				conn.postFatal(errors.New("unpack data fatal"))
				conn.Disconnect()
				break
			}

			conn.netComp.postRecived(conn, fragment, length)

			recvbuf.Reset()
			recvbuf.Write(fragment[length:])
			if recvbuf.Len() < 4096 && recvbuf.Cap() > 4096 {
				fragment = make([]byte, recvbuf.Len())
				copy(fragment, recvbuf.Bytes())
				recvbuf = bytes.NewBuffer(fragment)
			}
		}
	}
}

func (conn *tcpConnection) writeBytes() {
	var event connEvent
	for {
		if conn.isClosed() {
			return
		}

		event = <-conn.eventChan
		switch event.getType() {
		case connFatal:
			{
				conn.netComp.postFatal(conn, event.(*connFatalEvent).err)
			}
		case connClose:
			{
				conn.setClosed()
				conn.netComp.postClosed(conn)
			}
		default:
			{
				data := event.(*connSendEvent).data
				pos, len := 0, len(data)

				for pos < len {
					bytes, err := conn.netConn.Write(data[pos:])
					if err != nil {
						conn.postFatal(err)
						conn.Disconnect()
						break
					}

					pos += bytes
				}
			}
		}
	}
}

func (conn *tcpConnection) postSend(data []byte, len int) {
	slice := make([]byte, len)
	copy(slice, data[:len])

	evt := &connSendEvent{
		data: slice,
	}

	conn.eventChan <- evt
}

func (conn *tcpConnection) postFatal(err error) {
	conn.fatalOnce.Do(func() {
		evt := &connFatalEvent{
			err: err,
		}

		conn.eventChan <- evt
	})
}

func (conn *tcpConnection) postClosing() {
	evt := &connCloseEvent{}
	conn.eventChan <- evt
}

func (conn *tcpConnection) asyncDo(fn func()) {
	conn.netComp.waitGroup.Add(1)
	go func() {
		fn()
		conn.netComp.waitGroup.Done()
	}()
}

func (conn *tcpConnection) setClosed() {
	atomic.StoreInt32(&conn.netStatus, connClosed)
	conn.netConn.Close()
	close(conn.eventChan)
}

func (conn *tcpConnection) isClosed() bool {
	return atomic.LoadInt32(&conn.netStatus) == connClosed
}

func (conn *tcpConnection) SetData(data interface{}) {
	conn.customData = data
}

func (conn *tcpConnection) GetData() interface{} {
	return conn.customData
}

func (conn *tcpConnection) GetLocalAddr() string {
	return conn.netConn.LocalAddr().String()
}

func (conn *tcpConnection) GetRemoteAddr() string {
	return conn.netConn.RemoteAddr().String()
}

func (conn *tcpConnection) Send(data []byte, len int) bool {
	if conn.IsDisconnected() {
		return false
	}

	conn.postSend(data, len)
	return true
}

func (conn *tcpConnection) IsDialFatal() bool {
	return conn.netConn == nil
}

func (conn *tcpConnection) IsDisconnected() bool {
	return atomic.LoadInt32(&conn.netStatus) > connConnected
}

func (conn *tcpConnection) Disconnect() {
	conn.closeOnce.Do(func() {
		atomic.StoreInt32(&conn.netStatus, connClosing)
		conn.postClosing()
	})
}
