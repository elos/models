package models

import (
	"time"

	"github.com/elos/data"
)

type Action interface {
	data.Model
	data.Nameable
	data.Timeable

	Userable
	Actioned

	SetCompleted(bool)
	Completed() bool

	SetTask(Task) error
	Task(Store) (Task, error)

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
	DependenciesIter(Store) (data.ModelIterator, error)
	Dependencies(Store) ([]Task, error)
}

type Routine interface {
	Actionable
	ActionCount() int
	data.Nameable
	data.Timeable

	IncludeTask(Task) error
	ExcludeTask(Task) error
	TasksIter(Store) (data.ModelIterator, error)
	Tasks(Store) ([]Task, error)
	TaskIDs() []data.ID

	CompleteTask(Task) error
	UncompleteTask(Task) error
	CompletedTasksIter(Store) (data.ModelIterator, error)
	CompletedTasks(Store) ([]Task, error)
	CompletedTaskIDs() []data.ID

	IncompleteTaskIDs() []data.ID

	ActionIDs() []data.ID
	AddAction(Action) error
	DropAction(Action) error

	SetCurrentAction(Action)
	CurrentAction(Store) (Action, error)
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

type Ritual interface {
	data.Model
	data.Nameable

	CurrentStreak() (Streak, error)
	SetCurrentStreak(Streak) error

	IncludeStreak(Streak) error
	ExcludeStreak(Streak) error
}

type Streak interface {
	data.Model
	data.Timeable

	Ritual() (Ritual, error)
	SetRitual(Ritual) error

	Length() int
	IncludeDay(time.Time) error
	ExcludeDay(time.Time)
}
