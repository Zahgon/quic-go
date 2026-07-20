package http3

import (
	"context"
	"io"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"

	"github.com/quic-go/qpack"
)

type datagramStream interface {
	io.ReadWriteCloser
	CancelRead(quic.StreamErrorCode)
	CancelWrite(quic.StreamErrorCode)
	StreamID() quic.StreamID
	Context() context.Context
	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	SendDatagram(b []byte) error
	ReceiveDatagram(ctx context.Context) ([]byte, error)

	QUICStream() *quic.Stream
}

type Stream struct {
	datagramStream
	conn        *rawConn
	frameParser *frameParser

	buf []byte

	bytesRemainingInFrame uint64

	qlogger qlogwriter.Recorder

	parseTrailer  func(io.Reader, *headersFrame) error
	parsedTrailer bool
}

func newStream(
	str datagramStream,
	conn *rawConn,
	trace *httptrace.ClientTrace,
	parseTrailer func(io.Reader, *headersFrame) error,
	qlogger qlogwriter.Recorder,
) *Stream {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) hasMoreData() bool { _ = "STUB: not implemented"; return false }

func (s *Stream) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) writeUnframed(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) StreamID() quic.StreamID { _ = "STUB: not implemented"; return *new(quic.StreamID) }

func (s *Stream) SendDatagram(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) ReceiveDatagram(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RequestStream struct {
	str *Stream

	responseBody io.ReadCloser

	decoder            *qpack.Decoder
	requestWriter      *requestWriter
	maxHeaderBytes     int
	reqDone            chan<- struct{}
	disableCompression bool
	response           *http.Response

	sentRequest   bool
	requestedGzip bool
	isConnect     bool
}

func newRequestStream(
	str *Stream,
	requestWriter *requestWriter,
	reqDone chan<- struct{},
	decoder *qpack.Decoder,
	disableCompression bool,
	maxHeaderBytes int,
	rsp *http.Response,
) *RequestStream {
	_ = "STUB: not implemented"
	return nil
}

func (s *RequestStream) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *RequestStream) StreamID() quic.StreamID {
	_ = "STUB: not implemented"
	return *new(quic.StreamID)
}

func (s *RequestStream) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *RequestStream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *RequestStream) CancelRead(errorCode quic.StreamErrorCode) {
	_ = "STUB: not implemented"
	return
}

func (s *RequestStream) CancelWrite(errorCode quic.StreamErrorCode) {
	_ = "STUB: not implemented"
	return
}

func (s *RequestStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *RequestStream) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *RequestStream) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *RequestStream) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *RequestStream) SendDatagram(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *RequestStream) ReceiveDatagram(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RequestStream) SendRequestHeader(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RequestStream) sendRequestHeader(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RequestStream) sendRequestTrailer(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RequestStream) ReadResponse() (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type tracingReader struct {
	io.Reader
	readFirst bool
	trace     *httptrace.ClientTrace
}

func (r *tracingReader) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
