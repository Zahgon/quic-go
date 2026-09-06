package qlog

import (
	"net/netip"
	"time"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/qlogwriter/jsontext"
)

func milliseconds(dur time.Duration) float64 { _ = "STUB: not implemented"; return 0 }

type encoderHelper struct {
	enc *jsontext.Encoder
	err error
}

func (h *encoderHelper) WriteToken(t jsontext.Token) { _ = "STUB: not implemented"; return }

type versions []Version

func (v versions) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type RawInfo struct {
	Length        int
	PayloadLength int
}

func (i RawInfo) encode(enc *jsontext.Encoder) error { _ = "STUB: not implemented"; return nil }

type PathEndpointInfo struct {
	IPv4 netip.AddrPort
	IPv6 netip.AddrPort
}

func (p PathEndpointInfo) encode(enc *jsontext.Encoder) error {
	_ = "STUB: not implemented"
	return nil
}

type StartedConnection struct {
	Local  PathEndpointInfo
	Remote PathEndpointInfo
}

func (e StartedConnection) Name() string { _ = "STUB: not implemented"; return "" }

func (e StartedConnection) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type VersionInformation struct {
	ClientVersions, ServerVersions []Version
	ChosenVersion                  Version
}

func (e VersionInformation) Name() string { _ = "STUB: not implemented"; return "" }

func (e VersionInformation) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type ConnectionClosed struct {
	Initiator Initiator

	ConnectionError  *TransportErrorCode
	ApplicationError *ApplicationErrorCode

	Reason string

	Trigger ConnectionCloseTrigger
}

func (e ConnectionClosed) Name() string { _ = "STUB: not implemented"; return "" }

func (e ConnectionClosed) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type PacketSent struct {
	Header                  PacketHeader
	Raw                     RawInfo
	DatagramPayloadChecksum DatagramPayloadChecksum
	Frames                  []Frame
	ECN                     ECN
	IsCoalesced             bool
	Trigger                 string
	SupportedVersions       []Version
}

func (e PacketSent) Name() string { _ = "STUB: not implemented"; return "" }

func (e PacketSent) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type PacketReceived struct {
	Header                  PacketHeader
	Raw                     RawInfo
	DatagramPayloadChecksum DatagramPayloadChecksum
	Frames                  []Frame
	ECN                     ECN
	IsCoalesced             bool
	Trigger                 string
}

func (e PacketReceived) Name() string { _ = "STUB: not implemented"; return "" }

func (e PacketReceived) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type VersionNegotiationReceived struct {
	Header            PacketHeaderVersionNegotiation
	SupportedVersions []Version
}

func (e VersionNegotiationReceived) Name() string { _ = "STUB: not implemented"; return "" }

func (e VersionNegotiationReceived) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type VersionNegotiationSent struct {
	Header            PacketHeaderVersionNegotiation
	SupportedVersions []Version
}

func (e VersionNegotiationSent) Name() string { _ = "STUB: not implemented"; return "" }

func (e VersionNegotiationSent) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type PacketBuffered struct {
	Header                  PacketHeader
	Raw                     RawInfo
	DatagramPayloadChecksum DatagramPayloadChecksum
}

func (e PacketBuffered) Name() string { _ = "STUB: not implemented"; return "" }

func (e PacketBuffered) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type PacketDropped struct {
	Header                  PacketHeader
	Raw                     RawInfo
	DatagramPayloadChecksum DatagramPayloadChecksum
	Trigger                 PacketDropReason
}

func (e PacketDropped) Name() string { _ = "STUB: not implemented"; return "" }

func (e PacketDropped) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type StreamPriorityUpdated struct {
	StreamID    StreamID
	Urgency     int8
	Incremental bool
}

func (e StreamPriorityUpdated) Name() string { _ = "STUB: not implemented"; return "" }

func (e StreamPriorityUpdated) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type MTUUpdated struct {
	Value int
	Done  bool
}

func (e MTUUpdated) Name() string { _ = "STUB: not implemented"; return "" }

func (e MTUUpdated) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type MetricsUpdated struct {
	MinRTT           time.Duration
	SmoothedRTT      time.Duration
	LatestRTT        time.Duration
	RTTVariance      time.Duration
	CongestionWindow int
	BytesInFlight    int
	PacketsInFlight  int
}

func (e MetricsUpdated) Name() string { _ = "STUB: not implemented"; return "" }

func (e MetricsUpdated) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type PTOCountUpdated struct {
	PTOCount uint32
}

func (e PTOCountUpdated) Name() string { _ = "STUB: not implemented"; return "" }

func (e PTOCountUpdated) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type PacketLost struct {
	Header  PacketHeader
	Trigger PacketLossReason
}

func (e PacketLost) Name() string { _ = "STUB: not implemented"; return "" }

func (e PacketLost) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type SpuriousLoss struct {
	EncryptionLevel  protocol.EncryptionLevel
	PacketNumber     protocol.PacketNumber
	PacketReordering uint64
	TimeReordering   time.Duration
}

func (e SpuriousLoss) Name() string { _ = "STUB: not implemented"; return "" }

func (e SpuriousLoss) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type KeyUpdated struct {
	Trigger  KeyUpdateTrigger
	KeyType  KeyType
	KeyPhase KeyPhase
}

func (e KeyUpdated) Name() string { _ = "STUB: not implemented"; return "" }

func (e KeyUpdated) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type KeyDiscarded struct {
	KeyType  KeyType
	KeyPhase KeyPhase
}

func (e KeyDiscarded) Name() string { _ = "STUB: not implemented"; return "" }

func (e KeyDiscarded) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type ParametersSet struct {
	Restore                         bool
	Initiator                       Initiator
	SentBy                          protocol.Perspective
	OriginalDestinationConnectionID protocol.ConnectionID
	InitialSourceConnectionID       protocol.ConnectionID
	RetrySourceConnectionID         *protocol.ConnectionID
	StatelessResetToken             *protocol.StatelessResetToken
	DisableActiveMigration          bool
	MaxIdleTimeout                  time.Duration
	MaxUDPPayloadSize               protocol.ByteCount
	AckDelayExponent                uint8
	MaxAckDelay                     time.Duration
	ActiveConnectionIDLimit         uint64
	InitialMaxData                  protocol.ByteCount
	InitialMaxStreamDataBidiLocal   protocol.ByteCount
	InitialMaxStreamDataBidiRemote  protocol.ByteCount
	InitialMaxStreamDataUni         protocol.ByteCount
	InitialMaxStreamsBidi           int64
	InitialMaxStreamsUni            int64
	PreferredAddress                *PreferredAddress
	MaxDatagramFrameSize            protocol.ByteCount
	EnableResetStreamAt             bool
}

func (e ParametersSet) Name() string { _ = "STUB: not implemented"; return "" }

func (e ParametersSet) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type PreferredAddress struct {
	IPv4, IPv6          netip.AddrPort
	ConnectionID        protocol.ConnectionID
	StatelessResetToken protocol.StatelessResetToken
}

func (a PreferredAddress) encode(enc *jsontext.Encoder) error {
	_ = "STUB: not implemented"
	return nil
}

type LossTimerUpdated struct {
	Type      LossTimerUpdateType
	TimerType TimerType
	EncLevel  EncryptionLevel
	Time      time.Time
}

func (e LossTimerUpdated) Name() string { _ = "STUB: not implemented"; return "" }

func (e LossTimerUpdated) Encode(enc *jsontext.Encoder, t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type eventLossTimerCanceled struct{}

func (e eventLossTimerCanceled) Name() string { _ = "STUB: not implemented"; return "" }

func (e eventLossTimerCanceled) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type CongestionStateUpdated struct {
	State CongestionState
}

func (e CongestionStateUpdated) Name() string { _ = "STUB: not implemented"; return "" }

func (e CongestionStateUpdated) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type ECNStateUpdated struct {
	State   ECNState
	Trigger string
}

func (e ECNStateUpdated) Name() string { _ = "STUB: not implemented"; return "" }

func (e ECNStateUpdated) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type ALPNInformation struct {
	ChosenALPN string
}

func (e ALPNInformation) Name() string { _ = "STUB: not implemented"; return "" }

func (e ALPNInformation) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type DebugEvent struct {
	EventName string
	Message   string
}

func (e DebugEvent) Name() string { _ = "STUB: not implemented"; return "" }

func (e DebugEvent) Encode(enc *jsontext.Encoder, _ time.Time) error {
	_ = "STUB: not implemented"
	return nil
}
