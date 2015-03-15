package calendar

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

/*
	Returns this day in the elos canonical representation
	October 14th is 1014
*/
func canonDay(t time.Time) int {
	return 100*int(t.Month()) + t.Day()
}

func MergeSchedules(a data.Access, schedules ...models.Schedule) (s models.Schedule, err error) {
	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return
	}

	s, ok := m.(models.Schedule)
	if !ok {
		err = models.CastError(models.ScheduleKind)
		return
	}

	m, err = a.ModelFor(models.ScheduleKind)
	if err != nil {
		return
	}
	f, ok := m.(models.Fixture)
	if !ok {
		err = models.CastError(models.FixtureKind)
		return
	}

	for _, schedule := range schedules {
		iter, e := schedule.Fixtures(a)
		if e != nil {
			return
		}

		for iter.Next(f) {
			s.IncludeFixture(f)
		}

		err = iter.Close()
		if err != nil {
			break
		}
	}

	return
}

func MergedScheduleForTime(a data.Access, c models.Calendar, t time.Time) (s models.Schedule, err error) {
	base, err := c.BaseSchedule(a)
	if err != nil {
		return
	}

	weekday, err := c.WeekdaySchedule(a, t.Weekday())
	if err != nil {
		return
	}

	day, err := c.ScheduleForDay(a, t)
	if err != nil {
		return
	}

	return MergeSchedules(a, base, weekday, day)
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
