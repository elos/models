package action

import (
	"errors"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoAction struct {
	mongo.Model      `bson:",inline"`
	mongo.Named      `bson:",inline"`
	mongo.Timed      `bson:",inline"`
	models.UserOwned `bson:",inline"`

	ECompleted bool          `json:"completed" bson:"completed"`
	ETaskID    bson.ObjectId `json:"task_id" bson:"task_id,omitempty"`
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

func (a *mongoAction) SetUser(u models.User) error {
	return a.Schema().Link(a, u, User)
}

func (a *mongoAction) SetTask(t models.Task) error {
	return a.Schema().Link(a, t, Task)
}

func (a *mongoAction) Task(access data.Access) (models.Task, error) {
	m, err := access.ModelFor(models.TaskKind)
	if err != nil {
		return nil, err
	}
	t, ok := m.(models.Task)
	if !ok {
		return nil, errors.New("TODO")
	}

	t.SetID(a.ETaskID)

	err = access.PopulateByID(t)
	return t, err
}

func (a *mongoAction) Completed() bool {
	return a.ECompleted
}

func (a *mongoAction) Complete() {
	a.SetEndTime(time.Now())
	a.ECompleted = true
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
