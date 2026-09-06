package quic

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"time"

	"github.com/quic-go/quic-go/internal/handshake"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/qlogwriter"
)

type StreamID = protocol.StreamID

type Version = protocol.Version

const (
	Version1 = protocol.Version1

	Version2 = protocol.Version2
)

func SupportedVersions() []Version { _ = "STUB: not implemented"; return nil }

type ClientToken struct {
	data []byte
	rtt  time.Duration
}

type TokenStore interface {
	Pop(key string) (token *ClientToken)

	Put(key string, token *ClientToken)
}

var Err0RTTRejected = errors.New("0-RTT rejected")

var ErrWouldBlock = errors.New("operation would block")

var ErrWriteLimitReached = errors.New("write limit reached")

var QUICVersionContextKey = handshake.QUICVersionContextKey

type StatelessResetKey [32]byte

type TokenGeneratorKey = handshake.TokenProtectorKey

type ConnectionID = protocol.ConnectionID

func ConnectionIDFromBytes(b []byte) ConnectionID {
	_ = "STUB: not implemented"
	return *new(ConnectionID)
}

type ConnectionIDGenerator interface {
	GenerateConnectionID() (ConnectionID, error)

	ConnectionIDLen() int
}

type Config struct {
	GetConfigForClient func(info *ClientInfo) (*Config, error)

	Versions []Version

	HandshakeIdleTimeout time.Duration

	MaxIdleTimeout time.Duration

	TokenStore TokenStore

	InitialStreamReceiveWindow uint64

	MaxStreamReceiveWindow uint64

	InitialConnectionReceiveWindow uint64

	MaxConnectionReceiveWindow uint64

	AllowConnectionWindowIncrease func(conn *Conn, delta uint64) bool

	MaxIncomingStreams int64

	MaxIncomingUniStreams int64

	KeepAlivePeriod time.Duration

	InitialPacketSize uint16

	DisablePathMTUDiscovery bool

	Allow0RTT bool

	EnableDatagrams bool

	EnableStreamResetPartialDelivery bool

	Tracer func(ctx context.Context, isClient bool, connID ConnectionID) qlogwriter.Trace
}

type ClientInfo struct {
	RemoteAddr net.Addr

	AddrVerified bool
}

type ConnectionState struct {
	TLS tls.ConnectionState

	SupportsDatagrams struct {
		Remote, Local bool
	}

	SupportsStreamResetPartialDelivery struct {
		Remote, Local bool
	}

	Used0RTT bool

	Version Version

	GSO bool
}
