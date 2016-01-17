package models

import (
	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
)

func MongoDB(addr string) (data.DB, error) {
	db, err := mongo.New(&mongo.Opts{Addr: addr})

	if err != nil {
		return db, err
	}

	db.SetName("test")

	db.RegisterKind(ActionKind, "actions")

	db.RegisterKind(AttributeKind, "attributes")

	db.RegisterKind(CalendarKind, "calendars")

	db.RegisterKind(ContextKind, "contexts")

	db.RegisterKind(CredentialKind, "credentials")

	db.RegisterKind(DatumKind, "data")

	db.RegisterKind(EventKind, "events")

	db.RegisterKind(FixtureKind, "fixtures")

	db.RegisterKind(GroupKind, "groups")

	db.RegisterKind(LinkKind, "links")

	db.RegisterKind(LocationKind, "locations")

	db.RegisterKind(MediaKind, "medias")

	db.RegisterKind(ModelKind, "models")

	db.RegisterKind(NoteKind, "notes")

	db.RegisterKind(ObjectKind, "objects")

	db.RegisterKind(OntologyKind, "ontologies")

	db.RegisterKind(PersonKind, "persons")

	db.RegisterKind(ProfileKind, "profiles")

	db.RegisterKind(QuantityKind, "quantities")

	db.RegisterKind(RelationKind, "relations")

	db.RegisterKind(RoutineKind, "routines")

	db.RegisterKind(ScheduleKind, "schedules")

	db.RegisterKind(SessionKind, "sessions")

	db.RegisterKind(TagKind, "tags")

	db.RegisterKind(TaskKind, "tasks")

	db.RegisterKind(TraitKind, "traits")

	db.RegisterKind(UserKind, "users")

	return db, nil
}
