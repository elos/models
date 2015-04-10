package models

import "github.com/elos/data"

/*
	Userable is the interface for a user-owned
	model.
*/
type Userable interface {
	SetUser(User) error
	User(Store) (User, error)
	UserID() data.ID
}

/*
	ActionableOps is the interface for the operations
	an Actionable should support.

	It should be noted that ActionableOps and EventableOps
	can be combined to make an Actionable and Eventable
	model interface, without having the repeated data.Model
	and Userable interfaces. An example is Fixture.
*/
type ActionableOps interface {
	NextAction(Store) (Action, error)
	StartAction(Store, Action) error
	CompleteAction(Store, Action) error
}

/*
	EventableOps is the interface for the operations
	an Eventable should support.

	It should be noted that EventableOps and ActionableOps
	can be combined to make an Eventable and Actionable
	model interface, without having the repeated data.Model
	and Userable interfaces. An example is Fixture.
*/
type EventableOps interface {
	NextEvent(Store) (Event, error)
}

/*
	Actionable is the interface for a model which can
	be treated like an actionable.

	The data.Model requirements are used to extract the
	data.Kind and data.ID from the model which fulfills
	this interface.
*/
type Actionable interface {
	data.Model
	Userable
	ActionableOps
}

/*
	Eventable is the interface for a model which can
	be treated like an eventable.

	The data.Model requirements are used to extract the
	data.Kind and data.ID from the model which fulfills
	this interface
*/
type Eventable interface {
	data.Model
	Userable
	EventableOps
}

type Actioned interface {
	SetActionable(Actionable)
	Actionable(Store) (Actionable, error)
	DropActionable()
	HasActionable() bool
}

type Evented interface {
	SetEventable(Eventable)
	Eventable(Store) (Eventable, error)
	DropEventable()
	HasEventable() bool
}
