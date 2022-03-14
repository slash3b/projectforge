package util

import (
	"time"
)

type Timer struct {
	Started   int64 `json:"started"`
	Completed int64 `json:"complete"`
}

func TimerStart() *Timer {
	return &Timer{Started: time.Now().UnixNano()}
}

func (t *Timer) End() int {
	t.Completed = time.Now().UnixNano()
	return t.Elapsed()
}

func (t *Timer) EndString() string {
	t.End()
	return t.String()
}

func (t *Timer) Elapsed() int {
	if t.Completed == 0 {
		return int((time.Now().UnixNano() - t.Started) / int64(time.Microsecond))
	}
	return int((t.Completed - t.Started) / int64(time.Microsecond))
}

func (t *Timer) String() string {
	return MicrosToMillis(t.Elapsed())
}
