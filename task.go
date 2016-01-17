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
type Task struct {
	Complete         bool        `json:"complete" bson:"complete"`
	CreatedAt        time.Time   `json:"created_at" bson:"created_at"`
	Deadline         time.Time   `json:"deadline" bson:"deadline"`
	DeletedAt        time.Time   `json:"deleted_at" bson:"deleted_at"`
	Id               string      `json:"id" bson:"_id,omitempty"`
	Name             string      `json:"name" bson:"name"`
	OwnerId          string      `json:"owner_id" bson:"owner_id"`
	PrerequisitesIds []string    `json:"prerequisites_ids" bson:"prerequisites_ids"`
	Stages           []time.Time `json:"stages" bson:"stages"`
	UpdatedAt        time.Time   `json:"updated_at" bson:"updated_at"`
}

func NewTask() *Task {
	return &Task{}
}

func FindTask(db data.DB, id data.ID) (*Task, error) {

	task := NewTask()
	task.SetID(id)

	return task, db.PopulateByID(task)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (task *Task) Kind() data.Kind {
	return TaskKind
}

// just returns itself for now
func (task *Task) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = task.ID()
	return foo
}

func (task *Task) SetID(id data.ID) {
	task.Id = id.String()
}

func (task *Task) ID() data.ID {
	return data.ID(task.Id)
}

func (task *Task) SetOwner(userArgument *User) error {
	task.OwnerId = userArgument.ID().String()
	return nil
}

func (task *Task) Owner(db data.DB) (*User, error) {
	if task.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(task.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (task *Task) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := task.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := task.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(task); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

func (task *Task) IncludePrerequisite(prerequisite *Task) {
	otherID := prerequisite.ID().String()
	for i := range task.PrerequisitesIds {
		if task.PrerequisitesIds[i] == otherID {
			return
		}
	}
	task.PrerequisitesIds = append(task.PrerequisitesIds, otherID)
}

func (task *Task) ExcludePrerequisite(prerequisite *Task) {
	tmp := make([]string, 0)
	id := prerequisite.ID().String()
	for _, s := range task.PrerequisitesIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	task.PrerequisitesIds = tmp
}

func (task *Task) PrerequisitesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(task.PrerequisitesIds), db), nil
}

func (task *Task) Prerequisites(db data.DB) (prerequisites []*Task, err error) {
	prerequisites = make([]*Task, len(task.PrerequisitesIds))
	prerequisite := NewTask()
	for i, id := range task.PrerequisitesIds {
		prerequisite.Id = id
		if err = db.PopulateByID(prerequisite); err != nil {
			return
		}

		prerequisites[i] = prerequisite
		prerequisite = NewTask()
	}

	return
}

// BSON {{{
func (task *Task) GetBSON() (interface{}, error) {

	return struct {
		Complete bool `json:"complete" bson:"complete"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		Deadline time.Time `json:"deadline" bson:"deadline"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		Stages []time.Time `json:"stages" bson:"stages"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PrerequisitesIds []string `json:"prerequisites_ids" bson:"prerequisites_ids"`
	}{

		Complete: task.Complete,

		CreatedAt: task.CreatedAt,

		Deadline: task.Deadline,

		DeletedAt: task.DeletedAt,

		Name: task.Name,

		Stages: task.Stages,

		UpdatedAt: task.UpdatedAt,

		OwnerId: task.OwnerId,

		PrerequisitesIds: task.PrerequisitesIds,
	}, nil

}

func (task *Task) SetBSON(raw bson.Raw) error {

	tmp := struct {
		Complete bool `json:"complete" bson:"complete"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		Deadline time.Time `json:"deadline" bson:"deadline"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		Stages []time.Time `json:"stages" bson:"stages"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PrerequisitesIds []string `json:"prerequisites_ids" bson:"prerequisites_ids"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	task.Complete = tmp.Complete

	task.CreatedAt = tmp.CreatedAt

	task.Deadline = tmp.Deadline

	task.DeletedAt = tmp.DeletedAt

	task.Id = tmp.Id.Hex()

	task.Name = tmp.Name

	task.Stages = tmp.Stages

	task.UpdatedAt = tmp.UpdatedAt

	task.OwnerId = tmp.OwnerId

	task.PrerequisitesIds = tmp.PrerequisitesIds

	return nil

}

// BSON }}}

func (task *Task) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["deleted_at"]; ok {
		task.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["name"]; ok {
		task.Name = val.(string)
	}

	if val, ok := structure["deadline"]; ok {
		task.Deadline = val.(time.Time)
	}

	if val, ok := structure["stages"]; ok {
		task.Stages = val.([]time.Time)
	}

	if val, ok := structure["complete"]; ok {
		task.Complete = val.(bool)
	}

	if val, ok := structure["id"]; ok {
		task.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		task.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		task.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["prerequisites_ids"]; ok {
		task.PrerequisitesIds = val.([]string)
	}

	if val, ok := structure["owner_id"]; ok {
		task.OwnerId = val.(string)
	}

}

var TaskStructure = map[string]metis.Primitive{

	"deadline": 4,

	"stages": 8,

	"complete": 0,

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"name": 3,

	"owner_id": 9,

	"prerequisites_ids": 10,
}
