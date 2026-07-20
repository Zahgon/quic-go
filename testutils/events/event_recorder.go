package events

import (
	"sync"
	"time"

	"github.com/quic-go/quic-go/qlogwriter"
)

type Event struct {
	Time  time.Time
	Event qlogwriter.Event
}

type Trace struct {
	Recorder qlogwriter.Recorder
}

var _ qlogwriter.Trace = &Trace{}

func (t *Trace) AddProducer() qlogwriter.Recorder {
	_ = "STUB: not implemented"
	return *new(qlogwriter.Recorder)
}

func (t *Trace) SupportsSchemas(string) bool { _ = "STUB: not implemented"; return false }

type Recorder struct {
	mx     sync.Mutex
	events []Event
}

var _ qlogwriter.Recorder = &Recorder{}

func (r *Recorder) RecordEvent(ev qlogwriter.Event) { _ = "STUB: not implemented"; return }

func (r *Recorder) Events(filter ...qlogwriter.Event) []qlogwriter.Event {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recorder) EventsWithTime(filter ...qlogwriter.Event) []Event {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recorder) Clear() { _ = "STUB: not implemented"; return }

func (r *Recorder) Close() error { _ = "STUB: not implemented"; return nil }
