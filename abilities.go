package models

import (
	"github.com/elos/data"
)

type Userable interface {
	SetUser(User) error
	User(data.Access) (User, error)
	UserID() data.ID
}

type ActionableOps interface {
	NextAction(data.Access) (Action, error)
	StartAction(data.Access, Action) error
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
