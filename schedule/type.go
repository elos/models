package schedule

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User     data.LinkName = models.ScheduleUser
	Fixtures data.LinkName = models.ScheduleFixtures
)

var (
	kind    data.Kind   = models.ScheduleKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Schedule, error) {
	switch s.Type() {
	case mongo.DBType:
		return &mongoSchedule{}, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}
