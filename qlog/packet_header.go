package qlog

import (
	"github.com/quic-go/quic-go/qlogwriter/jsontext"
)

type Token struct {
	Raw []byte
}

func (t Token) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type PacketHeader struct {
	PacketType       PacketType
	KeyPhaseBit      KeyPhaseBit
	PacketNumber     PacketNumber
	Version          Version
	SrcConnectionID  ConnectionID
	DestConnectionID ConnectionID
	Token            *Token
}

func (h PacketHeader) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type PacketHeaderVersionNegotiation struct {
	SrcConnectionID  ArbitraryLenConnectionID
	DestConnectionID ArbitraryLenConnectionID
}

func (h PacketHeaderVersionNegotiation) encode(enc *jsontext.Encoder) error {
	_ = "STUB: not implemented"
	return nil
}
