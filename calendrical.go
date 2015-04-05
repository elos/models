package models

import (
	"time"

	"github.com/elos/data"
)

// See: https://github.com/elos/documentation/blob/master/data/models/calendar.md
type Calendar interface {
	Actionable

	SetBaseSchedule(Schedule) error
	BaseSchedule(data.Access) (Schedule, error)

	SetWeekdaySchedule(Schedule, time.Weekday) error
	WeekdaySchedule(data.Access, time.Weekday) (Schedule, error)

	SetYeardaySchedule(Schedule, time.Time) error
	YeardaySchedule(data.Access, time.Time) (Schedule, error)

	SetCurrentFixture(Fixture) error
	CurrentFixture(data.Access) (Fixture, error)

	NextFixture(data.Access) (Fixture, error)
	IntegratedSchedule(data.Access, time.Time) (Schedule, error)
}

// See: https://github.com/elos/documentation/blob/master/data/models/schedule.md
type Schedule interface {
	data.Model
	data.Timeable

	IncludeFixture(Fixture) error
	ExcludeFixture(Fixture) error
	FixturesIter(data.Access) (data.ModelIterator, error)
	Fixtures(data.Access) ([]Fixture, error)

	FirstFixture(data.Access) (Fixture, error)
	FirstFixtureSince(data.Access, time.Time) (Fixture, error)
	OrderedFixtures(data.Access) ([]Fixture, error)
}

// See: https://github.com/elos/documentation/blob/master/data/models/fixture.md
type Fixture interface {
	data.Model
	data.Nameable
	data.Timeable
	Userable
	ActionableOps
	EventableOps
	Evented
	Actioned

	SetDescription(string)
	Description() string

	SetRank(int)
	Rank() int

	SetLabel(bool)
	Label() bool
	AllDay() bool

	SetExpires(time.Time)
	Expires() time.Time
	Expired() bool

	AddDateException(time.Time)
	DateExceptions() []time.Time
	ShouldOmitOnDate(t time.Time) bool

	SetSchedule(Schedule) error
	Schedule(data.Access) (Schedule, error)

	IncludeAction(Action) error
	ExcludeAction(Action) error
	ActionsIter(data.Access) (data.ModelIterator, error)
	Actions(data.Access) ([]Action, error)

	IncludeEvent(Event) error
	ExcludeEvent(Event) error
	EventsIter(data.Access) (data.ModelIterator, error)
	Events(data.Access) ([]Event, error)

	Conflicts(Fixture) bool
	Order(Fixture) (Fixture, Fixture)
	Before(Fixture) bool
}
