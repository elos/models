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
type Action struct {
	ActionableId   string    `json:"actionable_id" bson:"actionable_id"`
	ActionableKind string    `json:"actionable_kind" bson:"actionable_kind"`
	Completed      bool      `json:"completed" bson:"completed"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	EndTime        time.Time `json:"end_time" bson:"end_time"`
	Id             string    `json:"id" bson:"_id,omitempty"`
	Name           string    `json:"name" bson:"name"`
	OwnerId        string    `json:"owner_id" bson:"owner_id"`
	PersonId       string    `json:"person_id" bson:"person_id"`
	StartTime      time.Time `json:"start_time" bson:"start_time"`
	TaskId         string    `json:"task_id" bson:"task_id"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}

func NewAction() *Action {
	return &Action{}
}

func FindAction(db data.DB, id data.ID) (*Action, error) {

	action := NewAction()
	action.SetID(id)

	return action, db.PopulateByID(action)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (action *Action) Kind() data.Kind {
	return ActionKind
}

// just returns itself for now
func (action *Action) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = action.ID()
	return foo
}

func (action *Action) SetID(id data.ID) {
	action.Id = id.String()
}

func (action *Action) ID() data.ID {
	return data.ID(action.Id)
}

func (action *Action) SetActionable(actionableArgument Actionable) error {
	action.ActionableId = actionableArgument.ID().String()
	return nil
}

func (action *Action) Actionable(db data.DB) (Actionable, error) {
	if action.ActionableId == "" {
		return nil, ErrEmptyLink
	}

	m := ModelFor(data.Kind(action.ActionableKind))
	actionable := m.(Actionable)

	pid, _ := mongo.ParseObjectID(action.ActionableId)

	actionable.SetID(data.ID(pid.Hex()))
	return actionable, db.PopulateByID(actionable)

}

func (action *Action) SetOwner(userArgument *User) error {
	action.OwnerId = userArgument.ID().String()
	return nil
}

func (action *Action) Owner(db data.DB) (*User, error) {
	if action.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	pid, _ := mongo.ParseObjectID(action.OwnerId)
	userArgument.SetID(data.ID(pid.Hex()))
	return userArgument, db.PopulateByID(userArgument)

}

func (action *Action) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := action.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := action.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(action); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

func (action *Action) SetPerson(personArgument *Person) error {
	action.PersonId = personArgument.ID().String()
	return nil
}

func (action *Action) Person(db data.DB) (*Person, error) {
	if action.PersonId == "" {
		return nil, ErrEmptyLink
	}

	personArgument := NewPerson()
	pid, _ := mongo.ParseObjectID(action.PersonId)
	personArgument.SetID(data.ID(pid.Hex()))
	return personArgument, db.PopulateByID(personArgument)

}

func (action *Action) PersonOrCreate(db data.DB) (*Person, error) {
	person, err := action.Person(db)

	if err == ErrEmptyLink {
		person := NewPerson()
		person.SetID(db.NewID())
		if err := action.SetPerson(person); err != nil {
			return nil, err
		}

		if err := db.Save(person); err != nil {
			return nil, err
		}

		if err := db.Save(action); err != nil {
			return nil, err
		}

		return person, nil
	} else {
		return person, err
	}
}

func (action *Action) SetTask(taskArgument *Task) error {
	action.TaskId = taskArgument.ID().String()
	return nil
}

func (action *Action) Task(db data.DB) (*Task, error) {
	if action.TaskId == "" {
		return nil, ErrEmptyLink
	}

	taskArgument := NewTask()
	pid, _ := mongo.ParseObjectID(action.TaskId)
	taskArgument.SetID(data.ID(pid.Hex()))
	return taskArgument, db.PopulateByID(taskArgument)

}

func (action *Action) TaskOrCreate(db data.DB) (*Task, error) {
	task, err := action.Task(db)

	if err == ErrEmptyLink {
		task := NewTask()
		task.SetID(db.NewID())
		if err := action.SetTask(task); err != nil {
			return nil, err
		}

		if err := db.Save(task); err != nil {
			return nil, err
		}

		if err := db.Save(action); err != nil {
			return nil, err
		}

		return task, nil
	} else {
		return task, err
	}
}

// BSON {{{
func (action *Action) GetBSON() (interface{}, error) {

	return struct {
		Completed bool `json:"completed" bson:"completed"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ActionableId string `json:"actionable_id" bson:"actionable_id"`

		ActionableKind string `json:"actionable_kind" bson:"actionable_kind"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PersonId string `json:"person_id" bson:"person_id"`

		TaskId string `json:"task_id" bson:"task_id"`
	}{

		Completed: action.Completed,

		CreatedAt: action.CreatedAt,

		EndTime: action.EndTime,

		Name: action.Name,

		StartTime: action.StartTime,

		UpdatedAt: action.UpdatedAt,

		ActionableId: action.ActionableId,

		ActionableKind: action.ActionableKind,

		OwnerId: action.OwnerId,

		PersonId: action.PersonId,

		TaskId: action.TaskId,
	}, nil

}

func (action *Action) SetBSON(raw bson.Raw) error {

	tmp := struct {
		Completed bool `json:"completed" bson:"completed"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ActionableId string `json:"actionable_id" bson:"actionable_id"`

		ActionableKind string `json:"actionable_kind" bson:"actionable_kind"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PersonId string `json:"person_id" bson:"person_id"`

		TaskId string `json:"task_id" bson:"task_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	action.Completed = tmp.Completed

	action.CreatedAt = tmp.CreatedAt

	action.EndTime = tmp.EndTime

	action.Id = tmp.Id.Hex()

	action.Name = tmp.Name

	action.StartTime = tmp.StartTime

	action.UpdatedAt = tmp.UpdatedAt

	action.ActionableId = tmp.ActionableId

	action.ActionableKind = tmp.ActionableKind

	action.OwnerId = tmp.OwnerId

	action.PersonId = tmp.PersonId

	action.TaskId = tmp.TaskId

	return nil

}

// BSON }}}

func (action *Action) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["updated_at"]; ok {
		action.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["name"]; ok {
		action.Name = val.(string)
	}

	if val, ok := structure["start_time"]; ok {
		action.StartTime = val.(time.Time)
	}

	if val, ok := structure["end_time"]; ok {
		action.EndTime = val.(time.Time)
	}

	if val, ok := structure["completed"]; ok {
		action.Completed = val.(bool)
	}

	if val, ok := structure["id"]; ok {
		action.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		action.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["task_id"]; ok {
		action.TaskId = val.(string)
	}

	if val, ok := structure["owner_id"]; ok {
		action.OwnerId = val.(string)
	}

	if val, ok := structure["person_id"]; ok {
		action.PersonId = val.(string)
	}

	if val, ok := structure["actionable_id"]; ok {
		action.ActionableId = val.(string)
	}

	if val, ok := structure["actionable_kind"]; ok {
		action.ActionableKind = val.(string)
	}

}

var ActionStructure = map[string]metis.Primitive{

	"created_at": 4,

	"updated_at": 4,

	"name": 3,

	"start_time": 4,

	"end_time": 4,

	"completed": 0,

	"id": 9,

	"owner_id": 9,

	"person_id": 9,

	"actionable_id": 9,

	"actionable_kind": 3,

	"task_id": 9,
}
