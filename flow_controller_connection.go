package quic

import (
	"sync"

	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
)

type connectionFlowController struct {
	receiveFlowController

	sendMutex     sync.Mutex
	bytesSent     protocol.ByteCount
	sendWindow    protocol.ByteCount
	lastBlockedAt protocol.ByteCount
}

func newConnectionFlowController(
	receiveWindow protocol.ByteCount,
	maxReceiveWindow protocol.ByteCount,
	allowWindowIncrease func(size protocol.ByteCount) bool,
	rttStats *utils.RTTStats,
	logger utils.Logger,
) *connectionFlowController {
	_ = "STUB: not implemented"
	return nil
}

func (c *connectionFlowController) IncrementHighestReceived(increment protocol.ByteCount, now monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *connectionFlowController) AddBytesRead(n protocol.ByteCount) (hasWindowUpdate bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *connectionFlowController) TryAddBytesSent(n protocol.ByteCount) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *connectionFlowController) UpdateSendWindow(offset protocol.ByteCount) (updated bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *connectionFlowController) SendWindowSize() protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (c *connectionFlowController) IsNewlyBlocked() (bool, protocol.ByteCount) {
	_ = "STUB: not implemented"
	return false, *new(protocol.ByteCount)
}

func (c *connectionFlowController) GetWindowUpdate(now monotime.Time) protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (c *connectionFlowController) EnsureMinimumWindowSize(inc protocol.ByteCount, now monotime.Time) {
	_ = "STUB: not implemented"
	return
}

func (c *connectionFlowController) Reset() error { _ = "STUB: not implemented"; return nil }
