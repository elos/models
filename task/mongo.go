package task

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoTask struct {
	models.MongoID     `bson:",inline"`
	models.Named       `bson:",inline"`
	models.Timestamped `bson:",inline"`
	models.Timed       `bson:",inline"`
	models.UserOwned   `bson:",inline"`

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

func (t *mongoTask) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = t.UserID()
	return a
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

// Accessors

func (t *mongoTask) SetUser(u models.User) error {
	return t.Schema().Link(t, u, User)
}

func (t *mongoTask) AddDependency(other models.Task) error {
	return t.Schema().Link(t, other, Dependencies)
}

func (t *mongoTask) DropDependency(other models.Task) error {
	return t.Schema().Unlink(t, other, Dependencies)
}

func (t *mongoTask) Dependencies(a data.Access) (data.RecordIterator, error) {
	if t.CanRead(a.Client) {
		return mongo.NewIDIter(t.TaskIDs, a.Store), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (t *mongoTask) CanRead(c data.Client) bool {
	if c.Kind() != models.UserKind {
		return false
	}

	if t.UserID().Valid() && c.ID() != t.UserID() {
		return false
	}

	return true
}

func (t *mongoTask) CanWrite(c data.Client) bool {
	return t.CanRead(c)
}
