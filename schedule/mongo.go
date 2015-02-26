package schedule

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoSchedule struct {
	mongo.Model      `bson:",inline"`
	mongo.Named      `bson:",inline"`
	mongo.Timed      `bson:",inline"`
	models.UserOwned `bson:",inline"`

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

func (s *mongoSchedule) SetUser(u models.User) error {
	return s.Schema().Link(s, u, User)
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
