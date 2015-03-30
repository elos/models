package schedule

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/structures"
)

func FirstFixture(s models.Schedule, a data.Access) (first models.Fixture, err error) {
	return EarliestSince(s, a, *new(time.Time))
}

func EarliestSince(s models.Schedule, a data.Access, start time.Time) (models.Fixture, error) {
	iter, _ := s.FixturesIter(a)
	fixtures, _ := OrderFixtures(a, iter)

	for _, f := range fixtures {
		if start.Before(f.StartTime()) {
			return f, nil
		}
	}

	return nil, data.ErrNotFound
}

func OrderedFixtures(s models.Schedule, a data.Access) ([]models.Fixture, error) {
	iter, err := s.FixturesIter(a)
	if err != nil {
		return nil, err
	}
	return OrderFixtures(a, iter)
}

func OrderFixtures(a data.Access, iter data.ModelIterator) ([]models.Fixture, error) {
	m, _ := a.ModelFor(models.FixtureKind)
	f := m.(models.Fixture)
	tree := new(structures.TimeableTree)
	fixtures := make([]models.Fixture, 0)
	c := 0

	for iter.Next(f) {
		tree.Add(f)
		c++
		m, _ = a.ModelFor(models.FixtureKind)
		f = m.(models.Fixture)
	}

	if err := iter.Close(); err != nil {
		return fixtures, err
	}

	s := tree.Stream()
	for i := 0; i < c; i++ {
		of := <-s
		fixtures = append(fixtures, of.(models.Fixture))
	}

	return fixtures, nil
}
