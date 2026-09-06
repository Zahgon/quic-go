package http3

import (
	"log/slog"
	"net/http"
	"time"
)

type HTTPStreamer interface {
	HTTPStream() *Stream
}

const maxSmallResponseSize = 4096

type responseWriter struct {
	str *Stream

	conn     *rawConn
	header   http.Header
	trailers map[string]struct{}
	buf      []byte
	status   int

	smallResponseBuf []byte

	contentLen     int64
	numWritten     int64
	headerComplete bool
	headerWritten  bool
	isHead         bool
	trailerWritten bool

	hijacked bool

	logger *slog.Logger
}

var (
	_ http.ResponseWriter = &responseWriter{}
	_ http.Flusher        = &responseWriter{}
	_ Settingser          = &responseWriter{}
	_ HTTPStreamer        = &responseWriter{}

	_ interface {
		SetReadDeadline(time.Time) error
		SetWriteDeadline(time.Time) error
		Flush()
		FlushError() error
	} = &responseWriter{}
)

func newResponseWriter(str *Stream, conn *rawConn, isHead bool, logger *slog.Logger) *responseWriter {
	_ = "STUB: not implemented"
	return nil
}

func (w *responseWriter) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (w *responseWriter) WriteHeader(status int) { _ = "STUB: not implemented"; return }

func (w *responseWriter) sniffContentType(p []byte) { _ = "STUB: not implemented"; return }

func (w *responseWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *responseWriter) doWrite(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *responseWriter) writeHeader(status int) error { _ = "STUB: not implemented"; return nil }

func (w *responseWriter) FlushError() error { _ = "STUB: not implemented"; return nil }

func (w *responseWriter) flushTrailers() { _ = "STUB: not implemented"; return }

func (w *responseWriter) Flush() { _ = "STUB: not implemented"; return }

func (w *responseWriter) declareTrailer(k string) { _ = "STUB: not implemented"; return }

func (w *responseWriter) writeTrailers() error { _ = "STUB: not implemented"; return nil }

func (w *responseWriter) HTTPStream() *Stream { _ = "STUB: not implemented"; return nil }

func (w *responseWriter) wasStreamHijacked() bool { _ = "STUB: not implemented"; return false }

func (w *responseWriter) ReceivedSettings() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (w *responseWriter) Settings() *Settings { _ = "STUB: not implemented"; return nil }

func (w *responseWriter) SetReadDeadline(deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *responseWriter) SetWriteDeadline(deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func bodyAllowedForStatus(status int) bool { _ = "STUB: not implemented"; return false }
