package calendar

import (
	"errors"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoCalendar struct {
	mongo.Model      `bson:",inline"`
	models.UserOwned `bson:",inline"`

	EBaseScheduleID   bson.ObjectId                  `json:"base_schedule_id" bson:"base_schedule_id,omitempty"`
	EWeekdaySchedules map[time.Weekday]bson.ObjectId `json:"weekday_schedules" bson:"weekday_schedules"`
	ESchedules        map[int]bson.ObjectId          `json:"schedules" bson:"schedules"`

	ECurrentFixtureID bson.ObjectId `json:"current_fixture_id", bson:"current_fixture_id,omitempty"`
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

func (c *mongoCalendar) SetWeekdaySchedule(s models.Schedule, t time.Weekday) error {
	c.EWeekdaySchedules[t] = s.ID().(bson.ObjectId)
	return nil
}

func (c *mongoCalendar) Base(a data.Access) (s models.Schedule, err error) {
	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return
	}

	s, ok := m.(models.Schedule)
	if !ok {
		err = errors.New("cast error")
		return
	}

	s.SetID(c.EBaseScheduleID)

	err = a.PopulateByID(s)
	return
}

func (c *mongoCalendar) WeekdaySchedule(a data.Access, t time.Weekday) (s models.Schedule, err error) {
	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return
	}

	s, ok := m.(models.Schedule)
	if !ok {
		err = errors.New("I NEED A CAST ERROR")
		return
	}

	id, ok := c.EWeekdaySchedules[t]
	if !ok {
		err = data.ErrNotFound
		return
	}

	s.SetID(id)

	err = a.PopulateByID(s)

	return
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

func (c *mongoCalendar) SetCurrentFixture(f models.Fixture) error {
	return c.Schema().Link(c, f, CurrentFixture)
}

func (c *mongoCalendar) FindNextFixture(a data.Access) (err error) {
	base, err := c.Base(a)
	if err != nil {
		return
	}

	first, err := base.FirstFixture(a)
	if err != nil {
		return
	}

	if weekday, e1 := c.WeekdaySchedule(a, time.Now().Weekday()); e1 != nil {
		wfirst, e2 := weekday.FirstFixture(a)
		if e2 != nil {
			err = e2
			return
		}

		if wfirst.Before(first) {
			first = wfirst
		}
	} else {
		err = e1
		return
	}

	if day, e1 := c.ScheduleForDay(a, time.Now()); e1 != nil {
		dfirst, e2 := day.FirstFixture(a)
		if e2 != nil {
			err = e2
			return
		}

		if dfirst.Before(first) {
			first = dfirst
		}
	} else {
		err = e1
		return
	}

	c.SetCurrentFixture(first)
	return a.Save(c)
}

func (c *mongoCalendar) NextAction(a data.Access) (action models.Action, err error) {
	current, err := c.CurrentFixture(a)
	if err != nil {
		return
	}

	action, err = current.NextAction(a)

	return
}

func (c *mongoCalendar) CurrentFixture(a data.Access) (models.Fixture, error) {
	m, err := a.ModelFor(models.FixtureKind)
	if err != nil {
		return nil, err
	}

	f, ok := m.(models.Fixture)
	if !ok {
		return nil, errors.New("TODO")
	}

	f.SetID(c.ECurrentFixtureID)
	err = a.PopulateByID(f)

	return f, err
}

func (c *mongoCalendar) CompleteAction(access data.Access, action models.Action) error {
	fixture, err := c.CurrentFixture(access)
	if err != nil {
		return err
	}

	return fixture.CompleteAction(access, action)
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
	case Schedules:
		s, ok := m.(models.Schedule)
		if !ok {
			return data.NewLinkError(c, m, l)
		}

		c.ESchedules[canonDay(s.StartTime())] = s.ID().(bson.ObjectId)
	case CurrentFixture:
		c.ECurrentFixtureID = m.ID().(bson.ObjectId)
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
	case Schedules:
		s, ok := m.(models.Schedule)
		if !ok {
			return data.NewLinkError(c, m, l)
		}

		delete(c.ESchedules, canonDay(s.StartTime()))
	case CurrentFixture:
		if c.ECurrentFixtureID == id {
			c.ECurrentFixtureID = *new(bson.ObjectId)
		}
	default:
		return data.NewLinkError(c, m, l)
	}
	return nil
}
