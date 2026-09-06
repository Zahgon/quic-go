package testutils

import (
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

func writePacket(hdr *wire.ExtendedHeader, data []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func packRawPayload(version protocol.Version, frames []wire.Frame) []byte {
	_ = "STUB: not implemented"
	return nil
}

func ComposeInitialPacket(
	srcConnID, destConnID, key protocol.ConnectionID,
	token []byte,
	frames []wire.Frame,
	sentBy protocol.Perspective,
	version protocol.Version,
) []byte {
	_ = "STUB: not implemented"
	return nil
}

func ComposeRetryPacket(
	srcConnID protocol.ConnectionID,
	destConnID protocol.ConnectionID,
	origDestConnID protocol.ConnectionID,
	token []byte,
	version protocol.Version,
) []byte {
	_ = "STUB: not implemented"
	return nil
}
