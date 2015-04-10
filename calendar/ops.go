package calendar

import (
	"time"

	"github.com/elos/models"
	"github.com/elos/models/schedule"
)

/*
	Returns this day in the elos canonical representation
	October 14th is 1014
*/
func Yearday(t time.Time) int {
	return 100*int(t.Month()) + t.Day()
}

func MergedScheduleForTime(store models.Store, c models.Calendar, t time.Time) (s models.Schedule, err error) {
	schedules := make([]models.Schedule, 0)

	base, err := c.BaseSchedule(store)
	if err != nil && err != models.ErrEmptyRelationship {
		return
	} else {
		schedules = append(schedules, base)
	}

	weekday, err := c.WeekdaySchedule(store, t.Weekday())
	if err != nil && err != models.ErrEmptyRelationship {
		return
	} else {
		schedules = append(schedules, weekday)
	}

	yearday, err := c.YeardaySchedule(store, t)
	if err != nil && err != models.ErrEmptyRelationship {
		return
	} else {
		schedules = append(schedules, yearday)
	}

	return schedule.Merge(store, schedules...)
}

func NextFixture(store models.Store, c models.Calendar) (first models.Fixture, err error) {
	s, err := MergedScheduleForTime(store, c, time.Now())
	if err != nil {
		return nil, err
	}

	return s.FirstFixture(store)
}

func CurrentFreeTime(store models.Store, c models.Calendar) time.Duration {
	f, err := NextFixture(store, c)
	if err != nil {
		return 0
	}

	return f.StartTime().Sub(time.Now())
}

func FreeTimeOnDay(store models.Store, c models.Calendar, t time.Time) time.Duration {
	/*
		s, _ := MergedScheduleForTime(a, c, t)
		iter, _ := s.Fixtures(a)

		m, _ := a.ModelFor(models.FixtureKind)
		f := m.(models.Fixture)
	*/

	return 0
}

func FreeTimeUntil(store models.Store, c models.Calendar, t time.Time) time.Duration {
	return 0
}
