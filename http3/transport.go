package http3

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"

	"github.com/quic-go/quic-go"
)

type Settings struct {
	EnableDatagrams bool

	EnableExtendedConnect bool

	Other map[uint64]uint64
}

type RoundTripOpt struct {
	OnlyCachedConn bool
}

type clientConn interface {
	OpenRequestStream(context.Context) (*RequestStream, error)
	RoundTrip(*http.Request) (*http.Response, error)
	handleUnidirectionalStream(*quic.ReceiveStream)
}

type roundTripperWithCount struct {
	cancel     context.CancelFunc
	dialing    chan struct{}
	dialErr    error
	conn       *quic.Conn
	clientConn clientConn

	useCount atomic.Int64
}

func (r *roundTripperWithCount) Close() error { _ = "STUB: not implemented"; return nil }

type Transport struct {
	TLSClientConfig *tls.Config

	QUICConfig *quic.Config

	Dial func(ctx context.Context, addr string, tlsCfg *tls.Config, cfg *quic.Config) (*quic.Conn, error)

	EnableDatagrams bool

	AdditionalSettings map[uint64]uint64

	MaxResponseHeaderBytes int

	DisableCompression bool

	Logger *slog.Logger

	mutex sync.Mutex

	initOnce sync.Once
	initErr  error

	newClientConn func(*quic.Conn) clientConn

	clients   map[string]*roundTripperWithCount
	transport *quic.Transport
	closed    bool
}

var (
	_ http.RoundTripper = &Transport{}
	_ io.Closer         = &Transport{}
)

var (
	ErrNoCachedConn = errors.New("http3: no cached connection was available")

	ErrTransportClosed = errors.New("http3: transport is closed")
)

func (t *Transport) init() error {
	if t.newClientConn == nil {
		t.newClientConn = func(conn *quic.Conn) clientConn {
			return newClientConn(
				conn,
				t.EnableDatagrams,
				t.AdditionalSettings,
				t.MaxResponseHeaderBytes,
				t.DisableCompression,
				t.Logger,
			)
		}
	}
	if t.QUICConfig == nil {
		t.QUICConfig = defaultQuicConfig.Clone()
		t.QUICConfig.EnableDatagrams = t.EnableDatagrams
	}
	if t.EnableDatagrams && !t.QUICConfig.EnableDatagrams {
		return errors.New("HTTP Datagrams enabled, but QUIC Datagrams disabled")
	}
	if len(t.QUICConfig.Versions) == 0 {
		t.QUICConfig = t.QUICConfig.Clone()
		t.QUICConfig.Versions = []quic.Version{quic.SupportedVersions()[0]}
	}
	if len(t.QUICConfig.Versions) != 1 {
		return errors.New("can only use a single QUIC version for dialing a HTTP/3 connection")
	}
	if t.QUICConfig.MaxIncomingStreams == 0 {
		t.QUICConfig.MaxIncomingStreams = -1
	}
	if t.Dial == nil {
		udpConn, err := net.ListenUDP("udp", nil)
		if err != nil {
			return err
		}
		t.transport = &quic.Transport{Conn: udpConn}
	}
	return nil
}

func (t *Transport) RoundTripOpt(req *http.Request, opt RoundTripOpt) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) roundTripOpt(req *http.Request, opt RoundTripOpt) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) doRoundTripOpt(req *http.Request, opt RoundTripOpt, isRetried bool) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func canRetryRequest(err error, req *http.Request) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) getClient(ctx context.Context, hostname string, onlyCached bool) (rtc *roundTripperWithCount, isReused bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (t *Transport) dial(ctx context.Context, hostname string) (*quic.Conn, clientConn, error) {
	_ = "STUB: not implemented"
	return nil, *new(clientConn), nil
}

func (t *Transport) resolveUDPAddr(ctx context.Context, network, addr string) (*net.UDPAddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) removeClient(hostname string) { _ = "STUB: not implemented"; return }

func (t *Transport) NewClientConn(conn *quic.Conn) *ClientConn {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transport) NewRawClientConn(conn *quic.Conn) *RawClientConn {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transport) Close() error { _ = "STUB: not implemented"; return nil }

func hostnameFromURL(url *url.URL) string { _ = "STUB: not implemented"; return "" }

func validMethod(method string) bool { _ = "STUB: not implemented"; return false }

func isNotToken(r rune) bool { _ = "STUB: not implemented"; return false }

func (t *Transport) CloseIdleConnections() { _ = "STUB: not implemented"; return }
