package persistence

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/action"
	"github.com/elos/models/calendar"
	"github.com/elos/models/class"
	"github.com/elos/models/event"
	"github.com/elos/models/fixture"
	"github.com/elos/models/object"
	"github.com/elos/models/ontology"
	"github.com/elos/models/routine"
	"github.com/elos/models/schedule"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
	"github.com/elos/mongo"
)

func MongoDB(addr string) (data.DB, error) {
	db := mongo.NewDB()
	if err := db.Connect(addr); err != nil {
		return nil, err
	}
	db.SetName("test")

	db.RegisterKind(models.UserKind, "users")
	db.RegisterKind(models.EventKind, "events")
	db.RegisterKind(models.TaskKind, "tasks")
	db.RegisterKind(models.RoutineKind, "routines")
	db.RegisterKind(models.ActionKind, "actions")
	db.RegisterKind(models.SetKind, "sets")
	db.RegisterKind(models.FixtureKind, "fixtures")
	db.RegisterKind(models.ScheduleKind, "schedules")
	db.RegisterKind(models.OntologyKind, "ontologies")
	db.RegisterKind(models.ClassKind, "classes")
	db.RegisterKind(models.ObjectKind, "objects")

	return db, nil
}

func MongoMemoryDB() data.DB {
	db := data.NewMemoryDBWithType(mongo.DBType)
	db.SetIDConstructor(func() data.ID {
		return mongo.NewObjectID()
	})
	return db
}

func Store(db data.DB) data.Store {
	s := data.NewStore(db, models.Schema)

	s.Register(models.UserKind, user.NewM)
	s.Register(models.EventKind, event.NewM)
	s.Register(models.TaskKind, task.NewM)
	s.Register(models.RoutineKind, routine.NewM)
	s.Register(models.ActionKind, action.NewM)
	s.Register(models.FixtureKind, fixture.NewM)
	s.Register(models.ScheduleKind, schedule.NewM)
	s.Register(models.OntologyKind, ontology.NewM)
	s.Register(models.ClassKind, class.NewM)
	s.Register(models.ObjectKind, object.NewM)
	s.Register(models.CalendarKind, calendar.NewM)

	return s
}
