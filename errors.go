package quic

import (
	"github.com/quic-go/quic-go/internal/qerr"
)

type (
	TransportError = qerr.TransportError

	ApplicationError = qerr.ApplicationError

	VersionNegotiationError = qerr.VersionNegotiationError

	StatelessResetError = qerr.StatelessResetError

	IdleTimeoutError = qerr.IdleTimeoutError

	HandshakeTimeoutError = qerr.HandshakeTimeoutError
)

type (
	TransportErrorCode = qerr.TransportErrorCode

	ApplicationErrorCode = qerr.ApplicationErrorCode

	StreamErrorCode = qerr.StreamErrorCode
)

const (
	NoError = qerr.NoError

	InternalError = qerr.InternalError

	ConnectionRefused = qerr.ConnectionRefused

	FlowControlError = qerr.FlowControlError

	StreamLimitError = qerr.StreamLimitError

	StreamStateError = qerr.StreamStateError

	FinalSizeError = qerr.FinalSizeError

	FrameEncodingError = qerr.FrameEncodingError

	TransportParameterError = qerr.TransportParameterError

	ConnectionIDLimitError = qerr.ConnectionIDLimitError

	ProtocolViolation = qerr.ProtocolViolation

	InvalidToken = qerr.InvalidToken

	ApplicationErrorErrorCode = qerr.ApplicationErrorErrorCode

	CryptoBufferExceeded = qerr.CryptoBufferExceeded

	KeyUpdateError = qerr.KeyUpdateError

	AEADLimitReached = qerr.AEADLimitReached

	NoViablePathError = qerr.NoViablePathError
)

type StreamError struct {
	StreamID  StreamID
	ErrorCode StreamErrorCode
	Remote    bool
}

func (e *StreamError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *StreamError) Error() string { _ = "STUB: not implemented"; return "" }

type DatagramTooLargeError struct {
	MaxDatagramPayloadSize int64
}

func (e *DatagramTooLargeError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *DatagramTooLargeError) Error() string { _ = "STUB: not implemented"; return "" }
