package qlog

import (
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter/jsontext"
)

type encoderHelper struct {
	enc *jsontext.Encoder
	err error
}

func (h *encoderHelper) WriteToken(t jsontext.Token) { _ = "STUB: not implemented"; return }

type RawInfo struct {
	Length        int
	PayloadLength int
}

func (i RawInfo) HasValues() bool { _ = "STUB: not implemented"; return false }

func (i RawInfo) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type FrameParsed struct {
	StreamID quic.StreamID
	Raw      RawInfo
	Frame    Frame
}

func (e FrameParsed) Name() string { _ = "STUB: not implemented"; return "" }

func (e FrameParsed) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type FrameCreated struct {
	StreamID quic.StreamID
	Raw      RawInfo
	Frame    Frame
}

func (e FrameCreated) Name() string { _ = "STUB: not implemented"; return "" }

func (e FrameCreated) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type DatagramCreated struct {
	QuarterStreamID uint64
	Raw             RawInfo
}

func (e DatagramCreated) Name() string { _ = "STUB: not implemented"; return "" }

func (e DatagramCreated) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type DatagramParsed struct {
	QuarterStreamID uint64
	Raw             RawInfo
}

func (e DatagramParsed) Name() string { _ = "STUB: not implemented"; return "" }

func (e DatagramParsed) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}
