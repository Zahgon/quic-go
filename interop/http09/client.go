package http09

import (
	"crypto/tls"
	"net/http"
	"sync"

	"github.com/quic-go/quic-go"
)

const MethodGet0RTT = "GET_0RTT"

type RoundTripper struct {
	mutex sync.Mutex

	TLSClientConfig *tls.Config
	QuicConfig      *quic.Config

	clients map[string]*client
}

var _ http.RoundTripper = &RoundTripper{}

func (r *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RoundTripper) Close() error { _ = "STUB: not implemented"; return nil }

type client struct {
	hostname string
	tlsConf  *tls.Config
	quicConf *quic.Config

	once    sync.Once
	conn    *quic.Conn
	dialErr error
}

func (c *client) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) doRequest(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) Close() error { _ = "STUB: not implemented"; return nil }

func hostnameFromRequest(req *http.Request) string { _ = "STUB: not implemented"; return "" }

func authorityAddr(scheme string, authority string) (addr string) {
	_ = "STUB: not implemented"
	return ""
}
