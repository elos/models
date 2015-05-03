package models

import (
	"github.com/elos/d"
	"github.com/elos/d/builtin/mongo"
)

func MongoDB(addr string) (d.DB, error) {
	db := mongo.NewDB()
	if err := db.Connect(addr); err != nil {
		return nil, err
	}
	db.SetName("test")

	db.RegisterKind(ActionKind, "actions")

	db.RegisterKind(AttributeKind, "attributes")

	db.RegisterKind(CalendarKind, "calendars")

	db.RegisterKind(ClassKind, "classes")

	db.RegisterKind(EventKind, "events")

	db.RegisterKind(FixtureKind, "fixtures")

	db.RegisterKind(LinkKind, "links")

	db.RegisterKind(ObjectKind, "objects")

	db.RegisterKind(OntologyKind, "ontologies")

	db.RegisterKind(RelationKind, "relations")

	db.RegisterKind(RoutineKind, "routines")

	db.RegisterKind(ScheduleKind, "schedules")

	db.RegisterKind(TaskKind, "tasks")

	db.RegisterKind(TraitKind, "traits")

	db.RegisterKind(UserKind, "users")

	return db, nil
}
