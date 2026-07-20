package quic

import (
	"errors"
	"sync"

	"github.com/quic-go/quic-go/internal/protocol"
	list "github.com/quic-go/quic-go/internal/utils/linkedlist"
)

type byteInterval struct {
	Start protocol.ByteCount
	End   protocol.ByteCount
}

var byteIntervalElementPool sync.Pool

func init() {
	byteIntervalElementPool = *list.NewPool[byteInterval]()
}

type frameSorterEntry struct {
	Data   []byte
	DoneCb func()
}

type frameSorter struct {
	queue   map[protocol.ByteCount]frameSorterEntry
	readPos protocol.ByteCount
	gaps    *list.List[byteInterval]
}

var errDuplicateStreamData = errors.New("duplicate stream data")

func newFrameSorter() *frameSorter { _ = "STUB: not implemented"; return nil }

func (s *frameSorter) Push(data []byte, offset protocol.ByteCount, doneCb func()) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *frameSorter) push(data []byte, offset protocol.ByteCount, doneCb func()) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *frameSorter) findStartGap(offset protocol.ByteCount) (*list.Element[byteInterval], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *frameSorter) findEndGap(startGap *list.Element[byteInterval], offset protocol.ByteCount) (*list.Element[byteInterval], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *frameSorter) deleteConsecutive(pos protocol.ByteCount) { _ = "STUB: not implemented"; return }

func (s *frameSorter) Pop() (protocol.ByteCount, []byte, func()) {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount), nil, nil
}

func (s *frameSorter) HasMoreData() bool { _ = "STUB: not implemented"; return false }

var errTooLittleData = errors.New("too little data")

func (s *frameSorter) Peek(offset protocol.ByteCount, p []byte) error {
	_ = "STUB: not implemented"
	return nil
}
