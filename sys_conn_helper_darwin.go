//go:build darwin

package quic

import (
	"net/netip"
	"syscall"

	"golang.org/x/sys/unix"
)

const (
	msgTypeIPTOS = unix.IP_RECVTOS
	ipv4PKTINFO  = unix.IP_RECVPKTINFO
)

const ecnIPv4DataLen = 4

const batchSize = 1

func parseIPv4PktInfo(body []byte) (ip netip.Addr, ifIndex uint32, ok bool) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), 0, false
}

func isGSOEnabled(syscall.RawConn) bool { _ = "STUB: not implemented"; return false }

func isECNEnabled() bool { _ = "STUB: not implemented"; return false }
