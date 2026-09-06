package quic

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/quic-go/quic-go/internal/protocol"
)

var generateConnectionIDForInitial = protocol.GenerateConnectionIDForInitial

func DialAddr(ctx context.Context, addr string, tlsConf *tls.Config, conf *Config) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DialAddrEarly(ctx context.Context, addr string, tlsConf *tls.Config, conf *Config) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DialEarly(ctx context.Context, c net.PacketConn, addr net.Addr, tlsConf *tls.Config, conf *Config) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Dial(ctx context.Context, c net.PacketConn, addr net.Addr, tlsConf *tls.Config, conf *Config) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setupTransport(c net.PacketConn, tlsConf *tls.Config, createdPacketConn bool) (*Transport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
