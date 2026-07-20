//go:build linux

package quic

import (
	"net/netip"
	"syscall"

	"golang.org/x/sys/unix"
)

const (
	msgTypeIPTOS = unix.IP_TOS
	ipv4PKTINFO  = unix.IP_PKTINFO
)

const ecnIPv4DataLen = 1

const batchSize = 8

var kernelVersionMajor int

func init() {
	kernelVersionMajor, _ = kernelVersion()
}

func forceSetReceiveBuffer(c syscall.RawConn, bytes int) error {
	_ = "STUB: not implemented"
	return nil
}

func forceSetSendBuffer(c syscall.RawConn, bytes int) error { _ = "STUB: not implemented"; return nil }

func parseIPv4PktInfo(body []byte) (ip netip.Addr, ifIndex uint32, ok bool) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), 0, false
}

func isGSOEnabled(conn syscall.RawConn) bool { _ = "STUB: not implemented"; return false }

func appendUDPSegmentSizeMsg(b []byte, size uint16) []byte { _ = "STUB: not implemented"; return nil }

func isGSOError(err error) bool { _ = "STUB: not implemented"; return false }

func isPermissionError(err error) bool { _ = "STUB: not implemented"; return false }

func isECNEnabled() bool { _ = "STUB: not implemented"; return false }

func kernelVersion() (major, minor int) { _ = "STUB: not implemented"; return 0, 0 }
