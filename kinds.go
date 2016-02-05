package models

import "github.com/elos/data"

const (
	ActionKind data.Kind = "action"

	AttributeKind data.Kind = "attribute"

	CalendarKind data.Kind = "calendar"

	ContextKind data.Kind = "context"

	CredentialKind data.Kind = "credential"

	DatumKind data.Kind = "datum"

	EventKind data.Kind = "event"

	FixtureKind data.Kind = "fixture"

	GroupKind data.Kind = "group"

	HabitKind data.Kind = "habit"

	LinkKind data.Kind = "link"

	LocationKind data.Kind = "location"

	MediaKind data.Kind = "media"

	ModelKind data.Kind = "model"

	NoteKind data.Kind = "note"

	ObjectKind data.Kind = "object"

	OntologyKind data.Kind = "ontology"

	PersonKind data.Kind = "person"

	ProfileKind data.Kind = "profile"

	QuantityKind data.Kind = "quantity"

	RecurrenceKind data.Kind = "recurrence"

	RelationKind data.Kind = "relation"

	RoutineKind data.Kind = "routine"

	ScheduleKind data.Kind = "schedule"

	SessionKind data.Kind = "session"

	TagKind data.Kind = "tag"

	TaskKind data.Kind = "task"

	TraitKind data.Kind = "trait"

	UserKind data.Kind = "user"
)

var Kinds = map[data.Kind]bool{

	ActionKind: true,

	AttributeKind: true,

	CalendarKind: true,

	ContextKind: true,

	CredentialKind: true,

	DatumKind: true,

	EventKind: true,

	FixtureKind: true,

	GroupKind: true,

	HabitKind: true,

	LinkKind: true,

	LocationKind: true,

	MediaKind: true,

	ModelKind: true,

	NoteKind: true,

	ObjectKind: true,

	OntologyKind: true,

	PersonKind: true,

	ProfileKind: true,

	QuantityKind: true,

	RecurrenceKind: true,

	RelationKind: true,

	RoutineKind: true,

	ScheduleKind: true,

	SessionKind: true,

	TagKind: true,

	TaskKind: true,

	TraitKind: true,

	UserKind: true,
}
