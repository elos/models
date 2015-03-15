/*
	The Models Package supplies the domain specific data interfaces elos relies on.

	It's subdirectories supply the implementation for the interfaces defined here.
*/
package models

import (
	"time"

	"github.com/elos/data"
)

type User interface {
	data.Model
	data.Nameable

	SetKey(string)
	Key() string

	SetCurrentAction(Action) error
	CurrentAction(data.Access) (Action, error)

	SetCurrentActionable(Actionable) error
	CurrentActionable(data.Access) (Actionable, error)
	ClearCurrentActionable()

	SetCalendar(Calendar) error
	Calendar(data.Access) (Calendar, error)

	SetOntology(Ontology) error
	Ontology(data.Access) (Ontology, error)

	IncludeEvent(Event) error
	ExcludeEvent(Event) error
	EventsIter(data.Access) (data.ModelIterator, error)
	Events(data.Access) ([]Event, error)

	IncludeTask(Task) error
	ExcludeTask(Task) error
	TasksIter(data.Access) (data.ModelIterator, error)
	Tasks(data.Access) ([]Task, error)

	IncludeRoutine(Routine) error
	ExcludeRoutine(Routine) error
	RoutinesIter(data.Access) (data.ModelIterator, error)
	Routines(data.Access) ([]Routine, error)
}

type Calendar interface {
	Actionable

	SetBaseSchedule(Schedule) error
	BaseSchedule(data.Access) (Schedule, error)

	SetWeekdaySchedule(Schedule, time.Weekday) error
	WeekdaySchedule(data.Access, time.Weekday) (Schedule, error)

	IncludeSchedule(Schedule) error
	ExcludeSchedule(Schedule) error
	Schedules(data.Access) (data.ModelIterator, error)
	ScheduleForDay(data.Access, time.Time) (Schedule, error)

	SetCurrentFixture(Fixture) error
	CurrentFixture(data.Access) (Fixture, error)

	NextFixture(data.Access) (Fixture, error)
}

type Action interface {
	data.Model
	data.Nameable
	data.Timeable
	Userable

	Completed() bool
	Complete()
	SetTask(Task) error
	Task(data.Access) (Task, error)

	SetActionable(Actionable)
	Actionable(data.Access) (Actionable, error)
}

type Event interface {
	data.Model
	data.Nameable
	data.Timeable
	Userable
}

type Task interface {
	data.Model
	data.Nameable
	data.Timeable
	Userable

	AddDependency(Task) error
	DropDependency(Task) error
	Dependencies(data.Access) (data.ModelIterator, error)
}

type Routine interface {
	Actionable
	ActionCount() int
	data.Nameable
	data.Timeable

	IncludeTask(Task) error
	ExcludeTask(Task) error
	Tasks(data.Access) (data.ModelIterator, error)
	TaskIDs() []data.ID

	CompleteTask(Task) error
	UncompleteTask(Task) error
	CompletedTasks(data.Access) (data.ModelIterator, error)
	CompletedTaskIDs() []data.ID

	IncompleteTaskIDs() []data.ID

	ActionIDs() []data.ID
	AddAction(Action) error
	DropAction(Action) error

	SetCurrentAction(Action)
	CurrentAction(data.Access, Action) error
}

type GeoPoint interface {
	data.Model
	Lat() float64
	Lon() float64

	Latitude() float64
	Longitude() float64
}

type Location interface {
	GeoPoint
	data.Nameable
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

type Schedule interface {
	data.Model
	data.Timeable

	IncludeFixture(Fixture) error
	ExcludeFixture(Fixture) error

	Fixtures(data.Access) (data.ModelIterator, error) // traditional model iterator for infinite (long) lists and memory
	OrderedFixtures(data.Access) ([]Fixture, error)   // if you are ordering them you are biting the memory bullet anyway

	FirstFixture(data.Access) (Fixture, error)
	FirstFixtureSince(data.Access, time.Time) (Fixture, error)
}

type Ritual interface {
	data.Model
	data.Nameable

	CurrentStreak() (Streak, error)
	SetCurrentStreak(Streak) error

	IncludeStreak(Streak) error
	ExcludeStreak(Streak) error
}

type Streak interface {
	data.Model
	data.Timeable

	Ritual() (Ritual, error)
	SetRitual(Ritual) error

	Length() int
	IncludeDay(time.Time) error
	ExcludeDay(time.Time)
}
