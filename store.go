package models

import "github.com/elos/data"

type Models interface {
	Action() Action
	Calendar() Calendar
	Class() Class
	Event() Event
	Fixture() Fixture
	Object() Object
	Ontology() Ontology
	Routine() Routine
	Schedule() Schedule
	Task() Task
	User() User
}

type Store interface {
	data.Store
	Models
}

type Access interface {
	data.Access
	Models
}
