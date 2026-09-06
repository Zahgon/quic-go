package simnet

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var ErrDeadlineExceeded = errors.New("deadline exceeded")

type PacketReceiver interface {
	RecvPacket(p Packet)
}

type Router interface {
	SendPacket(p Packet) error
	AddNode(addr net.Addr, receiver PacketReceiver)
}

type Packet struct {
	To   net.Addr
	From net.Addr

	Data []byte
}

type SimConn struct {
	mu              sync.Mutex
	closed          bool
	closedChan      chan struct{}
	deadlineUpdated chan struct{}

	packetsSent atomic.Uint64
	packetsRcvd atomic.Uint64
	bytesSent   atomic.Int64
	bytesRcvd   atomic.Int64

	router Router

	myAddr        *net.UDPAddr
	myLocalAddr   net.Addr
	packetsToRead chan Packet

	recvBackPressure bool

	readDeadline  time.Time
	writeDeadline time.Time
}

var _ net.PacketConn = &SimConn{}

func NewSimConn(addr *net.UDPAddr, rtr Router) *SimConn { _ = "STUB: not implemented"; return nil }

func NewBlockingSimConn(addr *net.UDPAddr, rtr Router) *SimConn {
	_ = "STUB: not implemented"
	return nil
}

func newSimConn(addr *net.UDPAddr, rtr Router, block bool) *SimConn {
	_ = "STUB: not implemented"
	return nil
}

type ConnStats struct {
	BytesSent   int
	BytesRcvd   int
	PacketsSent int
	PacketsRcvd int
}

func (c *SimConn) Stats() ConnStats { _ = "STUB: not implemented"; return *new(ConnStats) }

func (c *SimConn) SetReadBuffer(n int) error { _ = "STUB: not implemented"; return nil }

func (c *SimConn) SetWriteBuffer(n int) error { _ = "STUB: not implemented"; return nil }

func (c *SimConn) RecvPacket(p Packet) { _ = "STUB: not implemented"; return }

func (c *SimConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *SimConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func (c *SimConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *SimConn) UnicastAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *SimConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *SimConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *SimConn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *SimConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
