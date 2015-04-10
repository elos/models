package models

import (
	"time"

	"github.com/elos/data"
)

// See: https://github.com/elos/documentation/blob/master/data/models/calendar.md
type Calendar interface {
	Actionable

	SetBaseSchedule(Schedule) error
	BaseSchedule(Store) (Schedule, error)

	SetWeekdaySchedule(Schedule, time.Weekday) error
	WeekdaySchedule(Store, time.Weekday) (Schedule, error)

	SetYeardaySchedule(Schedule, time.Time) error
	YeardaySchedule(Store, time.Time) (Schedule, error)

	SetCurrentFixture(Fixture) error
	CurrentFixture(Store) (Fixture, error)

	NextFixture(Store) (Fixture, error)
	IntegratedSchedule(Store, time.Time) (Schedule, error)
}

// See: https://github.com/elos/documentation/blob/master/data/models/schedule.md
type Schedule interface {
	data.Model
	data.Timeable

	IncludeFixture(Fixture) error
	ExcludeFixture(Fixture) error
	FixturesIter(Store) (data.ModelIterator, error)
	Fixtures(Store) ([]Fixture, error)

	FirstFixture(Store) (Fixture, error)
	FirstFixtureSince(Store, time.Time) (Fixture, error)
	OrderedFixtures(Store) ([]Fixture, error)
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
	Schedule(Store) (Schedule, error)

	IncludeAction(Action) error
	ExcludeAction(Action) error
	ActionsIter(Store) (data.ModelIterator, error)
	Actions(Store) ([]Action, error)

	IncludeEvent(Event) error
	ExcludeEvent(Event) error
	EventsIter(Store) (data.ModelIterator, error)
	Events(Store) ([]Event, error)

	Conflicts(Fixture) bool
	Order(Fixture) (Fixture, Fixture)
	Before(Fixture) bool
}
