package quic

import (
	"context"
	"crypto/tls"
	"net"
	"sync"
	"time"

	"github.com/quic-go/quic-go/internal/handshake"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/qerr"
	"github.com/quic-go/quic-go/internal/utils"
	"github.com/quic-go/quic-go/internal/wire"
	"github.com/quic-go/quic-go/qlogwriter"
)

var ErrServerClosed = errServerClosed{}

type errServerClosed struct{}

func (errServerClosed) Error() string { _ = "STUB: not implemented"; return "" }
func (errServerClosed) Unwrap() error { _ = "STUB: not implemented"; return nil }

type packetHandler interface {
	handlePacket(receivedPacket)
	destroy(error)
	closeWithTransportError(qerr.TransportErrorCode)
}

type zeroRTTQueue struct {
	packets    []receivedPacket
	expiration monotime.Time
}

type rejectedPacket struct {
	receivedPacket
	hdr *wire.Header
}

type baseServer struct {
	tr                        *packetHandlerMap
	disableVersionNegotiation bool
	acceptEarlyConns          bool

	tlsConf *tls.Config
	config  *Config

	conn rawConn

	tokenGenerator *handshake.TokenGenerator
	maxTokenAge    time.Duration

	connIDGenerator   ConnectionIDGenerator
	statelessResetter *statelessResetter
	onClose           func()

	receivedPackets chan receivedPacket

	nextZeroRTTCleanup monotime.Time
	zeroRTTQueues      map[protocol.ConnectionID]*zeroRTTQueue

	connContext func(context.Context, *ClientInfo) (context.Context, error)

	newConn func(
		context.Context,
		context.CancelCauseFunc,
		sendConn,
		connRunner,
		protocol.ConnectionID,
		*protocol.ConnectionID,
		protocol.ConnectionID,
		protocol.ConnectionID,
		protocol.ConnectionID,
		ConnectionIDGenerator,
		*statelessResetter,
		*Config,
		*tls.Config,
		*handshake.TokenGenerator,
		bool,
		time.Duration,
		qlogwriter.Trace,
		utils.Logger,
		protocol.Version,
	) *wrappedConn

	closeMx sync.Mutex

	errorChan chan struct{}

	stopAccepting chan struct{}
	closeErr      error
	running       chan struct{}

	versionNegotiationQueue chan receivedPacket
	invalidTokenQueue       chan rejectedPacket
	connectionRefusedQueue  chan rejectedPacket
	retryQueue              chan rejectedPacket
	handshakingCount        sync.WaitGroup

	verifySourceAddress func(net.Addr) bool

	connQueue chan *Conn

	qlogger qlogwriter.Recorder

	logger utils.Logger
}

type Listener struct {
	baseServer *baseServer
}

func (l *Listener) Accept(ctx context.Context) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *Listener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *Listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

type EarlyListener struct {
	baseServer *baseServer
}

func (l *EarlyListener) Accept(ctx context.Context) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *EarlyListener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *EarlyListener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func ListenAddr(addr string, tlsConf *tls.Config, config *Config) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListenAddrEarly(addr string, tlsConf *tls.Config, config *Config) (*EarlyListener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listenUDP(addr string) (*net.UDPConn, error) { _ = "STUB: not implemented"; return nil, nil }

func Listen(conn net.PacketConn, tlsConf *tls.Config, config *Config) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListenEarly(conn net.PacketConn, tlsConf *tls.Config, config *Config) (*EarlyListener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newServer(
	conn rawConn,
	tr *packetHandlerMap,
	connIDGenerator ConnectionIDGenerator,
	statelessResetter *statelessResetter,
	connContext func(context.Context, *ClientInfo) (context.Context, error),
	tlsConf *tls.Config,
	config *Config,
	qlogger qlogwriter.Recorder,
	onClose func(),
	tokenGeneratorKey TokenGeneratorKey,
	maxTokenAge time.Duration,
	verifySourceAddress func(net.Addr) bool,
	disableVersionNegotiation bool,
	acceptEarly bool,
) *baseServer {
	_ = "STUB: not implemented"
	return nil
}

func (s *baseServer) run() { _ = "STUB: not implemented"; return }

func (s *baseServer) runSendQueue() { _ = "STUB: not implemented"; return }

func (s *baseServer) Accept(ctx context.Context) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *baseServer) accept(ctx context.Context) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *baseServer) Close() error { _ = "STUB: not implemented"; return nil }

func (s *baseServer) close(e error, transportClose bool) { _ = "STUB: not implemented"; return }

func (s *baseServer) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (s *baseServer) handlePacket(p receivedPacket) { _ = "STUB: not implemented"; return }

func (s *baseServer) handlePacketImpl(p receivedPacket) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *baseServer) handle0RTTPacket(p receivedPacket) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *baseServer) cleanupZeroRTTQueues(now monotime.Time) { _ = "STUB: not implemented"; return }

func (s *baseServer) validateToken(token *handshake.Token, addr net.Addr) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *baseServer) handleInitialImpl(p receivedPacket, hdr *wire.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *baseServer) refuseNewConn(p receivedPacket, hdr *wire.Header) {
	_ = "STUB: not implemented"
	return
}

func (s *baseServer) handleNewConn(conn *wrappedConn) { _ = "STUB: not implemented"; return }

func (s *baseServer) sendRetry(p rejectedPacket) { _ = "STUB: not implemented"; return }

func (s *baseServer) sendRetryPacket(p rejectedPacket) error { _ = "STUB: not implemented"; return nil }

func (s *baseServer) maybeSendInvalidToken(p rejectedPacket) { _ = "STUB: not implemented"; return }

func (s *baseServer) sendConnectionRefused(p rejectedPacket) { _ = "STUB: not implemented"; return }

func (s *baseServer) sendError(remoteAddr net.Addr, hdr *wire.Header, sealer handshake.LongHeaderSealer, errorCode qerr.TransportErrorCode, info packetInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *baseServer) enqueueVersionNegotiationPacket(p receivedPacket) (bufferInUse bool) {
	_ = "STUB: not implemented"
	return false
}

func (s *baseServer) maybeSendVersionNegotiationPacket(p receivedPacket) {
	_ = "STUB: not implemented"
	return
}
