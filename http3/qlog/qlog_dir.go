package qlog

import (
	"context"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
)

const EventSchema = "urn:ietf:params:qlog:events:http3-12"

func DefaultConnectionTracer(ctx context.Context, isClient bool, connID quic.ConnectionID) qlogwriter.Trace {
	_ = "STUB: not implemented"
	return *new(qlogwriter.Trace)
}
