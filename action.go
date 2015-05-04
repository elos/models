package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Action struct {
	ActionableID   string    `json:"actionable_id" bson:"actionable_id"`
	ActionableKind string    `json:"actionable_kind" bson:"actionable_kind"`
	Completed      bool      `json:"completed" bson:"completed"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	EndTime        time.Time `json:"end_time" bson:"end_time"`
	Id             string    `json:"id" bson:"_id,omitempty"`
	Name           string    `json:"name" bson:"name"`
	StartTime      time.Time `json:"start_time" bson:"start_time"`
	TaskID         string    `json:"task_id" bson:"task_id"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
	UserID         string    `json:"user_id" bson:"user_id"`
}

func NewAction() *Action {
	return &Action{}
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

func (action *Action) SetActionable(actionable Actionable) error {
	action.ActionableID = actionable.ID().String()
	return nil
}

func (action *Action) Actionable(db data.DB) (Actionable, error) {
	if action.ActionableID == "" {
		return nil, ErrEmptyLink
	}

	m := ModelFor(data.Kind(action.ActionableKind))
	actionable := m.(Actionable)

	pid, _ := mongo.ParseObjectID(action.ActionableID)

	actionable.SetID(data.ID(pid.Hex()))
	return actionable, db.PopulateByID(actionable)

}

func (action *Action) SetTask(task *Task) error {
	action.TaskID = task.ID().String()
	return nil
}

func (action *Action) Task(db data.DB) (*Task, error) {
	if action.TaskID == "" {
		return nil, ErrEmptyLink
	}

	task := NewTask()
	pid, _ := mongo.ParseObjectID(action.TaskID)
	task.SetID(data.ID(pid.Hex()))
	return task, db.PopulateByID(task)

}

func (action *Action) SetUser(user *User) error {
	action.UserID = user.ID().String()
	return nil
}

func (action *Action) User(db data.DB) (*User, error) {
	if action.UserID == "" {
		return nil, ErrEmptyLink
	}

	user := NewUser()
	pid, _ := mongo.ParseObjectID(action.UserID)
	user.SetID(data.ID(pid.Hex()))
	return user, db.PopulateByID(user)

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

		ActionableID string `json:"actionable_id" bson:"actionable_id"`

		ActionableKind string `json:"actionable_kind" bson:"actionable_kind"`

		TaskID string `json:"task_id" bson:"task_id"`

		UserID string `json:"user_id" bson:"user_id"`
	}{

		Completed: action.Completed,

		CreatedAt: action.CreatedAt,

		EndTime: action.EndTime,

		Name: action.Name,

		StartTime: action.StartTime,

		UpdatedAt: action.UpdatedAt,

		ActionableID: action.ActionableID,

		ActionableKind: action.ActionableKind,

		TaskID: action.TaskID,

		UserID: action.UserID,
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

		ActionableID string `json:"actionable_id" bson:"actionable_id"`

		ActionableKind string `json:"actionable_kind" bson:"actionable_kind"`

		TaskID string `json:"task_id" bson:"task_id"`

		UserID string `json:"user_id" bson:"user_id"`
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

	action.ActionableID = tmp.ActionableID

	action.ActionableKind = tmp.ActionableKind

	action.TaskID = tmp.TaskID

	action.UserID = tmp.UserID

	return nil

}

// BSON }}}
