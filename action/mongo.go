package action

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/shared"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoAction struct {
	mongo.Model           `bson:",inline"`
	mongo.Named           `bson:",inline"`
	mongo.Timed           `bson:",inline"`
	shared.MongoUserOwned `bson:",inline"`

	ActionableKind data.Kind     `json:"actionable_kind" bson:"actionable_kind"`
	ActionableID   bson.ObjectId `json:"actionable_id" bson:"actionable_id,omitempty"`

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

func (a *mongoAction) SetUser(u models.User) error {
	return a.Schema().Link(a, u, User)
}

func (a *mongoAction) SetActionable(actionable models.Actionable) {
	a.ActionableKind = actionable.Kind()
	a.ActionableID = actionable.ID().(bson.ObjectId)
}

func (a *mongoAction) Actionable(access data.Access) (models.Actionable, error) {
	if !a.HasActionable() {
		return nil, models.ErrEmptyRelationship
	}

	m, err := access.ModelFor(a.ActionableKind)
	if err != nil {
		return nil, err
	}

	m.SetID(a.ActionableID)
	if err = access.PopulateByID(m); err != nil {
		return nil, err
	}

	err = access.PopulateByID(m)

	return m.(models.Actionable), err
}

func (a *mongoAction) DropActionable() {
	a.ActionableKind = data.Kind("")
	a.ActionableID = *new(bson.ObjectId)
}

func (a *mongoAction) HasActionable() bool {
	return !mongo.EmptyID(a.ActionableID) && !data.EmptyKind(a.ActionableKind)
}

func (a *mongoAction) SetCompleted(b bool) {
	a.ECompleted = b
}

func (a *mongoAction) Completed() bool {
	return a.ECompleted
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
		return nil, models.CastError(models.TaskKind)
	}

	if !a.CanRead(access.Client()) {
		return nil, data.ErrAccessDenial
	}

	if mongo.EmptyID(a.ETaskID) {
		return nil, data.NewEmptyLinkError(a, models.RMap[models.ActionKind][Task])
	}

	t.SetID(a.ETaskID)

	return t, access.PopulateByID(t)
}

func (a *mongoAction) Complete() {
	a.SetEndTime(time.Now())
	a.SetCompleted(true)
}
