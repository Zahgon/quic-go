package qlogwriter

import (
	"io"
	"sync"
	"time"

	"github.com/quic-go/quic-go/qlogwriter/jsontext"
)

type Trace interface {
	AddProducer() Recorder

	SupportsSchemas(schema string) bool
}

type Recorder interface {
	RecordEvent(Event)

	io.Closer
}

type Event interface {
	Name() string

	Encode(encoder *jsontext.Encoder, eventTime time.Time) error
}

const RecordSeparator byte = 0x1e

var recordSeparator = []byte{RecordSeparator}

type event struct {
	Time  time.Time
	Event Event
}

const eventChanSize = 50

type FileSeq struct {
	w             io.WriteCloser
	enc           *jsontext.Encoder
	referenceTime time.Time

	runStopped chan struct{}
	encodeErr  error
	events     chan event
	done       chan struct{}

	mx        sync.Mutex
	producers int
	closed    bool

	eventSchemas []string
}

var _ Trace = &FileSeq{}

func NewFileSeq(w io.WriteCloser) *FileSeq { _ = "STUB: not implemented"; return nil }

func NewConnectionFileSeq(w io.WriteCloser, isClient bool, odcid ConnectionID, eventSchemas []string) *FileSeq {
	_ = "STUB: not implemented"
	return nil
}

func newFileSeq(w io.WriteCloser, pers string, odcid *ConnectionID, eventSchemas []string) *FileSeq {
	_ = "STUB: not implemented"
	return nil
}

func (t *FileSeq) SupportsSchemas(schema string) bool { _ = "STUB: not implemented"; return false }

func (t *FileSeq) AddProducer() Recorder { _ = "STUB: not implemented"; return *new(Recorder) }

func (t *FileSeq) record(eventTime time.Time, details Event) { _ = "STUB: not implemented"; return }

func (t *FileSeq) Run() { _ = "STUB: not implemented"; return }

func (t *FileSeq) encodeEvent(e event) { _ = "STUB: not implemented"; return }

func (t *FileSeq) removeProducer() { _ = "STUB: not implemented"; return }

type Writer struct {
	t *FileSeq
}

func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }

func (w *Writer) RecordEvent(ev Event) { _ = "STUB: not implemented"; return }
