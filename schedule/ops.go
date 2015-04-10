package schedule

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/structures"
)

func FirstFixture(s models.Schedule, store models.Store) (first models.Fixture, err error) {
	return EarliestSince(s, store, *new(time.Time))
}

func EarliestSince(s models.Schedule, store models.Store, start time.Time) (models.Fixture, error) {
	iter, err := s.FixturesIter(store)
	if err != nil {
		return nil, err
	}

	fixtures, err := OrderFixtures(store, iter)
	if err != nil {
		return nil, err
	}

	for _, f := range fixtures {
		if start.Before(f.StartTime()) {
			return f, nil
		}
	}

	return nil, data.ErrNotFound
}

func OrderedFixtures(s models.Schedule, store models.Store) ([]models.Fixture, error) {
	iter, err := s.FixturesIter(store)
	if err != nil {
		return nil, err
	}
	return OrderFixtures(store, iter)
}

func OrderFixtures(store models.Store, iter data.ModelIterator) ([]models.Fixture, error) {
	tree := new(structures.TimeableTree)
	c := 0

	f := store.Fixture()
	for iter.Next(f) {
		tree.Add(f)
		f = store.Fixture()
		c++
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	fixtures := make([]models.Fixture, 0)
	s := tree.Stream()
	for i := 0; i < c; i++ {
		of := <-s
		fixtures = append(fixtures, of.(models.Fixture))
	}

	return fixtures, nil
}
