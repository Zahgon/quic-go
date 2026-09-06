package http3

import (
	"net"
)

type addrList []net.IPAddr

func isIPv4(addr net.IPAddr) bool { _ = "STUB: not implemented"; return false }

func isNotIPv4(addr net.IPAddr) bool { _ = "STUB: not implemented"; return false }

func (addrs addrList) forResolve(network, addr string) net.IPAddr {
	_ = "STUB: not implemented"
	return *new(net.IPAddr)
}

func (addrs addrList) first(strategy func(net.IPAddr) bool) net.IPAddr {
	_ = "STUB: not implemented"
	return *new(net.IPAddr)
}
