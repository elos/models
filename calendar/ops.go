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
	base, err := c.BaseSchedule(a)
	if err != nil {
		return
	}

	weekday, err := c.WeekdaySchedule(a, t.Weekday())
	if err != nil {
		return
	}

	day, err := c.YeardaySchedule(a, t)
	if err != nil {
		return
	}

	return schedule.Merge(a, base, weekday, day)
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
