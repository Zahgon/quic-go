package quic

import (
	"context"
	"sync"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

type outgoingStream interface {
	updateSendWindow(protocol.ByteCount)
	enableResetStreamAt()
	closeForShutdown(error)
}

type outgoingStreamsMap[T outgoingStream] struct {
	mutex sync.RWMutex

	streamType protocol.StreamType
	streams    map[protocol.StreamID]T

	openQueue []chan struct{}

	nextStream  protocol.StreamID
	maxStream   protocol.StreamID
	blockedSent bool

	newStream            func(protocol.StreamID) T
	queueStreamIDBlocked func(*wire.StreamsBlockedFrame)

	closeErr error
}

func newOutgoingStreamsMap[T outgoingStream](
	streamType protocol.StreamType,
	newStream func(protocol.StreamID) T,
	queueControlFrame func(wire.Frame),
	pers protocol.Perspective,
) *outgoingStreamsMap[T] {
	_ = "STUB: not implemented"
	return nil
}

func (m *outgoingStreamsMap[T]) OpenStream() (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (m *outgoingStreamsMap[T]) OpenStreamSync(ctx context.Context) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (m *outgoingStreamsMap[T]) openStream() T { _ = "STUB: not implemented"; return *new(T) }

func (m *outgoingStreamsMap[T]) maybeSendBlockedFrame() { _ = "STUB: not implemented"; return }

func (m *outgoingStreamsMap[T]) GetStream(id protocol.StreamID) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (m *outgoingStreamsMap[T]) DeleteStream(id protocol.StreamID) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *outgoingStreamsMap[T]) SetMaxStream(id protocol.StreamID) {
	_ = "STUB: not implemented"
	return
}

func (m *outgoingStreamsMap[T]) UpdateSendWindow(limit protocol.ByteCount) {
	_ = "STUB: not implemented"
	return
}

func (m *outgoingStreamsMap[T]) EnableResetStreamAt() { _ = "STUB: not implemented"; return }

func (m *outgoingStreamsMap[T]) maybeUnblockOpenSync() { _ = "STUB: not implemented"; return }

func (m *outgoingStreamsMap[T]) CloseWithError(err error) { _ = "STUB: not implemented"; return }
