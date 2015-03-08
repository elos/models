package structures

import (
	"time"

	"github.com/elos/data"
)

type OverlapKind int

const (
	Before OverlapKind = iota
	StartBefore
	Contains
	Contained
	StartAfter
	After
)

type Interval struct {
	Start time.Time
	End   time.Time
}

func NewInterval(s time.Time, e time.Time) *Interval {
	return &Interval{
		Start: s,
		End:   e,
	}
}

func IntervalFor(t data.Timeable) *Interval {
	return NewInterval(t.StartTime(), t.EndTime())
}

func (i *Interval) Overlap(other *Interval) OverlapKind {
	if i.Contains(other) {
		return Contains
	}

	if other.Contains(i) {
		return Contained
	}

	if i.End.Before(other.Start) {
		return Before
	}

	if other.End.Before(i.Start) {
		return After
	}

	if i.Start.Before(other.Start) && other.Start.Before(i.End) {
		return StartBefore
	}

	if other.Start.Before(i.Start) && i.Start.Before(other.End) {
		return StartAfter
	}

	panic("Shouldn't be possible")
}

func (i *Interval) Contains(other *Interval) bool {
	return i.Overlap(other) == Contains
}

func (i *Interval) Dur() time.Duration {
	return i.Start.Sub(i.End)
}

type IntervalTree struct {
	Left  *IntervalTree
	Right *IntervalTree
	*Interval
}

var EmptyIntervalTree = new(IntervalTree)

func (t *IntervalTree) Empty() bool {
	return *t == *EmptyIntervalTree
}

func (t *IntervalTree) Start() time.Time {
	for !t.Left.Empty() {
		t = t.Left
	}

	return t.Interval.Start
}

func (t *IntervalTree) End() time.Time {
	for !t.Right.Empty() {
		t = t.Right
	}

	return t.Interval.End
}

func (t *IntervalTree) Add(other *Interval) {
	if t.Empty() {
		t.Interval = other
	}

	i := t.Interval

	switch i.Overlap(other) {
	case Contains:
		return
	case Contained:
		t.Interval = other
	case Before:
		t.Right.Add(other)
	case After:
		t.Left.Add(other)
	case StartBefore:
		i.End = other.End
	case StartAfter:
		i.Start = other.Start
	}
}

func (t *IntervalTree) DurWalk() <-chan time.Duration {
	c := make(chan time.Duration)
	go func(dc chan time.Duration) {
		t.rwalk(dc)
		close(dc)
	}(c)
	return c
}

func (t *IntervalTree) rwalk(c chan time.Duration) {
	if t.Empty() {
		return
	}

	t.Left.rwalk(c)
	c <- t.Interval.Dur()
	t.Right.rwalk(c)
}
