package qlogwriter

import (
	"runtime/debug"
	"time"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/qlogwriter/jsontext"
)

type ConnectionID = protocol.ConnectionID

var quicGoVersion = "(devel)"

func init() {
	if quicGoVersion != "(devel)" {
		return
	}
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	for _, d := range info.Deps {
		if d.Path == "github.com/quic-go/quic-go" {
			quicGoVersion = d.Version
			if d.Replace != nil {
				if len(d.Replace.Version) > 0 {
					quicGoVersion = d.Version
				} else {
					quicGoVersion += " (replaced)"
				}
			}
			break
		}
	}
}

type encoderHelper struct {
	enc *jsontext.Encoder
	err error
}

func (h *encoderHelper) WriteToken(t jsontext.Token) { _ = "STUB: not implemented"; return }

type traceHeader struct {
	VantagePointType string
	GroupID          *ConnectionID
	ReferenceTime    time.Time
	EventSchemas     []string
}

func (l traceHeader) Encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }
