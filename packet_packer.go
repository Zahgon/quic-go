package quic

import (
	"errors"
	"math/rand/v2"

	"github.com/quic-go/quic-go/internal/ackhandler"
	"github.com/quic-go/quic-go/internal/handshake"
	"github.com/quic-go/quic-go/internal/monotime"
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/qerr"
	"github.com/quic-go/quic-go/internal/wire"
)

var errNothingToPack = errors.New("nothing to pack")

type packer interface {
	PackCoalescedPacket(onlyAck bool, maxPacketSize protocol.ByteCount, now monotime.Time, v protocol.Version) (*coalescedPacket, error)
	PackAckOnlyPacket(maxPacketSize protocol.ByteCount, now monotime.Time, v protocol.Version) (shortHeaderPacket, *packetBuffer, error)
	AppendPacket(_ *packetBuffer, maxPacketSize protocol.ByteCount, now monotime.Time, v protocol.Version) (shortHeaderPacket, error)
	PackPTOProbePacket(_ protocol.EncryptionLevel, _ protocol.ByteCount, addPingIfEmpty bool, now monotime.Time, v protocol.Version) (*coalescedPacket, error)
	PackConnectionClose(*qerr.TransportError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)
	PackApplicationClose(*qerr.ApplicationError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)
	PackPathProbePacket(protocol.ConnectionID, []ackhandler.Frame, protocol.Version) (shortHeaderPacket, *packetBuffer, error)
	PackMTUProbePacket(ping ackhandler.Frame, size protocol.ByteCount, v protocol.Version) (shortHeaderPacket, *packetBuffer, error)

	SetToken([]byte)
}

type sealer interface {
	handshake.LongHeaderSealer
}

type payload struct {
	streamFrames []ackhandler.StreamFrame
	frames       []ackhandler.Frame
	ack          *wire.AckFrame
	length       protocol.ByteCount
}

type longHeaderPacket struct {
	header       *wire.ExtendedHeader
	ack          *wire.AckFrame
	frames       []ackhandler.Frame
	streamFrames []ackhandler.StreamFrame

	length protocol.ByteCount
}

type shortHeaderPacket struct {
	PacketNumber         protocol.PacketNumber
	Frames               []ackhandler.Frame
	StreamFrames         []ackhandler.StreamFrame
	Ack                  *wire.AckFrame
	Length               protocol.ByteCount
	IsPathMTUProbePacket bool
	IsPathProbePacket    bool

	DestConnID      protocol.ConnectionID
	PacketNumberLen protocol.PacketNumberLen
	KeyPhase        protocol.KeyPhaseBit
}

func (p *shortHeaderPacket) IsAckEliciting() bool { _ = "STUB: not implemented"; return false }

type coalescedPacket struct {
	buffer         *packetBuffer
	longHdrPackets []*longHeaderPacket
	shortHdrPacket *shortHeaderPacket
}

func (p *coalescedPacket) IsOnlyShortHeaderPacket() bool { _ = "STUB: not implemented"; return false }

func (p *longHeaderPacket) EncryptionLevel() protocol.EncryptionLevel {
	_ = "STUB: not implemented"
	//nolint:exhaustive // Will never be called for Retry packets (and they don't have encrypted data).
	return *new(protocol.EncryptionLevel)
}

func (p *longHeaderPacket) IsAckEliciting() bool { _ = "STUB: not implemented"; return false }

type packetNumberManager interface {
	PeekPacketNumber(protocol.EncryptionLevel) (protocol.PacketNumber, protocol.PacketNumberLen)
	PopPacketNumber(protocol.EncryptionLevel) protocol.PacketNumber
}

type sealingManager interface {
	GetInitialSealer() (handshake.LongHeaderSealer, error)
	GetHandshakeSealer() (handshake.LongHeaderSealer, error)
	Get0RTTSealer() (handshake.LongHeaderSealer, error)
	Get1RTTSealer() (handshake.ShortHeaderSealer, error)
}

type frameSource interface {
	HasData() bool
	Append([]ackhandler.Frame, []ackhandler.StreamFrame, protocol.ByteCount, monotime.Time, protocol.Version) ([]ackhandler.Frame, []ackhandler.StreamFrame, protocol.ByteCount)
}

type ackFrameSource interface {
	GetAckFrame(_ protocol.EncryptionLevel, now monotime.Time, onlyIfQueued bool) *wire.AckFrame
}

type packetPacker struct {
	srcConnID     protocol.ConnectionID
	getDestConnID func() protocol.ConnectionID

	perspective protocol.Perspective
	cryptoSetup sealingManager

	initialStream   *initialCryptoStream
	handshakeStream *cryptoStream

	token []byte

	pnManager           packetNumberManager
	framer              frameSource
	acks                ackFrameSource
	datagramQueue       *datagramQueue
	retransmissionQueue *retransmissionQueue
	rand                rand.Rand

	numNonAckElicitingAcks int
}

var _ packer = &packetPacker{}

func newPacketPacker(
	srcConnID protocol.ConnectionID,
	getDestConnID func() protocol.ConnectionID,
	initialStream *initialCryptoStream,
	handshakeStream *cryptoStream,
	packetNumberManager packetNumberManager,
	retransmissionQueue *retransmissionQueue,
	cryptoSetup sealingManager,
	framer frameSource,
	acks ackFrameSource,
	datagramQueue *datagramQueue,
	perspective protocol.Perspective,
) *packetPacker {
	_ = "STUB: not implemented"
	return nil
}

func (p *packetPacker) PackConnectionClose(e *qerr.TransportError, maxPacketSize protocol.ByteCount, v protocol.Version) (*coalescedPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packetPacker) PackApplicationClose(e *qerr.ApplicationError, maxPacketSize protocol.ByteCount, v protocol.Version) (*coalescedPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packetPacker) packConnectionClose(
	isApplicationError bool,
	errorCode uint64,
	frameType uint64,
	reason string,
	maxPacketSize protocol.ByteCount,
	v protocol.Version,
) (*coalescedPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packetPacker) longHeaderPacketLength(hdr *wire.ExtendedHeader, pl payload, v protocol.Version) protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (p *packetPacker) shortHeaderPacketLength(connID protocol.ConnectionID, pnLen protocol.PacketNumberLen, pl payload) protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (p *packetPacker) initialPaddingLen(frames []ackhandler.Frame, currentSize, maxPacketSize protocol.ByteCount) protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (p *packetPacker) PackCoalescedPacket(onlyAck bool, maxSize protocol.ByteCount, now monotime.Time, v protocol.Version) (*coalescedPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packetPacker) PackAckOnlyPacket(maxSize protocol.ByteCount, now monotime.Time, v protocol.Version) (shortHeaderPacket, *packetBuffer, error) {
	_ = "STUB: not implemented"
	return *new(shortHeaderPacket), nil, nil
}

func (p *packetPacker) AppendPacket(buf *packetBuffer, maxSize protocol.ByteCount, now monotime.Time, v protocol.Version) (shortHeaderPacket, error) {
	_ = "STUB: not implemented"
	return *new(shortHeaderPacket), nil
}

func (p *packetPacker) appendPacket(
	buf *packetBuffer,
	onlyAck bool,
	maxPacketSize protocol.ByteCount,
	now monotime.Time,
	v protocol.Version,
) (shortHeaderPacket, error) {
	_ = "STUB: not implemented"
	return *new(shortHeaderPacket), nil
}

func (p *packetPacker) maybeGetCryptoPacket(
	maxPacketSize protocol.ByteCount,
	encLevel protocol.EncryptionLevel,
	now monotime.Time,
	addPingIfEmpty bool,
	onlyAck bool,
	v protocol.Version,
) (*wire.ExtendedHeader, payload) {
	_ = "STUB: not implemented"
	return nil, *new(payload)
}

//nolint:exhaustive // Initial and Handshake are the only two encryption levels here.

func (p *packetPacker) maybeGetAppDataPacketFor0RTT(sealer sealer, maxSize protocol.ByteCount, now monotime.Time, v protocol.Version) (*wire.ExtendedHeader, payload) {
	_ = "STUB: not implemented"
	return nil, *new(payload)
}

func (p *packetPacker) maybeGetShortHeaderPacket(
	sealer handshake.ShortHeaderSealer,
	hdrLen, maxPacketSize protocol.ByteCount,
	onlyAck bool,
	now monotime.Time,
	v protocol.Version,
) payload {
	_ = "STUB: not implemented"
	return *new(payload)
}

func (p *packetPacker) maybeGetAppDataPacket(
	maxPayloadSize protocol.ByteCount,
	onlyAck, ackAllowed bool,
	now monotime.Time,
	v protocol.Version,
) payload {
	_ = "STUB: not implemented"
	return *new(payload)
}

func (p *packetPacker) composeNextPacket(
	maxPayloadSize protocol.ByteCount,
	onlyAck, ackAllowed bool,
	now monotime.Time,
	v protocol.Version,
) payload {
	_ = "STUB: not implemented"
	return *new(payload)
}

func (p *packetPacker) PackPTOProbePacket(
	encLevel protocol.EncryptionLevel,
	maxPacketSize protocol.ByteCount,
	addPingIfEmpty bool,
	now monotime.Time,
	v protocol.Version,
) (*coalescedPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:exhaustive // Probe packets are never sent for 0-RTT.

func (p *packetPacker) packPTOProbePacket1RTT(maxPacketSize protocol.ByteCount, addPingIfEmpty bool, now monotime.Time, v protocol.Version) (*coalescedPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packetPacker) PackMTUProbePacket(ping ackhandler.Frame, size protocol.ByteCount, v protocol.Version) (shortHeaderPacket, *packetBuffer, error) {
	_ = "STUB: not implemented"
	return *new(shortHeaderPacket), nil, nil
}

func (p *packetPacker) PackPathProbePacket(connID protocol.ConnectionID, frames []ackhandler.Frame, v protocol.Version) (shortHeaderPacket, *packetBuffer, error) {
	_ = "STUB: not implemented"
	return *new(shortHeaderPacket), nil, nil
}

func (p *packetPacker) getLongHeader(encLevel protocol.EncryptionLevel, v protocol.Version) *wire.ExtendedHeader {
	_ = "STUB: not implemented"
	return nil
}

//nolint:exhaustive // 1-RTT packets are not long header packets.

func (p *packetPacker) appendLongHeaderPacket(buffer *packetBuffer, header *wire.ExtendedHeader, pl payload, padding protocol.ByteCount, encLevel protocol.EncryptionLevel, sealer sealer, v protocol.Version) (*longHeaderPacket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packetPacker) appendShortHeaderPacket(
	buffer *packetBuffer,
	connID protocol.ConnectionID,
	pn protocol.PacketNumber,
	pnLen protocol.PacketNumberLen,
	kp protocol.KeyPhaseBit,
	pl payload,
	padding, maxPacketSize protocol.ByteCount,
	sealer sealer,
	isMTUProbePacket bool,
	v protocol.Version,
) (shortHeaderPacket, error) {
	_ = "STUB: not implemented"
	return *new(shortHeaderPacket), nil
}

func (p *packetPacker) appendPacketPayload(raw []byte, pl payload, paddingLen protocol.ByteCount, v protocol.Version) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packetPacker) encryptPacket(raw []byte, sealer sealer, pn protocol.PacketNumber, payloadOffset, pnLen protocol.ByteCount) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (p *packetPacker) SetToken(token []byte) { _ = "STUB: not implemented"; return }

type emptyHandler struct{}

var _ ackhandler.FrameHandler = emptyHandler{}

func (emptyHandler) OnAcked(wire.Frame) { _ = "STUB: not implemented"; return }
func (emptyHandler) OnLost(wire.Frame)  { _ = "STUB: not implemented"; return }
