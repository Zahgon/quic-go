package http3

import (
	"errors"
	"io"

	"github.com/quic-go/quic-go/quicvarint"
)

type CapsuleType uint64

const CapsuleProtocolHeader = "Capsule-Protocol"

type noCopy struct{}

func (*noCopy) Lock()   { _ = "STUB: not implemented"; return }
func (*noCopy) Unlock() { _ = "STUB: not implemented"; return }

type CapsuleParser struct {
	noCopy noCopy

	r quicvarint.Reader

	generation uint64
	remaining  uint64
}

func NewCapsuleParser(r io.Reader) *CapsuleParser { _ = "STUB: not implemented"; return nil }

var (
	errReaderInvalid      = errors.New("http3: capsule reader is no longer valid")
	errCapsuleNotConsumed = errors.New("http3: previous capsule was not fully consumed")
)

func (p *CapsuleParser) Next() (CapsuleType, CapsuleReader, error) {
	_ = "STUB: not implemented"
	return *new(CapsuleType), *new(CapsuleReader), nil
}

type CapsuleReader struct {
	parser     *CapsuleParser
	generation uint64
}

var _ quicvarint.Reader = CapsuleReader{}

func (r CapsuleReader) valid() bool { _ = "STUB: not implemented"; return false }

func (r CapsuleReader) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r CapsuleReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (r CapsuleReader) Remaining() int64 { _ = "STUB: not implemented"; return 0 }

func (r CapsuleReader) Discard() error { _ = "STUB: not implemented"; return nil }

func WriteCapsule(w quicvarint.Writer, ct CapsuleType, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}
