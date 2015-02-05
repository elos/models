package task

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

type mongoTask struct {
	EID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ECreatedAt time.Time     `json:"created_at" bson:"created_at"`
	EUpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	EName      string    `json:"name" bson:"name"`
	EStartTime time.Time `json:"start_time" bson:"end_time"`
	EEndTime   time.Time `json:"start_time" bson:"start_time"`

	UserID  bson.ObjectId `json:"user_id" bson:"user_id,omitempty"`
	TaskIDs mongo.IDSet   `json:"task_dependencies" bson:"task_dependencies"`
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

func (t *mongoTask) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = t.UserID
	return a
}

func (t *mongoTask) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		id, ok := m.ID().(bson.ObjectId)
		if !ok {
			return data.ErrInvalidID
		}

		t.UserID = id
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
		t.UserID = *new(bson.ObjectId)
	case Dependencies:
		t.TaskIDs = mongo.DropID(t.TaskIDs, m.ID().(bson.ObjectId))
	default:
		return data.ErrUndefinedLink
	}

	return nil
}

// Accessors

func (t *mongoTask) ID() data.ID {
	return t.EID
}

func (t *mongoTask) SetID(id data.ID) {
	t.EID = id.(bson.ObjectId)
}

func (t *mongoTask) CreatedAt() time.Time {
	return t.ECreatedAt
}

func (t *mongoTask) SetCreatedAt(c time.Time) {
	t.ECreatedAt = c
}

func (t *mongoTask) UpdatedAt() time.Time {
	return t.EUpdatedAt
}

func (t *mongoTask) SetUpdatedAt(u time.Time) {
	t.EUpdatedAt = u
}

func (t *mongoTask) Name() string {
	return t.EName
}

func (t *mongoTask) SetName(n string) {
	t.EName = n
}

func (t *mongoTask) SetUser(u models.User) error {
	return t.Schema().Link(t, u, User)
}

func (t *mongoTask) AddDependency(other models.Task) error {
	return t.Schema().Link(t, other, Dependencies)
}

func (t *mongoTask) RemoveDependency(other models.Task) error {
	return t.Schema().Unlink(t, other, Dependencies)
}

func (t *mongoTask) Dependencies(s data.Store) data.RecordIterator {
	return mongo.NewIDIter(t.TaskIDs, s)
}
