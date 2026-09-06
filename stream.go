package quic

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/quic-go/quic-go/internal/ackhandler"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

type deadlineError struct{}

func (deadlineError) Error() string   { _ = "STUB: not implemented"; return "" }
func (deadlineError) Temporary() bool { _ = "STUB: not implemented"; return false }
func (deadlineError) Timeout() bool   { _ = "STUB: not implemented"; return false }
func (deadlineError) Unwrap() error   { _ = "STUB: not implemented"; return nil }

var errDeadline net.Error = &deadlineError{}

type streamSender interface {
	onHasConnectionData()
	onHasStreamData(protocol.StreamID, *SendStream)
	onHasStreamRetransmission(protocol.StreamID, *SendStream)
	onHasStreamControlFrame(protocol.StreamID, streamControlFrameGetter)
	updateStreamPriority(protocol.StreamID)
	recordStreamPriorityUpdated(protocol.StreamID, int8, bool)

	onStreamCompleted(protocol.StreamID)
}

type uniStreamSender struct {
	streamSender
	onStreamCompletedImpl       func()
	onHasStreamControlFrameImpl func(protocol.StreamID, streamControlFrameGetter)
}

func (s *uniStreamSender) onHasStreamData(id protocol.StreamID, str *SendStream) {
	_ = "STUB: not implemented"
	return
}

func (s *uniStreamSender) onStreamCompleted(protocol.StreamID) { _ = "STUB: not implemented"; return }
func (s *uniStreamSender) onHasStreamControlFrame(id protocol.StreamID, str streamControlFrameGetter) {
	_ = "STUB: not implemented"
	return
}

var _ streamSender = &uniStreamSender{}

type Stream struct {
	receiveStr *ReceiveStream
	sendStr    *SendStream

	completedMutex         sync.Mutex
	sender                 streamSender
	receiveStreamCompleted bool
	sendStreamCompleted    bool
}

var (
	_ outgoingStream            = &Stream{}
	_ sendStreamFrameHandler    = &Stream{}
	_ receiveStreamFrameHandler = &Stream{}
)

func newStream(
	ctx context.Context,
	streamID protocol.StreamID,
	sender streamSender,
	flowController *streamFlowController,
	supportsResetStreamAt bool,
) *Stream {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) StreamID() StreamID { _ = "STUB: not implemented"; return *new(StreamID) }

func (s *Stream) SetPriority(urgency int8, incremental bool) { _ = "STUB: not implemented"; return }

func (s *Stream) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) Peek(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) WriteWithLimit(p []byte, limiter func(maxBytes int) int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Stream) TryWriteAll(p []byte) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) SetReliableBoundary() { _ = "STUB: not implemented"; return }

func (s *Stream) CancelWrite(errorCode StreamErrorCode) { _ = "STUB: not implemented"; return }

func (s *Stream) CancelRead(errorCode StreamErrorCode) { _ = "STUB: not implemented"; return }

func (s *Stream) SetReceiveFinalSizeCallback(callback func(int64)) {
	_ = "STUB: not implemented"
	return
}

func (s *Stream) Context() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

func (s *Stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Stream) handleResetStreamFrame(frame *wire.ResetStreamFrame, rcvTime monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) handleStreamFrame(frame *wire.StreamFrame, rcvTime monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) handleStopSendingFrame(frame *wire.StopSendingFrame) {
	_ = "STUB: not implemented"
	return
}

func (s *Stream) updateSendWindow(limit protocol.ByteCount) { _ = "STUB: not implemented"; return }

func (s *Stream) enableResetStreamAt() { _ = "STUB: not implemented"; return }

func (s *Stream) popStreamFrame(maxBytes protocol.ByteCount, v protocol.Version) (_ ackhandler.StreamFrame, _ *wire.StreamDataBlockedFrame, hasMore bool) {
	_ = "STUB: not implemented"
	return *new(ackhandler.StreamFrame), nil, false
}

func (s *Stream) getControlFrame(now monotime.Time) (_ ackhandler.Frame, ok, hasMore bool) {
	_ = "STUB: not implemented"
	return *new(ackhandler.Frame), false, false
}

func (s *Stream) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) closeForShutdown(err error) { _ = "STUB: not implemented"; return }

func (s *Stream) checkIfCompleted() { _ = "STUB: not implemented"; return }
