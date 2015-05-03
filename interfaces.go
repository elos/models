package models

import (
	"errors"

	"github.com/elos/data"
)

var ErrEmptyLink = errors.New("EMPTY LINK")

/*
	ActionableOps is the interface for the operations
	an Actionable should support.

	It should be noted that ActionableOps and EventableOps
	can be combined to make an Actionable and Eventable
	model interface, without having the repeated data.Model
	and Userable interfaces. An example is Fixture.
*/
type ActionableOps interface {
	NextAction(data.Store) (Action, error)
	StartAction(data.Store, Action) error
	CompleteAction(data.Store, Action) error
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
	NextEvent(data.Store) (Event, error)
}

type Actionable interface {
	data.Model
	ActionableOps
}

type Eventable interface {
	data.Model
	EventableOps
}
