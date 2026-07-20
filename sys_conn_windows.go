//go:build windows

package quic

import (
	"net/netip"
	"syscall"
)

func newConn(c OOBCapablePacketConn, supportsDF bool) (*basicConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inspectReadBuffer(c syscall.RawConn) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func inspectWriteBuffer(c syscall.RawConn) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type packetInfo struct {
	addr netip.Addr
}

func (i *packetInfo) OOB() []byte { _ = "STUB: not implemented"; return nil }
