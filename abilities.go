package models

import (
	"github.com/elos/data"
)

/*
	Userable is the interface for a user-owned
	model.
*/
type Userable interface {
	SetUser(User) error
	User(data.Access) (User, error)
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
	NextAction(data.Access) (Action, error)
	StartAction(data.Access, Action) error
	CompleteAction(data.Access, Action) error
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
	NextEvent(data.Access) (Event, error)
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
	Actionable(data.Access) (Actionable, error)
	DropActionable()
	HasActionable() bool
}

type Evented interface {
	SetEventable(Eventable)
	Eventable(data.Access) (Eventable, error)
	DropEventable()
	HasEventable() bool
}
