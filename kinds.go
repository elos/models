package models

import (
	"github.com/elos/data"
	"github.com/elos/metis"
)

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

	IntegrationKind data.Kind = "integration"

	LinkKind data.Kind = "link"

	LocationKind data.Kind = "location"

	MediaKind data.Kind = "media"

	ModelKind data.Kind = "model"

	NoteKind data.Kind = "note"

	OauthKind data.Kind = "oauth"

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

	IntegrationKind: true,

	LinkKind: true,

	LocationKind: true,

	MediaKind: true,

	ModelKind: true,

	NoteKind: true,

	OauthKind: true,

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

var Metis = map[data.Kind]*metis.Model{

	ActionKind: &metis.Model{
		Kind:    "action",
		Space:   "actions",
		Domains: []string{"actions"},
		Traits: map[string]*metis.Trait{"created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "start_time": &metis.Trait{
			Name: "start_time",
			Type: metis.Primitive(4),
		}, "end_time": &metis.Trait{
			Name: "end_time",
			Type: metis.Primitive(4),
		}, "completed": &metis.Trait{
			Name: "completed",
			Type: metis.Primitive(0),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}},
		Relations: map[string]*metis.Relation{"actionable": &metis.Relation{
			Name:         "actionable",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "actionable",
			Inverse:      "",
		}, "task": &metis.Relation{
			Name:         "task",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "tasks",
			Inverse:      "",
		}, "owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "person": &metis.Relation{
			Name:         "person",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "persons",
			Inverse:      "",
		}},
	},

	AttributeKind: &metis.Model{
		Kind:    "attribute",
		Space:   "attributes",
		Domains: []string{"attributes"},
		Traits: map[string]*metis.Trait{"value": &metis.Trait{
			Name: "value",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "object": &metis.Relation{
			Name:         "object",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "objects",
			Inverse:      "attributes",
		}, "trait": &metis.Relation{
			Name:         "trait",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "traits",
			Inverse:      "attributes",
		}},
	},

	CalendarKind: &metis.Model{
		Kind:    "calendar",
		Space:   "calendars",
		Domains: []string{"calendars"},
		Traits: map[string]*metis.Trait{"created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "weekday_schedules": &metis.Trait{
			Name: "weekday_schedules",
			Type: metis.Primitive(11),
		}, "yearday_schedules": &metis.Trait{
			Name: "yearday_schedules",
			Type: metis.Primitive(11),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}},
		Relations: map[string]*metis.Relation{"monthly_schedule": &metis.Relation{
			Name:         "monthly_schedule",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "schedules",
			Inverse:      "",
		}, "schedule": &metis.Relation{
			Name:         "schedule",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "schedules",
			Inverse:      "",
		}, "daily_schedule": &metis.Relation{
			Name:         "daily_schedule",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "schedules",
			Inverse:      "",
		}, "base_schedule": &metis.Relation{
			Name:         "base_schedule",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "schedules",
			Inverse:      "",
		}, "weekly_schedule": &metis.Relation{
			Name:         "weekly_schedule",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "schedules",
			Inverse:      "",
		}, "yearly_schedule": &metis.Relation{
			Name:         "yearly_schedule",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "schedules",
			Inverse:      "",
		}, "manifest_fixture": &metis.Relation{
			Name:         "manifest_fixture",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "fixtures",
			Inverse:      "",
		}, "fixtures": &metis.Relation{
			Name:         "fixtures",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "fixture",
			Codomain:     "fixtures",
			Inverse:      "",
		}, "owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	ContextKind: &metis.Model{
		Kind:    "context",
		Space:   "contexts",
		Domains: []string{"contexts"},
		Traits: map[string]*metis.Trait{"deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "domain": &metis.Trait{
			Name: "domain",
			Type: metis.Primitive(3),
		}, "ids": &metis.Trait{
			Name: "ids",
			Type: metis.Primitive(10),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	CredentialKind: &metis.Model{
		Kind:    "credential",
		Space:   "credentials",
		Domains: []string{"credentials"},
		Traits: map[string]*metis.Trait{"deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "public": &metis.Trait{
			Name: "public",
			Type: metis.Primitive(3),
		}, "private": &metis.Trait{
			Name: "private",
			Type: metis.Primitive(3),
		}, "spec": &metis.Trait{
			Name: "spec",
			Type: metis.Primitive(3),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "credentials",
		}, "sessions": &metis.Relation{
			Name:         "sessions",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "session",
			Codomain:     "sessions",
			Inverse:      "credential",
		}},
	},

	DatumKind: &metis.Model{
		Kind:    "datum",
		Space:   "data",
		Domains: []string{"data"},
		Traits: map[string]*metis.Trait{"unit": &metis.Trait{
			Name: "unit",
			Type: metis.Primitive(3),
		}, "tags": &metis.Trait{
			Name: "tags",
			Type: metis.Primitive(7),
		}, "context": &metis.Trait{
			Name: "context",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "value": &metis.Trait{
			Name: "value",
			Type: metis.Primitive(2),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "person": &metis.Relation{
			Name:         "person",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "persons",
			Inverse:      "",
		}, "event": &metis.Relation{
			Name:         "event",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "events",
			Inverse:      "",
		}},
	},

	EventKind: &metis.Model{
		Kind:    "event",
		Space:   "events",
		Domains: []string{"events"},
		Traits: map[string]*metis.Trait{"deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "time": &metis.Trait{
			Name: "time",
			Type: metis.Primitive(4),
		}, "data": &metis.Trait{
			Name: "data",
			Type: metis.Primitive(13),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"media": &metis.Relation{
			Name:         "media",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "medias",
			Inverse:      "",
		}, "owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "prior": &metis.Relation{
			Name:         "prior",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "events",
			Inverse:      "",
		}, "quantity": &metis.Relation{
			Name:         "quantity",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "quantities",
			Inverse:      "",
		}, "note": &metis.Relation{
			Name:         "note",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "notes",
			Inverse:      "",
		}, "location": &metis.Relation{
			Name:         "location",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "locations",
			Inverse:      "",
		}, "tags": &metis.Relation{
			Name:         "tags",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "tag",
			Codomain:     "tags",
			Inverse:      "",
		}},
	},

	FixtureKind: &metis.Model{
		Kind:    "fixture",
		Space:   "fixtures",
		Domains: []string{"fixtures", "actionable", "eventable"},
		Traits: map[string]*metis.Trait{"id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "start_offset": &metis.Trait{
			Name: "start_offset",
			Type: metis.Primitive(1),
		}, "exceptions": &metis.Trait{
			Name: "exceptions",
			Type: metis.Primitive(8),
		}, "end_offset": &metis.Trait{
			Name: "end_offset",
			Type: metis.Primitive(1),
		}, "rank": &metis.Trait{
			Name: "rank",
			Type: metis.Primitive(1),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "start_time": &metis.Trait{
			Name: "start_time",
			Type: metis.Primitive(4),
		}, "end_time": &metis.Trait{
			Name: "end_time",
			Type: metis.Primitive(4),
		}, "label": &metis.Trait{
			Name: "label",
			Type: metis.Primitive(0),
		}, "expires_at": &metis.Trait{
			Name: "expires_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "actionable": &metis.Relation{
			Name:         "actionable",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "actionable",
			Inverse:      "",
		}, "eventable": &metis.Relation{
			Name:         "eventable",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "eventable",
			Inverse:      "",
		}, "actions": &metis.Relation{
			Name:         "actions",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "action",
			Codomain:     "actions",
			Inverse:      "",
		}, "events": &metis.Relation{
			Name:         "events",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "event",
			Codomain:     "events",
			Inverse:      "",
		}},
	},

	GroupKind: &metis.Model{
		Kind:    "group",
		Space:   "groups",
		Domains: []string{"groups"},
		Traits: map[string]*metis.Trait{"created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "access": &metis.Trait{
			Name: "access",
			Type: metis.Primitive(1),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "groups",
		}, "grantees": &metis.Relation{
			Name:         "grantees",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "grantee",
			Codomain:     "users",
			Inverse:      "authorizations",
		}, "contexts": &metis.Relation{
			Name:         "contexts",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "context",
			Codomain:     "contexts",
			Inverse:      "",
		}},
	},

	HabitKind: &metis.Model{
		Kind:    "habit",
		Space:   "habits",
		Domains: []string{"habits"},
		Traits: map[string]*metis.Trait{"id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "tag": &metis.Relation{
			Name:         "tag",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "tags",
			Inverse:      "",
		}, "checkins": &metis.Relation{
			Name:         "checkins",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "checkin",
			Codomain:     "events",
			Inverse:      "",
		}},
	},

	IntegrationKind: &metis.Model{
		Kind:    "integration",
		Space:   "integrations",
		Domains: []string{"integrations"},
		Traits: map[string]*metis.Trait{"domain": &metis.Trait{
			Name: "domain",
			Type: metis.Primitive(3),
		}, "vendor": &metis.Trait{
			Name: "vendor",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "integration_credential": &metis.Relation{
			Name:         "integration_credential",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "integration_credentials",
			Inverse:      "",
		}},
	},

	LinkKind: &metis.Model{
		Kind:    "link",
		Space:   "links",
		Domains: []string{"links"},
		Traits: map[string]*metis.Trait{"updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "ids": &metis.Trait{
			Name: "ids",
			Type: metis.Primitive(12),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "object": &metis.Relation{
			Name:         "object",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "objects",
			Inverse:      "links",
		}, "relation": &metis.Relation{
			Name:         "relation",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "relations",
			Inverse:      "links",
		}},
	},

	LocationKind: &metis.Model{
		Kind:    "location",
		Space:   "locations",
		Domains: []string{"locations"},
		Traits: map[string]*metis.Trait{"id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "latitude": &metis.Trait{
			Name: "latitude",
			Type: metis.Primitive(2),
		}, "longitude": &metis.Trait{
			Name: "longitude",
			Type: metis.Primitive(2),
		}, "altitude": &metis.Trait{
			Name: "altitude",
			Type: metis.Primitive(2),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	MediaKind: &metis.Model{
		Kind:    "media",
		Space:   "medias",
		Domains: []string{"medias"},
		Traits: map[string]*metis.Trait{"created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "content": &metis.Trait{
			Name: "content",
			Type: metis.Primitive(3),
		}, "codec": &metis.Trait{
			Name: "codec",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	ModelKind: &metis.Model{
		Kind:    "model",
		Space:   "models",
		Domains: []string{"models"},
		Traits: map[string]*metis.Trait{"name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "traits": &metis.Relation{
			Name:         "traits",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "trait",
			Codomain:     "traits",
			Inverse:      "model",
		}, "relations": &metis.Relation{
			Name:         "relations",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "relation",
			Codomain:     "relations",
			Inverse:      "model",
		}, "ontology": &metis.Relation{
			Name:         "ontology",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "ontologies",
			Inverse:      "models",
		}, "objects": &metis.Relation{
			Name:         "objects",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "object",
			Codomain:     "objects",
			Inverse:      "model",
		}},
	},

	NoteKind: &metis.Model{
		Kind:    "note",
		Space:   "notes",
		Domains: []string{"notes"},
		Traits: map[string]*metis.Trait{"id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "text": &metis.Trait{
			Name: "text",
			Type: metis.Primitive(3),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	OauthKind: &metis.Model{
		Kind:    "oauth",
		Space:   "oauths",
		Domains: []string{"oauths", "integration_credentials"},
		Traits: map[string]*metis.Trait{"id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "access_token": &metis.Trait{
			Name: "access_token",
			Type: metis.Primitive(3),
		}, "token_type": &metis.Trait{
			Name: "token_type",
			Type: metis.Primitive(3),
		}, "refresh_token": &metis.Trait{
			Name: "refresh_token",
			Type: metis.Primitive(3),
		}, "expiry": &metis.Trait{
			Name: "expiry",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	ObjectKind: &metis.Model{
		Kind:    "object",
		Space:   "objects",
		Domains: []string{"objects"},
		Traits: map[string]*metis.Trait{"id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"model": &metis.Relation{
			Name:         "model",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "models",
			Inverse:      "objects",
		}, "ontology": &metis.Relation{
			Name:         "ontology",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "ontologies",
			Inverse:      "objects",
		}, "owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "attributes": &metis.Relation{
			Name:         "attributes",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "attribute",
			Codomain:     "attributes",
			Inverse:      "object",
		}, "links": &metis.Relation{
			Name:         "links",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "link",
			Codomain:     "links",
			Inverse:      "object",
		}},
	},

	OntologyKind: &metis.Model{
		Kind:    "ontology",
		Space:   "ontologies",
		Domains: []string{"ontologies"},
		Traits: map[string]*metis.Trait{"created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "models": &metis.Relation{
			Name:         "models",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "model",
			Codomain:     "models",
			Inverse:      "ontology",
		}, "objects": &metis.Relation{
			Name:         "objects",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "object",
			Codomain:     "objects",
			Inverse:      "ontology",
		}},
	},

	PersonKind: &metis.Model{
		Kind:    "person",
		Space:   "persons",
		Domains: []string{"persons"},
		Traits: map[string]*metis.Trait{"name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "first_name": &metis.Trait{
			Name: "first_name",
			Type: metis.Primitive(3),
		}, "last_name": &metis.Trait{
			Name: "last_name",
			Type: metis.Primitive(3),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "notes": &metis.Relation{
			Name:         "notes",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "note",
			Codomain:     "notes",
			Inverse:      "",
		}},
	},

	ProfileKind: &metis.Model{
		Kind:    "profile",
		Space:   "profiles",
		Domains: []string{"profiles"},
		Traits: map[string]*metis.Trait{"updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "phone": &metis.Trait{
			Name: "phone",
			Type: metis.Primitive(3),
		}, "email": &metis.Trait{
			Name: "email",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"ontology": &metis.Relation{
			Name:         "ontology",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "ontologies",
			Inverse:      "",
		}, "current_actionable": &metis.Relation{
			Name:         "current_actionable",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "actionable",
			Inverse:      "",
		}, "calendar": &metis.Relation{
			Name:         "calendar",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "calendars",
			Inverse:      "",
		}, "owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "data": &metis.Relation{
			Name:         "data",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "datum",
			Codomain:     "data",
			Inverse:      "",
		}, "actions": &metis.Relation{
			Name:         "actions",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "action",
			Codomain:     "actions",
			Inverse:      "",
		}, "tasks": &metis.Relation{
			Name:         "tasks",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "task",
			Codomain:     "tasks",
			Inverse:      "",
		}, "routines": &metis.Relation{
			Name:         "routines",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "routine",
			Codomain:     "routines",
			Inverse:      "",
		}, "location": &metis.Relation{
			Name:         "location",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "locations",
			Inverse:      "",
		}, "events": &metis.Relation{
			Name:         "events",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "event",
			Codomain:     "events",
			Inverse:      "",
		}, "current_action": &metis.Relation{
			Name:         "current_action",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "actions",
			Inverse:      "",
		}},
	},

	QuantityKind: &metis.Model{
		Kind:    "quantity",
		Space:   "quantities",
		Domains: []string{"quantities"},
		Traits: map[string]*metis.Trait{"id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "value": &metis.Trait{
			Name: "value",
			Type: metis.Primitive(2),
		}, "unit": &metis.Trait{
			Name: "unit",
			Type: metis.Primitive(3),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	RecurrenceKind: &metis.Model{
		Kind:    "recurrence",
		Space:   "recurrences",
		Domains: []string{"recurrences"},
		Traits: map[string]*metis.Trait{"inclusions": &metis.Trait{
			Name: "inclusions",
			Type: metis.Primitive(8),
		}, "limit": &metis.Trait{
			Name: "limit",
			Type: metis.Primitive(4),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "frequency": &metis.Trait{
			Name: "frequency",
			Type: metis.Primitive(3),
		}, "interval": &metis.Trait{
			Name: "interval",
			Type: metis.Primitive(1),
		}, "by_month_day": &metis.Trait{
			Name: "by_month_day",
			Type: metis.Primitive(6),
		}, "by_week_num": &metis.Trait{
			Name: "by_week_num",
			Type: metis.Primitive(6),
		}, "by_month_num": &metis.Trait{
			Name: "by_month_num",
			Type: metis.Primitive(6),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "exclusions": &metis.Trait{
			Name: "exclusions",
			Type: metis.Primitive(8),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "until": &metis.Trait{
			Name: "until",
			Type: metis.Primitive(4),
		}, "count": &metis.Trait{
			Name: "count",
			Type: metis.Primitive(1),
		}, "by_hour": &metis.Trait{
			Name: "by_hour",
			Type: metis.Primitive(6),
		}, "start": &metis.Trait{
			Name: "start",
			Type: metis.Primitive(4),
		}, "by_second": &metis.Trait{
			Name: "by_second",
			Type: metis.Primitive(6),
		}, "by_day": &metis.Trait{
			Name: "by_day",
			Type: metis.Primitive(6),
		}, "by_year_day": &metis.Trait{
			Name: "by_year_day",
			Type: metis.Primitive(6),
		}, "by_set_pos": &metis.Trait{
			Name: "by_set_pos",
			Type: metis.Primitive(6),
		}, "week_start": &metis.Trait{
			Name: "week_start",
			Type: metis.Primitive(1),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	RelationKind: &metis.Model{
		Kind:    "relation",
		Space:   "relations",
		Domains: []string{"relations"},
		Traits: map[string]*metis.Trait{"inverse": &metis.Trait{
			Name: "inverse",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "multiplicity": &metis.Trait{
			Name: "multiplicity",
			Type: metis.Primitive(3),
		}, "codomain": &metis.Trait{
			Name: "codomain",
			Type: metis.Primitive(3),
		}},
		Relations: map[string]*metis.Relation{"model": &metis.Relation{
			Name:         "model",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "models",
			Inverse:      "relations",
		}, "links": &metis.Relation{
			Name:         "links",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "link",
			Codomain:     "links",
			Inverse:      "relation",
		}, "owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	RoutineKind: &metis.Model{
		Kind:    "routine",
		Space:   "routines",
		Domains: []string{"routines", "actionable"},
		Traits: map[string]*metis.Trait{"updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "start_time": &metis.Trait{
			Name: "start_time",
			Type: metis.Primitive(4),
		}, "end_time": &metis.Trait{
			Name: "end_time",
			Type: metis.Primitive(4),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"current_action": &metis.Relation{
			Name:         "current_action",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "actions",
			Inverse:      "",
		}, "owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "tasks": &metis.Relation{
			Name:         "tasks",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "task",
			Codomain:     "tasks",
			Inverse:      "",
		}, "completed_tasks": &metis.Relation{
			Name:         "completed_tasks",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "completed_task",
			Codomain:     "tasks",
			Inverse:      "",
		}, "actions": &metis.Relation{
			Name:         "actions",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "action",
			Codomain:     "actions",
			Inverse:      "",
		}},
	},

	ScheduleKind: &metis.Model{
		Kind:    "schedule",
		Space:   "schedules",
		Domains: []string{"schedules"},
		Traits: map[string]*metis.Trait{"id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "start_time": &metis.Trait{
			Name: "start_time",
			Type: metis.Primitive(4),
		}, "end_time": &metis.Trait{
			Name: "end_time",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"fixtures": &metis.Relation{
			Name:         "fixtures",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "fixture",
			Codomain:     "fixtures",
			Inverse:      "",
		}, "owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	SessionKind: &metis.Model{
		Kind:    "session",
		Space:   "sessions",
		Domains: []string{"sessions"},
		Traits: map[string]*metis.Trait{"expires_after": &metis.Trait{
			Name: "expires_after",
			Type: metis.Primitive(1),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "token": &metis.Trait{
			Name: "token",
			Type: metis.Primitive(3),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "sessions",
		}, "credential": &metis.Relation{
			Name:         "credential",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "credentials",
			Inverse:      "sessions",
		}},
	},

	TagKind: &metis.Model{
		Kind:    "tag",
		Space:   "tags",
		Domains: []string{"tags"},
		Traits: map[string]*metis.Trait{"name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}},
	},

	TaskKind: &metis.Model{
		Kind:    "task",
		Space:   "tasks",
		Domains: []string{"tasks"},
		Traits: map[string]*metis.Trait{"created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "deadline": &metis.Trait{
			Name: "deadline",
			Type: metis.Primitive(4),
		}, "stages": &metis.Trait{
			Name: "stages",
			Type: metis.Primitive(8),
		}, "completed_at": &metis.Trait{
			Name: "completed_at",
			Type: metis.Primitive(4),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}},
		Relations: map[string]*metis.Relation{"tags": &metis.Relation{
			Name:         "tags",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "tag",
			Codomain:     "tags",
			Inverse:      "",
		}, "owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "prerequisites": &metis.Relation{
			Name:         "prerequisites",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "prerequisite",
			Codomain:     "tasks",
			Inverse:      "",
		}},
	},

	TraitKind: &metis.Model{
		Kind:    "trait",
		Space:   "traits",
		Domains: []string{"traits"},
		Traits: map[string]*metis.Trait{"name": &metis.Trait{
			Name: "name",
			Type: metis.Primitive(3),
		}, "primitive": &metis.Trait{
			Name: "primitive",
			Type: metis.Primitive(3),
		}, "id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}},
		Relations: map[string]*metis.Relation{"owner": &metis.Relation{
			Name:         "owner",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "users",
			Inverse:      "",
		}, "model": &metis.Relation{
			Name:         "model",
			Multiplicity: metis.Multiplicity(1),
			Singular:     "",
			Codomain:     "models",
			Inverse:      "traits",
		}, "attributes": &metis.Relation{
			Name:         "attributes",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "attribute",
			Codomain:     "attributes",
			Inverse:      "trait",
		}},
	},

	UserKind: &metis.Model{
		Kind:    "user",
		Space:   "users",
		Domains: []string{"users"},
		Traits: map[string]*metis.Trait{"id": &metis.Trait{
			Name: "id",
			Type: metis.Primitive(9),
		}, "created_at": &metis.Trait{
			Name: "created_at",
			Type: metis.Primitive(4),
		}, "updated_at": &metis.Trait{
			Name: "updated_at",
			Type: metis.Primitive(4),
		}, "deleted_at": &metis.Trait{
			Name: "deleted_at",
			Type: metis.Primitive(4),
		}, "password": &metis.Trait{
			Name: "password",
			Type: metis.Primitive(3),
		}},
		Relations: map[string]*metis.Relation{"credentials": &metis.Relation{
			Name:         "credentials",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "credential",
			Codomain:     "credentials",
			Inverse:      "owner",
		}, "groups": &metis.Relation{
			Name:         "groups",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "group",
			Codomain:     "groups",
			Inverse:      "owner",
		}, "authorizations": &metis.Relation{
			Name:         "authorizations",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "authorization",
			Codomain:     "groups",
			Inverse:      "grantees",
		}, "sessions": &metis.Relation{
			Name:         "sessions",
			Multiplicity: metis.Multiplicity(0),
			Singular:     "session",
			Codomain:     "sessions",
			Inverse:      "owner",
		}},
	},
}
