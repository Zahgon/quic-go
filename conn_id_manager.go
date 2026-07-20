package quic

import (
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
	"github.com/quic-go/quic-go/internal/wire"
)

type newConnID struct {
	SequenceNumber      uint64
	ConnectionID        protocol.ConnectionID
	StatelessResetToken protocol.StatelessResetToken
}

type connIDManager struct {
	queue []newConnID

	highestProbingID uint64
	pathProbing      map[pathID]newConnID

	handshakeComplete         bool
	activeSequenceNumber      uint64
	highestRetired            uint64
	activeConnectionID        protocol.ConnectionID
	activeStatelessResetToken *protocol.StatelessResetToken

	rand                   utils.Rand
	packetsSinceLastChange uint32
	packetsPerConnectionID uint32

	addStatelessResetToken    func(protocol.StatelessResetToken)
	removeStatelessResetToken func(protocol.StatelessResetToken)
	queueControlFrame         func(wire.Frame)

	closed bool
}

func newConnIDManager(
	initialDestConnID protocol.ConnectionID,
	addStatelessResetToken func(protocol.StatelessResetToken),
	removeStatelessResetToken func(protocol.StatelessResetToken),
	queueControlFrame func(wire.Frame),
) *connIDManager {
	_ = "STUB: not implemented"
	return nil
}

func (h *connIDManager) AddFromPreferredAddress(connID protocol.ConnectionID, resetToken protocol.StatelessResetToken) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *connIDManager) Add(f *wire.NewConnectionIDFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *connIDManager) add(f *wire.NewConnectionIDFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *connIDManager) addConnectionID(seq uint64, connID protocol.ConnectionID, resetToken protocol.StatelessResetToken) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *connIDManager) updateConnectionID() { _ = "STUB: not implemented"; return }

func (h *connIDManager) Close() { _ = "STUB: not implemented"; return }

func (h *connIDManager) ChangeInitialConnID(newConnID protocol.ConnectionID) {
	_ = "STUB: not implemented"
	return
}

func (h *connIDManager) SetStatelessResetToken(token protocol.StatelessResetToken) {
	_ = "STUB: not implemented"
	return
}

func (h *connIDManager) SentPacket() { _ = "STUB: not implemented"; return }

func (h *connIDManager) shouldUpdateConnID() bool { _ = "STUB: not implemented"; return false }

func (h *connIDManager) Get() protocol.ConnectionID {
	_ = "STUB: not implemented"
	return *new(protocol.ConnectionID)
}

func (h *connIDManager) SetHandshakeComplete() { _ = "STUB: not implemented"; return }

func (h *connIDManager) GetConnIDForPath(id pathID) (protocol.ConnectionID, bool) {
	_ = "STUB: not implemented"
	return *new(protocol.ConnectionID), false
}

func (h *connIDManager) RetireConnIDForPath(pathID pathID) { _ = "STUB: not implemented"; return }

func (h *connIDManager) IsActiveStatelessResetToken(token protocol.StatelessResetToken) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *connIDManager) assertNotClosed() { _ = "STUB: not implemented"; return }
