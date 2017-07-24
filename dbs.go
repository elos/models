package models

import (
	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"

	xmodels "github.com/elos/x/models/proto"
)

func MongoDB(addr string) (data.DB, error) {
	db, err := mongo.New(
		&mongo.Opts{
			Addr: addr,
			Name: "test",
		},
	)

	if err != nil {
		return db, err
	}

	db.RegisterKind(ActionKind, "actions")

	db.RegisterKind(AttributeKind, "attributes")

	db.RegisterKind(CalendarKind, "calendars")

	db.RegisterKind(ContextKind, "contexts")

	db.RegisterKind(CredentialKind, "credentials")

	db.RegisterKind(DatumKind, "data")

	db.RegisterKind(EventKind, "events")

	db.RegisterKind(FixtureKind, "fixtures")

	db.RegisterKind(GroupKind, "groups")

	db.RegisterKind(HabitKind, "habits")

	db.RegisterKind(IntegrationKind, "integrations")

	db.RegisterKind(LinkKind, "links")

	db.RegisterKind(LocationKind, "locations")

	db.RegisterKind(MediaKind, "medias")

	db.RegisterKind(ModelKind, "models")

	db.RegisterKind(NoteKind, "notes")

	db.RegisterKind(OauthKind, "oauths")

	db.RegisterKind(ObjectKind, "objects")

	db.RegisterKind(OntologyKind, "ontologies")

	db.RegisterKind(PersonKind, "persons")

	db.RegisterKind(ProfileKind, "profiles")

	db.RegisterKind(QuantityKind, "quantities")

	db.RegisterKind(RecurrenceKind, "recurrences")

	db.RegisterKind(RelationKind, "relations")

	db.RegisterKind(RoutineKind, "routines")

	db.RegisterKind(ScheduleKind, "schedules")

	db.RegisterKind(SessionKind, "sessions")

	db.RegisterKind(TagKind, "tags")

	db.RegisterKind(TaskKind, "tasks")

	db.RegisterKind(TraitKind, "traits")

	db.RegisterKind(UserKind, "users")

	// manaully added (singular to not coincide)
	// INSERTKIND
	db.RegisterKind(data.Kind(xmodels.Kind_USER.String()), "user")
	db.RegisterKind(data.Kind(xmodels.Kind_CREDENTIAL.String()), "credential")
	db.RegisterKind(data.Kind(xmodels.Kind_SESSION.String()), "session")
	db.RegisterKind(data.Kind(xmodels.Kind_EVENT.String()), "event")
	db.RegisterKind(data.Kind(xmodels.Kind_GRANT.String()), "grant")
	db.RegisterKind(data.Kind(xmodels.Kind_SERVICE.String()), "service")
	db.RegisterKind(data.Kind(xmodels.Kind_INTEGRATION.String()), "integration")
	db.RegisterKind(data.Kind(xmodels.Kind_PROFILE.String()), "profile")
	db.RegisterKind(data.Kind(xmodels.Kind_PERSON.String()), "person")
	db.RegisterKind(data.Kind(xmodels.Kind_CONTACT.String()), "contact")
	db.RegisterKind(data.Kind(xmodels.Kind_FIXTURE.String()), "fixture")
	db.RegisterKind(data.Kind(xmodels.Kind_CALENDAR.String()), "calendar")
	db.RegisterKind(data.Kind(xmodels.Kind_FEATURE.String()), "feature")
	db.RegisterKind(data.Kind(xmodels.Kind_BOUNTY.String()), "bounty")
	db.RegisterKind(data.Kind(xmodels.Kind_REWARD.String()), "reward")
	db.RegisterKind(data.Kind(xmodels.Kind_ACTION.String()), "action")
	db.RegisterKind(data.Kind(xmodels.Kind_TASK.String()), "task")
	db.RegisterKind(data.Kind(xmodels.Kind_STRUCTURE.String()), "structure")
	db.RegisterKind(data.Kind(xmodels.Kind_TAG.String()), "tag")
	db.RegisterKind(data.Kind(xmodels.Kind_CLASS.String()), "class")

	return db, nil
}
