//go:build freebsd

package quic

import (
	"net/netip"
	"syscall"

	"golang.org/x/sys/unix"
)

const (
	msgTypeIPTOS = unix.IP_RECVTOS
	ipv4PKTINFO  = 0x7
)

const ecnIPv4DataLen = 1

const batchSize = 8

func parseIPv4PktInfo(body []byte) (ip netip.Addr, _ uint32, ok bool) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), 0, false
}

func isGSOEnabled(syscall.RawConn) bool { _ = "STUB: not implemented"; return false }

func isECNEnabled() bool { _ = "STUB: not implemented"; return false }
