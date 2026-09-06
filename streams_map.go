package quic

import (
	"context"
	"sync"

	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

type StreamLimitReachedError struct{}

func (e StreamLimitReachedError) Error() string { _ = "STUB: not implemented"; return "" }

type streamsMap struct {
	ctx         context.Context
	perspective protocol.Perspective

	maxIncomingBidiStreams uint64
	maxIncomingUniStreams  uint64

	sender            streamSender
	queueControlFrame func(wire.Frame)
	newFlowController func(protocol.StreamID) *streamFlowController

	mutex                 sync.Mutex
	outgoingBidiStreams   *outgoingStreamsMap[*Stream]
	outgoingUniStreams    *outgoingStreamsMap[*SendStream]
	incomingBidiStreams   *incomingStreamsMap[*Stream]
	incomingUniStreams    *incomingStreamsMap[*ReceiveStream]
	reset                 bool
	supportsResetStreamAt bool
}

func newStreamsMap(
	ctx context.Context,
	sender streamSender,
	queueControlFrame func(wire.Frame),
	newFlowController func(protocol.StreamID) *streamFlowController,
	maxIncomingBidiStreams uint64,
	maxIncomingUniStreams uint64,
	perspective protocol.Perspective,
) *streamsMap {
	_ = "STUB: not implemented"
	return nil
}

func (m *streamsMap) initMaps() { _ = "STUB: not implemented"; return }

func (m *streamsMap) OpenStream() (*Stream, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *streamsMap) OpenStreamSync(ctx context.Context) (*Stream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *streamsMap) OpenUniStream() (*SendStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *streamsMap) OpenUniStreamSync(ctx context.Context) (*SendStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *streamsMap) AcceptStream(ctx context.Context) (*Stream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *streamsMap) AcceptUniStream(ctx context.Context) (*ReceiveStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *streamsMap) DeleteStream(id protocol.StreamID) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *streamsMap) HandleMaxStreamsFrame(f *wire.MaxStreamsFrame) {
	_ = "STUB: not implemented"
	return
}

type sendStreamFrameHandler interface {
	updateSendWindow(protocol.ByteCount)
	handleStopSendingFrame(*wire.StopSendingFrame)
}

func (m *streamsMap) getSendStream(id protocol.StreamID) (sendStreamFrameHandler, error) {
	_ = "STUB: not implemented"
	return *new(sendStreamFrameHandler), nil
}

func (m *streamsMap) HandleMaxStreamDataFrame(f *wire.MaxStreamDataFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *streamsMap) HandleStopSendingFrame(f *wire.StopSendingFrame) error {
	_ = "STUB: not implemented"
	return nil
}

type receiveStreamFrameHandler interface {
	handleResetStreamFrame(*wire.ResetStreamFrame, monotime.Time) error
	handleStreamFrame(*wire.StreamFrame, monotime.Time) error
}

func (m *streamsMap) getReceiveStream(id protocol.StreamID) (receiveStreamFrameHandler, error) {
	_ = "STUB: not implemented"
	return *new(receiveStreamFrameHandler), nil
}

func (m *streamsMap) HandleStreamDataBlockedFrame(f *wire.StreamDataBlockedFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *streamsMap) HandleResetStreamFrame(f *wire.ResetStreamFrame, rcvTime monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *streamsMap) HandleStreamFrame(f *wire.StreamFrame, rcvTime monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *streamsMap) HandleTransportParameters(p *wire.TransportParameters) {
	_ = "STUB: not implemented"
	return
}

func (m *streamsMap) CloseWithError(err error) { _ = "STUB: not implemented"; return }

func (m *streamsMap) ResetFor0RTT() { _ = "STUB: not implemented"; return }

func (m *streamsMap) UseResetMaps() { _ = "STUB: not implemented"; return }
