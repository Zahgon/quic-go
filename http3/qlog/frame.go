package qlog

import (
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter/jsontext"
)

type Frame struct {
	Frame any
}

func (f Frame) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type DataFrame struct{}

func (f *DataFrame) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type HeaderField struct {
	Name  string
	Value string
}

type HeadersFrame struct {
	HeaderFields []HeaderField
}

func (f *HeadersFrame) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type GoAwayFrame struct {
	StreamID quic.StreamID
}

func (f *GoAwayFrame) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type SettingsFrame struct {
	MaxFieldSectionSize int64
	Datagram            *bool
	ExtendedConnect     *bool
	Other               map[uint64]uint64
}

func (f *SettingsFrame) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type PushPromiseFrame struct{}

func (f *PushPromiseFrame) encode(enc *jsontext.Encoder) error {
	_ = "STUB: not implemented"
	return nil
}

type CancelPushFrame struct{}

func (f *CancelPushFrame) encode(enc *jsontext.Encoder) error {
	_ = "STUB: not implemented"
	return nil
}

type MaxPushIDFrame struct{}

func (f *MaxPushIDFrame) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type ReservedFrame struct {
	Type uint64
}

func (f *ReservedFrame) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type UnknownFrame struct {
	Type uint64
}

func (f *UnknownFrame) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }
