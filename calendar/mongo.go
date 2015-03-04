package calendar

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoCalendar struct {
	mongo.Model      `bson:",inline"`
	models.UserOwned `bson:",inline"`

	EBaseScheduleID bson.ObjectId         `json:"base_schedule_id" bson:"base_schedule_id,omitempty"`
	EMonScheduleID  bson.ObjectId         `json:"monday_schedule_id" bson:"monday_schedule_id,omitempty"`
	ETueScheduleID  bson.ObjectId         `json:"tuesday_schedule_id" bson:"tuesday_schedule_id,omitempty"`
	EWedScheduleID  bson.ObjectId         `json:"wednesday_schedule_id" bson"wednesday_schedule_id,omitempty"`
	EThuScheduleID  bson.ObjectId         `json:"thursday_schedule_id" bson"thursday_schedule_id,omitempty"`
	EFriScheduleID  bson.ObjectId         `json:"friday_schedule_id" bson:"friday_schedule_id,omitempty"`
	ESatScheduleID  bson.ObjectId         `json:"saturday_schedule_id" bson:"saturday_schedule_id,omitempty"`
	ESunScheduleID  bson.ObjectId         `json:"sunday_schedule_id" bson:"sunday_schedule_id,omitempty"`
	ESchedules      map[int]bson.ObjectId `json:"schedules" bson:"schedules"`
}

func (c *mongoCalendar) Kind() data.Kind {
	return kind
}

func (c *mongoCalendar) Version() int {
	return version
}

func (c *mongoCalendar) Schema() data.Schema {
	return schema
}

func (c *mongoCalendar) SetUser(u models.User) error {
	return c.Schema().Link(c, u, User)
}

func (c *mongoCalendar) SetBase(s models.Schedule) error {
	return c.Schema().Link(c, s, Base)
}

func (c *mongoCalendar) SetMonday(s models.Schedule) error {
	return c.Schema().Link(c, s, Mon)
}

func (c *mongoCalendar) SetTuesday(s models.Schedule) error {
	return c.Schema().Link(c, s, Tue)
}

func (c *mongoCalendar) SetWednesday(s models.Schedule) error {
	return c.Schema().Link(c, s, Wed)
}

func (c *mongoCalendar) SetThursday(s models.Schedule) error {
	return c.Schema().Link(c, s, Thu)
}

func (c *mongoCalendar) SetFriday(s models.Schedule) error {
	return c.Schema().Link(c, s, Fri)
}

func (c *mongoCalendar) SetSaturday(s models.Schedule) error {
	return c.Schema().Link(c, s, Sat)
}

func (c *mongoCalendar) SetSunday(s models.Schedule) error {
	return c.Schema().Link(c, s, Sun)
}

func (c *mongoCalendar) Base(a data.Access, s models.Schedule) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	s.SetID(c.EBaseScheduleID)
	return a.PopulateByID(s)
}

func (c *mongoCalendar) Monday(a data.Access, s models.Schedule) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	s.SetID(c.EMonScheduleID)
	return a.PopulateByID(s)
}

func (c *mongoCalendar) Tuesday(a data.Access, s models.Schedule) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	s.SetID(c.ETueScheduleID)
	return a.PopulateByID(s)
}

func (c *mongoCalendar) Wednesday(a data.Access, s models.Schedule) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	s.SetID(c.EWedScheduleID)
	return a.PopulateByID(s)
}

func (c *mongoCalendar) Thursday(a data.Access, s models.Schedule) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	s.SetID(c.EThuScheduleID)
	return a.PopulateByID(s)
}

func (c *mongoCalendar) Friday(a data.Access, s models.Schedule) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	s.SetID(c.EFriScheduleID)
	return a.PopulateByID(s)
}

func (c *mongoCalendar) Saturday(a data.Access, s models.Schedule) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	s.SetID(c.ESatScheduleID)
	return a.PopulateByID(s)
}

func (c *mongoCalendar) Sunday(a data.Access, s models.Schedule) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	s.SetID(c.ESunScheduleID)
	return a.PopulateByID(s)
}

func (c *mongoCalendar) IncludeSchedule(s models.Schedule) error {
	return c.Schema().Link(c, s, Schedules)
}

func (c *mongoCalendar) ExcludeSchedule(s models.Schedule) error {
	return c.Schema().Unlink(c, s, Schedules)
}

func (c *mongoCalendar) Schedules(a data.Access) (data.ModelIterator, error) {
	ids := make(mongo.IDSet, 0)
	for _, id := range c.ESchedules {
		ids = mongo.AddID(ids, id)
	}

	if c.CanRead(a.Client()) {
		return mongo.NewIDIter(ids, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (c *mongoCalendar) ScheduleForDay(a data.Access, t time.Time) (models.Schedule, error) {
	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return nil, err
	}

	s, ok := m.(models.Schedule)

	id, ok := c.ESchedules[canonDay(t)]
	if !ok {
		return nil, data.ErrNotFound
	}

	s.SetID(id)
	err = a.PopulateByID(s)

	return s, err
}

func canonDay(t time.Time) int {
	return 100*int(t.Month()) + t.Day()
}

func (c *mongoCalendar) Link(m data.Model, l data.Link) error {
	if !data.Compatible(c, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		return c.SetUserID(m.ID())
	case Base:
		c.EBaseScheduleID = m.ID().(bson.ObjectId)
	case Mon:
		c.EMonScheduleID = m.ID().(bson.ObjectId)
	case Tue:
		c.ETueScheduleID = m.ID().(bson.ObjectId)
	case Wed:
		c.EWedScheduleID = m.ID().(bson.ObjectId)
	case Thu:
		c.EThuScheduleID = m.ID().(bson.ObjectId)
	case Fri:
		c.EFriScheduleID = m.ID().(bson.ObjectId)
	case Sat:
		c.ESatScheduleID = m.ID().(bson.ObjectId)
	case Sun:
		c.ESunScheduleID = m.ID().(bson.ObjectId)
	case Schedules:
		s, ok := m.(models.Schedule)
		if !ok {
			return data.NewLinkError(c, m, l)
		}

		c.ESchedules[canonDay(s.StartTime())] = s.ID().(bson.ObjectId)
	default:
		return data.NewLinkError(c, m, l)
	}

	return nil
}

func (c *mongoCalendar) Unlink(m data.Model, l data.Link) error {
	if !data.Compatible(c, m) {
		return data.ErrIncompatibleModels
	}

	id := m.ID().(bson.ObjectId)

	switch l.Name {
	case User:
		c.DropUserID()
	case Base:
		if c.EBaseScheduleID == id {
			c.EBaseScheduleID = *new(bson.ObjectId)
		}
	case Mon:
		if c.EMonScheduleID == id {
			c.EMonScheduleID = *new(bson.ObjectId)
		}
	case Tue:
		if c.ETueScheduleID == id {
			c.ETueScheduleID = *new(bson.ObjectId)
		}
	case Wed:
		if c.EWedScheduleID == id {
			c.EWedScheduleID = *new(bson.ObjectId)
		}
	case Thu:
		if c.EThuScheduleID == id {
			c.EThuScheduleID = *new(bson.ObjectId)
		}
	case Fri:
		if c.EFriScheduleID == id {
			c.EFriScheduleID = *new(bson.ObjectId)
		}
	case Sat:
		if c.ESatScheduleID == id {
			c.ESatScheduleID = *new(bson.ObjectId)
		}
	case Sun:
		if c.ESunScheduleID == id {
			c.ESunScheduleID = *new(bson.ObjectId)
		}
	case Schedules:
		s, ok := m.(models.Schedule)
		if !ok {
			return data.NewLinkError(c, m, l)
		}

		delete(c.ESchedules, canonDay(s.StartTime()))
	default:
		return data.NewLinkError(c, m, l)
	}
	return nil
}
