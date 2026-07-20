package quicproxy

import (
	"net"
	"sync"
	"time"

	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/utils"
)

type connection struct {
	ClientAddr *net.UDPAddr
	ServerAddr *net.UDPAddr

	mx         sync.Mutex
	ServerConn *net.UDPConn

	incomingPackets chan packetEntry

	Incoming *queue
	Outgoing *queue
}

func (c *connection) queuePacket(t monotime.Time, b []byte) { _ = "STUB: not implemented"; return }

func (c *connection) SwitchConn(conn *net.UDPConn) { _ = "STUB: not implemented"; return }

func (c *connection) GetServerConn() *net.UDPConn { _ = "STUB: not implemented"; return nil }

type Direction int

const (
	DirectionIncoming Direction = iota

	DirectionOutgoing

	DirectionBoth
)

type packetEntry struct {
	Time monotime.Time
	Raw  []byte
}

type queue struct {
	sync.Mutex

	timer   *time.Timer
	Packets []packetEntry
}

func newQueue() *queue { _ = "STUB: not implemented"; return nil }

func (q *queue) Add(e packetEntry) { _ = "STUB: not implemented"; return }

func (q *queue) Get() []byte { _ = "STUB: not implemented"; return nil }

func (q *queue) Timer() <-chan time.Time { _ = "STUB: not implemented"; return nil }

func (q *queue) Close() { _ = "STUB: not implemented"; return }

func (d Direction) String() string { _ = "STUB: not implemented"; return "" }

func (d Direction) Is(dir Direction) bool { _ = "STUB: not implemented"; return false }

type DropCallback func(dir Direction, from, to net.Addr, packet []byte) bool

type DelayCallback func(dir Direction, from, to net.Addr, packet []byte) time.Duration

type Proxy struct {
	Conn *net.UDPConn

	ServerAddr *net.UDPAddr

	DropPacket DropCallback

	DelayPacket DelayCallback

	closeChan chan struct{}
	logger    utils.Logger

	mutex      sync.Mutex
	clientDict map[string]*connection
}

func (p *Proxy) Start() error { _ = "STUB: not implemented"; return nil }

func (p *Proxy) SwitchConn(clientAddr *net.UDPAddr, conn *net.UDPConn) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Proxy) Close() error { _ = "STUB: not implemented"; return nil }

func (p *Proxy) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (p *Proxy) newConnection(cliAddr *net.UDPAddr) (*connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Proxy) runProxy() error { _ = "STUB: not implemented"; return nil }

func (p *Proxy) runOutgoingConnection(conn *connection) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Proxy) runIncomingConnection(conn *connection) error {
	_ = "STUB: not implemented"
	return nil
}
