package quic

import (
	"context"
	"sync"

	"github.com/quic-go/quic-go/internal/utils"
	"github.com/quic-go/quic-go/internal/utils/ringbuffer"
	"github.com/quic-go/quic-go/internal/wire"
)

const (
	maxDatagramSendQueueLen = 32
	maxDatagramRcvQueueLen  = 128
)

type datagramQueue struct {
	sendMx    sync.Mutex
	sendQueue ringbuffer.RingBuffer[*wire.DatagramFrame]
	sent      chan struct{}

	rcvMx    sync.Mutex
	rcvQueue [][]byte
	rcvd     chan struct{}

	closeErr error
	closed   chan struct{}

	hasData func()

	logger utils.Logger
}

func newDatagramQueue(hasData func(), logger utils.Logger) *datagramQueue {
	_ = "STUB: not implemented"
	return nil
}

func (h *datagramQueue) Add(f *wire.DatagramFrame) error { _ = "STUB: not implemented"; return nil }

func (h *datagramQueue) Peek() *wire.DatagramFrame { _ = "STUB: not implemented"; return nil }

func (h *datagramQueue) Pop() { _ = "STUB: not implemented"; return }

func (h *datagramQueue) HandleDatagramFrame(f *wire.DatagramFrame) {
	_ = "STUB: not implemented"
	return
}

func (h *datagramQueue) Receive(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *datagramQueue) CloseWithError(e error) { _ = "STUB: not implemented"; return }
