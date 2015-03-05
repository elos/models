package calendar

import (
	"errors"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

func MergedScheduleForTime(a data.Access, c models.Calendar, t time.Time) (s models.Schedule, err error) {
	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return nil, err
	}

	s, ok := m.(models.Schedule)
	if !ok {
		return nil, errors.New("calnedar ops line 19")
	}

	base, err := c.Base(a)
	if err != nil {
		return nil, err
	}

	iter, err := base.Fixtures(a)
	if err != nil {
		return nil, err
	}

	m, err = a.ModelFor(models.FixtureKind)
	if err != nil {
		return nil, err
	}

	f, ok := m.(models.Fixture)
	if !ok {
		return nil, errors.New("adsf")
	}

	for iter.Next(f) {
		s.IncludeFixture(f)
	}

	weekday, err := c.WeekdaySchedule(a, t.Weekday())
	if err != nil {
		return nil, err
	}

	iter, err = weekday.Fixtures(a)
	if err != nil {
		return nil, err
	}

	for iter.Next(f) {
		s.IncludeFixture(f)
	}

	day, err := c.ScheduleForDay(a, t)
	if err != nil {
		return nil, err
	}

	iter, err = day.Fixtures(a)
	if err != nil {
		return nil, err
	}

	for iter.Next(f) {
		s.IncludeFixture(f)
	}

	return s, nil

}

func NextFixture(a data.Access, c models.Calendar) (first models.Fixture, err error) {
	s, err := MergedScheduleForTime(a, c, time.Now())
	if err != nil {
		return nil, err
	}

	return s.FirstFixture(a)
}

func CurrentFreeTime(a data.Access, c models.Calendar) time.Duration {
	f, err := NextFixture(a, c)
	if err != nil {
		return 0
	}

	return f.StartTime().Sub(time.Now())
}

func FreeTimeOnDay(a data.Access, c models.Calendar, t time.Time) time.Duration {
	/*
		s, _ := MergedScheduleForTime(a, c, t)
		iter, _ := s.Fixtures(a)

		m, _ := a.ModelFor(models.FixtureKind)
		f := m.(models.Fixture)
	*/

	return 0
}

func FreeTimeUntil(a data.Access, c models.Calendar, t time.Time) time.Duration {
	return 0
}
