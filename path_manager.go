package quic

import (
	"net"
	"time"

	"github.com/quic-go/quic-go/internal/ackhandler"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
	"github.com/quic-go/quic-go/internal/wire"
)

type pathID int64

const invalidPathID pathID = -1

const maxPaths = 3

const pathTimeout = 5 * time.Second

type path struct {
	id             pathID
	addr           net.Addr
	lastPacketTime monotime.Time
	pathChallenge  [8]byte
	validated      bool
	rcvdNonProbing bool
}

type pathManager struct {
	nextPathID pathID

	paths []*path

	getConnID    func(pathID) (_ protocol.ConnectionID, ok bool)
	retireConnID func(pathID)

	logger utils.Logger
}

func newPathManager(
	getConnID func(pathID) (_ protocol.ConnectionID, ok bool),
	retireConnID func(pathID),
	logger utils.Logger,
) *pathManager {
	_ = "STUB: not implemented"
	return nil
}

func (pm *pathManager) HandlePacket(
	remoteAddr net.Addr,
	t monotime.Time,
	pathChallenge *wire.PathChallengeFrame,
	isNonProbing bool,
) (_ protocol.ConnectionID, _ []ackhandler.Frame, shouldSwitch bool) {
	_ = "STUB: not implemented"
	return *new(protocol.ConnectionID), nil, false
}

func (pm *pathManager) HandlePathResponseFrame(f *wire.PathResponseFrame) {
	_ = "STUB: not implemented"
	return
}

func (pm *pathManager) SwitchToPath(addr net.Addr) { _ = "STUB: not implemented"; return }

type pathManagerAckHandler pathManager

var _ ackhandler.FrameHandler = &pathManagerAckHandler{}

func (pm *pathManagerAckHandler) OnAcked(f wire.Frame) { _ = "STUB: not implemented"; return }

func (pm *pathManagerAckHandler) OnLost(f wire.Frame) { _ = "STUB: not implemented"; return }

func addrsEqual(addr1, addr2 net.Addr) bool { _ = "STUB: not implemented"; return false }
