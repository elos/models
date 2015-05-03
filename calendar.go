package models

// THIS FILE GENERATED BY METIS

import (
	"time"

	"github.com/elos/d"
	"github.com/elos/d/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// this type def generated by metis
type Calendar struct {
	BaseScheduleID   string            `json:"base_schedule_id" bson:"base_schedule_id"`
	CreatedAt        *time.Time        `json:"created_at" bson:"created_at"`
	CurrentFixtureID string            `json:"current_fixture_id" bson:"current_fixture_id"`
	Id               string            `json:"id" bson:"_id,omitempty"`
	Name             string            `json:"name" bson:"name"`
	UpdatedAt        *time.Time        `json:"updated_at" bson:"updated_at"`
	UserID           string            `json:"user_id" bson:"user_id"`
	WeekdaySchedules map[string]string `json:"weekday_schedules" bson:"weekday_schedules"`
	YeardaySchedules map[string]string `json:"yearday_schedules" bson:"yearday_schedules"`
}

func NewCalendar() *Calendar {
	return &Calendar{}
}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (calendar *Calendar) Kind() d.Kind {
	return CalendarKind
}

// just returns itself for now
func (calendar *Calendar) Concerned() []d.ID {
	foo := make([]d.ID, 1)
	foo[0] = calendar.ID()
	return foo
}

func (calendar *Calendar) SetID(id d.ID) {
	calendar.Id = id.String()
}

func (calendar *Calendar) ID() d.ID {
	return d.ID(calendar.Id)
}

func (calendar *Calendar) SetBaseSchedule(schedule *Schedule) error {
	calendar.BaseScheduleID = schedule.ID().String()
	return nil
}

func (calendar *Calendar) BaseSchedule(store d.Store) (*Schedule, error) {
	if calendar.BaseScheduleID == "" {
		return nil, ErrEmptyLink
	}

	schedule := NewSchedule()
	pid, _ := mongo.ParseObjectID(calendar.BaseScheduleID)
	schedule.SetID(d.ID(pid.Hex()))
	return schedule, store.PopulateByID(schedule)

}

func (calendar *Calendar) SetCurrentFixture(fixture *Fixture) error {
	calendar.CurrentFixtureID = fixture.ID().String()
	return nil
}

func (calendar *Calendar) CurrentFixture(store d.Store) (*Fixture, error) {
	if calendar.CurrentFixtureID == "" {
		return nil, ErrEmptyLink
	}

	fixture := NewFixture()
	pid, _ := mongo.ParseObjectID(calendar.CurrentFixtureID)
	fixture.SetID(d.ID(pid.Hex()))
	return fixture, store.PopulateByID(fixture)

}

func (calendar *Calendar) SetUser(user *User) error {
	calendar.UserID = user.ID().String()
	return nil
}

func (calendar *Calendar) User(store d.Store) (*User, error) {
	if calendar.UserID == "" {
		return nil, ErrEmptyLink
	}

	user := NewUser()
	pid, _ := mongo.ParseObjectID(calendar.UserID)
	user.SetID(d.ID(pid.Hex()))
	return user, store.PopulateByID(user)

}

func (calendar *Calendar) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt *time.Time `json:"created_at" bson:"created_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`

		WeekdaySchedules map[string]string `json:"weekday_schedules" bson:"weekday_schedules"`

		YeardaySchedules map[string]string `json:"yearday_schedules" bson:"yearday_schedules"`

		BaseScheduleID string `json:"base_schedule_id" bson:"base_schedule_id"`

		CurrentFixtureID string `json:"current_fixture_id" bson:"current_fixture_id"`

		UserID string `json:"user_id" bson:"user_id"`
	}{

		CreatedAt: calendar.CreatedAt,

		Name: calendar.Name,

		UpdatedAt: calendar.UpdatedAt,

		WeekdaySchedules: calendar.WeekdaySchedules,

		YeardaySchedules: calendar.YeardaySchedules,

		BaseScheduleID: calendar.BaseScheduleID,

		CurrentFixtureID: calendar.CurrentFixtureID,

		UserID: calendar.UserID,
	}, nil

}

func (calendar *Calendar) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt *time.Time `json:"created_at" bson:"created_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`

		WeekdaySchedules map[string]string `json:"weekday_schedules" bson:"weekday_schedules"`

		YeardaySchedules map[string]string `json:"yearday_schedules" bson:"yearday_schedules"`

		BaseScheduleID string `json:"base_schedule_id" bson:"base_schedule_id"`

		CurrentFixtureID string `json:"current_fixture_id" bson:"current_fixture_id"`

		UserID string `json:"user_id" bson:"user_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	calendar.CreatedAt = tmp.CreatedAt

	calendar.Id = tmp.Id.Hex()

	calendar.Name = tmp.Name

	calendar.UpdatedAt = tmp.UpdatedAt

	calendar.WeekdaySchedules = tmp.WeekdaySchedules

	calendar.YeardaySchedules = tmp.YeardaySchedules

	calendar.BaseScheduleID = tmp.BaseScheduleID

	calendar.CurrentFixtureID = tmp.CurrentFixtureID

	calendar.UserID = tmp.UserID

	return nil

}
