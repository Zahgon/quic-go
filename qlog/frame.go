package qlog

import (
	"github.com/quic-go/quic-go/internal/wire"
	"github.com/quic-go/quic-go/qlogwriter/jsontext"
)

type Frame struct {
	Frame any
}

type frames []Frame

type (
	AckFrame = wire.AckFrame

	ConnectionCloseFrame = wire.ConnectionCloseFrame

	DataBlockedFrame = wire.DataBlockedFrame

	HandshakeDoneFrame = wire.HandshakeDoneFrame

	MaxDataFrame = wire.MaxDataFrame

	MaxStreamDataFrame = wire.MaxStreamDataFrame

	MaxStreamsFrame = wire.MaxStreamsFrame

	NewConnectionIDFrame = wire.NewConnectionIDFrame

	NewTokenFrame = wire.NewTokenFrame

	PathChallengeFrame = wire.PathChallengeFrame

	PathResponseFrame = wire.PathResponseFrame

	PingFrame = wire.PingFrame

	ResetStreamFrame = wire.ResetStreamFrame

	RetireConnectionIDFrame = wire.RetireConnectionIDFrame

	StopSendingFrame = wire.StopSendingFrame

	StreamsBlockedFrame = wire.StreamsBlockedFrame

	StreamDataBlockedFrame = wire.StreamDataBlockedFrame

	AckFrequencyFrame = wire.AckFrequencyFrame

	ImmediateAckFrame = wire.ImmediateAckFrame
)

type AckRange = wire.AckRange

type CryptoFrame struct {
	Offset int64
	Length int64
}

type StreamFrame struct {
	StreamID StreamID
	Offset   int64
	Length   int64
	Fin      bool
}

type DatagramFrame struct {
	Length int64
}

func (fs frames) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

func (f Frame) Encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

func encodePingFrame(enc *jsontext.Encoder, _ *PingFrame) error {
	_ = "STUB: not implemented"
	return nil
}

type ackRanges []wire.AckRange

func (ars ackRanges) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type ackRange wire.AckRange

func (ar ackRange) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

func encodeAckFrame(enc *jsontext.Encoder, f *AckFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeResetStreamFrame(enc *jsontext.Encoder, f *ResetStreamFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeStopSendingFrame(enc *jsontext.Encoder, f *StopSendingFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeCryptoFrame(enc *jsontext.Encoder, f *CryptoFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeNewTokenFrame(enc *jsontext.Encoder, f *NewTokenFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeStreamFrame(enc *jsontext.Encoder, f *StreamFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeMaxDataFrame(enc *jsontext.Encoder, f *MaxDataFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeMaxStreamDataFrame(enc *jsontext.Encoder, f *MaxStreamDataFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeMaxStreamsFrame(enc *jsontext.Encoder, f *MaxStreamsFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeDataBlockedFrame(enc *jsontext.Encoder, f *DataBlockedFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeStreamDataBlockedFrame(enc *jsontext.Encoder, f *StreamDataBlockedFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeStreamsBlockedFrame(enc *jsontext.Encoder, f *StreamsBlockedFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeNewConnectionIDFrame(enc *jsontext.Encoder, f *NewConnectionIDFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeRetireConnectionIDFrame(enc *jsontext.Encoder, f *RetireConnectionIDFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodePathChallengeFrame(enc *jsontext.Encoder, f *PathChallengeFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodePathResponseFrame(enc *jsontext.Encoder, f *PathResponseFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeConnectionCloseFrame(enc *jsontext.Encoder, f *ConnectionCloseFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeHandshakeDoneFrame(enc *jsontext.Encoder, _ *HandshakeDoneFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeDatagramFrame(enc *jsontext.Encoder, f *DatagramFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeAckFrequencyFrame(enc *jsontext.Encoder, f *AckFrequencyFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeImmediateAckFrame(enc *jsontext.Encoder, _ *ImmediateAckFrame) error {
	_ = "STUB: not implemented"
	return nil
}
