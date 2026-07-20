package quic

import (
	"net"
	"sync/atomic"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
)

type sendConn interface {
	Write(b []byte, gsoSize uint16, ecn protocol.ECN) error
	WriteTo([]byte, net.Addr, packetInfo) error
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	ChangeRemoteAddr(addr net.Addr, info packetInfo)

	capabilities() connCapabilities
}

type remoteAddrInfo struct {
	addr net.Addr
	oob  []byte
}

type sconn struct {
	rawConn

	localAddr net.Addr

	remoteAddrInfo atomic.Pointer[remoteAddrInfo]

	logger utils.Logger

	gotGSOError bool

	wroteFirstPacket bool
}

var _ sendConn = &sconn{}

func newSendConn(c rawConn, remote net.Addr, info packetInfo, logger utils.Logger) *sconn {
	_ = "STUB: not implemented"
	return nil
}

func (c *sconn) Write(p []byte, gsoSize uint16, ecn protocol.ECN) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *sconn) writePacket(p []byte, addr net.Addr, oob []byte, gsoSize uint16, ecn protocol.ECN) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *sconn) WriteTo(b []byte, addr net.Addr, info packetInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *sconn) capabilities() connCapabilities {
	_ = "STUB: not implemented"
	return *new(connCapabilities)
}

func (c *sconn) ChangeRemoteAddr(addr net.Addr, info packetInfo) { _ = "STUB: not implemented"; return }

func (c *sconn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
func (c *sconn) LocalAddr() net.Addr  { _ = "STUB: not implemented"; return *new(net.Addr) }
