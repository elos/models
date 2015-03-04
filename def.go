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

	IncludeEvent(Event) error
	ExcludeEvent(Event) error

	IncludeTask(Task) error
	ExcludeTask(Task) error

	IncludeRoutine(Routine) error
	ExcludeRoutine(Routine) error

	Events(data.Access) (data.ModelIterator, error)
	Tasks(data.Access) (data.ModelIterator, error)
	Routines(data.Access) (data.ModelIterator, error)

	SetCalendar(Calendar) error
	Calendar(data.Access) (Calendar, error)

	SetCurrentAction(Action)
	CurrentAction(data.Access) (Action, error)
	SetCurrentActionable(Actionable)
	CurrentActionable(data.Access) (Actionable, error)
	ClearCurrentActionable()

	SetOntology(Ontology) error
	Ontology(data.Access) (Ontology, error)
}

type Userable interface {
	SetUser(User) error
	User(data.Access, User) error
	UserID() data.ID
	SetUserID(data.ID) error
}

type ActionableOps interface {
	NextAction(data.Access) (Action, error)
	CompleteAction(data.Access, Action) error
}

type EventableOps interface {
	Event(data.Access) (Event, error)
}

type Actionable interface {
	data.Model
	Userable
	ActionableOps
}

type Eventable interface {
	data.Model
	Userable
	EventableOps
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

type Set interface {
	data.Model

	IncludeModel(data.Model) error
	ExcludeModel(data.Model) error
	ElementKind() data.Kind
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
	Userable
	ActionableOps
	EventableOps

	data.Nameable
	data.Timeable

	SetSchedule(Schedule) error
	Schedule(data.Access, Schedule) error

	SetDescription(string)
	Description() string

	SetExpires(time.Time)
	Expires() time.Time
	Expired() bool

	AddDateException(time.Time)
	DateExceptions() []time.Time
	ShouldOmitOnDate(t time.Time) bool

	IncludeAction(Action) error
	ExcludeAction(Action) error
	IncludeEvent(Event) error
	ExcludeEvent(Event) error

	Conflicts(Fixture) bool
	Rank(Fixture) (Fixture, Fixture)
	Before(Fixture) bool
}

type Schedule interface {
	data.Model
	data.Timeable

	IncludeFixture(Fixture) error
	ExcludeFixture(Fixture) error

	Fixtures(data.Access) (data.ModelIterator, error)

	FirstFixture(data.Access) (Fixture, error)
	FirstFixtureSince(data.Access, time.Time) (Fixture, error)
}

type Calendar interface {
	Actionable

	Base(data.Access) (Schedule, error)
	WeekdaySchedule(data.Access, time.Weekday) (Schedule, error)

	SetBase(Schedule) error
	SetWeekdaySchedule(Schedule, time.Weekday) error

	IncludeSchedule(Schedule) error
	ExcludeSchedule(Schedule) error
	Schedules(data.Access) (data.ModelIterator, error)
	ScheduleForDay(data.Access, time.Time) (Schedule, error)

	SetCurrentFixture(Fixture) error
	CurrentFixture(data.Access) (Fixture, error)

	NextFixture(data.Access) (Fixture, error)
}

// Experimental

type Ontology interface {
	data.Model
	Userable

	IncludeClass(Class) error
	ExcludeClass(Class) error

	IncludeObject(Object) error
	ExcludeObject(Object) error

	Classes(data.Access) (data.ModelIterator, error)
	Objects(data.Access) (data.ModelIterator, error)
}

type Class interface {
	data.Model
	data.Nameable
	Userable

	SetOntology(Ontology) error
	Ontology(data.Access) (Ontology, error)

	IncludeTrait(*Trait) error
	ExcludeTrait(*Trait) error
	Traits() []*Trait

	IncludeRelationship(*Relationship) error
	ExcludeRelationship(*Relationship) error
	Relationships() []*Relationship

	IncludeObject(Object) error
	ExcludeObject(Object) error
	Objects(data.Access) (data.ModelIterator, error)

	Trait(string) (*Trait, bool)
	Relationship(string) (*Relationship, bool)

	NewObject(a data.Access) Object
}

type Trait struct {
	Name string
	Type string
}

type Relationship struct {
	Name    string
	Other   string
	Inverse string
}

type Object interface {
	data.Model
	data.Nameable

	SetOntology(Ontology) error
	Ontology(data.Access) (Ontology, error)

	SetClass(Class) error
	Class(data.Access) (Class, error)

	SetTrait(data.Access, string, string) error
	AddRelationship(data.Access, string, Object) error
	DropRelationship(data.Access, string, Object) error
}
