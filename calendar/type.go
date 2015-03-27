package calendar

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

// variables shared among implementations
var (
	kind    data.Kind   = models.CalendarKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion

	user             data.LinkName = models.CalendarUser
	baseSchedule     data.LinkName = models.CalendarBaseSchedule
	weekdaySchedules data.LinkName = models.CalendarWeekdaySchedules
	schedules        data.LinkName = models.CalendarSchedules
	currentFixture   data.LinkName = models.CalendarCurrentFixture
)

/*
	NewM has the same signature as New except that
	it loosens the type of the return to data.Model
	rather than models.Calendar. Use it to register
	the models.CalendarKind with the store.

		var s data.Store
		s.Register(models.CalendarKind, calendar.NewM)
*/
func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

/*
	New returns an instantiated calendar. It sets
	the id using the NewID() method on the store.

		var s data.Store
		c, err := calendar.New(s)
		if err != nil {
			// don't use c, it is nil
		}

	Currently supports the mongo db type.
*/
func New(s data.Store) (models.Calendar, error) {
	var c models.Calendar

	switch s.Type() {
	case mongo.DBType:
		c = newMongoCalendar()
	default:
		return nil, data.ErrInvalidDBType
	}

	c.SetID(s.NewID())
	return c, nil
}

func Create(s data.Store) (models.Calendar, error) {
	c, err := New(s)
	if err != nil {
		return c, err
	}

	return c, s.Save(c)
}
