package simnet

import (
	"net"
	"sync"
	"time"
)

type packetWithDeliveryTime struct {
	Packet
	DeliveryTime time.Time
}

type LinkSettings struct {
	MTU int
}

type SimulatedLink struct {
	wg sync.WaitGroup

	downstreamQueue *queue
	upstreamQueue   *queue

	UplinkSettings   LinkSettings
	DownlinkSettings LinkSettings

	Latency time.Duration

	LatencyFunc func(Packet) time.Duration

	UploadPacket   Router
	downloadPacket PacketReceiver
}

func (l *SimulatedLink) AddNode(addr net.Addr, receiver PacketReceiver) {
	_ = "STUB: not implemented"
	return
}

func (l *SimulatedLink) Start() { _ = "STUB: not implemented"; return }

func (l *SimulatedLink) Close() error { _ = "STUB: not implemented"; return nil }

func (l *SimulatedLink) backgroundDownlink() { _ = "STUB: not implemented"; return }

func (l *SimulatedLink) backgroundUplink() { _ = "STUB: not implemented"; return }

func (l *SimulatedLink) SendPacket(p Packet) error { _ = "STUB: not implemented"; return nil }

func (l *SimulatedLink) RecvPacket(p Packet) { _ = "STUB: not implemented"; return }
