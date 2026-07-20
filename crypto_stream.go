package quic

import (
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/wire"
)

const disableClientHelloScramblingEnv = "QUIC_GO_DISABLE_CLIENTHELLO_SCRAMBLING"

type baseCryptoStream struct {
	queue frameSorter

	highestOffset protocol.ByteCount
	finished      bool

	writeOffset protocol.ByteCount
	writeBuf    []byte
}

func newCryptoStream() *cryptoStream { _ = "STUB: not implemented"; return nil }

func (s *baseCryptoStream) HandleCryptoFrame(f *wire.CryptoFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *baseCryptoStream) GetCryptoData() []byte { _ = "STUB: not implemented"; return nil }

func (s *baseCryptoStream) Finish() error { _ = "STUB: not implemented"; return nil }

func (s *baseCryptoStream) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *baseCryptoStream) HasData() bool { _ = "STUB: not implemented"; return false }

func (s *baseCryptoStream) PopCryptoFrame(maxLen protocol.ByteCount) *wire.CryptoFrame {
	_ = "STUB: not implemented"
	return nil
}

type cryptoStream struct {
	baseCryptoStream
}

type clientHelloCut struct {
	start protocol.ByteCount
	end   protocol.ByteCount
}

type initialCryptoStream struct {
	baseCryptoStream

	scramble bool
	end      protocol.ByteCount
	cuts     [2]clientHelloCut
}

func newInitialCryptoStream(isClient bool) *initialCryptoStream {
	_ = "STUB: not implemented"
	return nil
}

func (s *initialCryptoStream) HasData() bool { _ = "STUB: not implemented"; return false }

func (s *initialCryptoStream) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *initialCryptoStream) PopCryptoFrame(maxLen protocol.ByteCount) *wire.CryptoFrame {
	_ = "STUB: not implemented"
	return nil
}
