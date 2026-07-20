package qlog

import (
	"context"

	"github.com/quic-go/quic-go/qlogwriter"
)

const EventSchema = "urn:ietf:params:qlog:events:quic-12"

func DefaultConnectionTracer(_ context.Context, isClient bool, connID ConnectionID) qlogwriter.Trace {
	_ = "STUB: not implemented"
	return *new(qlogwriter.Trace)
}

func DefaultConnectionTracerWithSchemas(_ context.Context, isClient bool, connID ConnectionID, eventSchemas []string) qlogwriter.Trace {
	_ = "STUB: not implemented"
	return *new(qlogwriter.Trace)
}

func defaultConnectionTracerWithSchemas(isClient bool, connID ConnectionID, eventSchemas []string) qlogwriter.Trace {
	_ = "STUB: not implemented"
	return *new(qlogwriter.Trace)
}
