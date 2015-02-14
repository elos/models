package action

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoAction struct {
	models.MongoID     `bson:",inline"`
	models.Named       `bson:",inline"`
	models.Timestamped `bson:",inline"`
	models.Timed       `bson:",inline"`
	models.UserOwned   `bson:",inline"`

	ETaskID bson.ObjectId `json:"task_id" bson:"task_id,omitempty"`
}

func (a *mongoAction) Kind() data.Kind {
	return kind
}

func (a *mongoAction) Schema() data.Schema {
	return schema
}

func (a *mongoAction) Version() int {
	return version
}

func (a *mongoAction) DBType() data.DBType {
	return mongo.DBType
}

func (a *mongoAction) Valid() bool {
	return true
}

func (a *mongoAction) SetUser(u models.User) error {
	return a.Schema().Link(a, u, User)
}

func (a *mongoAction) SetTask(t models.Task) error {
	return a.Schema().Link(a, t, Task)
}

func (a *mongoAction) Task(access *data.Access, t models.Task) error {
	t.SetID(a.ETaskID)
	return access.PopulateByID(t)
}

func (a *mongoAction) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		return a.SetUserID(m.ID())
	case Task:
		a.ETaskID = m.ID().(bson.ObjectId)
	default:
		return data.ErrUndefinedLink
	}

	return nil
}

func (a *mongoAction) Unlink(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		a.DropUserID()
	case Task:
		a.ETaskID = ""
	default:
		return data.ErrUndefinedLink
	}

	return nil
}
