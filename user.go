package models

import "github.com/elos/data"

// See: https://github.com/elos/documentation/blob/master/data/models/user.md
type User interface {
	data.Model
	data.Nameable

	SetKey(string)
	Key() string

	SetCurrentAction(Action) error
	CurrentAction(data.Access) (Action, error)

	SetCurrentActionable(Actionable) error
	CurrentActionable(data.Access) (Actionable, error)
	ClearCurrentActionable()

	SetCalendar(Calendar) error
	Calendar(data.Access) (Calendar, error)

	SetOntology(Ontology) error
	Ontology(data.Access) (Ontology, error)

	IncludeAction(Action) error
	ExcludeAction(Action) error
	ActionsIter(data.Access) (data.ModelIterator, error)
	Actions(data.Access) ([]Action, error)

	IncludeEvent(Event) error
	ExcludeEvent(Event) error
	EventsIter(data.Access) (data.ModelIterator, error)
	Events(data.Access) ([]Event, error)

	IncludeTask(Task) error
	ExcludeTask(Task) error
	TasksIter(data.Access) (data.ModelIterator, error)
	Tasks(data.Access) ([]Task, error)

	IncludeRoutine(Routine) error
	ExcludeRoutine(Routine) error
	RoutinesIter(data.Access) (data.ModelIterator, error)
	Routines(data.Access) ([]Routine, error)
}
