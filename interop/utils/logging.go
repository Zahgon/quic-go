package utils

import (
	"context"
	"io"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
)

func GetSSLKeyLog() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func NewQLOGConnectionTracer(_ context.Context, isClient bool, connID quic.ConnectionID) qlogwriter.Trace {
	_ = "STUB: not implemented"
	return *new(qlogwriter.Trace)
}
