package models

import (
	"time"

	"github.com/elos/data"
)

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
}

type Schedule interface {
	data.Model
	data.Timeable

	IncludeFixture(Fixture) error
	ExcludeFixture(Fixture) error
	FixturesIter(data.Access) (data.ModelIterator, error)
	Fixtures(data.Access) ([]Fixture, error)
	OrderedFixtures(data.Access) ([]Fixture, error)

	FirstFixture(data.Access) (Fixture, error)
	FirstFixtureSince(data.Access, time.Time) (Fixture, error)
}

type Fixture interface {
	data.Model
	data.Nameable
	data.Timeable

	Userable
	ActionableOps
	EventableOps

	Evented
	Actioned

	IncludeAction(Action) error
	ExcludeAction(Action) error

	IncludeEvent(Event) error
	ExcludeEvent(Event) error

	SetDescription(string)
	Description() string

	SetExpires(time.Time)
	Expires() time.Time
	Expired() bool

	SetRank(int)
	Rank() int

	SetLabel(bool)
	Label() bool
	AllDay() bool

	AddDateException(time.Time)
	DateExceptions() []time.Time
	ShouldOmitOnDate(t time.Time) bool

	SetSchedule(Schedule) error
	Schedule(data.Access, Schedule) error

	Conflicts(Fixture) bool
	Order(Fixture) (Fixture, Fixture)
	Before(Fixture) bool
}
