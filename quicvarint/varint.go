package quicvarint

import (
	"io"
)

const (
	Min = 0

	Max = maxVarInt8

	maxVarInt1 = 63
	maxVarInt2 = 16383
	maxVarInt4 = 1073741823
	maxVarInt8 = 4611686018427387903
)

type varintLengthError struct {
	Num uint64
}

func (e *varintLengthError) Error() string { _ = "STUB: not implemented"; return "" }

func Read(r io.ByteReader) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func Parse(b []byte) (uint64, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func Append(b []byte, i uint64) []byte { _ = "STUB: not implemented"; return nil }

func AppendWithLen(b []byte, i uint64, length int) []byte { _ = "STUB: not implemented"; return nil }

//gcassert:inline
func Len(i uint64) int { _ = "STUB: not implemented"; return 0 }
