package models

import (
	"github.com/elos/data"
)

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

type MockActionable struct {
	*data.NM
	UserOwned
	nextAction Action
}

func NewMockActionable() *MockActionable {
	return &MockActionable{
		NM: data.NewNullModel(),
	}
}

func (a *MockActionable) SetNextAction(act Action) {
	a.nextAction = act
}

func (a *MockActionable) NextAction(data.Access) (Action, error) {
	if a.nextAction == nil {
		return nil, data.ErrNotFound
	} else {
		return a.nextAction, nil
	}
}

func (a *MockActionable) CompleteAction(access data.Access, act Action) error {
	if a.nextAction == act {
		return nil
	} else {
		return nil
	}
}
