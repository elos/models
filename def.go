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

	Events(data.Access) (data.ModelIterator, error)
	Tasks(data.Access) (data.ModelIterator, error)
	Routines(data.Access) (data.ModelIterator, error)

	SetCalendar(Calendar) error
	Calendar(data.Access, Calendar) error

	SetCurrentAction(Action)
	CurrentAction(data.Access, Action) error
	SetCurrentActionable(Actionable)
	CurrentActionable(data.Access) (Actionable, error)

	SetOntology(Ontology) error
	Ontology(data.Access) (Ontology, error)
}

type Userable interface {
	SetUser(User) error
	User(data.Access, User) error
	UserID() data.ID
	SetUserID(data.ID) error
}

type Actionable interface {
	data.Model
	Userable
	ActionCount() int
	NextAction(data.Access) (Action, bool)
	CompleteAction(data.Access, Action)
}

type Action interface {
	data.Model
	data.Nameable
	data.Timeable
	Userable

	SetTask(Task) error
	Task(data.Access, Task) error
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
	Dependencies(data.Access) (data.ModelIterator, error)
}

type Routine interface {
	Actionable
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

	Base(data.Access, Schedule) error
	Monday(data.Access, Schedule) error
	Tuesday(data.Access, Schedule) error
	Wednesday(data.Access, Schedule) error
	Thursday(data.Access, Schedule) error
	Friday(data.Access, Schedule) error
	Saturday(data.Access, Schedule) error
	Sunday(data.Access, Schedule) error

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
	Schedules(data.Access) (data.ModelIterator, error)
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
	Ontology(data.Access, Ontology) error

	IncludeTrait(Trait) error
	ExcludeTrait(Trait) error
	Traits() []*Trait

	IncludeRelationship(Relationship) error
	ExcludeRelationship(Relationship) error
	Relationships() []*Relationship

	IncludeObject(Object) error
	ExcludeObject(Object) error
	Objects(data.Access) (data.ModelIterator, error)

	Trait(string) (*Trait, bool)
	Relationship(string) (*Relationship, bool)
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
	Ontology(data.Access, Ontology) error

	SetClass(Class) error
	Class(data.Access) (Class, error)

	SetTrait(data.Access, string, string) error
	AddRelationship(data.Access, string, Object) error
	DropRelationship(data.Access, string, Object) error
}
