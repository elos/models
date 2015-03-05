package calendar

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

type interval struct {
	Start time.Time
	End   time.Time
}

func (i *interval) Contains(other *interval) bool {
	return i.Start.Before(other.Start) && i.End.After(other.End)
}

func (i *interval) Dur() time.Duration {
	return i.Start.Sub(i.End)
}

func intervalFor(f data.Timeable) *interval {
	return &interval{
		Start: f.StartTime(),
		End:   f.EndTime(),
	}
}

type IntervalTree struct {
	Left  *IntervalTree
	Right *IntervalTree
	*interval
}

func (t *IntervalTree) Empty() bool {
	return t == new(IntervalTree)
}

func (t *IntervalTree) Start() time.Time {
	for !t.Left.Empty() {
		t = t.Left
	}

	return t.interval.Start
}

func (t *IntervalTree) End() time.Time {
	for !t.Right.Empty() {
		t = t.Right
	}

	return t.interval.End
}

func (t *IntervalTree) Add(f models.Fixture) {
	tf := intervalFor(f)
	tt := t.interval

	if t.Empty() {
		t.interval = tf
	}

	if tt.Contains(tf) {
		return
	}

	if tf.Contains(t.interval) {
		tt.Start = tf.Start
		tt.End = tf.End
	}

	if tt.End.Before(tf.Start) {
		t.Right.Add(f)
	}

	if tf.End.Before(tt.Start) {
		t.Left.Add(f)
	}

	if tt.Start.Before(tf.Start) && tf.Start.Before(tt.End) {
		tt.End = tf.End
	}

	if tf.Start.Before(tt.Start) && tt.Start.Before(tf.End) {
		tt.Start = tf.Start
	}
}

func (t *IntervalTree) DurWalk() <-chan time.Duration {
	c := make(chan time.Duration)
	go t.rwalk(c)
	return c
}

func (t *IntervalTree) rwalk(c chan time.Duration) {
	if t.Empty() {
		return
	}

	t.Left.rwalk(c)
	c <- t.interval.Dur()
	t.Right.rwalk(c)
}
