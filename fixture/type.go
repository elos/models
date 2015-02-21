package fixture

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User     data.LinkName = models.FixtureUser
	Schedule data.LinkName = models.FixtureSchedule
)

var (
	kind    data.Kind   = models.FixtureKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Fixture, error) {
	switch s.Type() {
	case mongo.DBType:
		return &mongoFixture{}, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}
