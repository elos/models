package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Person struct {
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

func NewPerson() *Person {
	return &Person{}
}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (person *Person) Kind() data.Kind {
	return PersonKind
}

// just returns itself for now
func (person *Person) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = person.ID()
	return foo
}

func (person *Person) SetID(id data.ID) {
	person.Id = id.String()
}

func (person *Person) ID() data.ID {
	return data.ID(person.Id)
}

func (person *Person) IncludeAction(action *Action) {
	person.ActionsIDs = append(person.ActionsIDs, action.ID().String())
}

func (person *Person) ExcludeAction(action *Action) {
	tmp := make([]string, 0)
	id := action.ID().String()
	for _, s := range person.ActionsIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	person.ActionsIDs = tmp
}

func (person *Person) ActionsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(person.ActionsIDs), db), nil
}

func (person *Person) Actions(db data.DB) ([]*Action, error) {

	actions := make([]*Action, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(person.ActionsIDs), db)
	action := NewAction()
	for iter.Next(action) {
		actions = append(actions, action)
		action = NewAction()
	}
	return actions, nil
}

func (person *Person) SetCalendar(calendar *Calendar) error {
	person.CalendarID = calendar.ID().String()
	return nil
}

func (person *Person) Calendar(db data.DB) (*Calendar, error) {
	if person.CalendarID == "" {
		return nil, ErrEmptyLink
	}

	calendar := NewCalendar()
	pid, _ := mongo.ParseObjectID(person.CalendarID)
	calendar.SetID(data.ID(pid.Hex()))
	return calendar, db.PopulateByID(calendar)

}

func (person *Person) CalendarOrCreate(db data.DB) (*Calendar, error) {
	calendar, err := person.Calendar(db)

	if err == ErrEmptyLink {
		calendar := NewCalendar()
		calendar.SetID(db.NewID())
		if err := person.SetCalendar(calendar); err != nil {
			return nil, err
		}

		if err := db.Save(calendar); err != nil {
			return nil, err
		}

		if err := db.Save(person); err != nil {
			return nil, err
		}

		return calendar, nil
	} else {
		return calendar, err
	}
}

func (person *Person) SetCurrentAction(action *Action) error {
	person.CurrentActionID = action.ID().String()
	return nil
}

func (person *Person) CurrentAction(db data.DB) (*Action, error) {
	if person.CurrentActionID == "" {
		return nil, ErrEmptyLink
	}

	action := NewAction()
	pid, _ := mongo.ParseObjectID(person.CurrentActionID)
	action.SetID(data.ID(pid.Hex()))
	return action, db.PopulateByID(action)

}

func (person *Person) CurrentActionOrCreate(db data.DB) (*Action, error) {
	action, err := person.CurrentAction(db)

	if err == ErrEmptyLink {
		action := NewAction()
		action.SetID(db.NewID())
		if err := person.SetCurrentAction(action); err != nil {
			return nil, err
		}

		if err := db.Save(action); err != nil {
			return nil, err
		}

		if err := db.Save(person); err != nil {
			return nil, err
		}

		return action, nil
	} else {
		return action, err
	}
}

func (person *Person) SetCurrentActionable(actionable Actionable) error {
	person.CurrentActionableID = actionable.ID().String()
	return nil
}

func (person *Person) CurrentActionable(db data.DB) (Actionable, error) {
	if person.CurrentActionableID == "" {
		return nil, ErrEmptyLink
	}

	m := ModelFor(data.Kind(person.CurrentActionableKind))
	actionable := m.(Actionable)

	pid, _ := mongo.ParseObjectID(person.CurrentActionableID)

	actionable.SetID(data.ID(pid.Hex()))
	return actionable, db.PopulateByID(actionable)

}

func (person *Person) IncludeEvent(event *Event) {
	person.EventsIDs = append(person.EventsIDs, event.ID().String())
}

func (person *Person) ExcludeEvent(event *Event) {
	tmp := make([]string, 0)
	id := event.ID().String()
	for _, s := range person.EventsIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	person.EventsIDs = tmp
}

func (person *Person) EventsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(person.EventsIDs), db), nil
}

func (person *Person) Events(db data.DB) ([]*Event, error) {

	events := make([]*Event, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(person.EventsIDs), db)
	event := NewEvent()
	for iter.Next(event) {
		events = append(events, event)
		event = NewEvent()
	}
	return events, nil
}

func (person *Person) SetOntology(ontology *Ontology) error {
	person.OntologyID = ontology.ID().String()
	return nil
}

func (person *Person) Ontology(db data.DB) (*Ontology, error) {
	if person.OntologyID == "" {
		return nil, ErrEmptyLink
	}

	ontology := NewOntology()
	pid, _ := mongo.ParseObjectID(person.OntologyID)
	ontology.SetID(data.ID(pid.Hex()))
	return ontology, db.PopulateByID(ontology)

}

func (person *Person) OntologyOrCreate(db data.DB) (*Ontology, error) {
	ontology, err := person.Ontology(db)

	if err == ErrEmptyLink {
		ontology := NewOntology()
		ontology.SetID(db.NewID())
		if err := person.SetOntology(ontology); err != nil {
			return nil, err
		}

		if err := db.Save(ontology); err != nil {
			return nil, err
		}

		if err := db.Save(person); err != nil {
			return nil, err
		}

		return ontology, nil
	} else {
		return ontology, err
	}
}

func (person *Person) IncludeRoutine(routine *Routine) {
	person.RoutinesIDs = append(person.RoutinesIDs, routine.ID().String())
}

func (person *Person) ExcludeRoutine(routine *Routine) {
	tmp := make([]string, 0)
	id := routine.ID().String()
	for _, s := range person.RoutinesIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	person.RoutinesIDs = tmp
}

func (person *Person) RoutinesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(person.RoutinesIDs), db), nil
}

func (person *Person) Routines(db data.DB) ([]*Routine, error) {

	routines := make([]*Routine, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(person.RoutinesIDs), db)
	routine := NewRoutine()
	for iter.Next(routine) {
		routines = append(routines, routine)
		routine = NewRoutine()
	}
	return routines, nil
}

func (person *Person) IncludeTask(task *Task) {
	person.TasksIDs = append(person.TasksIDs, task.ID().String())
}

func (person *Person) ExcludeTask(task *Task) {
	tmp := make([]string, 0)
	id := task.ID().String()
	for _, s := range person.TasksIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	person.TasksIDs = tmp
}

func (person *Person) TasksIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(person.TasksIDs), db), nil
}

func (person *Person) Tasks(db data.DB) ([]*Task, error) {

	tasks := make([]*Task, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(person.TasksIDs), db)
	task := NewTask()
	for iter.Next(task) {
		tasks = append(tasks, task)
		task = NewTask()
	}
	return tasks, nil
}

// BSON {{{
func (person *Person) GetBSON() (interface{}, error) {

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

		CreatedAt: person.CreatedAt,

		Key: person.Key,

		Name: person.Name,

		PublicKeys: person.PublicKeys,

		UpdatedAt: person.UpdatedAt,

		ActionsIDs: person.ActionsIDs,

		CalendarID: person.CalendarID,

		CurrentActionID: person.CurrentActionID,

		CurrentActionableID: person.CurrentActionableID,

		CurrentActionableKind: person.CurrentActionableKind,

		EventsIDs: person.EventsIDs,

		OntologyID: person.OntologyID,

		RoutinesIDs: person.RoutinesIDs,

		TasksIDs: person.TasksIDs,
	}, nil

}

func (person *Person) SetBSON(raw bson.Raw) error {

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

	person.CreatedAt = tmp.CreatedAt

	person.Id = tmp.Id.Hex()

	person.Key = tmp.Key

	person.Name = tmp.Name

	person.PublicKeys = tmp.PublicKeys

	person.UpdatedAt = tmp.UpdatedAt

	person.ActionsIDs = tmp.ActionsIDs

	person.CalendarID = tmp.CalendarID

	person.CurrentActionID = tmp.CurrentActionID

	person.CurrentActionableID = tmp.CurrentActionableID

	person.CurrentActionableKind = tmp.CurrentActionableKind

	person.EventsIDs = tmp.EventsIDs

	person.OntologyID = tmp.OntologyID

	person.RoutinesIDs = tmp.RoutinesIDs

	person.TasksIDs = tmp.TasksIDs

	return nil

}

// BSON }}}
