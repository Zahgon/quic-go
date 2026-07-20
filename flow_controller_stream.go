package quic

import (
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
)

type streamFlowController struct {
	receiveFlowController

	bytesSent     protocol.ByteCount
	sendWindow    protocol.ByteCount
	lastBlockedAt protocol.ByteCount

	streamID protocol.StreamID

	connection *connectionFlowController

	receivedFinalOffset bool
}

func newStreamFlowController(
	streamID protocol.StreamID,
	cfc *connectionFlowController,
	receiveWindow protocol.ByteCount,
	maxReceiveWindow protocol.ByteCount,
	initialSendWindow protocol.ByteCount,
	rttStats *utils.RTTStats,
	logger utils.Logger,
) *streamFlowController {
	_ = "STUB: not implemented"
	return nil
}

func (c *streamFlowController) UpdateHighestReceived(offset protocol.ByteCount, final bool, now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *streamFlowController) AddBytesRead(n protocol.ByteCount) (hasStreamWindowUpdate, hasConnWindowUpdate bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (c *streamFlowController) Abandon() { _ = "STUB: not implemented"; return }

func (c *streamFlowController) UpdateSendWindow(offset protocol.ByteCount) (updated bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *streamFlowController) TryAddBytesSent(n protocol.ByteCount) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *streamFlowController) SendWindowSize() protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (c *streamFlowController) IsNewlyBlocked() bool { _ = "STUB: not implemented"; return false }

func (c *streamFlowController) isNewlyBlocked() (bool, protocol.ByteCount) {
	_ = "STUB: not implemented"
	return false, *new(protocol.ByteCount)
}

func (c *streamFlowController) shouldQueueWindowUpdate() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *streamFlowController) GetWindowUpdate(now monotime.Time) protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}
