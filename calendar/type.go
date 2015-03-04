package calendar

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User             data.LinkName = models.CalendarUser
	Base             data.LinkName = models.CalendarBase
	WeekdaySchedules data.LinkName = models.CalendarWeekdaySchedules
	Schedules        data.LinkName = models.CalendarSchedules
	CurrentFixture   data.LinkName = models.CalendarCurrentFixture
)

var (
	kind    data.Kind   = models.CalendarKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Calendar, error) {
	switch s.Type() {
	case mongo.DBType:
		c := &mongoCalendar{}
		c.SetID(s.NewID())
		return c, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}
