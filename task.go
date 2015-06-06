package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Task struct {
	CreatedAt           time.Time `json:"created_at" bson:"created_at"`
	EndTime             time.Time `json:"end_time" bson:"end_time"`
	Id                  string    `json:"id" bson:"_id,omitempty"`
	Name                string    `json:"name" bson:"name"`
	OwnerID             string    `json:"owner_id" bson:"owner_id"`
	PersonID            string    `json:"person_id" bson:"person_id"`
	StartTime           time.Time `json:"start_time" bson:"start_time"`
	TaskDependenciesIDs []string  `json:"task_dependencies_ids" bson:"task_dependencies_ids"`
	UpdatedAt           time.Time `json:"updated_at" bson:"updated_at"`
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

func (task *Task) SetOwner(user *User) error {
	task.OwnerID = user.ID().String()
	return nil
}

func (task *Task) Owner(db data.DB) (*User, error) {
	if task.OwnerID == "" {
		return nil, ErrEmptyLink
	}

	user := NewUser()
	pid, _ := mongo.ParseObjectID(task.OwnerID)
	user.SetID(data.ID(pid.Hex()))
	return user, db.PopulateByID(user)

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

func (task *Task) SetPerson(person *Person) error {
	task.PersonID = person.ID().String()
	return nil
}

func (task *Task) Person(db data.DB) (*Person, error) {
	if task.PersonID == "" {
		return nil, ErrEmptyLink
	}

	person := NewPerson()
	pid, _ := mongo.ParseObjectID(task.PersonID)
	person.SetID(data.ID(pid.Hex()))
	return person, db.PopulateByID(person)

}

func (task *Task) PersonOrCreate(db data.DB) (*Person, error) {
	person, err := task.Person(db)

	if err == ErrEmptyLink {
		person := NewPerson()
		person.SetID(db.NewID())
		if err := task.SetPerson(person); err != nil {
			return nil, err
		}

		if err := db.Save(person); err != nil {
			return nil, err
		}

		if err := db.Save(task); err != nil {
			return nil, err
		}

		return person, nil
	} else {
		return person, err
	}
}

func (task *Task) IncludeTaskDependency(taskDependency *Task) {
	task.TaskDependenciesIDs = append(task.TaskDependenciesIDs, taskDependency.ID().String())
}

func (task *Task) ExcludeTaskDependency(taskDependency *Task) {
	tmp := make([]string, 0)
	id := taskDependency.ID().String()
	for _, s := range task.TaskDependenciesIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	task.TaskDependenciesIDs = tmp
}

func (task *Task) TaskDependenciesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(task.TaskDependenciesIDs), db), nil
}

func (task *Task) TaskDependencies(db data.DB) ([]*Task, error) {

	task_dependencies := make([]*Task, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(task.TaskDependenciesIDs), db)
	task_dependency := NewTask()
	for iter.Next(task_dependency) {
		task_dependencies = append(task_dependencies, task_dependency)
		task_dependency = NewTask()
	}
	return task_dependencies, nil
}

// BSON {{{
func (task *Task) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerID string `json:"owner_id" bson:"owner_id"`

		PersonID string `json:"person_id" bson:"person_id"`

		TaskDependenciesIDs []string `json:"task_dependencies_ids" bson:"task_dependencies_ids"`
	}{

		CreatedAt: task.CreatedAt,

		EndTime: task.EndTime,

		Name: task.Name,

		StartTime: task.StartTime,

		UpdatedAt: task.UpdatedAt,

		OwnerID: task.OwnerID,

		PersonID: task.PersonID,

		TaskDependenciesIDs: task.TaskDependenciesIDs,
	}, nil

}

func (task *Task) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerID string `json:"owner_id" bson:"owner_id"`

		PersonID string `json:"person_id" bson:"person_id"`

		TaskDependenciesIDs []string `json:"task_dependencies_ids" bson:"task_dependencies_ids"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	task.CreatedAt = tmp.CreatedAt

	task.EndTime = tmp.EndTime

	task.Id = tmp.Id.Hex()

	task.Name = tmp.Name

	task.StartTime = tmp.StartTime

	task.UpdatedAt = tmp.UpdatedAt

	task.OwnerID = tmp.OwnerID

	task.PersonID = tmp.PersonID

	task.TaskDependenciesIDs = tmp.TaskDependenciesIDs

	return nil

}

// BSON }}}
