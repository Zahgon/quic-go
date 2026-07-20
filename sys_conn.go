package quic

import (
	"io"
	"net"
	"syscall"
	"time"

	"github.com/quic-go/quic-go/internal/protocol"
)

type connCapabilities struct {
	DF bool

	GSO bool

	ECN bool
}

type rawConn interface {
	ReadPacket() (receivedPacket, error)

	WritePacket(b []byte, addr net.Addr, packetInfoOOB []byte, gsoSize uint16, ecn protocol.ECN) (int, error)
	LocalAddr() net.Addr
	SetReadDeadline(time.Time) error
	io.Closer

	capabilities() connCapabilities
}

type OOBCapablePacketConn interface {
	net.PacketConn
	SyscallConn() (syscall.RawConn, error)
	SetReadBuffer(int) error
	ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *net.UDPAddr, err error)
	WriteMsgUDP(b, oob []byte, addr *net.UDPAddr) (n, oobn int, err error)
}

var _ OOBCapablePacketConn = &net.UDPConn{}

func wrapConn(pc net.PacketConn) (rawConn, error) {
	_ = "STUB: not implemented"
	return *new(rawConn), nil
}

type basicConn struct {
	net.PacketConn
	supportsDF bool
}

var _ rawConn = &basicConn{}

func (c *basicConn) ReadPacket() (receivedPacket, error) {
	_ = "STUB: not implemented"
	return *new(receivedPacket), nil
}

func (c *basicConn) WritePacket(b []byte, addr net.Addr, _ []byte, gsoSize uint16, ecn protocol.ECN) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *basicConn) capabilities() connCapabilities {
	_ = "STUB: not implemented"
	return *new(connCapabilities)
}
