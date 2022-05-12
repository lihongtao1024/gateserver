package networks

import (
	"gateserver/pkg"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	comConnected = 1
	comFatal     = 2
	comClosed    = 3
	comMessage   = 4
)

const (
	comAcceptTimeout = time.Millisecond
)

type comEvent interface {
	getType() int
}

type comConnectedEvent struct {
	listener *net.TCPListener
	conn     *tcpConnection
}

func (evt *comConnectedEvent) getType() int {
	return comConnected
}

type comFatalEvent struct {
	err  error
	conn *tcpConnection
}

func (evt *comFatalEvent) getType() int {
	return comFatal
}

type comClosedEvent struct {
	conn *tcpConnection
}

func (evt *comClosedEvent) getType() int {
	return comClosed
}

type comMessageEvent struct {
	data []byte
	conn *tcpConnection
}

func (evt *comMessageEvent) getType() int {
	return comMessage
}

type tcpComponent struct {
	comFragment pkg.TcpFragment
	comDispatch pkg.TcpDispatcher
	waitGroup   *sync.WaitGroup
	finishChan  chan struct{}
	eventChan   chan comEvent
}

func NewComponent(dispatch pkg.TcpDispatcher,
	fragment pkg.TcpFragment) pkg.TcpComponent {
	comp := &tcpComponent{
		comFragment: fragment,
		comDispatch: dispatch,
		waitGroup:   &sync.WaitGroup{},
		finishChan:  make(chan struct{}),
		eventChan:   make(chan comEvent, 409600),
	}

	return comp
}

func (comp *tcpComponent) onAccpeted(listener *net.TCPListener) {
	defer func() {
		listener.Close()
		comp.waitGroup.Done()
	}()

	for {
		select {
		case <-comp.finishChan:
			return
		default:
		}

		err := listener.SetDeadline(time.Now().Add(comAcceptTimeout))
		if err != nil {
			continue
		}

		var conn *net.TCPConn
		conn, err = listener.AcceptTCP()
		if err != nil {
			continue
		}

		NewConnection(comp, conn).(*tcpConnection).onAccpeted(listener)
	}
}

func (comp *tcpComponent) postConnected(listener *net.TCPListener, conn *tcpConnection) {
	evt := &comConnectedEvent{
		listener: listener,
		conn:     conn,
	}
	comp.eventChan <- evt
}

func (comp *tcpComponent) postRecived(conn *tcpConnection, data []byte, len int) {
	slice := make([]byte, len)
	copy(slice, data[:len])

	evt := &comMessageEvent{
		data: slice,
		conn: conn,
	}

	comp.eventChan <- evt
}

func (comp *tcpComponent) postFatal(conn *tcpConnection, err error) {
	evt := &comFatalEvent{
		err:  err,
		conn: conn,
	}
	comp.eventChan <- evt
}

func (comp *tcpComponent) postClosed(conn *tcpConnection) {
	evt := &comClosedEvent{
		conn: conn,
	}
	comp.eventChan <- evt
}

func (comp *tcpComponent) Listen(ip string, port uint16) pkg.TcpListener {
	addr, err := net.ResolveTCPAddr("tcp4", ip+":"+strconv.Itoa(int(port)))
	if err != nil {
		return nil
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil
	}

	comp.waitGroup.Add(1)
	go comp.onAccpeted(listener)

	return listener
}

func (comp *tcpComponent) Connect(ip string, port uint16) pkg.TcpConnection {
	connection := NewConnection(comp, nil)
	comp.waitGroup.Add(1)

	go func() {
		defer comp.waitGroup.Done()

		addr, err := net.ResolveTCPAddr("tcp4", ip+":"+strconv.Itoa(int(port)))
		if err != nil {
			comp.postFatal(connection.(*tcpConnection), err)
			return
		}

		conn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			comp.postFatal(connection.(*tcpConnection), err)
			return
		}

		connection.(*tcpConnection).onConnected(conn)
	}()

	return connection
}

func (comp *tcpComponent) Do() bool {
	busy := false

	select {
	case event := <-comp.eventChan:
		{
			switch event.getType() {
			case comConnected:
				{
					evt := event.(*comConnectedEvent)
					comp.comDispatch.OnConnected(evt.listener, evt.conn)
				}
			case comFatal:
				{
					evt := event.(*comFatalEvent)
					comp.comDispatch.OnFatal(evt.err, evt.conn)
				}
			case comClosed:
				{
					conn := event.(*comClosedEvent).conn
					conn.close()
					comp.comDispatch.OnClosed(conn)
				}
			default:
				{
					evt := event.(*comMessageEvent)
					comp.comDispatch.OnReceived(evt.data, evt.conn)
				}
			}
			busy = true
		}
	default:
	}
	return busy
}

func (comp *tcpComponent) Close() {
	close(comp.finishChan)
	comp.waitGroup.Wait()
}
