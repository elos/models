package schedule

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/shared"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoSchedule struct {
	mongo.Model           `bson:",inline"`
	mongo.Named           `bson:",inline"`
	mongo.Timed           `bson:",inline"`
	shared.MongoUserOwned `bson:",inline"`

	EFixtureIDs mongo.IDSet `json:"fixture_ids" bson:"fixture_ids"`
}

func (s *mongoSchedule) Kind() data.Kind {
	return kind
}

func (s *mongoSchedule) Version() int {
	return version
}

func (s *mongoSchedule) Schema() data.Schema {
	return schema
}

func (s *mongoSchedule) IncludeFixture(f models.Fixture) error {
	return s.Schema().Link(s, f, Fixtures)
}

func (s *mongoSchedule) ExcludeFixture(f models.Fixture) error {
	return s.Schema().Unlink(s, f, Fixtures)
}

func (s *mongoSchedule) FixturesIter(a data.Access) (data.ModelIterator, error) {
	if !s.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	return mongo.NewIDIter(s.EFixtureIDs, a), nil
}

func (s *mongoSchedule) Fixtures(a data.Access) ([]models.Fixture, error) {
	if !s.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	fixtures := make([]models.Fixture, 0)

	iter, err := s.FixturesIter(a)
	if err != nil {
		return fixtures, err
	}

	m, err := a.ModelFor(models.FixtureKind)
	if err != nil {
		return fixtures, err
	}

	for iter.Next(m) {
		e, ok := m.(models.Fixture)
		if !ok {
			return fixtures, models.CastError(models.FixtureKind)
		}

		fixtures = append(fixtures, e)

		m, err = a.ModelFor(models.FixtureKind)
		if err != nil {
			return fixtures, err
		}
	}

	return fixtures, nil

}

func (s *mongoSchedule) OrderedFixtures(a data.Access) ([]models.Fixture, error) {
	iter, _ := s.FixturesIter(a)
	return OrderFixtures(a, iter)
}

func (s *mongoSchedule) SetUser(u models.User) error {
	return s.Schema().Link(s, u, User)
}

func (s *mongoSchedule) FirstFixture(a data.Access) (models.Fixture, error) {
	return FirstFixture(a, s)
}

func (s *mongoSchedule) FirstFixtureSince(a data.Access, t time.Time) (models.Fixture, error) {
	return EarliestSince(a, s, t)
}

func (s *mongoSchedule) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case Fixtures:
		s.EFixtureIDs = mongo.AddID(s.EFixtureIDs, m.ID().(bson.ObjectId))
		return nil
	case User:
		return s.SetUserID(m.ID())
	default:
		return data.NewLinkError(s, m, l)
	}
}

func (s *mongoSchedule) Unlink(m data.Model, l data.Link) error {
	switch l.Name {
	case Fixtures:
		s.EFixtureIDs = mongo.DropID(s.EFixtureIDs, m.ID().(bson.ObjectId))
		return nil
	case User:
		s.DropUserID()
		return nil
	default:
		return data.NewLinkError(s, m, l)
	}
}
