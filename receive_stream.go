package quic

import (
	"sync"
	"time"

	"github.com/quic-go/quic-go/internal/ackhandler"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/qerr"
	"github.com/quic-go/quic-go/internal/wire"
)

type ReceiveStream struct {
	mutex sync.Mutex

	streamID protocol.StreamID

	sender streamSender

	frameQueue               *frameSorter
	finalOffset              protocol.ByteCount
	receiveFinalSizeCallback func(int64)

	currentFrame       []byte
	currentFrameDone   func()
	readPosInFrame     int
	currentFrameIsLast bool

	queuedStopSending   bool
	queuedMaxStreamData bool

	errorRead           bool
	completed           bool
	cancelledRemotely   bool
	cancelledLocally    bool
	cancelErr           *StreamError
	closeForShutdownErr error

	readPos      protocol.ByteCount
	reliableSize protocol.ByteCount

	readChan chan struct{}
	readOnce chan struct{}
	deadline monotime.Time

	flowController *streamFlowController
}

var (
	_ streamControlFrameGetter  = &ReceiveStream{}
	_ receiveStreamFrameHandler = &ReceiveStream{}
)

func newReceiveStream(
	streamID protocol.StreamID,
	sender streamSender,
	flowController *streamFlowController,
) *ReceiveStream {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReceiveStream) StreamID() StreamID { _ = "STUB: not implemented"; return *new(StreamID) }

func (s *ReceiveStream) SetReceiveFinalSizeCallback(callback func(int64)) {
	_ = "STUB: not implemented"
	return
}

func (s *ReceiveStream) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *ReceiveStream) isNewlyCompleted() bool { _ = "STUB: not implemented"; return false }

func (s *ReceiveStream) readImpl(p []byte) (hasStreamWindowUpdate bool, hasConnWindowUpdate bool, _ int, _ error) {
	_ = "STUB: not implemented"
	return false, false, 0, nil
}

func (s *ReceiveStream) isRemoteCancellationEffective() bool {
	_ = "STUB: not implemented"
	return false
}

func (s *ReceiveStream) Peek(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *ReceiveStream) peekImpl(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *ReceiveStream) dequeueNextFrame() { _ = "STUB: not implemented"; return }

func (s *ReceiveStream) CancelRead(errorCode StreamErrorCode) { _ = "STUB: not implemented"; return }

func (s *ReceiveStream) cancelReadImpl(errorCode qerr.StreamErrorCode) (queuedNewControlFrame bool) {
	_ = "STUB: not implemented"
	return false
}

func (s *ReceiveStream) handleStreamFrame(frame *wire.StreamFrame, now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReceiveStream) handleStreamFrameImpl(frame *wire.StreamFrame, now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReceiveStream) handleResetStreamFrame(frame *wire.ResetStreamFrame, now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReceiveStream) handleResetStreamFrameImpl(frame *wire.ResetStreamFrame, now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReceiveStream) takeReceiveFinalSizeCallback() (int64, func(int64)) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *ReceiveStream) getControlFrame(now monotime.Time) (_ ackhandler.Frame, ok, hasMore bool) {
	_ = "STUB: not implemented"
	return *new(ackhandler.Frame), false, false
}

func (s *ReceiveStream) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *ReceiveStream) closeForShutdown(err error) { _ = "STUB: not implemented"; return }

func (s *ReceiveStream) signalRead() { _ = "STUB: not implemented"; return }
