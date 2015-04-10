package models

import "github.com/elos/data"

// See: https://github.com/elos/documentation/blob/master/data/models/user.md
type User interface {
	data.Model
	data.Nameable

	SetKey(string)
	Key() string

	SetCurrentAction(Action) error
	CurrentAction(Store) (Action, error)

	SetCurrentActionable(Actionable) error
	CurrentActionable(Store) (Actionable, error)
	ClearCurrentActionable()

	SetCalendar(Calendar) error
	Calendar(Store) (Calendar, error)

	SetOntology(Ontology) error
	Ontology(Store) (Ontology, error)

	IncludeAction(Action) error
	ExcludeAction(Action) error
	ActionsIter(Store) (data.ModelIterator, error)
	Actions(Store) ([]Action, error)

	IncludeEvent(Event) error
	ExcludeEvent(Event) error
	EventsIter(Store) (data.ModelIterator, error)
	Events(Store) ([]Event, error)

	IncludeTask(Task) error
	ExcludeTask(Task) error
	TasksIter(Store) (data.ModelIterator, error)
	Tasks(Store) ([]Task, error)

	IncludeRoutine(Routine) error
	ExcludeRoutine(Routine) error
	RoutinesIter(Store) (data.ModelIterator, error)
	Routines(Store) ([]Routine, error)
}
