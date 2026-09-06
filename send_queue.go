package quic

import (
	"net"

	"github.com/quic-go/quic-go/internal/protocol"
)

type sender interface {
	Send(p *packetBuffer, gsoSize uint16, ecn protocol.ECN)
	SendProbe(*packetBuffer, net.Addr, packetInfo)
	Run() error
	WouldBlock() bool
	Available() <-chan struct{}
	Close()
}

type queueEntry struct {
	buf     *packetBuffer
	gsoSize uint16
	ecn     protocol.ECN
}

type sendQueue struct {
	queue       chan queueEntry
	closeCalled chan struct{}
	runStopped  chan struct{}
	available   chan struct{}
	conn        sendConn
}

var _ sender = &sendQueue{}

const sendQueueCapacity = 8

func newSendQueue(conn sendConn) sender { _ = "STUB: not implemented"; return *new(sender) }

func (h *sendQueue) Send(p *packetBuffer, gsoSize uint16, ecn protocol.ECN) {
	_ = "STUB: not implemented"
	return
}

func (h *sendQueue) SendProbe(p *packetBuffer, addr net.Addr, info packetInfo) {
	_ = "STUB: not implemented"
	return
}

func (h *sendQueue) WouldBlock() bool { _ = "STUB: not implemented"; return false }

func (h *sendQueue) Available() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (h *sendQueue) Run() error { _ = "STUB: not implemented"; return nil }

func (h *sendQueue) Close() { _ = "STUB: not implemented"; return }
