package quic

import (
	"github.com/quic-go/quic-go/internal/ackhandler"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
	"github.com/quic-go/quic-go/internal/wire"
	"github.com/quic-go/quic-go/qlogwriter"
)

const (
	maxMTUDiff protocol.ByteCount = 20

	mtuProbeDelay = 5

	maxLostMTUProbes = 3
)

type mtuFinder struct {
	lastProbeTime monotime.Time

	rttStats *utils.RTTStats

	inFlight protocol.ByteCount
	min      protocol.ByteCount

	lost             [maxLostMTUProbes]protocol.ByteCount
	lastProbeWasLost bool

	generation uint8

	qlogger qlogwriter.Recorder
}

func newMTUDiscoverer(
	rttStats *utils.RTTStats,
	start, max protocol.ByteCount,
	qlogger qlogwriter.Recorder,
) *mtuFinder {
	_ = "STUB: not implemented"
	return nil
}

func (f *mtuFinder) init(start, max protocol.ByteCount) {
	f.min = start
	for i := range f.lost {
		if i == 0 {
			f.lost[i] = max
			continue
		}
		f.lost[i] = protocol.InvalidByteCount
	}
}

func (f *mtuFinder) done() bool { _ = "STUB: not implemented"; return false }

func (f *mtuFinder) max() protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (f *mtuFinder) Start(now monotime.Time) { _ = "STUB: not implemented"; return }

func (f *mtuFinder) ShouldSendProbe(now monotime.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *mtuFinder) GetPing(now monotime.Time) (ackhandler.Frame, protocol.ByteCount) {
	_ = "STUB: not implemented"
	return *new(ackhandler.Frame), *new(protocol.ByteCount)
}

func (f *mtuFinder) CurrentSize() protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (f *mtuFinder) Reset(now monotime.Time, start, max protocol.ByteCount) {
	_ = "STUB: not implemented"
	return
}

type mtuFinderAckHandler struct {
	*mtuFinder
	generation uint8
}

var _ ackhandler.FrameHandler = &mtuFinderAckHandler{}

func (h *mtuFinderAckHandler) OnAcked(wire.Frame) { _ = "STUB: not implemented"; return }

func (h *mtuFinderAckHandler) OnLost(wire.Frame) { _ = "STUB: not implemented"; return }
