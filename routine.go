package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Routine struct {
	ActionsIds        []string  `json:"actions_ids" bson:"actions_ids"`
	CompletedTasksIds []string  `json:"completed_tasks_ids" bson:"completed_tasks_ids"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
	CurrentActionId   string    `json:"current_action_id" bson:"current_action_id"`
	EndTime           time.Time `json:"end_time" bson:"end_time"`
	Id                string    `json:"id" bson:"_id,omitempty"`
	Name              string    `json:"name" bson:"name"`
	OwnerId           string    `json:"owner_id" bson:"owner_id"`
	PersonId          string    `json:"person_id" bson:"person_id"`
	StartTime         time.Time `json:"start_time" bson:"start_time"`
	TasksIds          []string  `json:"tasks_ids" bson:"tasks_ids"`
	UpdatedAt         time.Time `json:"updated_at" bson:"updated_at"`
}

func NewRoutine() *Routine {
	return &Routine{}
}

func FindRoutine(db data.DB, id data.ID) (*Routine, error) {

	routine := NewRoutine()
	routine.SetID(id)

	return routine, db.PopulateByID(routine)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (routine *Routine) Kind() data.Kind {
	return RoutineKind
}

// just returns itself for now
func (routine *Routine) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = routine.ID()
	return foo
}

func (routine *Routine) SetID(id data.ID) {
	routine.Id = id.String()
}

func (routine *Routine) ID() data.ID {
	return data.ID(routine.Id)
}

func (routine *Routine) IncludeAction(action *Action) {
	routine.ActionsIds = append(routine.ActionsIds, action.ID().String())
}

func (routine *Routine) ExcludeAction(action *Action) {
	tmp := make([]string, 0)
	id := action.ID().String()
	for _, s := range routine.ActionsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	routine.ActionsIds = tmp
}

func (routine *Routine) ActionsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(routine.ActionsIds), db), nil
}

func (routine *Routine) Actions(db data.DB) ([]*Action, error) {

	actions := make([]*Action, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(routine.ActionsIds), db)
	action := NewAction()
	for iter.Next(action) {
		actions = append(actions, action)
		action = NewAction()
	}
	return actions, nil
}

func (routine *Routine) IncludeCompletedTask(completedTask *Task) {
	routine.CompletedTasksIds = append(routine.CompletedTasksIds, completedTask.ID().String())
}

func (routine *Routine) ExcludeCompletedTask(completedTask *Task) {
	tmp := make([]string, 0)
	id := completedTask.ID().String()
	for _, s := range routine.CompletedTasksIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	routine.CompletedTasksIds = tmp
}

func (routine *Routine) CompletedTasksIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(routine.CompletedTasksIds), db), nil
}

func (routine *Routine) CompletedTasks(db data.DB) ([]*Task, error) {

	completed_tasks := make([]*Task, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(routine.CompletedTasksIds), db)
	completed_task := NewTask()
	for iter.Next(completed_task) {
		completed_tasks = append(completed_tasks, completed_task)
		completed_task = NewTask()
	}
	return completed_tasks, nil
}

func (routine *Routine) SetCurrentAction(actionArgument *Action) error {
	routine.CurrentActionId = actionArgument.ID().String()
	return nil
}

func (routine *Routine) CurrentAction(db data.DB) (*Action, error) {
	if routine.CurrentActionId == "" {
		return nil, ErrEmptyLink
	}

	actionArgument := NewAction()
	pid, _ := mongo.ParseObjectID(routine.CurrentActionId)
	actionArgument.SetID(data.ID(pid.Hex()))
	return actionArgument, db.PopulateByID(actionArgument)

}

func (routine *Routine) CurrentActionOrCreate(db data.DB) (*Action, error) {
	action, err := routine.CurrentAction(db)

	if err == ErrEmptyLink {
		action := NewAction()
		action.SetID(db.NewID())
		if err := routine.SetCurrentAction(action); err != nil {
			return nil, err
		}

		if err := db.Save(action); err != nil {
			return nil, err
		}

		if err := db.Save(routine); err != nil {
			return nil, err
		}

		return action, nil
	} else {
		return action, err
	}
}

func (routine *Routine) SetOwner(userArgument *User) error {
	routine.OwnerId = userArgument.ID().String()
	return nil
}

func (routine *Routine) Owner(db data.DB) (*User, error) {
	if routine.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	pid, _ := mongo.ParseObjectID(routine.OwnerId)
	userArgument.SetID(data.ID(pid.Hex()))
	return userArgument, db.PopulateByID(userArgument)

}

func (routine *Routine) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := routine.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := routine.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(routine); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

func (routine *Routine) SetPerson(personArgument *Person) error {
	routine.PersonId = personArgument.ID().String()
	return nil
}

func (routine *Routine) Person(db data.DB) (*Person, error) {
	if routine.PersonId == "" {
		return nil, ErrEmptyLink
	}

	personArgument := NewPerson()
	pid, _ := mongo.ParseObjectID(routine.PersonId)
	personArgument.SetID(data.ID(pid.Hex()))
	return personArgument, db.PopulateByID(personArgument)

}

func (routine *Routine) PersonOrCreate(db data.DB) (*Person, error) {
	person, err := routine.Person(db)

	if err == ErrEmptyLink {
		person := NewPerson()
		person.SetID(db.NewID())
		if err := routine.SetPerson(person); err != nil {
			return nil, err
		}

		if err := db.Save(person); err != nil {
			return nil, err
		}

		if err := db.Save(routine); err != nil {
			return nil, err
		}

		return person, nil
	} else {
		return person, err
	}
}

func (routine *Routine) IncludeTask(task *Task) {
	routine.TasksIds = append(routine.TasksIds, task.ID().String())
}

func (routine *Routine) ExcludeTask(task *Task) {
	tmp := make([]string, 0)
	id := task.ID().String()
	for _, s := range routine.TasksIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	routine.TasksIds = tmp
}

func (routine *Routine) TasksIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(routine.TasksIds), db), nil
}

func (routine *Routine) Tasks(db data.DB) ([]*Task, error) {

	tasks := make([]*Task, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(routine.TasksIds), db)
	task := NewTask()
	for iter.Next(task) {
		tasks = append(tasks, task)
		task = NewTask()
	}
	return tasks, nil
}

// BSON {{{
func (routine *Routine) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ActionsIds []string `json:"actions_ids" bson:"actions_ids"`

		CompletedTasksIds []string `json:"completed_tasks_ids" bson:"completed_tasks_ids"`

		CurrentActionId string `json:"current_action_id" bson:"current_action_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PersonId string `json:"person_id" bson:"person_id"`

		TasksIds []string `json:"tasks_ids" bson:"tasks_ids"`
	}{

		CreatedAt: routine.CreatedAt,

		EndTime: routine.EndTime,

		Name: routine.Name,

		StartTime: routine.StartTime,

		UpdatedAt: routine.UpdatedAt,

		ActionsIds: routine.ActionsIds,

		CompletedTasksIds: routine.CompletedTasksIds,

		CurrentActionId: routine.CurrentActionId,

		OwnerId: routine.OwnerId,

		PersonId: routine.PersonId,

		TasksIds: routine.TasksIds,
	}, nil

}

func (routine *Routine) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ActionsIds []string `json:"actions_ids" bson:"actions_ids"`

		CompletedTasksIds []string `json:"completed_tasks_ids" bson:"completed_tasks_ids"`

		CurrentActionId string `json:"current_action_id" bson:"current_action_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PersonId string `json:"person_id" bson:"person_id"`

		TasksIds []string `json:"tasks_ids" bson:"tasks_ids"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	routine.CreatedAt = tmp.CreatedAt

	routine.EndTime = tmp.EndTime

	routine.Id = tmp.Id.Hex()

	routine.Name = tmp.Name

	routine.StartTime = tmp.StartTime

	routine.UpdatedAt = tmp.UpdatedAt

	routine.ActionsIds = tmp.ActionsIds

	routine.CompletedTasksIds = tmp.CompletedTasksIds

	routine.CurrentActionId = tmp.CurrentActionId

	routine.OwnerId = tmp.OwnerId

	routine.PersonId = tmp.PersonId

	routine.TasksIds = tmp.TasksIds

	return nil

}

// BSON }}}

func (routine *Routine) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		routine.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		routine.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		routine.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["name"]; ok {
		routine.Name = val.(string)
	}

	if val, ok := structure["start_time"]; ok {
		routine.StartTime = val.(time.Time)
	}

	if val, ok := structure["end_time"]; ok {
		routine.EndTime = val.(time.Time)
	}

	if val, ok := structure["completed_tasks_ids"]; ok {
		routine.CompletedTasksIds = val.([]string)
	}

	if val, ok := structure["actions_ids"]; ok {
		routine.ActionsIds = val.([]string)
	}

	if val, ok := structure["current_action_id"]; ok {
		routine.CurrentActionId = val.(string)
	}

	if val, ok := structure["owner_id"]; ok {
		routine.OwnerId = val.(string)
	}

	if val, ok := structure["person_id"]; ok {
		routine.PersonId = val.(string)
	}

	if val, ok := structure["tasks_ids"]; ok {
		routine.TasksIds = val.([]string)
	}

}

var RoutineStructure = map[string]metis.Primitive{

	"created_at": 4,

	"updated_at": 4,

	"name": 3,

	"start_time": 4,

	"end_time": 4,

	"id": 9,

	"current_action_id": 9,

	"owner_id": 9,

	"person_id": 9,

	"tasks_ids": 10,

	"completed_tasks_ids": 10,

	"actions_ids": 10,
}
