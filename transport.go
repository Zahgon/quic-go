package quic

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
	"github.com/quic-go/quic-go/qlogwriter"
)

var ErrTransportClosed = &errTransportClosed{}

type errTransportClosed struct {
	err error
}

func (e *errTransportClosed) Unwrap() []error { _ = "STUB: not implemented"; return nil }

func (e *errTransportClosed) Error() string { _ = "STUB: not implemented"; return "" }

func (e *errTransportClosed) Is(target error) bool { _ = "STUB: not implemented"; return false }

var errListenerAlreadySet = errors.New("listener already set")

type closePacket struct {
	payload []byte
	addr    net.Addr
	info    packetInfo
}

type Transport struct {
	Conn net.PacketConn

	ConnectionIDLength int

	ConnectionIDGenerator ConnectionIDGenerator

	StatelessResetKey *StatelessResetKey

	TokenGeneratorKey *TokenGeneratorKey

	MaxTokenAge time.Duration

	DisableVersionNegotiationPackets bool

	VerifySourceAddress func(net.Addr) bool

	ConnContext func(context.Context, *ClientInfo) (context.Context, error)

	Tracer qlogwriter.Recorder

	mutex       sync.Mutex
	handlers    map[protocol.ConnectionID]packetHandler
	resetTokens map[protocol.StatelessResetToken]packetHandler

	initOnce sync.Once
	initErr  error

	connIDLen int

	connIDGenerator   ConnectionIDGenerator
	statelessResetter *statelessResetter

	server *baseServer

	conn rawConn

	closeQueue          chan closePacket
	statelessResetQueue chan receivedPacket

	listening   chan struct{}
	closeErr    error
	createdConn bool
	isSingleUse bool

	readingNonQUICPackets atomic.Bool
	nonQUICPackets        chan receivedPacket

	logger utils.Logger
}

func (t *Transport) Listen(tlsConf *tls.Config, conf *Config) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) ListenEarly(tlsConf *tls.Config, conf *Config) (*EarlyListener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) createServer(tlsConf *tls.Config, conf *Config, allow0RTT bool) (*baseServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) Dial(ctx context.Context, addr net.Addr, tlsConf *tls.Config, conf *Config) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) DialEarly(ctx context.Context, addr net.Addr, tlsConf *tls.Config, conf *Config) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) dial(ctx context.Context, addr net.Addr, host string, tlsConf *tls.Config, conf *Config, use0RTT bool) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) doDial(
	ctx context.Context,
	sendConn sendConn,
	tlsConf *tls.Config,
	config *Config,
	initialPacketNumber protocol.PacketNumber,
	hasNegotiatedVersion bool,
	use0RTT bool,
	version protocol.Version,
) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) init(allowZeroLengthConnIDs bool) error {
	t.initOnce.Do(func() {
		var conn rawConn
		if c, ok := t.Conn.(rawConn); ok {
			conn = c
		} else {
			var err error
			conn, err = wrapConn(t.Conn)
			if err != nil {
				t.initErr = err
				return
			}
		}

		t.logger = utils.DefaultLogger
		t.conn = conn
		t.handlers = make(map[protocol.ConnectionID]packetHandler)
		t.resetTokens = make(map[protocol.StatelessResetToken]packetHandler)
		t.listening = make(chan struct{})

		t.closeQueue = make(chan closePacket, 4)
		t.statelessResetQueue = make(chan receivedPacket, 4)
		if t.TokenGeneratorKey == nil {
			var key TokenGeneratorKey
			if _, err := rand.Read(key[:]); err != nil {
				t.initErr = err
				return
			}
			t.TokenGeneratorKey = &key
		}

		if t.ConnectionIDGenerator != nil {
			t.connIDGenerator = t.ConnectionIDGenerator
			t.connIDLen = t.ConnectionIDGenerator.ConnectionIDLen()
		} else {
			connIDLen := t.ConnectionIDLength
			if t.ConnectionIDLength == 0 && !allowZeroLengthConnIDs {
				connIDLen = protocol.DefaultConnectionIDLength
			}
			t.connIDLen = connIDLen
			t.connIDGenerator = &protocol.DefaultConnectionIDGenerator{ConnLen: t.connIDLen}
		}
		t.statelessResetter = newStatelessResetter(t.StatelessResetKey)

		go func() {
			defer close(t.listening)
			t.listen(conn)

			if t.createdConn {
				conn.Close()
			}
		}()
		go t.runSendQueue()
	})
	return t.initErr
}

func (t *Transport) WriteTo(b []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t *Transport) runSendQueue() { _ = "STUB: not implemented"; return }

func (t *Transport) Close() error { _ = "STUB: not implemented"; return nil }

func (t *Transport) closeServer() { _ = "STUB: not implemented"; return }

func (t *Transport) close(e error) { _ = "STUB: not implemented"; return }

var setBufferWarningOnce sync.Once

func (t *Transport) listen(conn rawConn) { _ = "STUB: not implemented"; return }

//nolint:staticcheck // SA1019 ignore this!

func (t *Transport) maybeStopListening() { _ = "STUB: not implemented"; return }

func (t *Transport) handlePacket(p receivedPacket) { _ = "STUB: not implemented"; return }

func (t *Transport) maybeSendStatelessReset(p receivedPacket) (statelessResetQueued bool) {
	_ = "STUB: not implemented"
	return false
}

func (t *Transport) sendStatelessReset(p receivedPacket) { _ = "STUB: not implemented"; return }

func (t *Transport) maybeHandleStatelessReset(data []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Transport) handleNonQUICPacket(p receivedPacket) { _ = "STUB: not implemented"; return }

const maxQueuedNonQUICPackets = 32

func (t *Transport) ReadNonQUICPacket(ctx context.Context, b []byte) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func setTLSConfigServerName(tlsConf *tls.Config, addr net.Addr, host string) {
	_ = "STUB: not implemented"
	return
}

type packetHandlerMap Transport

var _ connRunner = &packetHandlerMap{}

func (h *packetHandlerMap) Add(id protocol.ConnectionID, handler packetHandler) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *packetHandlerMap) Get(connID protocol.ConnectionID) (packetHandler, bool) {
	_ = "STUB: not implemented"
	return *new(packetHandler), false
}

func (h *packetHandlerMap) AddResetToken(token protocol.StatelessResetToken, handler packetHandler) {
	_ = "STUB: not implemented"
	return
}

func (h *packetHandlerMap) RemoveResetToken(token protocol.StatelessResetToken) {
	_ = "STUB: not implemented"
	return
}

func (h *packetHandlerMap) AddWithConnID(clientDestConnID, newConnID protocol.ConnectionID, handler packetHandler) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *packetHandlerMap) Remove(id protocol.ConnectionID) { _ = "STUB: not implemented"; return }

func (h *packetHandlerMap) ReplaceWithClosed(ids []protocol.ConnectionID, connClosePacket []byte, expiry time.Duration) {
	_ = "STUB: not implemented"
	return
}
