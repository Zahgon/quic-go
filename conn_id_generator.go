package quic

import (
	"time"

	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

type connRunnerCallbacks struct {
	AddConnectionID    func(protocol.ConnectionID)
	RemoveConnectionID func(protocol.ConnectionID)
	ReplaceWithClosed  func([]protocol.ConnectionID, []byte, time.Duration)
}

type connRunners map[connRunner]connRunnerCallbacks

func (cr connRunners) AddConnectionID(id protocol.ConnectionID) { _ = "STUB: not implemented"; return }

func (cr connRunners) RemoveConnectionID(id protocol.ConnectionID) {
	_ = "STUB: not implemented"
	return
}

func (cr connRunners) ReplaceWithClosed(ids []protocol.ConnectionID, b []byte, expiry time.Duration) {
	_ = "STUB: not implemented"
	return
}

type connIDToRetire struct {
	t      monotime.Time
	connID protocol.ConnectionID
}

type connIDGenerator struct {
	generator   ConnectionIDGenerator
	highestSeq  uint64
	connRunners connRunners

	activeSrcConnIDs        map[uint64]protocol.ConnectionID
	connIDsToRetire         []connIDToRetire
	initialClientDestConnID *protocol.ConnectionID

	statelessResetter *statelessResetter

	queueControlFrame func(wire.Frame)
}

func newConnIDGenerator(
	runner connRunner,
	initialConnectionID protocol.ConnectionID,
	initialClientDestConnID *protocol.ConnectionID,
	statelessResetter *statelessResetter,
	callbacks connRunnerCallbacks,
	queueControlFrame func(wire.Frame),
	generator ConnectionIDGenerator,
) *connIDGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (m *connIDGenerator) SetMaxActiveConnIDs(limit uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *connIDGenerator) Retire(seq uint64, sentWithDestConnID protocol.ConnectionID, expiry monotime.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *connIDGenerator) queueConnIDForRetiring(connID protocol.ConnectionID, expiry monotime.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *connIDGenerator) issueNewConnID() error { _ = "STUB: not implemented"; return nil }

func (m *connIDGenerator) SetHandshakeComplete(connIDExpiry monotime.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *connIDGenerator) RemoveRetiredConnIDs(now monotime.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *connIDGenerator) RemoveAll() { _ = "STUB: not implemented"; return }

func (m *connIDGenerator) ReplaceWithClosed(connClose []byte, expiry time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *connIDGenerator) AddConnRunner(runner connRunner, r connRunnerCallbacks) {
	_ = "STUB: not implemented"
	return
}
