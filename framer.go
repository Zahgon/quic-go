package quic

import (
	"sync"

	"github.com/quic-go/quic-go/internal/ackhandler"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils/ringbuffer"
	"github.com/quic-go/quic-go/internal/wire"
)

const (
	maxPathResponses = 256
	maxControlFrames = 16 << 10
)

const maxStreamControlFrameSize = 25

type streamFrameGetter interface {
	popStreamFrame(protocol.ByteCount, protocol.Version) (ackhandler.StreamFrame, *wire.StreamDataBlockedFrame, bool)
}

type streamControlFrameGetter interface {
	getControlFrame(monotime.Time) (_ ackhandler.Frame, ok, hasMore bool)
}

type framer struct {
	mutex sync.Mutex

	activeStreams            map[protocol.StreamID]streamFrameGetter
	streamQueue              ringbuffer.RingBuffer[protocol.StreamID]
	streamsWithControlFrames map[protocol.StreamID]streamControlFrameGetter

	controlFrameMutex          sync.Mutex
	controlFrames              []wire.Frame
	pathResponses              []*wire.PathResponseFrame
	connFlowController         *connectionFlowController
	queuedTooManyControlFrames bool
}

func newFramer(connFlowController *connectionFlowController) *framer {
	_ = "STUB: not implemented"
	return nil
}

func (f *framer) HasData() bool { _ = "STUB: not implemented"; return false }

func (f *framer) QueueControlFrame(frame wire.Frame) { _ = "STUB: not implemented"; return }

func (f *framer) Append(
	frames []ackhandler.Frame,
	streamFrames []ackhandler.StreamFrame,
	maxLen protocol.ByteCount,
	now monotime.Time,
	v protocol.Version,
) ([]ackhandler.Frame, []ackhandler.StreamFrame, protocol.ByteCount) {
	_ = "STUB: not implemented"
	return nil, nil, *new(protocol.ByteCount)
}

func (f *framer) appendControlFrames(
	frames []ackhandler.Frame,
	maxLen protocol.ByteCount,
	now monotime.Time,
	v protocol.Version,
) ([]ackhandler.Frame, protocol.ByteCount) {
	_ = "STUB: not implemented"
	return nil, *new(protocol.ByteCount)
}

func (f *framer) QueuedTooManyControlFrames() bool { _ = "STUB: not implemented"; return false }

func (f *framer) AddActiveStream(id protocol.StreamID, str streamFrameGetter) {
	_ = "STUB: not implemented"
	return
}

func (f *framer) AddStreamWithControlFrames(id protocol.StreamID, str streamControlFrameGetter) {
	_ = "STUB: not implemented"
	return
}

func (f *framer) RemoveActiveStream(id protocol.StreamID) { _ = "STUB: not implemented"; return }

func (f *framer) getNextStreamFrame(maxLen protocol.ByteCount, v protocol.Version) (ackhandler.StreamFrame, *wire.StreamDataBlockedFrame) {
	_ = "STUB: not implemented"
	return *new(ackhandler.StreamFrame), nil
}

func (f *framer) Handle0RTTRejection() { _ = "STUB: not implemented"; return }
