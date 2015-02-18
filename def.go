/*
	The Models Package supplies the domain specific data interfaces elos relies on.

	It's subdirectories supply the implementation for the interfaces defined here.
*/
package models

import "github.com/elos/data"

type User interface {
	data.Model
	data.Nameable

	SetKey(string)
	Key() string

	AddEvent(Event) error
	DropEvent(Event) error

	AddTask(Task) error
	DropTask(Task) error

	Events(*data.Access) (data.RecordIterator, error)
	Tasks(*data.Access) (data.RecordIterator, error)

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

// Experimental

type Ontology interface {
	AddClass(Class) error
	DropClass(Class) error

	AddObject(Object) error
	DropObject(Object) error
}

type Class interface {
	AddTrait()
}

type Trait interface {
	Name()
	Type()
}

type Object interface {
	AddAttribute(string, string)
	AddRelationship(Relationship)
}

type Attribute interface {
	Trait()
	Value()
}

type Relationship interface {
	Trait()
	Tail()
}
