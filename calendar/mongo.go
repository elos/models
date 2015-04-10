package calendar

import (
	"strings"
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

func (c *mongoCalendar) BaseSchedule(store models.Store) (models.Schedule, error) {
	if !store.Compatible(c) {
		return nil, data.ErrInvalidDBType
	}

	if mongo.EmptyID(c.EBaseScheduleID) {
		return nil, models.ErrEmptyRelationship
	}

	schedule := store.Schedule()
	schedule.SetID(c.EBaseScheduleID)
	return schedule, store.PopulateByID(schedule)
}

func (c *mongoCalendar) SetWeekdaySchedule(s models.Schedule, t time.Weekday) error {
	if !data.Compatible(c, s) {
		return data.ErrIncompatibleModels
	}

	id, ok := s.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	c.EWeekdaySchedules[strings.ToLower(t.String())] = id
	return nil
}

func (c *mongoCalendar) WeekdaySchedule(s models.Store, t time.Weekday) (models.Schedule, error) {
	if !s.Compatible(c) {
		return nil, data.ErrInvalidDBType
	}

	id, ok := c.EWeekdaySchedules[strings.ToLower(t.String())]
	if !ok {
		return nil, models.ErrEmptyRelationship
	}

	schedule := s.Schedule()
	schedule.SetID(id)
	return schedule, s.PopulateByID(schedule)
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

func (c *mongoCalendar) YeardaySchedule(store models.Store, t time.Time) (models.Schedule, error) {
	if !store.Compatible(c) {
		return nil, data.ErrInvalidDBType
	}

	id, ok := c.EYeardaySchedules[string(Yearday(t))]
	if !ok {
		return nil, models.ErrEmptyRelationship
	}

	schedule := store.Schedule()
	schedule.SetID(id)
	return schedule, store.PopulateByID(schedule)
}

func (c *mongoCalendar) SetCurrentFixture(f models.Fixture) error {
	return c.Schema().Link(c, f, currentFixture)
}

func (c *mongoCalendar) CurrentFixture(store models.Store) (models.Fixture, error) {
	if !store.Compatible(c) {
		return nil, data.ErrInvalidDBType
	}

	if mongo.EmptyID(c.ECurrentFixtureID) {
		return nil, models.ErrEmptyRelationship
	}

	fixture := store.Fixture()
	fixture.SetID(c.ECurrentFixtureID)
	return fixture, store.PopulateByID(fixture)
}

func (c *mongoCalendar) NextFixture(store models.Store) (first models.Fixture, err error) {
	return NextFixture(store, c)
}

func (c *mongoCalendar) NextAction(s models.Store) (action models.Action, err error) {
	return NextAction(c, s)
}

func (c *mongoCalendar) StartAction(s models.Store, action models.Action) error {
	return StartAction(c, s, action)
}

func (c *mongoCalendar) CompleteAction(s models.Store, action models.Action) error {
	return CompleteAction(c, s, action)
}

func (c *mongoCalendar) IntegratedSchedule(s models.Store, t time.Time) (models.Schedule, error) {
	return MergedScheduleForTime(s, c, t)
}
