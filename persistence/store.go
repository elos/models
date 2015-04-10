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
)

type modelsStore struct {
	data.Store
}

func ModelsStore(s data.Store) models.Store {
	return &modelsStore{
		Store: s,
	}
}

func (ms *modelsStore) Action() models.Action {
	return action.New(ms.Store)
}

func (ms *modelsStore) Calendar() models.Calendar {
	return calendar.New(ms.Store)
}

func (ms *modelsStore) Class() models.Class {
	return class.New(ms.Store)
}

func (ms *modelsStore) Event() models.Event {
	return event.New(ms.Store)
}

func (ms *modelsStore) Fixture() models.Fixture {
	return fixture.New(ms.Store)
}

func (ms *modelsStore) Object() models.Object {
	return object.New(ms.Store)
}

func (ms *modelsStore) Ontology() models.Ontology {
	return ontology.New(ms.Store)
}

func (ms *modelsStore) Routine() models.Routine {
	return routine.New(ms.Store)
}

func (ms *modelsStore) Schedule() models.Schedule {
	return schedule.New(ms.Store)
}

func (ms *modelsStore) Task() models.Task {
	return task.New(ms.Store)
}

func (ms *modelsStore) User() models.User {
	return user.New(ms.Store)
}
