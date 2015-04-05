package calendar

import (
	"time"

	"github.com/elos/data"
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

func MergedScheduleForTime(a data.Access, c models.Calendar, t time.Time) (s models.Schedule, err error) {
	schedules := make([]models.Schedule, 0)

	base, err := c.BaseSchedule(a)
	if err != nil {
		// We don't mind if it's an empty link error
		if _, ok := err.(*data.EmptyLinkError); !ok {
			return
		}
	} else {
		schedules = append(schedules, base)
	}

	weekday, err := c.WeekdaySchedule(a, t.Weekday())
	if err != nil {
		if _, ok := err.(*data.EmptyLinkError); !ok {
			return
		}
	} else {
		schedules = append(schedules, weekday)
	}

	yearday, err := c.YeardaySchedule(a, t)
	if err != nil {
		if _, ok := err.(*data.EmptyLinkError); !ok {
			return
		}
	} else {
		schedules = append(schedules, yearday)
	}

	return schedule.Merge(a, schedules...)
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
