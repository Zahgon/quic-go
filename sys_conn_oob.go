//go:build darwin || linux || freebsd

package quic

import (
	"net"
	"net/netip"
	"sync"
	"syscall"

	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"

	"github.com/quic-go/quic-go/internal/protocol"
)

const (
	ecnMask       = 0x3
	oobBufferSize = 128
)

var _ ipv4.Message = ipv6.Message{}

type batchConn interface {
	ReadBatch(ms []ipv4.Message, flags int) (int, error)
}

func inspectReadBuffer(c syscall.RawConn) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func inspectWriteBuffer(c syscall.RawConn) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func isECNDisabledUsingEnv() bool { _ = "STUB: not implemented"; return false }

type oobConn struct {
	OOBCapablePacketConn
	batchConn batchConn

	readPos uint8

	messages []ipv4.Message
	buffers  [batchSize]*packetBuffer

	cap connCapabilities
}

var _ rawConn = &oobConn{}

func newConn(c OOBCapablePacketConn, supportsDF bool) (*oobConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var invalidCmsgOnceV4, invalidCmsgOnceV6 sync.Once

func (c *oobConn) ReadPacket() (receivedPacket, error) {
	_ = "STUB: not implemented"
	return *new(receivedPacket), nil
}

func (c *oobConn) WritePacket(b []byte, addr net.Addr, packetInfoOOB []byte, gsoSize uint16, ecn protocol.ECN) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *oobConn) capabilities() connCapabilities {
	_ = "STUB: not implemented"
	return *new(connCapabilities)
}

type packetInfo struct {
	addr    netip.Addr
	ifIndex uint32
}

func (info *packetInfo) OOB() []byte { _ = "STUB: not implemented"; return nil }

func appendIPv4ECNMsg(b []byte, val protocol.ECN) []byte { _ = "STUB: not implemented"; return nil }

func appendIPv6ECNMsg(b []byte, val protocol.ECN) []byte { _ = "STUB: not implemented"; return nil }
