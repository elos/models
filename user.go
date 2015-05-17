package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type User struct {
	ActionsIDs            []string  `json:"actions_ids" bson:"actions_ids"`
	CalendarID            string    `json:"calendar_id" bson:"calendar_id"`
	CreatedAt             time.Time `json:"created_at" bson:"created_at"`
	CurrentActionID       string    `json:"current_action_id" bson:"current_action_id"`
	CurrentActionableID   string    `json:"current_actionable_id" bson:"current_actionable_id"`
	CurrentActionableKind string    `json:"current_actionable_kind" bson:"current_actionable_kind"`
	EventsIDs             []string  `json:"events_ids" bson:"events_ids"`
	Id                    string    `json:"id" bson:"_id,omitempty"`
	Key                   string    `json:"key" bson:"key"`
	Name                  string    `json:"name" bson:"name"`
	OntologyID            string    `json:"ontology_id" bson:"ontology_id"`
	PublicKeys            []string  `json:"public_keys" bson:"public_keys"`
	RoutinesIDs           []string  `json:"routines_ids" bson:"routines_ids"`
	TasksIDs              []string  `json:"tasks_ids" bson:"tasks_ids"`
	UpdatedAt             time.Time `json:"updated_at" bson:"updated_at"`
}

func NewUser() *User {
	return &User{}
}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (user *User) Kind() data.Kind {
	return UserKind
}

// just returns itself for now
func (user *User) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = user.ID()
	return foo
}

func (user *User) SetID(id data.ID) {
	user.Id = id.String()
}

func (user *User) ID() data.ID {
	return data.ID(user.Id)
}

func (user *User) IncludeAction(action *Action) {
	user.ActionsIDs = append(user.ActionsIDs, action.ID().String())
}

func (user *User) ExcludeAction(action *Action) {
	tmp := make([]string, 0)
	id := action.ID().String()
	for _, s := range user.ActionsIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	user.ActionsIDs = tmp
}

func (user *User) ActionsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(user.ActionsIDs), db), nil
}

func (user *User) Actions(db data.DB) ([]*Action, error) {

	actions := make([]*Action, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(user.ActionsIDs), db)
	action := NewAction()
	for iter.Next(action) {
		actions = append(actions, action)
		action = NewAction()
	}
	return actions, nil
}

func (user *User) SetCalendar(calendar *Calendar) error {
	user.CalendarID = calendar.ID().String()
	return nil
}

func (user *User) Calendar(db data.DB) (*Calendar, error) {
	if user.CalendarID == "" {
		return nil, ErrEmptyLink
	}

	calendar := NewCalendar()
	pid, _ := mongo.ParseObjectID(user.CalendarID)
	calendar.SetID(data.ID(pid.Hex()))
	return calendar, db.PopulateByID(calendar)

}

func (user *User) SetCurrentAction(action *Action) error {
	user.CurrentActionID = action.ID().String()
	return nil
}

func (user *User) CurrentAction(db data.DB) (*Action, error) {
	if user.CurrentActionID == "" {
		return nil, ErrEmptyLink
	}

	action := NewAction()
	pid, _ := mongo.ParseObjectID(user.CurrentActionID)
	action.SetID(data.ID(pid.Hex()))
	return action, db.PopulateByID(action)

}

func (user *User) SetCurrentActionable(actionable Actionable) error {
	user.CurrentActionableID = actionable.ID().String()
	return nil
}

func (user *User) CurrentActionable(db data.DB) (Actionable, error) {
	if user.CurrentActionableID == "" {
		return nil, ErrEmptyLink
	}

	m := ModelFor(data.Kind(user.CurrentActionableKind))
	actionable := m.(Actionable)

	pid, _ := mongo.ParseObjectID(user.CurrentActionableID)

	actionable.SetID(data.ID(pid.Hex()))
	return actionable, db.PopulateByID(actionable)

}

func (user *User) IncludeEvent(event *Event) {
	user.EventsIDs = append(user.EventsIDs, event.ID().String())
}

func (user *User) ExcludeEvent(event *Event) {
	tmp := make([]string, 0)
	id := event.ID().String()
	for _, s := range user.EventsIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	user.EventsIDs = tmp
}

func (user *User) EventsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(user.EventsIDs), db), nil
}

func (user *User) Events(db data.DB) ([]*Event, error) {

	events := make([]*Event, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(user.EventsIDs), db)
	event := NewEvent()
	for iter.Next(event) {
		events = append(events, event)
		event = NewEvent()
	}
	return events, nil
}

func (user *User) SetOntology(ontology *Ontology) error {
	user.OntologyID = ontology.ID().String()
	return nil
}

func (user *User) Ontology(db data.DB) (*Ontology, error) {
	if user.OntologyID == "" {
		return nil, ErrEmptyLink
	}

	ontology := NewOntology()
	pid, _ := mongo.ParseObjectID(user.OntologyID)
	ontology.SetID(data.ID(pid.Hex()))
	return ontology, db.PopulateByID(ontology)

}

func (user *User) IncludeRoutine(routine *Routine) {
	user.RoutinesIDs = append(user.RoutinesIDs, routine.ID().String())
}

func (user *User) ExcludeRoutine(routine *Routine) {
	tmp := make([]string, 0)
	id := routine.ID().String()
	for _, s := range user.RoutinesIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	user.RoutinesIDs = tmp
}

func (user *User) RoutinesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(user.RoutinesIDs), db), nil
}

func (user *User) Routines(db data.DB) ([]*Routine, error) {

	routines := make([]*Routine, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(user.RoutinesIDs), db)
	routine := NewRoutine()
	for iter.Next(routine) {
		routines = append(routines, routine)
		routine = NewRoutine()
	}
	return routines, nil
}

func (user *User) IncludeTask(task *Task) {
	user.TasksIDs = append(user.TasksIDs, task.ID().String())
}

func (user *User) ExcludeTask(task *Task) {
	tmp := make([]string, 0)
	id := task.ID().String()
	for _, s := range user.TasksIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	user.TasksIDs = tmp
}

func (user *User) TasksIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(user.TasksIDs), db), nil
}

func (user *User) Tasks(db data.DB) ([]*Task, error) {

	tasks := make([]*Task, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(user.TasksIDs), db)
	task := NewTask()
	for iter.Next(task) {
		tasks = append(tasks, task)
		task = NewTask()
	}
	return tasks, nil
}

// BSON {{{
func (user *User) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Key string `json:"key" bson:"key"`

		Name string `json:"name" bson:"name"`

		PublicKeys []string `json:"public_keys" bson:"public_keys"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ActionsIDs []string `json:"actions_ids" bson:"actions_ids"`

		CalendarID string `json:"calendar_id" bson:"calendar_id"`

		CurrentActionID string `json:"current_action_id" bson:"current_action_id"`

		CurrentActionableID string `json:"current_actionable_id" bson:"current_actionable_id"`

		CurrentActionableKind string `json:"current_actionable_kind" bson:"current_actionable_kind"`

		EventsIDs []string `json:"events_ids" bson:"events_ids"`

		OntologyID string `json:"ontology_id" bson:"ontology_id"`

		RoutinesIDs []string `json:"routines_ids" bson:"routines_ids"`

		TasksIDs []string `json:"tasks_ids" bson:"tasks_ids"`
	}{

		CreatedAt: user.CreatedAt,

		Key: user.Key,

		Name: user.Name,

		PublicKeys: user.PublicKeys,

		UpdatedAt: user.UpdatedAt,

		ActionsIDs: user.ActionsIDs,

		CalendarID: user.CalendarID,

		CurrentActionID: user.CurrentActionID,

		CurrentActionableID: user.CurrentActionableID,

		CurrentActionableKind: user.CurrentActionableKind,

		EventsIDs: user.EventsIDs,

		OntologyID: user.OntologyID,

		RoutinesIDs: user.RoutinesIDs,

		TasksIDs: user.TasksIDs,
	}, nil

}

func (user *User) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Key string `json:"key" bson:"key"`

		Name string `json:"name" bson:"name"`

		PublicKeys []string `json:"public_keys" bson:"public_keys"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ActionsIDs []string `json:"actions_ids" bson:"actions_ids"`

		CalendarID string `json:"calendar_id" bson:"calendar_id"`

		CurrentActionID string `json:"current_action_id" bson:"current_action_id"`

		CurrentActionableID string `json:"current_actionable_id" bson:"current_actionable_id"`

		CurrentActionableKind string `json:"current_actionable_kind" bson:"current_actionable_kind"`

		EventsIDs []string `json:"events_ids" bson:"events_ids"`

		OntologyID string `json:"ontology_id" bson:"ontology_id"`

		RoutinesIDs []string `json:"routines_ids" bson:"routines_ids"`

		TasksIDs []string `json:"tasks_ids" bson:"tasks_ids"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	user.CreatedAt = tmp.CreatedAt

	user.Id = tmp.Id.Hex()

	user.Key = tmp.Key

	user.Name = tmp.Name

	user.PublicKeys = tmp.PublicKeys

	user.UpdatedAt = tmp.UpdatedAt

	user.ActionsIDs = tmp.ActionsIDs

	user.CalendarID = tmp.CalendarID

	user.CurrentActionID = tmp.CurrentActionID

	user.CurrentActionableID = tmp.CurrentActionableID

	user.CurrentActionableKind = tmp.CurrentActionableKind

	user.EventsIDs = tmp.EventsIDs

	user.OntologyID = tmp.OntologyID

	user.RoutinesIDs = tmp.RoutinesIDs

	user.TasksIDs = tmp.TasksIDs

	return nil

}

// BSON }}}
