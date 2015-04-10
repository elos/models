package calendar

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

type FixtureStream struct {
	iter data.ModelIterator
	t    time.Time

	models.Store
	models.Calendar

	err error
}

func NewFixtureStream(store models.Store, c models.Calendar, start time.Time) {

}

func (s FixtureStream) Next(m data.Model) bool {
	if s.err != nil {
		return false
	}

	if s.iter == nil {
		return false
	}

	try := s.iter.Next(m)
	if !try {
		if err := s.Close(); err != nil {
			s.err = err
			return false
		}
		s.advanceDay()
		return s.Next(m)
	} else {
		return true
	}
}

func (s FixtureStream) Close() error {
	if s.err != nil {
		return s.err
	} else if s.iter == nil {
		return nil
	} else {
		return s.iter.Close()
	}
}

func (f FixtureStream) advanceDay() {
	diff := time.Date(f.t.Year(), f.t.Month(), f.t.Day(), 0, 0, 0, 0, time.UTC).Sub(f.t)
	fDiff := 24*time.Hour - diff
	f.t = f.t.Add(fDiff).Round(time.Hour)

	s, e := MergedScheduleForTime(f.Store, f.Calendar, f.t)
	if e != nil {
		f.err = e
		return
	}

	iter, e := s.FixturesIter(f.Store)
	if e != nil {
		f.err = e
		return
	}

	f.iter = iter
}
