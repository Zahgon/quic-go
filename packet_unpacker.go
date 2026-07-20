package quic

import (
	"github.com/quic-go/quic-go/internal/handshake"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

type headerDecryptor interface {
	DecryptHeader(sample []byte, firstByte *byte, pnBytes []byte)
}

type headerParseError struct {
	err error
}

func (e *headerParseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *headerParseError) Error() string { _ = "STUB: not implemented"; return "" }

type unpackedPacket struct {
	hdr             *wire.ExtendedHeader
	encryptionLevel protocol.EncryptionLevel
	data            []byte
}

type packetUnpacker struct {
	cs handshake.CryptoSetup

	shortHdrConnIDLen int
}

var _ unpacker = &packetUnpacker{}

func newPacketUnpacker(cs handshake.CryptoSetup, shortHdrConnIDLen int) *packetUnpacker {
	_ = "STUB: not implemented"
	return nil
}

func (u *packetUnpacker) UnpackLongHeader(hdr *wire.Header, data []byte) (*unpackedPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:exhaustive // Retry packets can't be unpacked.

func (u *packetUnpacker) UnpackShortHeader(rcvTime monotime.Time, data []byte) (protocol.PacketNumber, protocol.PacketNumberLen, protocol.KeyPhaseBit, []byte, error) {
	_ = "STUB: not implemented"
	return *new(protocol.PacketNumber), *new(protocol.PacketNumberLen), *new(protocol.KeyPhaseBit), nil, nil
}

func (u *packetUnpacker) unpackLongHeaderPacket(opener handshake.LongHeaderOpener, hdr *wire.Header, data []byte) (*wire.ExtendedHeader, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (u *packetUnpacker) unpackShortHeaderPacket(opener handshake.ShortHeaderOpener, rcvTime monotime.Time, data []byte) (protocol.PacketNumber, protocol.PacketNumberLen, protocol.KeyPhaseBit, []byte, error) {
	_ = "STUB: not implemented"
	return *new(protocol.PacketNumber), *new(protocol.PacketNumberLen), *new(protocol.KeyPhaseBit), nil, nil
}

func (u *packetUnpacker) unpackShortHeader(hd headerDecryptor, data []byte) (int, protocol.PacketNumber, protocol.PacketNumberLen, protocol.KeyPhaseBit, error) {
	_ = "STUB: not implemented"
	return 0, *new(protocol.PacketNumber), *new(protocol.PacketNumberLen), *new(protocol.KeyPhaseBit), nil
}

func (u *packetUnpacker) unpackLongHeader(hd headerDecryptor, hdr *wire.Header, data []byte) (*wire.ExtendedHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unpackLongHeader(hd headerDecryptor, hdr *wire.Header, data []byte) (*wire.ExtendedHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
