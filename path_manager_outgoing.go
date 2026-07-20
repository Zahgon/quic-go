package quic

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/quic-go/quic-go/internal/ackhandler"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

var (
	ErrPathClosed = errors.New("path closed")

	ErrPathNotValidated = errors.New("path not yet validated")
)

var errPathDoesNotExist = errors.New("path does not exist")

type Path struct {
	id          pathID
	pathManager *pathManagerOutgoing
	tr          *Transport
	initialRTT  time.Duration

	enablePath func()
	validated  atomic.Bool
	abandon    chan struct{}
}

func (p *Path) Probe(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *Path) Switch() error { _ = "STUB: not implemented"; return nil }

func (p *Path) Close() error { _ = "STUB: not implemented"; return nil }

type pathOutgoing struct {
	pathChallenges [][8]byte
	tr             *Transport
	isValidated    bool
	probeSent      chan struct{}
	validated      chan struct{}
	enablePath     func()
}

func (p *pathOutgoing) ProbeSent() <-chan struct{} { _ = "STUB: not implemented"; return nil }
func (p *pathOutgoing) Validated() <-chan struct{} { _ = "STUB: not implemented"; return nil }

type pathManagerOutgoing struct {
	getConnID       func(pathID) (_ protocol.ConnectionID, ok bool)
	retireConnID    func(pathID)
	scheduleSending func()

	mx             sync.Mutex
	activePath     pathID
	pathsToProbe   []pathID
	paths          map[pathID]*pathOutgoing
	nextPathID     pathID
	pathToSwitchTo *pathOutgoing
}

func newPathManagerOutgoing(
	getConnID func(pathID) (_ protocol.ConnectionID, ok bool),
	retireConnID func(pathID),
	scheduleSending func(),
) *pathManagerOutgoing {
	_ = "STUB: not implemented"
	return nil
}

func (pm *pathManagerOutgoing) addPath(p *Path, enablePath func()) *pathOutgoing {
	_ = "STUB: not implemented"
	return nil
}

func (pm *pathManagerOutgoing) enqueueProbe(p *Path) { _ = "STUB: not implemented"; return }

func (pm *pathManagerOutgoing) removePath(id pathID) error { _ = "STUB: not implemented"; return nil }

func (pm *pathManagerOutgoing) removePathImpl(id pathID) error {
	_ = "STUB: not implemented"
	return nil
}

func (pm *pathManagerOutgoing) switchToPath(id pathID) error { _ = "STUB: not implemented"; return nil }

func (pm *pathManagerOutgoing) NewPath(t *Transport, initialRTT time.Duration, enablePath func()) *Path {
	_ = "STUB: not implemented"
	return nil
}

func (pm *pathManagerOutgoing) NextPathToProbe() (_ protocol.ConnectionID, _ ackhandler.Frame, _ *Transport, hasPath bool) {
	_ = "STUB: not implemented"
	return *new(protocol.ConnectionID), *new(ackhandler.Frame), nil, false
}

func (pm *pathManagerOutgoing) HandlePathResponseFrame(f *wire.PathResponseFrame) {
	_ = "STUB: not implemented"
	return
}

func (pm *pathManagerOutgoing) ShouldSwitchPath() (*Transport, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type pathManagerOutgoingAckHandler pathManagerOutgoing

var _ ackhandler.FrameHandler = &pathManagerOutgoingAckHandler{}

func (pm *pathManagerOutgoingAckHandler) OnAcked(wire.Frame) { _ = "STUB: not implemented"; return }

func (pm *pathManagerOutgoingAckHandler) OnLost(wire.Frame) { _ = "STUB: not implemented"; return }
