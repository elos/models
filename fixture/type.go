package fixture

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User     data.LinkName = models.FixtureUser
	Schedule data.LinkName = models.FixtureSchedule
	Actions  data.LinkName = models.FixtureActions
	Events   data.LinkName = models.FixtureEvents
)

var (
	kind    data.Kind   = models.FixtureKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) data.Model {
	return New(s)
}

func New(s data.Store) models.Fixture {
	switch s.Type() {
	case mongo.DBType:
		f := &mongoFixture{}
		f.SetID(s.NewID())

		return f
	default:
		panic(data.ErrInvalidDBType)
	}
}

func Create(s data.Store) (models.Fixture, error) {
	f := New(s)
	return f, s.Save(f)
}
