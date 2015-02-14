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

	Events(data.Access) (data.RecordIterator, error)
	Tasks(data.Access) (data.RecordIterator, error)

	SetCurrentAction(Action)
	CurrentAction(data.Access, Action) error

	SetCurrentActionable(Actionable)
	CurrentActionable(data.Access) (Actionable, error)
}

type Actionable interface {
	data.Model
	Action() Action
}

type Action interface {
	data.Model
	data.Nameable
	data.Timeable

	SetUser(User) error
	User(data.Access, User) error
}

type Event interface {
	data.Model
	data.Nameable
	data.Timeable

	SetUser(User) error
}

type Task interface {
	data.Model
	data.Nameable
	data.Timeable

	User(data.Access, User) error
	SetUser(User) error

	AddDependency(Task) error
	DropDependency(Task) error
	Dependencies(data.Access) (data.RecordIterator, error)
}

type Routine interface {
	Actionable
	data.Nameable
	data.Timeable

	User(data.Access, User) error
	SetUser(User) error

	IncludeTask(Task) error
	ExcludeTask(Task) error
	Tasks(data.Access) (data.RecordIterator, error)
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
