package quic

import (
	"github.com/quic-go/quic-go/internal/ackhandler"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

type framesToRetransmit struct {
	crypto []*wire.CryptoFrame
	other  []wire.Frame
}

type retransmissionQueue struct {
	initial   *framesToRetransmit
	handshake *framesToRetransmit
	appData   framesToRetransmit
}

func newRetransmissionQueue() *retransmissionQueue { _ = "STUB: not implemented"; return nil }

func (q *retransmissionQueue) addInitial(f wire.Frame) { _ = "STUB: not implemented"; return }

func (q *retransmissionQueue) addHandshake(f wire.Frame) { _ = "STUB: not implemented"; return }

func (q *retransmissionQueue) addAppData(f wire.Frame) { _ = "STUB: not implemented"; return }

func (q *retransmissionQueue) HasData(encLevel protocol.EncryptionLevel) bool {
	_ = "STUB: not implemented"
	//nolint:exhaustive // 0-RTT data is retransmitted in 1-RTT packets.
	return false
}

func (q *retransmissionQueue) GetFrame(encLevel protocol.EncryptionLevel, maxLen protocol.ByteCount, v protocol.Version) wire.Frame {
	_ = "STUB: not implemented"
	return *

	//nolint:exhaustive // 0-RTT data is retransmitted in 1-RTT packets.
	new(wire.Frame)
}

func (q *retransmissionQueue) DropPackets(encLevel protocol.EncryptionLevel) {
	_ = "STUB: not implemented"
	//nolint:exhaustive // Can only drop Initial and Handshake packet number space.
	return
}

func (q *retransmissionQueue) AckHandler(encLevel protocol.EncryptionLevel) ackhandler.FrameHandler {
	_ = "STUB: not implemented"
	return *new(ackhandler.FrameHandler)
}

type retransmissionQueueInitialAckHandler retransmissionQueue

func (q *retransmissionQueueInitialAckHandler) OnAcked(wire.Frame) {
	_ = "STUB: not implemented"
	return
}
func (q *retransmissionQueueInitialAckHandler) OnLost(f wire.Frame) {
	_ = "STUB: not implemented"
	return
}

type retransmissionQueueHandshakeAckHandler retransmissionQueue

func (q *retransmissionQueueHandshakeAckHandler) OnAcked(wire.Frame) {
	_ = "STUB: not implemented"
	return
}
func (q *retransmissionQueueHandshakeAckHandler) OnLost(f wire.Frame) {
	_ = "STUB: not implemented"
	return
}

type retransmissionQueueAppDataAckHandler retransmissionQueue

func (q *retransmissionQueueAppDataAckHandler) OnAcked(wire.Frame) {
	_ = "STUB: not implemented"
	return
}
func (q *retransmissionQueueAppDataAckHandler) OnLost(f wire.Frame) {
	_ = "STUB: not implemented"
	return
}
