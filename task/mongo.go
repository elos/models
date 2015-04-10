package task

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/shared"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoTask struct {
	mongo.Model           `bson:",inline"`
	mongo.Named           `bson:",inline"`
	mongo.Timed           `bson:",inline"`
	shared.MongoUserOwned `bson:",inline"`

	TaskIDs mongo.IDSet `json:"task_dependencies" bson:"task_dependencies"`
}

func (t *mongoTask) DBType() data.DBType {
	return mongo.DBType
}

func (t *mongoTask) Kind() data.Kind {
	return kind
}

func (t *mongoTask) Schema() data.Schema {
	return schema
}

func (t *mongoTask) Version() int {
	return version
}

func (t *mongoTask) Valid() bool {
	return true
}

func (t *mongoTask) Save(s data.Store) error {
	return s.Save(t)
}

func (t *mongoTask) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		t.SetUserID(m.ID())
	case Dependencies:
		t.TaskIDs = mongo.AddID(t.TaskIDs, m.ID().(bson.ObjectId))
	default:
		return data.ErrUndefinedLink
	}

	return nil
}

func (t *mongoTask) Unlink(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		t.DropUserID()
	case Dependencies:
		t.TaskIDs = mongo.DropID(t.TaskIDs, m.ID().(bson.ObjectId))
	default:
		return data.ErrUndefinedLink
	}

	return nil
}

func (t *mongoTask) SetUser(u models.User) error {
	return t.Schema().Link(t, u, User)
}

func (t *mongoTask) AddDependency(other models.Task) error {
	return t.Schema().Link(t, other, Dependencies)
}

func (t *mongoTask) DropDependency(other models.Task) error {
	return t.Schema().Unlink(t, other, Dependencies)
}

func (t *mongoTask) DependenciesIter(store models.Store) (data.ModelIterator, error) {
	if !store.Compatible(t) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(t.TaskIDs, store), nil
}

func (t *mongoTask) Dependencies(store models.Store) ([]models.Task, error) {
	if !store.Compatible(t) {
		return nil, data.ErrInvalidDBType
	}

	tasks := make([]models.Task, 0)
	iter := mongo.NewIDIter(t.TaskIDs, store)
	task := store.Task()
	for iter.Next(task) {
		tasks = append(tasks, task)
		task = store.Task()
	}

	return tasks, iter.Close()
}
