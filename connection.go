package quic

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/quic-go/quic-go/internal/ackhandler"
	"github.com/quic-go/quic-go/internal/handshake"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
	"github.com/quic-go/quic-go/internal/utils/ringbuffer"
	"github.com/quic-go/quic-go/internal/wire"
	"github.com/quic-go/quic-go/qlog"
	"github.com/quic-go/quic-go/qlogwriter"
)

type unpacker interface {
	UnpackLongHeader(hdr *wire.Header, data []byte) (*unpackedPacket, error)
	UnpackShortHeader(rcvTime monotime.Time, data []byte) (protocol.PacketNumber, protocol.PacketNumberLen, protocol.KeyPhaseBit, []byte, error)
}

type cryptoStreamHandler interface {
	StartHandshake(context.Context) error
	ChangeConnectionID(protocol.ConnectionID)
	SetLargest1RTTAcked(protocol.PacketNumber) error
	SetHandshakeConfirmed()
	GetSessionTicket() ([]byte, error)
	NextEvent() handshake.Event
	DiscardInitialKeys()
	HandleMessage([]byte, protocol.EncryptionLevel) error
	io.Closer
	ConnectionState() handshake.ConnectionState
}

type receivedPacket struct {
	buffer *packetBuffer

	remoteAddr net.Addr
	rcvTime    monotime.Time
	data       []byte

	ecn protocol.ECN

	info packetInfo
}

type receivedPacketWithDatagramID struct {
	receivedPacket
	datagramID qlog.DatagramID
}

func (p *receivedPacket) Size() protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (p *receivedPacket) Clone() *receivedPacket { _ = "STUB: not implemented"; return nil }

type connRunner interface {
	Add(protocol.ConnectionID, packetHandler) bool
	Remove(protocol.ConnectionID)
	ReplaceWithClosed([]protocol.ConnectionID, []byte, time.Duration)
	AddResetToken(protocol.StatelessResetToken, packetHandler)
	RemoveResetToken(protocol.StatelessResetToken)
}

type closeError struct {
	err       error
	immediate bool
}

type errCloseForRecreating struct {
	nextPacketNumber protocol.PacketNumber
	nextVersion      protocol.Version
}

func (e *errCloseForRecreating) Error() string { _ = "STUB: not implemented"; return "" }

var deadlineSendImmediately = monotime.Time(42 * time.Millisecond)

type blockMode uint8

const (
	blockModeNone blockMode = iota

	blockModeCongestionLimited

	blockModeHardBlocked
)

type Conn struct {
	handshakeDestConnID protocol.ConnectionID

	origDestConnID protocol.ConnectionID
	retrySrcConnID *protocol.ConnectionID

	srcConnIDLen int

	perspective protocol.Perspective
	version     protocol.Version
	config      *Config

	conn      sendConn
	sendQueue sender

	pathManager         *pathManager
	largestRcvdAppData  protocol.PacketNumber
	pathManagerOutgoing atomic.Pointer[pathManagerOutgoing]

	streamsMap      *streamsMap
	connIDManager   *connIDManager
	connIDGenerator *connIDGenerator

	rttStats  *utils.RTTStats
	connStats utils.ConnectionStats

	cryptoStreamManager   *cryptoStreamManager
	sentPacketHandler     ackhandler.SentPacketHandler
	receivedPacketHandler ackhandler.ReceivedPacketHandler
	retransmissionQueue   *retransmissionQueue
	framer                *framer
	connFlowController    *connectionFlowController
	tokenStoreKey         string
	tokenGenerator        *handshake.TokenGenerator

	unpacker      unpacker
	frameParser   wire.FrameParser
	packer        packer
	mtuDiscoverer mtuDiscoverer

	maxPayloadSizeEstimate atomic.Uint32

	initialStream       *initialCryptoStream
	handshakeStream     *cryptoStream
	oneRTTStream        *cryptoStream
	cryptoStreamHandler cryptoStreamHandler

	notifyReceivedPacket chan struct{}
	sendingScheduled     chan struct{}
	receivedPacketMx     sync.Mutex
	receivedPackets      ringbuffer.RingBuffer[receivedPacket]

	closeChan chan struct{}
	closeErr  atomic.Pointer[closeError]

	ctx                   context.Context
	ctxCancel             context.CancelCauseFunc
	handshakeCompleteChan chan struct{}

	undecryptablePackets          []receivedPacketWithDatagramID
	undecryptablePacketsToProcess []receivedPacketWithDatagramID

	earlyConnReadyChan chan struct{}
	sentFirstPacket    bool
	droppedInitialKeys bool
	handshakeComplete  bool
	handshakeConfirmed bool

	receivedRetry       bool
	versionNegotiated   bool
	receivedFirstPacket bool

	blocked blockMode

	idleTimeout  time.Duration
	creationTime monotime.Time

	lastPacketReceivedTime monotime.Time

	firstAckElicitingPacketAfterIdleSentTime monotime.Time

	pacingDeadline monotime.Time

	peerParams *wire.TransportParameters

	timer *time.Timer

	keepAlivePingSent bool
	keepAliveInterval time.Duration

	datagramQueue *datagramQueue

	connStateMutex sync.Mutex
	connState      ConnectionState

	logID     string
	qlogTrace qlogwriter.Trace
	qlogger   qlogwriter.Recorder
	logger    utils.Logger
}

var _ streamSender = &Conn{}

type connTestHooks struct {
	run                     func() error
	earlyConnReady          func() <-chan struct{}
	context                 func() context.Context
	handshakeComplete       func() <-chan struct{}
	closeWithTransportError func(TransportErrorCode)
	destroy                 func(error)
	handlePacket            func(receivedPacket)
}

type wrappedConn struct {
	testHooks *connTestHooks
	*Conn
}

var newConnection = func(
	ctx context.Context,
	ctxCancel context.CancelCauseFunc,
	conn sendConn,
	runner connRunner,
	origDestConnID protocol.ConnectionID,
	retrySrcConnID *protocol.ConnectionID,
	clientDestConnID protocol.ConnectionID,
	destConnID protocol.ConnectionID,
	srcConnID protocol.ConnectionID,
	connIDGenerator ConnectionIDGenerator,
	statelessResetter *statelessResetter,
	conf *Config,
	tlsConf *tls.Config,
	tokenGenerator *handshake.TokenGenerator,
	clientAddressValidated bool,
	rtt time.Duration,
	qlogTrace qlogwriter.Trace,
	logger utils.Logger,
	v protocol.Version,
) *wrappedConn {
	s := &Conn{
		ctx:                 ctx,
		ctxCancel:           ctxCancel,
		conn:                conn,
		config:              conf,
		handshakeDestConnID: destConnID,
		srcConnIDLen:        srcConnID.Len(),
		tokenGenerator:      tokenGenerator,
		oneRTTStream:        newCryptoStream(),
		perspective:         protocol.PerspectiveServer,
		qlogTrace:           qlogTrace,
		logger:              logger,
		version:             v,
	}
	if qlogTrace != nil {
		s.qlogger = qlogTrace.AddProducer()
	}
	if origDestConnID.Len() > 0 {
		s.logID = origDestConnID.String()
	} else {
		s.logID = destConnID.String()
	}
	s.connIDManager = newConnIDManager(
		destConnID,
		func(token protocol.StatelessResetToken) { runner.AddResetToken(token, s) },
		runner.RemoveResetToken,
		s.queueControlFrame,
	)
	s.connIDGenerator = newConnIDGenerator(
		runner,
		srcConnID,
		&clientDestConnID,
		statelessResetter,
		connRunnerCallbacks{
			AddConnectionID:    func(connID protocol.ConnectionID) { runner.Add(connID, s) },
			RemoveConnectionID: runner.Remove,
			ReplaceWithClosed:  runner.ReplaceWithClosed,
		},
		s.queueControlFrame,
		connIDGenerator,
	)
	s.preSetup()
	s.rttStats.SetInitialRTT(rtt)
	s.sentPacketHandler = ackhandler.NewSentPacketHandler(
		0,
		protocol.ByteCount(s.config.InitialPacketSize),
		s.rttStats,
		&s.connStats,
		clientAddressValidated,
		s.conn.capabilities().ECN,
		s.receivedPacketHandler.IgnorePacketsBelow,
		s.perspective,
		s.qlogger,
		s.logger,
	)
	s.maxPayloadSizeEstimate.Store(uint32(estimateMaxPayloadSize(protocol.ByteCount(s.config.InitialPacketSize))))
	statelessResetToken := statelessResetter.GetStatelessResetToken(srcConnID)
	params := &wire.TransportParameters{
		InitialMaxStreamDataBidiLocal:   protocol.ByteCount(s.config.InitialStreamReceiveWindow),
		InitialMaxStreamDataBidiRemote:  protocol.ByteCount(s.config.InitialStreamReceiveWindow),
		InitialMaxStreamDataUni:         protocol.ByteCount(s.config.InitialStreamReceiveWindow),
		InitialMaxData:                  protocol.ByteCount(s.config.InitialConnectionReceiveWindow),
		MaxIdleTimeout:                  s.config.MaxIdleTimeout,
		MaxBidiStreamNum:                protocol.StreamNum(s.config.MaxIncomingStreams),
		MaxUniStreamNum:                 protocol.StreamNum(s.config.MaxIncomingUniStreams),
		MaxAckDelay:                     protocol.MaxAckDelayInclGranularity,
		AckDelayExponent:                protocol.AckDelayExponent,
		MaxUDPPayloadSize:               protocol.MaxPacketBufferSize,
		StatelessResetToken:             &statelessResetToken,
		OriginalDestinationConnectionID: origDestConnID,

		ActiveConnectionIDLimit:   protocol.MaxActiveConnectionIDs,
		InitialSourceConnectionID: srcConnID,
		RetrySourceConnectionID:   retrySrcConnID,
		EnableResetStreamAt:       conf.EnableStreamResetPartialDelivery,
	}
	if s.config.EnableDatagrams {
		params.MaxDatagramFrameSize = wire.MaxDatagramSize
	} else {
		params.MaxDatagramFrameSize = protocol.InvalidByteCount
	}
	if s.qlogger != nil {
		s.qlogTransportParameters(params, protocol.PerspectiveServer, false)
	}
	cs := handshake.NewCryptoSetupServer(
		clientDestConnID,
		conn.LocalAddr(),
		conn.RemoteAddr(),
		params,
		tlsConf,
		conf.Allow0RTT,
		s.rttStats,
		s.qlogger,
		logger,
		s.version,
	)
	s.cryptoStreamHandler = cs
	s.packer = newPacketPacker(srcConnID, s.connIDManager.Get, s.initialStream, s.handshakeStream, s.sentPacketHandler, s.retransmissionQueue, cs, s.framer, &s.receivedPacketHandler, s.datagramQueue, s.perspective)
	s.unpacker = newPacketUnpacker(cs, s.srcConnIDLen)
	s.cryptoStreamManager = newCryptoStreamManager(s.initialStream, s.handshakeStream, s.oneRTTStream)
	return &wrappedConn{Conn: s}
}

var newClientConnection = func(
	ctx context.Context,
	conn sendConn,
	runner connRunner,
	destConnID protocol.ConnectionID,
	srcConnID protocol.ConnectionID,
	connIDGenerator ConnectionIDGenerator,
	statelessResetter *statelessResetter,
	conf *Config,
	tlsConf *tls.Config,
	initialPacketNumber protocol.PacketNumber,
	enable0RTT bool,
	hasNegotiatedVersion bool,
	qlogTrace qlogwriter.Trace,
	logger utils.Logger,
	v protocol.Version,
) *wrappedConn {
	s := &Conn{
		conn:                conn,
		config:              conf,
		origDestConnID:      destConnID,
		handshakeDestConnID: destConnID,
		srcConnIDLen:        srcConnID.Len(),
		perspective:         protocol.PerspectiveClient,
		logID:               destConnID.String(),
		logger:              logger,
		qlogTrace:           qlogTrace,
		versionNegotiated:   hasNegotiatedVersion,
		version:             v,
	}
	if qlogTrace != nil {
		s.qlogger = qlogTrace.AddProducer()
	}
	if s.qlogger != nil {
		var srcAddr, destAddr *net.UDPAddr
		if addr, ok := conn.LocalAddr().(*net.UDPAddr); ok {
			srcAddr = addr
		}
		if addr, ok := conn.RemoteAddr().(*net.UDPAddr); ok {
			destAddr = addr
		}
		s.qlogger.RecordEvent(startedConnectionEvent(srcAddr, destAddr))
	}
	s.connIDManager = newConnIDManager(
		destConnID,
		func(token protocol.StatelessResetToken) { runner.AddResetToken(token, s) },
		runner.RemoveResetToken,
		s.queueControlFrame,
	)
	s.connIDGenerator = newConnIDGenerator(
		runner,
		srcConnID,
		nil,
		statelessResetter,
		connRunnerCallbacks{
			AddConnectionID:    func(connID protocol.ConnectionID) { runner.Add(connID, s) },
			RemoveConnectionID: runner.Remove,
			ReplaceWithClosed:  runner.ReplaceWithClosed,
		},
		s.queueControlFrame,
		connIDGenerator,
	)
	s.ctx, s.ctxCancel = context.WithCancelCause(ctx)
	s.preSetup()
	s.sentPacketHandler = ackhandler.NewSentPacketHandler(
		initialPacketNumber,
		protocol.ByteCount(s.config.InitialPacketSize),
		s.rttStats,
		&s.connStats,
		false,
		s.conn.capabilities().ECN,
		s.receivedPacketHandler.IgnorePacketsBelow,
		s.perspective,
		s.qlogger,
		s.logger,
	)
	s.maxPayloadSizeEstimate.Store(uint32(estimateMaxPayloadSize(protocol.ByteCount(s.config.InitialPacketSize))))
	oneRTTStream := newCryptoStream()
	params := &wire.TransportParameters{
		InitialMaxStreamDataBidiRemote: protocol.ByteCount(s.config.InitialStreamReceiveWindow),
		InitialMaxStreamDataBidiLocal:  protocol.ByteCount(s.config.InitialStreamReceiveWindow),
		InitialMaxStreamDataUni:        protocol.ByteCount(s.config.InitialStreamReceiveWindow),
		InitialMaxData:                 protocol.ByteCount(s.config.InitialConnectionReceiveWindow),
		MaxIdleTimeout:                 s.config.MaxIdleTimeout,
		MaxBidiStreamNum:               protocol.StreamNum(s.config.MaxIncomingStreams),
		MaxUniStreamNum:                protocol.StreamNum(s.config.MaxIncomingUniStreams),
		MaxAckDelay:                    protocol.MaxAckDelayInclGranularity,
		MaxUDPPayloadSize:              protocol.MaxPacketBufferSize,
		AckDelayExponent:               protocol.AckDelayExponent,

		ActiveConnectionIDLimit:   protocol.MaxActiveConnectionIDs,
		InitialSourceConnectionID: srcConnID,
		EnableResetStreamAt:       conf.EnableStreamResetPartialDelivery,
	}
	if s.config.EnableDatagrams {
		params.MaxDatagramFrameSize = wire.MaxDatagramSize
	} else {
		params.MaxDatagramFrameSize = protocol.InvalidByteCount
	}
	if s.qlogger != nil {
		s.qlogTransportParameters(params, protocol.PerspectiveClient, false)
	}
	cs := handshake.NewCryptoSetupClient(
		destConnID,
		params,
		tlsConf,
		enable0RTT,
		s.rttStats,
		s.qlogger,
		logger,
		s.version,
	)
	s.cryptoStreamHandler = cs
	s.cryptoStreamManager = newCryptoStreamManager(s.initialStream, s.handshakeStream, oneRTTStream)
	s.unpacker = newPacketUnpacker(cs, s.srcConnIDLen)
	s.packer = newPacketPacker(srcConnID, s.connIDManager.Get, s.initialStream, s.handshakeStream, s.sentPacketHandler, s.retransmissionQueue, cs, s.framer, &s.receivedPacketHandler, s.datagramQueue, s.perspective)
	if len(tlsConf.ServerName) > 0 {
		s.tokenStoreKey = tlsConf.ServerName
	} else {
		s.tokenStoreKey = conn.RemoteAddr().String()
	}
	if s.config.TokenStore != nil {
		if token := s.config.TokenStore.Pop(s.tokenStoreKey); token != nil {
			s.packer.SetToken(token.data)
			s.rttStats.SetInitialRTT(token.rtt)
		}
	}
	return &wrappedConn{Conn: s}
}

func (c *Conn) preSetup() { _ = "STUB: not implemented"; return }

func (c *Conn) run() (err error) { _ = "STUB: not implemented"; return nil }

func (c *Conn) earlyConnReady() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *Conn) Context() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

func (c *Conn) supportsDatagrams() bool { _ = "STUB: not implemented"; return false }

func (c *Conn) ConnectionState() ConnectionState {
	_ = "STUB: not implemented"
	return *new(ConnectionState)
}

type ConnectionStats struct {
	MinRTT time.Duration

	LatestRTT time.Duration

	SmoothedRTT time.Duration

	MeanDeviation time.Duration

	BytesSent uint64

	PacketsSent uint64

	BytesReceived uint64

	PacketsReceived uint64

	BytesLost uint64

	PacketsLost uint64
}

func (c *Conn) ConnectionStats() ConnectionStats {
	_ = "STUB: not implemented"
	return *new(ConnectionStats)
}

func (c *Conn) nextIdleTimeoutTime() monotime.Time {
	_ = "STUB: not implemented"
	return *new(monotime.Time)
}

func (c *Conn) nextKeepAliveTime() monotime.Time {
	_ = "STUB: not implemented"
	return *new(monotime.Time)
}

func (c *Conn) maybeResetTimer() { _ = "STUB: not implemented"; return }

func (c *Conn) idleTimeoutStartTime() monotime.Time {
	_ = "STUB: not implemented"
	return *new(monotime.Time)
}

func (c *Conn) switchToNewPath(tr *Transport, now monotime.Time) { _ = "STUB: not implemented"; return }

func (c *Conn) handleHandshakeComplete(now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleHandshakeConfirmed(now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

const maxPacketsToProcess = 32

func (c *Conn) handlePackets() (wasProcessed bool, _ error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Conn) handleOnePacket(rp receivedPacket, datagramID qlog.DatagramID) (wasProcessed bool, _ error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Conn) handleShortHeaderPacket(
	p receivedPacket,
	isCoalesced bool,
	datagramID qlog.DatagramID,
) (wasProcessed bool, _ error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Conn) handleLongHeaderPacket(p receivedPacket, hdr *wire.Header, datagramID qlog.DatagramID) (wasProcessed bool, _ error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Conn) handleUnpackError(err error, p receivedPacket, pt qlog.PacketType, datagramID qlog.DatagramID) (wasQueued bool, _ error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Conn) handleRetryPacket(hdr *wire.Header, data []byte, rcvTime monotime.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Conn) handleVersionNegotiationPacket(p receivedPacket) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleUnpackedLongHeaderPacket(
	packet *unpackedPacket,
	ecn protocol.ECN,
	rcvTime monotime.Time,
	datagramID qlog.DatagramID,
	packetSize protocol.ByteCount,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleUnpackedShortHeaderPacket(
	destConnID protocol.ConnectionID,
	pn protocol.PacketNumber,
	data []byte,
	ecn protocol.ECN,
	rcvTime monotime.Time,
	log func([]qlog.Frame),
) (isNonProbing bool, pathChallenge *wire.PathChallengeFrame, _ error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (c *Conn) handleFrames(
	data []byte,
	destConnID protocol.ConnectionID,
	encLevel protocol.EncryptionLevel,
	log func([]qlog.Frame),
	rcvTime monotime.Time,
) (isAckEliciting, isNonProbing bool, pathChallenge *wire.PathChallengeFrame, _ error) {
	_ = "STUB: not implemented"
	return false, false, nil, nil
}

func (c *Conn) handleFrame(
	f wire.Frame,
	encLevel protocol.EncryptionLevel,
	destConnID protocol.ConnectionID,
	rcvTime monotime.Time,
) (pathChallenge *wire.PathChallengeFrame, _ error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) handlePacket(p receivedPacket) { _ = "STUB: not implemented"; return }

func (c *Conn) handleConnectionCloseFrame(frame *wire.ConnectionCloseFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleCryptoFrame(frame *wire.CryptoFrame, encLevel protocol.EncryptionLevel, rcvTime monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleHandshakeEvents(now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handlePathChallengeFrame(f *wire.PathChallengeFrame) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) handlePathResponseFrame(f *wire.PathResponseFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handlePathResponseFrameClient(f *wire.PathResponseFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handlePathResponseFrameServer(f *wire.PathResponseFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleNewTokenFrame(frame *wire.NewTokenFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleHandshakeDoneFrame(rcvTime monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleAckFrame(frame *wire.AckFrame, encLevel protocol.EncryptionLevel, rcvTime monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleDatagramFrame(f *wire.DatagramFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) setCloseError(e *closeError) { _ = "STUB: not implemented"; return }

func (c *Conn) closeLocal(e error) { _ = "STUB: not implemented"; return }

func (c *Conn) destroy(e error) { _ = "STUB: not implemented"; return }

func (c *Conn) destroyImpl(e error) { _ = "STUB: not implemented"; return }

func (c *Conn) CloseWithError(code ApplicationErrorCode, desc string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) closeWithTransportError(code TransportErrorCode) { _ = "STUB: not implemented"; return }

func (c *Conn) handleCloseError(closeErr *closeError) { _ = "STUB: not implemented"; return }

func (c *Conn) dropEncryptionLevel(encLevel protocol.EncryptionLevel, now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:exhaustive // only Initial and 0-RTT need special treatment

func (c *Conn) restoreTransportParameters(params *wire.TransportParameters) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) handleTransportParameters(params *wire.TransportParameters) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) checkTransportParameters(params *wire.TransportParameters) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) applyTransportParameters() { _ = "STUB: not implemented"; return }

func (c *Conn) triggerSending(now monotime.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) sendPackets(now monotime.Time) error { _ = "STUB: not implemented"; return nil }

//nolint:exhaustive // only need to handle pacing-related events here

func (c *Conn) sendPacketsWithoutGSO(now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) sendPacketsWithGSO(now monotime.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) resetPacingDeadline() { _ = "STUB: not implemented"; return }

func (c *Conn) maybeSendAckOnlyPacket(now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) sendProbePacket(sendMode ackhandler.SendMode, now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:exhaustive // We only need to handle the PTO send modes here.

func (c *Conn) appendOneShortHeaderPacket(buf *packetBuffer, maxSize protocol.ByteCount, ecn protocol.ECN, now monotime.Time) (protocol.ByteCount, error) {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount), nil
}

func (c *Conn) registerPackedShortHeaderPacket(p shortHeaderPacket, ecn protocol.ECN, now monotime.Time) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) sendPackedCoalescedPacket(packet *coalescedPacket, ecn protocol.ECN, now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) sendConnectionClose(e error) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) maxPacketSize() protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (c *Conn) AcceptStream(ctx context.Context) (*Stream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) AcceptUniStream(ctx context.Context) (*ReceiveStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) OpenStream() (*Stream, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) OpenStreamSync(ctx context.Context) (*Stream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) OpenUniStream() (*SendStream, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) OpenUniStreamSync(ctx context.Context) (*SendStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) newFlowController(id protocol.StreamID) *streamFlowController {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) scheduleSending() { _ = "STUB: not implemented"; return }

func (c *Conn) tryQueueingUndecryptablePacket(p receivedPacket, pt qlog.PacketType, datagramID qlog.DatagramID) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) queueControlFrame(f wire.Frame) { _ = "STUB: not implemented"; return }

func (c *Conn) onHasConnectionData() { _ = "STUB: not implemented"; return }

func (c *Conn) onHasStreamData(id protocol.StreamID, str *SendStream) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) onHasStreamControlFrame(id protocol.StreamID, str streamControlFrameGetter) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) onStreamCompleted(id protocol.StreamID) { _ = "STUB: not implemented"; return }

func (c *Conn) SendDatagram(p []byte) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) ReceiveDatagram(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) getPathManager() *pathManagerOutgoing { _ = "STUB: not implemented"; return nil }

func (c *Conn) AddPath(t *Transport) (*Path, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) HandshakeComplete() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *Conn) QlogTrace() qlogwriter.Trace {
	_ = "STUB: not implemented"
	return *new(qlogwriter.Trace)
}

func (c *Conn) NextConnection(ctx context.Context) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func estimateMaxPayloadSize(mtu protocol.ByteCount) protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}
