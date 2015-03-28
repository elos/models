package calendar

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/shared"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoCalendar struct {
	mongo.Model           `bson:",inline"`
	shared.MongoUserOwned `bson:",inline"`

	EBaseScheduleID   bson.ObjectId            `json:"base_schedule_id"  bson:"base_schedule_id,omitempty"`
	EWeekdaySchedules map[string]bson.ObjectId `json:"weekday_schedules" bson:"weekday_schedules"`
	EYeardaySchedules map[string]bson.ObjectId `json:"schedules"         bson:"schedules"`

	ECurrentFixtureID bson.ObjectId `json:"current_fixture_id" bson:"current_fixture_id,omitempty"`
}

func newMongoCalendar() *mongoCalendar {
	c := &mongoCalendar{}
	c.EWeekdaySchedules = make(map[string]bson.ObjectId)
	c.EYeardaySchedules = make(map[string]bson.ObjectId)
	return c
}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (c *mongoCalendar) Kind() data.Kind {
	return kind
}

// Schema is derived from the models package and is
// defined in type.go, shared among implementations
func (c *mongoCalendar) Schema() data.Schema {
	return schema
}

// Version is derieved from teh models package and is
// defined in type.go, shared among implementations
func (c *mongoCalendar) Version() int {
	return version
}

func (c *mongoCalendar) Link(m data.Model, l data.Link) error {
	if !data.Compatible(c, m) {
		return data.ErrIncompatibleModels
	}

	id, ok := m.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	switch l.Name {
	case user:
		return c.SetUserID(id)
	case baseSchedule:
		c.EBaseScheduleID = id
	case yeardaySchedules:
		panic("link schedules not implemented")
	case currentFixture:
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

	id, ok := m.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	switch l.Name {
	case user:
		if c.UserID().String() == id.String() {
			c.DropUserID()
		}
	case baseSchedule:
		if c.EBaseScheduleID.String() == id.String() {
			c.EBaseScheduleID = mongo.NewEmptyID()
		}
	case weekdaySchedules:

	case yeardaySchedules:
		panic("link schedules not implemented")
	case currentFixture:
		if c.ECurrentFixtureID == id {
			c.ECurrentFixtureID = *new(bson.ObjectId)
		}
	default:
		return data.NewLinkError(c, m, l)
	}
	return nil
}

func (c *mongoCalendar) SetUser(u models.User) error {
	return c.Schema().Link(c, u, user)
}

func (c *mongoCalendar) SetBaseSchedule(s models.Schedule) error {
	return c.Schema().Link(c, s, baseSchedule)
}

func (c *mongoCalendar) BaseSchedule(a data.Access) (models.Schedule, error) {
	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return nil, err
	}

	s, ok := m.(models.Schedule)
	if !ok {
		return nil, models.CastError(models.ScheduleKind)
	}

	if mongo.EmptyID(c.EBaseScheduleID) {
		return nil, data.NewEmptyLinkError(c, models.RMap[models.CalendarKind][baseSchedule])
	}

	s.SetID(c.EBaseScheduleID)

	if !c.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	return s, a.PopulateByID(s)
}

func (c *mongoCalendar) SetWeekdaySchedule(s models.Schedule, t time.Weekday) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	id, ok := s.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	c.EWeekdaySchedules[t.String()] = id
	return nil
}

func (c *mongoCalendar) WeekdaySchedule(a data.Access, t time.Weekday) (models.Schedule, error) {
	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return nil, err
	}

	s, ok := m.(models.Schedule)
	if !ok {
		return nil, models.CastError(models.ScheduleKind)
	}

	id, ok := c.EWeekdaySchedules[t.String()]
	if !ok {
		return nil, data.NewEmptyLinkError(c, models.RMap[models.CalendarKind][weekdaySchedules])
	}

	if !c.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	s.SetID(id)

	return s, a.PopulateByID(s)
}

func (c *mongoCalendar) SetYeardaySchedule(s models.Schedule, t time.Time) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	id, ok := s.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	c.EYeardaySchedules[string(Yearday(t))] = id
	return nil
}

func (c *mongoCalendar) YeardaySchedule(a data.Access, t time.Time) (models.Schedule, error) {
	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return nil, err
	}

	s, ok := m.(models.Schedule)

	id, ok := c.EYeardaySchedules[string(Yearday(t))]
	if !ok {
		return nil, data.NewEmptyLinkError(c, models.RMap[models.CalendarKind][yeardaySchedules])
	}

	s.SetID(id)
	err = a.PopulateByID(s)

	return s, err
}

func (c *mongoCalendar) SetCurrentFixture(f models.Fixture) error {
	return c.Schema().Link(c, f, currentFixture)
}

func (c *mongoCalendar) CurrentFixture(a data.Access) (models.Fixture, error) {
	if mongo.EmptyID(c.ECurrentFixtureID) {
		return nil, data.NewEmptyLinkError(c, models.RMap[models.CalendarKind][currentFixture])
	}

	m, err := a.ModelFor(models.FixtureKind)
	if err != nil {
		return nil, err
	}

	f, ok := m.(models.Fixture)
	if !ok {
		return nil, models.CastError(models.FixtureKind)
	}

	if !c.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	f.SetID(c.ECurrentFixtureID)
	return f, a.PopulateByID(f)
}

func (c *mongoCalendar) NextFixture(a data.Access) (first models.Fixture, err error) {
	return NextFixture(a, c)
}

func (c *mongoCalendar) NextAction(a data.Access) (action models.Action, err error) {
	return NextAction(c, a)
}

func (c *mongoCalendar) StartAction(access data.Access, action models.Action) error {
	return StartAction(c, access, action)
}

func (c *mongoCalendar) CompleteAction(access data.Access, action models.Action) error {
	return CompleteAction(c, access, action)
}
