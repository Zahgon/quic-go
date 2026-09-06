package quic

import (
	"sync"

	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
)

type receiveFlowController struct {
	//nolint:structcheck // The mutex is used both by the stream and the connection flow controller
	mutex                sync.Mutex
	bytesRead            protocol.ByteCount
	highestReceived      protocol.ByteCount
	receiveWindow        protocol.ByteCount
	receiveWindowSize    protocol.ByteCount
	maxReceiveWindowSize protocol.ByteCount

	allowWindowIncrease func(size protocol.ByteCount) bool

	epochStartTime   monotime.Time
	epochStartOffset protocol.ByteCount
	rttStats         *utils.RTTStats

	logger utils.Logger
}

func (c *receiveFlowController) addBytesRead(n protocol.ByteCount) {
	_ = "STUB: not implemented"
	return
}

func (c *receiveFlowController) hasWindowUpdate() bool { _ = "STUB: not implemented"; return false }

func (c *receiveFlowController) getWindowUpdate(now monotime.Time) protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (c *receiveFlowController) maybeAdjustWindowSize(now monotime.Time) {
	_ = "STUB: not implemented"
	return
}

func (c *receiveFlowController) startNewAutoTuningEpoch(now monotime.Time) {
	_ = "STUB: not implemented"
	return
}

func (c *receiveFlowController) checkFlowControlViolation() bool {
	_ = "STUB: not implemented"
	return false
}
