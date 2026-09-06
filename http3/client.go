package http3

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/quic-go/qpack"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
)

const (
	MethodGet0RTT = "GET_0RTT"

	MethodHead0RTT = "HEAD_0RTT"
)

const (
	defaultUserAgent              = "quic-go HTTP/3"
	defaultMaxResponseHeaderBytes = 10 * 1 << 20
)

var errGoAway = errors.New("connection in graceful shutdown")

type errConnUnusable struct{ e error }

func (e *errConnUnusable) Unwrap() error { _ = "STUB: not implemented"; return nil }
func (e *errConnUnusable) Error() string { _ = "STUB: not implemented"; return "" }

const max1xxResponses = 5

var defaultQuicConfig = &quic.Config{
	MaxIncomingStreams: -1,
	KeepAlivePeriod:    10 * time.Second,
}

type ClientConn struct {
	conn    *quic.Conn
	rawConn *rawConn

	decoder *qpack.Decoder

	additionalSettings map[uint64]uint64

	maxResponseHeaderBytes int

	disableCompression bool

	streamMx     sync.Mutex
	maxStreamID  quic.StreamID
	goAwayCtx    context.Context
	goAwayCancel context.CancelFunc

	qlogger qlogwriter.Recorder
	logger  *slog.Logger

	requestWriter *requestWriter
}

var _ http.RoundTripper = &ClientConn{}

func newClientConn(
	conn *quic.Conn,
	enableDatagrams bool,
	additionalSettings map[uint64]uint64,
	maxResponseHeaderBytes int,
	disableCompression bool,
	logger *slog.Logger,
) *ClientConn {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClientConn) OpenRequestStream(ctx context.Context) (*RequestStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientConn) openRequestStream(
	ctx context.Context,
	requestWriter *requestWriter,
	reqDone chan<- struct{},
	disableCompression bool,
	maxHeaderBytes int,
) (*RequestStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientConn) handleUnidirectionalStream(str *quic.ReceiveStream) {
	_ = "STUB: not implemented"
	return
}

func (c *ClientConn) handleControlStream(str *quic.ReceiveStream, fp *frameParser) {
	_ = "STUB: not implemented"
	return
}

func (c *ClientConn) onStreamsEmpty() { _ = "STUB: not implemented"; return }

func (c *ClientConn) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientConn) roundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientConn) ReceivedSettings() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *ClientConn) Settings() *Settings { _ = "STUB: not implemented"; return nil }

func (c *ClientConn) CloseWithError(code quic.ApplicationErrorCode, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClientConn) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type cancelingReader struct {
	r   io.Reader
	str *RequestStream
}

func (r *cancelingReader) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *ClientConn) sendRequestBody(str *RequestStream, body io.ReadCloser, contentLength int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClientConn) doRequest(req *http.Request, str *RequestStream) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RawClientConn struct {
	*ClientConn
}

func (c *RawClientConn) HandleUnidirectionalStream(str *quic.ReceiveStream) {
	_ = "STUB: not implemented"
	return
}

func (c *ClientConn) HandleBidirectionalStream(str *quic.Stream) { _ = "STUB: not implemented"; return }
