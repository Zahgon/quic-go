package http09

import (
	"io"
	"net/http"

	"github.com/quic-go/quic-go"
)

const NextProto = "hq-interop"

type responseWriter struct {
	io.Writer
	headers http.Header
}

var _ http.ResponseWriter = &responseWriter{}

func (w *responseWriter) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (w *responseWriter) WriteHeader(int) { _ = "STUB: not implemented"; return }

type Server struct {
	Handler *http.ServeMux
}

func (s *Server) ServeListener(ln *quic.EarlyListener) error { _ = "STUB: not implemented"; return nil }

func (s *Server) handleConn(conn *quic.Conn) { _ = "STUB: not implemented"; return }

func (s *Server) handleStream(str *quic.Stream) error { _ = "STUB: not implemented"; return nil }
