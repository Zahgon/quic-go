package tools

import (
	"context"
	"io"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
)

func QlogTracer(logger io.Writer) qlogwriter.Trace {
	_ = "STUB: not implemented"
	return *new(qlogwriter.Trace)
}

func NewQlogConnectionTracer(logger io.Writer) func(ctx context.Context, isClient bool, connID quic.ConnectionID) qlogwriter.Trace {
	_ = "STUB: not implemented"
	return nil
}
