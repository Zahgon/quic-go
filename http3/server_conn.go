package http3

import (
	"context"
	"log/slog"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/quic-go/qpack"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
)

type RawServerConn struct {
	rawConn rawConn

	idleTimeout time.Duration
	idleTimer   *time.Timer

	serverContext  context.Context
	requestHandler http.Handler
	maxHeaderBytes int

	decoder       *qpack.Decoder
	priorityAware atomic.Bool

	qlogger qlogwriter.Recorder
	logger  *slog.Logger
}

func newRawServerConn(
	conn *quic.Conn,
	enableDatagrams bool,
	idleTimeout time.Duration,
	qlogger qlogwriter.Recorder,
	logger *slog.Logger,
	serverContext context.Context,
	requestHandler http.Handler,
	maxHeaderBytes int,
) *RawServerConn {
	_ = "STUB: not implemented"
	return nil
}

func (c *RawServerConn) onStreamsEmpty() { _ = "STUB: not implemented"; return }

func (c *RawServerConn) CloseWithError(code quic.ApplicationErrorCode, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RawServerConn) HandleRequestStream(str *quic.Stream) { _ = "STUB: not implemented"; return }

func (c *RawServerConn) requestMaxHeaderBytes() int { _ = "STUB: not implemented"; return 0 }

func (c *RawServerConn) openControlStream(settings *settingsFrame) (*quic.SendStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *RawServerConn) handleRequestStream(str *stateTrackingStream) {
	_ = "STUB: not implemented"
	return
}

func (c *RawServerConn) handleControlStream(_ *quic.ReceiveStream, fp *frameParser) {
	_ = "STUB: not implemented"
	return
}

func (c *RawServerConn) rejectWithHeaderFieldsTooLarge(str *stateTrackingStream) {
	_ = "STUB: not implemented"
	return
}

func (c *RawServerConn) HandleUnidirectionalStream(str *quic.ReceiveStream) {
	_ = "STUB: not implemented"
	return
}
