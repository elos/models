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
	RemoveEvent(Event) error

	AddTask(Task) error
	RemoveTask(Task) error
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

	SetUser(User) error

	AddDependency(Task) error
	RemoveDependency(Task) error
	Dependencies(data.Store) data.RecordIterator
}

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
