package http3

import (
	"errors"
	"io"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
	"github.com/quic-go/quic-go/quicvarint"
)

type FrameType uint64

type frame any

var errPriorityUpdateForPush = errors.New("http3: PRIORITY_UPDATE frame for push")

const frameHeaderLen = 16

type countingByteReader struct {
	quicvarint.Reader
	NumRead int
}

func (r *countingByteReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *countingByteReader) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *countingByteReader) Reset() { _ = "STUB: not implemented"; return }

type frameParser struct {
	r         io.Reader
	streamID  quic.StreamID
	closeConn func(quic.ApplicationErrorCode, string) error
}

func (p *frameParser) ParseNext(qlogger qlogwriter.Recorder) (frame, error) {
	_ = "STUB: not implemented"
	return *new(frame), nil
}

type dataFrame struct {
	Length uint64
}

func (f *dataFrame) Append(b []byte) []byte { _ = "STUB: not implemented"; return nil }

type headersFrame struct {
	Length    uint64
	headerLen int
}

func (f *headersFrame) Append(b []byte) []byte { _ = "STUB: not implemented"; return nil }

const (
	settingMaxFieldSectionSize = 0x6

	settingExtendedConnect = 0x8

	settingDatagram = 0x33
)

type settingsFrame struct {
	MaxFieldSectionSize int64

	Datagram        bool
	ExtendedConnect bool
	Other           map[uint64]uint64
}

const maxSettingsFrameSize = 8 << 10

func parseSettingsFrame(r *countingByteReader, l uint64, streamID quic.StreamID, qlogger qlogwriter.Recorder) (*settingsFrame, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *settingsFrame) Append(b []byte) []byte { _ = "STUB: not implemented"; return nil }

type goAwayFrame struct {
	StreamID quic.StreamID
}

func parseGoAwayFrame(r *countingByteReader, l uint64, streamID quic.StreamID, qlogger qlogwriter.Recorder) (*goAwayFrame, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *goAwayFrame) Append(b []byte) []byte { _ = "STUB: not implemented"; return nil }

type priorityUpdateFrame struct {
	ElementID          uint64
	PriorityFieldValue string
}

const maxPriorityUpdateFrameSize = 4 << 10

func parsePriorityUpdateFrame(r *countingByteReader, l uint64, streamID quic.StreamID, qlogger qlogwriter.Recorder) (*priorityUpdateFrame, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
