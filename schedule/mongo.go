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

func (s *mongoSchedule) SetUser(u models.User) error {
	return s.Schema().Link(s, u, User)
}

func (s *mongoSchedule) IncludeFixture(f models.Fixture) error {
	return s.Schema().Link(s, f, Fixtures)
}

func (s *mongoSchedule) ExcludeFixture(f models.Fixture) error {
	return s.Schema().Unlink(s, f, Fixtures)
}

func (s *mongoSchedule) FixturesIter(store models.Store) (data.ModelIterator, error) {
	if !store.Compatible(s) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(s.EFixtureIDs, store), nil
}

func (s *mongoSchedule) Fixtures(store models.Store) ([]models.Fixture, error) {
	if !store.Compatible(s) {
		return nil, data.ErrInvalidDBType
	}

	fixtures := make([]models.Fixture, 0)
	iter := mongo.NewIDIter(s.EFixtureIDs, store)
	fixture := store.Fixture()
	for iter.Next(fixture) {
		fixtures = append(fixtures, fixture)
		fixture = store.Fixture()
	}

	return fixtures, iter.Close()

}

func (s *mongoSchedule) FirstFixture(store models.Store) (models.Fixture, error) {
	return FirstFixture(s, store)
}

func (s *mongoSchedule) FirstFixtureSince(store models.Store, t time.Time) (models.Fixture, error) {
	return EarliestSince(s, store, t)
}

func (s *mongoSchedule) OrderedFixtures(store models.Store) ([]models.Fixture, error) {
	return OrderedFixtures(s, store)
}
