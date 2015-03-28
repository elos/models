package schedule

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

func Merge(a data.Access, schedules ...models.Schedule) (s models.Schedule, err error) {
	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return
	}

	s, ok := m.(models.Schedule)
	if !ok {
		err = models.CastError(models.ScheduleKind)
		return
	}

	m, err = a.ModelFor(models.FixtureKind)
	if err != nil {
		return
	}

	f, ok := m.(models.Fixture)
	if !ok {
		err = models.CastError(models.FixtureKind)
		return
	}

	for _, schedule := range schedules {
		iter, e := schedule.FixturesIter(a)
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
