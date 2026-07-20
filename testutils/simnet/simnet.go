package simnet

import (
	"net"
	"time"
)

type Simnet struct {
	Router Router

	links []*SimulatedLink
}

type NodeBiDiLinkSettings struct {
	Downlink LinkSettings

	Uplink LinkSettings

	Latency time.Duration

	LatencyFunc func(Packet) time.Duration
}

func (n *Simnet) Start() error { _ = "STUB: not implemented"; return nil }

func (n *Simnet) Close() error { _ = "STUB: not implemented"; return nil }

func (n *Simnet) NewEndpoint(addr *net.UDPAddr, linkSettings NodeBiDiLinkSettings) *SimConn {
	_ = "STUB: not implemented"
	return nil
}
