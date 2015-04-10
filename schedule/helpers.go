package schedule

import "github.com/elos/models"

func Merge(store models.Store, schedules ...models.Schedule) (models.Schedule, error) {
	s := store.Schedule()

	for _, schedule := range schedules {
		iter, err := schedule.FixturesIter(store)
		if err != nil {
			return nil, err
		}

		f := store.Fixture()
		for iter.Next(f) {
			s.IncludeFixture(f)
			f = store.Fixture()
		}

		if err := iter.Close(); err != nil {
			return nil, err
		}
	}

	return s, nil
}
