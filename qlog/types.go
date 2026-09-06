package qlog

import (
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/qerr"
)

type (
	ConnectionID             = protocol.ConnectionID
	ArbitraryLenConnectionID = protocol.ArbitraryLenConnectionID
	Version                  = protocol.Version
	PacketNumber             = protocol.PacketNumber
	EncryptionLevel          = protocol.EncryptionLevel
	KeyPhaseBit              = protocol.KeyPhaseBit
	KeyPhase                 = protocol.KeyPhase
	StreamID                 = protocol.StreamID
	TransportErrorCode       = qerr.TransportErrorCode
	ApplicationErrorCode     = qerr.ApplicationErrorCode
)

const (
	KeyPhaseZero = protocol.KeyPhaseZero

	KeyPhaseOne = protocol.KeyPhaseOne
)

type ECN string

const (
	ECNUnsupported ECN = ""

	ECTNot ECN = "Not-ECT"

	ECT0 ECN = "ECT(0)"

	ECT1 ECN = "ECT(1)"

	ECNCE ECN = "CE"
)

type Initiator string

const (
	InitiatorLocal  Initiator = "local"
	InitiatorRemote Initiator = "remote"
)

type streamType protocol.StreamType

func (s streamType) String() string { _ = "STUB: not implemented"; return "" }

type version protocol.Version

func (v version) String() string { _ = "STUB: not implemented"; return "" }

func encLevelToPacketNumberSpace(encLevel protocol.EncryptionLevel) string {
	_ = "STUB: not implemented"
	return ""
}

type KeyType string

const (
	KeyTypeServerInitial KeyType = "server_initial_secret"

	KeyTypeClientInitial KeyType = "client_initial_secret"

	KeyTypeServerHandshake KeyType = "server_handshake_secret"

	KeyTypeClientHandshake KeyType = "client_handshake_secret"

	KeyTypeServer0RTT KeyType = "server_0rtt_secret"

	KeyTypeClient0RTT KeyType = "client_0rtt_secret"

	KeyTypeServer1RTT KeyType = "server_1rtt_secret"

	KeyTypeClient1RTT KeyType = "client_1rtt_secret"
)

type KeyUpdateTrigger string

const (
	KeyUpdateTLS KeyUpdateTrigger = "tls"

	KeyUpdateRemote KeyUpdateTrigger = "remote_update"

	KeyUpdateLocal KeyUpdateTrigger = "local_update"
)

type transportError uint64

func (e transportError) String() string { _ = "STUB: not implemented"; return "" }

type PacketType string

const (
	PacketTypeInitial PacketType = "initial"

	PacketTypeHandshake PacketType = "handshake"

	PacketTypeRetry PacketType = "retry"

	PacketType0RTT PacketType = "0RTT"

	PacketTypeVersionNegotiation PacketType = "version_negotiation"

	PacketTypeStatelessReset PacketType = "stateless_reset"

	PacketType1RTT PacketType = "1RTT"
)

func EncryptionLevelToPacketType(l EncryptionLevel) PacketType {
	_ = "STUB: not implemented"
	return *new(PacketType)
}

type PacketLossReason string

const (
	PacketLossReorderingThreshold PacketLossReason = "reordering_threshold"

	PacketLossTimeThreshold PacketLossReason = "time_threshold"
)

type PacketDropReason string

const (
	PacketDropKeyUnavailable PacketDropReason = "key_unavailable"

	PacketDropUnknownConnectionID PacketDropReason = "unknown_connection_id"

	PacketDropHeaderParseError PacketDropReason = "header_parse_error"

	PacketDropPayloadDecryptError PacketDropReason = "payload_decrypt_error"

	PacketDropProtocolViolation PacketDropReason = "protocol_violation"

	PacketDropDOSPrevention PacketDropReason = "dos_prevention"

	PacketDropUnsupportedVersion PacketDropReason = "unsupported_version"

	PacketDropUnexpectedPacket PacketDropReason = "unexpected_packet"

	PacketDropUnexpectedSourceConnectionID PacketDropReason = "unexpected_source_connection_id"

	PacketDropUnexpectedVersion PacketDropReason = "unexpected_version"

	PacketDropDuplicate PacketDropReason = "duplicate"
)

type LossTimerUpdateType string

const (
	LossTimerUpdateTypeSet       LossTimerUpdateType = "set"
	LossTimerUpdateTypeExpired   LossTimerUpdateType = "expired"
	LossTimerUpdateTypeCancelled LossTimerUpdateType = "cancelled"
)

type TimerType string

const (
	TimerTypeACK TimerType = "ack"

	TimerTypePTO TimerType = "pto"

	TimerTypePathProbe TimerType = "path_probe"
)

type CongestionState string

const (
	CongestionStateSlowStart CongestionState = "slow_start"

	CongestionStateCongestionAvoidance CongestionState = "congestion_avoidance"

	CongestionStateRecovery CongestionState = "recovery"

	CongestionStateApplicationLimited CongestionState = "application_limited"
)

func (s CongestionState) String() string { _ = "STUB: not implemented"; return "" }

type ECNState string

const (
	ECNStateTesting ECNState = "testing"

	ECNStateUnknown ECNState = "unknown"

	ECNStateFailed ECNState = "failed"

	ECNStateCapable ECNState = "capable"
)

type ConnectionCloseTrigger string

const (
	ConnectionCloseTriggerIdleTimeout ConnectionCloseTrigger = "idle_timeout"

	ConnectionCloseTriggerApplication ConnectionCloseTrigger = "application"

	ConnectionCloseTriggerVersionMismatch ConnectionCloseTrigger = "version_mismatch"

	ConnectionCloseTriggerStatelessReset ConnectionCloseTrigger = "stateless_reset"
)

type DatagramPayloadChecksum uint32

func CalculateDatagramPayloadChecksum(payload []byte) DatagramPayloadChecksum {
	_ = "STUB: not implemented"
	return *new(DatagramPayloadChecksum)
}
