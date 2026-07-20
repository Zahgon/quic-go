package quic

import (
	"context"
	"sync"
	"time"

	"github.com/quic-go/quic-go/internal/ackhandler"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

type SendStream struct {
	mutex sync.Mutex

	numOutstandingFrames int64
	retransmissionQueue  []*wire.StreamFrame

	ctx       context.Context
	ctxCancel context.CancelCauseFunc

	streamID protocol.StreamID
	sender   streamSender

	reliableSize protocol.ByteCount
	writeOffset  protocol.ByteCount

	shutdownErr            error
	resetErr               *StreamError
	queuedResetStreamFrame *wire.ResetStreamFrame

	dataForWriting []byte
	nextFrame      *wire.StreamFrame

	nextFrameReserved bool

	supportsResetStreamAt bool
	finishedWriting       bool
	finSent               bool

	cancellationFlagged bool
	completed           bool

	writeChan chan struct{}
	writeOnce chan struct{}
	deadline  monotime.Time

	flowController *streamFlowController
}

var (
	_ streamControlFrameGetter = &SendStream{}
	_ outgoingStream           = &SendStream{}
	_ sendStreamFrameHandler   = &SendStream{}
)

func newSendStream(
	ctx context.Context,
	streamID protocol.StreamID,
	sender streamSender,
	flowController *streamFlowController,
	supportsResetStreamAt bool,
) *SendStream {
	_ = "STUB: not implemented"
	return nil
}

func (s *SendStream) StreamID() StreamID { _ = "STUB: not implemented"; return *new(StreamID) }

func (s *SendStream) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *SendStream) TryWriteAll(p []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SendStream) tryWriteAll(p []byte) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func (s *SendStream) write(p []byte) (bool, int, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (s *SendStream) canBufferStreamFrame() bool { _ = "STUB: not implemented"; return false }

func (s *SendStream) popStreamFrame(maxBytes protocol.ByteCount, v protocol.Version) (_ ackhandler.StreamFrame, _ *wire.StreamDataBlockedFrame, hasMore bool) {
	_ = "STUB: not implemented"
	return *new(ackhandler.StreamFrame), nil, false
}

func (s *SendStream) popNewOrRetransmittedStreamFrame(maxBytes protocol.ByteCount, v protocol.Version) (_ *wire.StreamFrame, _ *wire.StreamDataBlockedFrame, hasMoreData bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (s *SendStream) popNewStreamFrame(maxDataLen protocol.ByteCount) (_ *wire.StreamFrame, hasMoreData bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *SendStream) maybeGetRetransmission(maxBytes protocol.ByteCount, v protocol.Version) (*wire.StreamFrame, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *SendStream) getDataForWriting(f *wire.StreamFrame, maxBytes protocol.ByteCount) {
	_ = "STUB: not implemented"
	return
}

func (s *SendStream) isNewlyCompleted() bool { _ = "STUB: not implemented"; return false }

func (s *SendStream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *SendStream) SetReliableBoundary() { _ = "STUB: not implemented"; return }

func (s *SendStream) returnFramesToPool() { _ = "STUB: not implemented"; return }

func (s *SendStream) CancelWrite(errorCode StreamErrorCode) { _ = "STUB: not implemented"; return }

func (s *SendStream) enableResetStreamAt() { _ = "STUB: not implemented"; return }

func (s *SendStream) updateSendWindow(limit protocol.ByteCount) { _ = "STUB: not implemented"; return }

func (s *SendStream) handleStopSendingFrame(f *wire.StopSendingFrame) {
	_ = "STUB: not implemented"
	return
}

func (s *SendStream) getControlFrame(monotime.Time) (_ ackhandler.Frame, ok, hasMore bool) {
	_ = "STUB: not implemented"
	return *new(ackhandler.Frame), false, false
}

func (s *SendStream) reliableOffset() protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (s *SendStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *SendStream) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *SendStream) closeForShutdown(err error) { _ = "STUB: not implemented"; return }

func (s *SendStream) signalWrite() { _ = "STUB: not implemented"; return }

type sendStreamAckHandler SendStream

var _ ackhandler.FrameHandler = &sendStreamAckHandler{}

func (s *sendStreamAckHandler) OnAcked(f wire.Frame) { _ = "STUB: not implemented"; return }

func (s *sendStreamAckHandler) OnLost(f wire.Frame) { _ = "STUB: not implemented"; return }

type sendStreamResetStreamHandler SendStream

var _ ackhandler.FrameHandler = &sendStreamResetStreamHandler{}

func (s *sendStreamResetStreamHandler) OnAcked(f wire.Frame) { _ = "STUB: not implemented"; return }

func (s *sendStreamResetStreamHandler) OnLost(f wire.Frame) { _ = "STUB: not implemented"; return }
