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

func NewM(s data.Store) data.Model {
	return New(s)
}

func New(s data.Store) models.Schedule {
	switch s.Type() {
	case mongo.DBType:
		sched := &mongoSchedule{}
		sched.SetID(s.NewID())
		return sched
	default:
		panic(data.ErrInvalidDBType)
	}
}

func Create(s data.Store) (models.Schedule, error) {
	sched := New(s)
	return sched, s.Save(sched)
}
