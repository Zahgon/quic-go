package quic

import (
	"sync"

	"github.com/quic-go/quic-go/internal/protocol"
)

type packetBuffer struct {
	Data []byte

	refCount int
}

func (b *packetBuffer) Split() { _ = "STUB: not implemented"; return }

func (b *packetBuffer) Decrement() { _ = "STUB: not implemented"; return }

func (b *packetBuffer) MaybeRelease() { _ = "STUB: not implemented"; return }

func (b *packetBuffer) Release() { _ = "STUB: not implemented"; return }

func (b *packetBuffer) Len() protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}
func (b *packetBuffer) Cap() protocol.ByteCount {
	_ = "STUB: not implemented"
	return *new(protocol.ByteCount)
}

func (b *packetBuffer) putBack() { _ = "STUB: not implemented"; return }

var bufferPool, largeBufferPool sync.Pool

func getPacketBuffer() *packetBuffer { _ = "STUB: not implemented"; return nil }

func getLargePacketBuffer() *packetBuffer { _ = "STUB: not implemented"; return nil }

func init() {
	bufferPool.New = func() any {
		return &packetBuffer{Data: make([]byte, 0, protocol.MaxPacketBufferSize)}
	}
	largeBufferPool.New = func() any {
		return &packetBuffer{Data: make([]byte, 0, protocol.MaxLargePacketBufferSize)}
	}
}
