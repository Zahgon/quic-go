//go:build !darwin && !linux && !freebsd && !windows

package quic

import (
	"net"
	"net/netip"
)

func newConn(c net.PacketConn, supportsDF bool) (*basicConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inspectReadBuffer(any) (int, error)  { _ = "STUB: not implemented"; return 0, nil }
func inspectWriteBuffer(any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type packetInfo struct {
	addr netip.Addr
}

func (i *packetInfo) OOB() []byte { _ = "STUB: not implemented"; return nil }
