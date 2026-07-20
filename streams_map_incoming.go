package quic

import (
	"context"
	"sync"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

type incomingStream interface {
	closeForShutdown(error)
}

type incomingStreamEntry[T incomingStream] struct {
	stream       T
	shouldDelete bool
}

type incomingStreamsMap[T incomingStream] struct {
	mutex         sync.RWMutex
	newStreamChan chan struct{}

	streamType protocol.StreamType
	streams    map[protocol.StreamID]incomingStreamEntry[T]

	nextStreamToAccept protocol.StreamID
	nextStreamToOpen   protocol.StreamID
	maxStream          protocol.StreamID
	maxNumStreams      uint64

	newStream        func(protocol.StreamID) T
	queueMaxStreamID func(*wire.MaxStreamsFrame)

	closeErr error
}

func newIncomingStreamsMap[T incomingStream](
	streamType protocol.StreamType,
	newStream func(protocol.StreamID) T,
	maxStreams uint64,
	queueControlFrame func(wire.Frame),
	pers protocol.Perspective,
) *incomingStreamsMap[T] {
	_ = "STUB: not implemented"
	return nil
}

func (m *incomingStreamsMap[T]) AcceptStream(ctx context.Context) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (m *incomingStreamsMap[T]) GetOrOpenStream(id protocol.StreamID) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (m *incomingStreamsMap[T]) DeleteStream(id protocol.StreamID) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *incomingStreamsMap[T]) deleteStream(id protocol.StreamID) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *incomingStreamsMap[T]) CloseWithError(err error) { _ = "STUB: not implemented"; return }
