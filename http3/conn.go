package http3

import (
	"log/slog"
	"net"
	"sync"
	"sync/atomic"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
)

const maxQuarterStreamID = 1<<60 - 1

const invalidStreamID = quic.StreamID(-1)

type rawConn struct {
	conn *quic.Conn

	logger *slog.Logger

	enableDatagrams bool

	streamMx sync.Mutex
	streams  map[quic.StreamID]*stateTrackingStream

	rcvdControlStr      atomic.Bool
	rcvdQPACKEncoderStr atomic.Bool
	rcvdQPACKDecoderStr atomic.Bool
	controlStrHandler   func(*quic.ReceiveStream, *frameParser)

	onStreamsEmpty func()

	settings         *Settings
	receivedSettings chan struct{}

	qlogger   qlogwriter.Recorder
	qloggerWG sync.WaitGroup
}

func newRawConn(
	quicConn *quic.Conn,
	enableDatagrams bool,
	onStreamsEmpty func(),
	controlStrHandler func(*quic.ReceiveStream, *frameParser),
	qlogger qlogwriter.Recorder,
	logger *slog.Logger,
) *rawConn {
	_ = "STUB: not implemented"
	return nil
}

func (c *rawConn) OpenUniStream() (*quic.SendStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *rawConn) openControlStream(settings *settingsFrame) (*quic.SendStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *rawConn) TrackStream(str *quic.Stream) *stateTrackingStream {
	_ = "STUB: not implemented"
	return nil
}

func (c *rawConn) UpdateStreamPriority(id quic.StreamID, urgency int8, incremental bool) {
	_ = "STUB: not implemented"
	return
}

func (c *rawConn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *rawConn) ConnectionState() quic.ConnectionState {
	_ = "STUB: not implemented"
	return *new(quic.ConnectionState)
}

func (c *rawConn) clearStream(id quic.StreamID) { _ = "STUB: not implemented"; return }

func (c *rawConn) hasActiveStreams() bool { _ = "STUB: not implemented"; return false }

func (c *rawConn) CloseWithError(code quic.ApplicationErrorCode, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *rawConn) handleUnidirectionalStream(str *quic.ReceiveStream, isServer bool) {
	_ = "STUB: not implemented"
	return
}

func (c *rawConn) handleControlStream(str *quic.ReceiveStream) { _ = "STUB: not implemented"; return }

func (c *rawConn) sendDatagram(streamID quic.StreamID, b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *rawConn) receiveDatagrams() error { _ = "STUB: not implemented"; return nil }

func (c *rawConn) ReceivedSettings() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *rawConn) Settings() *Settings { _ = "STUB: not implemented"; return nil }

func (c *rawConn) closeQlogger() { _ = "STUB: not implemented"; return }
