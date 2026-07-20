package http3

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
)

const NextProtoH3 = "h3"

type StreamType uint64

const (
	streamTypeControlStream      = 0
	streamTypePushStream         = 1
	streamTypeQPACKEncoderStream = 2
	streamTypeQPACKDecoderStream = 3
)

type QUICListener interface {
	Accept(context.Context) (*quic.Conn, error)
	Addr() net.Addr
	io.Closer
}

var _ QUICListener = &quic.EarlyListener{}

func ConfigureTLSConfig(tlsConf *tls.Config) *tls.Config { _ = "STUB: not implemented"; return nil }

type contextKey struct {
	name string
}

func (k *contextKey) String() string { _ = "STUB: not implemented"; return "" }

var ServerContextKey = &contextKey{"http3-server"}

var RemoteAddrContextKey = &contextKey{"remote-addr"}

type listener struct {
	ln   *QUICListener
	port int

	createdLocally bool
}

type Server struct {
	Addr string

	Port int

	TLSConfig *tls.Config

	QUICConfig *quic.Config

	Handler http.Handler

	EnableDatagrams bool

	MaxHeaderBytes int

	AdditionalSettings map[uint64]uint64

	IdleTimeout time.Duration

	ConnContext func(ctx context.Context, c *quic.Conn) context.Context

	Logger *slog.Logger

	mutex     sync.RWMutex
	listeners []listener

	closed           bool
	closeCtx         context.Context
	closeCancel      context.CancelFunc
	graceCtx         context.Context
	graceCancel      context.CancelFunc
	connCount        atomic.Int64
	connHandlingDone chan struct{}

	altSvcHeader string
}

func (s *Server) ListenAndServe() error { _ = "STUB: not implemented"; return nil }

func (s *Server) ListenAndServeTLS(certFile, keyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) Serve(conn net.PacketConn) error { _ = "STUB: not implemented"; return nil }

func (s *Server) init() {
	if s.closeCtx == nil {
		s.closeCtx, s.closeCancel = context.WithCancel(context.Background())
		s.graceCtx, s.graceCancel = context.WithCancel(s.closeCtx)
	}
	s.connHandlingDone = make(chan struct{}, 1)
}

func (s *Server) decreaseConnCount() { _ = "STUB: not implemented"; return }

func (s *Server) ServeQUICConn(conn *quic.Conn) error { _ = "STUB: not implemented"; return nil }

func (s *Server) ServeListener(ln QUICListener) error { _ = "STUB: not implemented"; return nil }

func (s *Server) serveListener(ln QUICListener) error { _ = "STUB: not implemented"; return nil }

var errServerWithoutTLSConfig = errors.New("use of http3.Server without TLSConfig")

func (s *Server) setupListenerForConn(tlsConf *tls.Config, conn net.PacketConn) (*QUICListener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractPort(addr string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Server) generateAltSvcHeader() { _ = "STUB: not implemented"; return }

func (s *Server) addListener(l *QUICListener, createdLocally bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) removeListener(l *QUICListener) { _ = "STUB: not implemented"; return }

func (s *Server) NewRawServerConn(conn *quic.Conn) (*RawServerConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) newRawServerConn(conn *quic.Conn) (*RawServerConn, *quic.SendStream, qlogwriter.Recorder, error) {
	_ = "STUB: not implemented"
	return nil, nil, *new(qlogwriter.Recorder), nil
}

func (s *Server) handleConn(conn *quic.Conn) error { _ = "STUB: not implemented"; return nil }

func (s *Server) maxHeaderBytes() int { _ = "STUB: not implemented"; return 0 }

func (s *Server) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Server) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var ErrNoAltSvcPort = errors.New("no port can be announced, specify it explicitly using Server.Port or Server.Addr")

func (s *Server) SetQUICHeaders(hdr http.Header) error { _ = "STUB: not implemented"; return nil }

func ListenAndServeQUIC(addr, certFile, keyFile string, handler http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}
