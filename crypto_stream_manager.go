package quic

import (
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

type cryptoStreamManager struct {
	initialStream   *initialCryptoStream
	handshakeStream *cryptoStream
	oneRTTStream    *cryptoStream
}

func newCryptoStreamManager(
	initialStream *initialCryptoStream,
	handshakeStream *cryptoStream,
	oneRTTStream *cryptoStream,
) *cryptoStreamManager {
	_ = "STUB: not implemented"
	return nil
}

func (m *cryptoStreamManager) HandleCryptoFrame(frame *wire.CryptoFrame, encLevel protocol.EncryptionLevel) error {
	_ = "STUB: not implemented"
	//nolint:exhaustive // CRYPTO frames cannot be sent in 0-RTT packets.
	return nil
}

func (m *cryptoStreamManager) GetCryptoData(encLevel protocol.EncryptionLevel) []byte {
	_ = "STUB: not implemented"
	//nolint:exhaustive // CRYPTO frames cannot be sent in 0-RTT packets.
	return nil
}

func (m *cryptoStreamManager) GetPostHandshakeData(maxSize protocol.ByteCount) *wire.CryptoFrame {
	_ = "STUB: not implemented"
	return nil
}

func (m *cryptoStreamManager) Drop(encLevel protocol.EncryptionLevel) error {
	_ = "STUB: not implemented"
	//nolint:exhaustive // 1-RTT keys should never get dropped.
	return nil
}
