package jsontext

import (
	"io"
)

type kind uint8

const (
	kindString kind = iota
	kindInt
	kindUint
	kindFloat
	kindBool
	kindNull
	kindObjectStart
	kindObjectEnd
	kindArrayStart
	kindArrayEnd
)

type Token struct {
	kind kind
	str  string
	i64  int64
	u64  uint64
	f64  float64
	b    bool
}

func String(s string) Token { _ = "STUB: not implemented"; return *new(Token) }

func Int(i int64) Token { _ = "STUB: not implemented"; return *new(Token) }

func Uint(u uint64) Token { _ = "STUB: not implemented"; return *new(Token) }

func Float(f float64) Token { _ = "STUB: not implemented"; return *new(Token) }

func Bool(b bool) Token { _ = "STUB: not implemented"; return *new(Token) }

var Null Token = Token{kind: kindNull}

var BeginObject Token = Token{kind: kindObjectStart}

var EndObject Token = Token{kind: kindObjectEnd}

var BeginArray Token = Token{kind: kindArrayStart}

var EndArray Token = Token{kind: kindArrayEnd}

var True Token = Bool(true)

var False Token = Bool(false)

var hexDigits = [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

var (
	commaByte       = []byte(",")
	quoteByte       = []byte(`"`)
	colonByte       = []byte(":")
	trueByte        = []byte("true")
	falseByte       = []byte("false")
	nullByte        = []byte("null")
	openObjectByte  = []byte("{")
	closeObjectByte = []byte("}")
	openArrayByte   = []byte("[")
	closeArrayByte  = []byte("]")
	newlineByte     = []byte("\n")
	escapeQuote     = []byte(`\"`)
	escapeBackslash = []byte(`\\`)
	escapeBackspace = []byte(`\b`)
	escapeFormfeed  = []byte(`\f`)
	escapeNewline   = []byte(`\n`)
	escapeCarriage  = []byte(`\r`)
	escapeTab       = []byte(`\t`)
	escapeUnicode   = []byte(`\u00`)
)

type context struct {
	isObject   bool
	needsComma bool
	expectKey  bool
}

type Encoder struct {
	w     io.Writer
	buf   [64]byte
	stack []context
}

func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

func (e *Encoder) WriteToken(t Token) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) afterValue() { _ = "STUB: not implemented"; return }

func stringToBytes(s string) []byte { _ = "STUB: not implemented"; return nil }
