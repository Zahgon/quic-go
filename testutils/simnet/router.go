package simnet

import (
	"net"
	"sync"
	"time"
)

type ipPortKey struct {
	ip    string
	port  uint16
	isUDP bool
}

func (k *ipPortKey) FromNetAddr(addr net.Addr) error { _ = "STUB: not implemented"; return nil }

type addrMap[V any] struct {
	mu    sync.Mutex
	nodes map[ipPortKey]V
}

func (m *addrMap[V]) Get(addr net.Addr) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (m *addrMap[V]) Set(addr net.Addr, v V) error { _ = "STUB: not implemented"; return nil }

func (m *addrMap[V]) Delete(addr net.Addr) error { _ = "STUB: not implemented"; return nil }

type PerfectRouter struct {
	nodes addrMap[PacketReceiver]
}

func (r *PerfectRouter) SendPacket(p Packet) error { _ = "STUB: not implemented"; return nil }

func (r *PerfectRouter) AddNode(addr net.Addr, conn PacketReceiver) {
	_ = "STUB: not implemented"
	return
}

func (r *PerfectRouter) RemoveNode(addr net.Addr) { _ = "STUB: not implemented"; return }

var _ Router = &PerfectRouter{}

type DelayedPacketReceiver struct {
	inner PacketReceiver
	delay time.Duration
}

func (r *DelayedPacketReceiver) RecvPacket(p Packet) { _ = "STUB: not implemented"; return }

type FixedLatencyRouter struct {
	PerfectRouter
	latency time.Duration
}

func (r *FixedLatencyRouter) SendPacket(p Packet) error { _ = "STUB: not implemented"; return nil }

func (r *FixedLatencyRouter) AddNode(addr net.Addr, conn PacketReceiver) {
	_ = "STUB: not implemented"
	return
}

var _ Router = &FixedLatencyRouter{}
