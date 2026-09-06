package quicvarint

import (
	"bytes"
	"io"
)

type Reader interface {
	io.ByteReader
	io.Reader
}

var _ Reader = &bytes.Reader{}

type Peeker interface {
	Peek(b []byte) (int, error)
}

func Peek(p Peeker) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

type byteReader struct {
	io.Reader
}

var _ Reader = &byteReader{}

func NewReader(r io.Reader) Reader { _ = "STUB: not implemented"; return *new(Reader) }

func (r *byteReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

type Writer interface {
	io.ByteWriter
	io.Writer
}

var _ Writer = &bytes.Buffer{}

type byteWriter struct {
	io.Writer
}

var _ Writer = &byteWriter{}

func NewWriter(w io.Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func (w *byteWriter) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }
