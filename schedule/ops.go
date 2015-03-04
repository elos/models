package schedule

import (
	"errors"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

func FirstFixture(a data.Access, s models.Schedule) (first models.Fixture, err error) {
	return EarliestSince(a, s, *new(time.Time))
}

var castError error = errors.New("wow I need an errors for cast failures")

func EarliestSince(a data.Access, s models.Schedule, start time.Time) (next models.Fixture, err error) {
	iter, err := s.Fixtures(a)
	if err != nil {
		return
	}

	m, err := a.ModelFor(models.FixtureKind)
	if err != nil {
		return
	}

	earliest := *new(time.Time)

	for iter.Next(m) {
		f, ok := m.(models.Fixture)

		if !ok {
			err = castError
			return
		}

		if start.Before(f.StartTime()) {
			next = f
			earliest = next.StartTime()
			m, _ = a.ModelFor(models.FixtureKind)
		}
	}

	if earliest.IsZero() {
		err = errors.New("No fixutres since start")
		return
	}

	// Now that we have the first valid fixture
	// lets see if anything can beat it
	for iter.Next(m) {
		f, ok := m.(models.Fixture)

		if !ok {
			err = castError
			return
		}

		if start.Before(f.StartTime()) && f.StartTime().Before(earliest) {
			next = f
			m, _ = (a.ModelFor(models.FixtureKind))
		}
	}

	if err = iter.Close(); err != nil {
		return
	}

	// atm kinda funny double rturn
	// but may in the future be other things to do
	return

}
