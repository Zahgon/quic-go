package http3

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/quic-go/quic-go"
)

type Settingser interface {
	ReceivedSettings() <-chan struct{}

	Settings() *Settings
}

var errTooMuchData = errors.New("peer sent too much data")

type body struct {
	str *Stream

	remainingContentLength int64
	violatedContentLength  bool
	hasContentLength       bool
}

func newBody(str *Stream, contentLength int64) *body { _ = "STUB: not implemented"; return nil }

func (r *body) StreamID() quic.StreamID { _ = "STUB: not implemented"; return *new(quic.StreamID) }

func (r *body) checkContentLengthViolation() error { _ = "STUB: not implemented"; return nil }

func (r *body) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *body) Close() error { _ = "STUB: not implemented"; return nil }

type requestBody struct {
	body
	connCtx      context.Context
	rcvdSettings <-chan struct{}
	getSettings  func() *Settings
}

var _ io.ReadCloser = &requestBody{}

func newRequestBody(str *Stream, contentLength int64, connCtx context.Context, rcvdSettings <-chan struct{}, getSettings func() *Settings) *requestBody {
	_ = "STUB: not implemented"
	return nil
}

type hijackableBody struct {
	body body

	reqDone     chan<- struct{}
	reqDoneOnce sync.Once
}

var _ io.ReadCloser = &hijackableBody{}

func newResponseBody(str *Stream, contentLength int64, done chan<- struct{}) *hijackableBody {
	_ = "STUB: not implemented"
	return nil
}

func (r *hijackableBody) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *hijackableBody) requestDone() { _ = "STUB: not implemented"; return }

func (r *hijackableBody) Close() error { _ = "STUB: not implemented"; return nil }
