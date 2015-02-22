/*
	The Models Package supplies the domain specific data interfaces elos relies on.

	It's subdirectories supply the implementation for the interfaces defined here.
*/
package models

import (
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

	Events(*data.Access) (data.RecordIterator, error)
	Tasks(*data.Access) (data.RecordIterator, error)
	Routines(*data.Access) (data.RecordIterator, error)

	SetCalendar(Calendar) error
	Calendar(*data.Access, Calendar) error

	SetCurrentAction(Action)
	CurrentAction(*data.Access, Action) error
	SetCurrentActionable(Actionable)
	CurrentActionable(*data.Access) (Actionable, error)
}

type Userable interface {
	SetUser(User) error
	User(*data.Access, User) error
	UserID() data.ID
	SetUserID(data.ID) error
}

type Actionable interface {
	data.Model
	Userable
	ActionCount() int
	NextAction(*data.Access) (Action, bool)
	CompleteAction(*data.Access, Action)
}

type Action interface {
	data.Model
	data.Nameable
	data.Timeable
	Userable

	SetTask(Task) error
	Task(*data.Access, Task) error
	Completed() bool
	Complete()
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
	Dependencies(*data.Access) (data.RecordIterator, error)
}

type Routine interface {
	Actionable
	data.Nameable
	data.Timeable

	IncludeTask(Task) error
	ExcludeTask(Task) error
	Tasks(*data.Access) (data.RecordIterator, error)
	TaskIDs() []data.ID

	CompleteTask(Task) error
	UncompleteTask(Task) error
	CompletedTasks(*data.Access) (data.RecordIterator, error)
	CompletedTaskIDs() []data.ID

	IncompleteTaskIDs() []data.ID

	ActionIDs() []data.ID
	AddAction(Action) error
	DropAction(Action) error

	SetCurrentAction(Action)
	CurrentAction(*data.Access, Action) error
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
	data.Timeable
	Userable

	SetSchedule(Schedule) error
	Schedule(data.Access, Schedule) error

	SetDescription(string)
	Description() string
}

type Schedule interface {
	data.Model
	data.Timeable
	Userable

	IncludeFixture(Fixture) error
	ExcludeFixture(Fixture) error
}

type Calendar interface {
	data.Model
	Userable

	Base(*data.Access, Schedule) error
	Monday(*data.Access, Schedule) error
	Tuesday(*data.Access, Schedule) error
	Wednesday(*data.Access, Schedule) error
	Thursday(*data.Access, Schedule) error
	Friday(*data.Access, Schedule) error
	Saturday(*data.Access, Schedule) error
	Sunday(*data.Access, Schedule) error

	SetBase(Schedule) error
	SetMonday(Schedule) error
	SetTuesday(Schedule) error
	SetWednesday(Schedule) error
	SetThursday(Schedule) error
	SetFriday(Schedule) error
	SetSaturday(Schedule) error
	SetSunday(Schedule) error

	IncludeSchedule(Schedule) error
	ExcludeSchedule(Schedule) error
	Schedules(*data.Access) (data.RecordIterator, error)
}

// Experimental

type Ontology interface {
	AddClass(Class) error
	DropClass(Class) error

	AddObject(Object) error
	DropObject(Object) error
}

type Class interface {
	data.Nameable

	IncludeTrait(Trait)
	ExcludeTrait(Trait)

	IncludeLink(l Link)
	ExcludeLink(l Link)
}

type Trait interface {
	data.Model
	data.Nameable

	Type() string
	SetType(string)
}

type Link interface {
	data.Model
	data.Nameable

	LinkKind() string
	OtherKind() string
	Inverse() string

	SetLinkKind(string)
	SetOtherKind(string)
	SetInverse(string)
}

type Object interface {
	data.Model

	SetClass(Class)
	Class() Class

	AddAttribute(string, string)
	AddRelationship(Relationship)
}

type Attribute interface {
	data.Model

	Trait() Trait
	Value() string
}

type Relationship interface {
	data.Model
}
