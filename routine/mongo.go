package routine

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoRoutine struct {
	models.MongoID     `bson:",inline"`
	models.Named       `bson:",inline"`
	models.Timestamped `bson:",inline"`
	models.Timed       `bson:",inline"`
	models.UserOwned   `bson:",inline"`

	ETaskIDs          mongo.IDSet   `json:"task_ids" bson:"task_ids"`
	ECompletedTaskIDs mongo.IDSet   `json:"completed_task_ids" bson:"completed_task_ids"`
	EActionIDs        mongo.IDSet   `json:"action_ids" bson:"action_ids"`
	ECurrentActionID  bson.ObjectId `json:"current_action_id" bson:"current_action_id,omitempty"`
}

func (r *mongoRoutine) Kind() data.Kind {
	return kind
}

func (r *mongoRoutine) Schema() data.Schema {
	return schema
}

func (r *mongoRoutine) Version() int {
	return version
}

func (r *mongoRoutine) DBType() data.DBType {
	return mongo.DBType
}

func (r *mongoRoutine) Valid() bool {
	return true
}

func (r *mongoRoutine) SetUser(u models.User) error {
	return r.Schema().Link(r, u, User)
}

func (r *mongoRoutine) IncludeTask(t models.Task) error {
	return r.Schema().Link(r, t, Tasks)
}

func (r *mongoRoutine) ExcludeTask(t models.Task) error {
	return r.Schema().Unlink(r, t, Tasks)
}

func (r *mongoRoutine) Tasks(a *data.Access) (data.RecordIterator, error) {
	if r.CanRead(a.Client) {
		return mongo.NewIDIter(r.ETaskIDs, a.Store), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (r *mongoRoutine) CompleteTask(t models.Task) error {
	return r.Schema().Link(r, t, CompletedTasks)
}

func (r *mongoRoutine) UncompleteTask(t models.Task) error {
	return r.Schema().Unlink(r, t, CompletedTasks)
}

func (r *mongoRoutine) CompletedTasks(a *data.Access) (data.RecordIterator, error) {
	if r.CanRead(a.Client) {
		return mongo.NewIDIter(r.ECompletedTaskIDs, a.Store), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (r *mongoRoutine) ActionCount() int {
	return len(r.ETaskIDs) - len(r.ECompletedTaskIDs)
}

func (r *mongoRoutine) NextAction(a *data.Access) (models.Action, bool) {
	return NewActionRoutine(a, r).Next()
}

func (r *mongoRoutine) CompletedTaskIDs() []data.ID {
	return r.ECompletedTaskIDs.IDs()
}

func (r *mongoRoutine) TaskIDs() []data.ID {
	return r.ETaskIDs.IDs()
}

func (r *mongoRoutine) IncompleteTaskIDs() []data.ID {
	return r.ETaskIDs.NotIn(r.ECompletedTaskIDs).IDs()
}

func (r *mongoRoutine) ActionIDs() []data.ID {
	return r.EActionIDs.IDs()
}

func (r *mongoRoutine) AddAction(a models.Action) error {
	return r.Schema().Link(r, a, Actions)
}

func (r *mongoRoutine) DropAction(a models.Action) error {
	return r.Schema().Unlink(r, a, Actions)
}

func (r *mongoRoutine) SetCurrentAction(a models.Action) {
	r.Schema().Link(r, a, CurrentAction)
}

func (r *mongoRoutine) CurrentAction(a *data.Access, action models.Action) error {
	action.SetID(r.ECurrentActionID)
	return a.PopulateByID(action)
}

func (r *mongoRoutine) CompleteAction(access *data.Access, a models.Action) {
	if a.ID() == r.ECurrentActionID {
		r.ECurrentActionID = ""
	}

	a.Complete()
	m, _ := access.ModelFor(models.TaskKind)
	task := m.(models.Task)
	a.Task(access, task)
	r.CompleteTask(task)
	access.Save(a)
	access.Save(task)
	access.Save(r)
}

func (r *mongoRoutine) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		return r.SetUserID(m.ID())
	case Tasks:
		r.ETaskIDs = mongo.AddID(r.ETaskIDs, m.ID().(bson.ObjectId))
	case CompletedTasks:
		r.ECompletedTaskIDs = mongo.AddID(r.ECompletedTaskIDs, m.ID().(bson.ObjectId))
	case Actions:
		r.EActionIDs = mongo.AddID(r.EActionIDs, m.ID().(bson.ObjectId))
	default:
		return data.ErrUndefinedLink
	}

	return nil
}

func (r *mongoRoutine) Unlink(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		r.DropUserID()
	case Tasks:
		r.ETaskIDs = mongo.DropID(r.ETaskIDs, m.ID().(bson.ObjectId))
	case CompletedTasks:
		r.ECompletedTaskIDs = mongo.DropID(r.ECompletedTaskIDs, m.ID().(bson.ObjectId))
	case Actions:
		r.EActionIDs = mongo.DropID(r.EActionIDs, m.ID().(bson.ObjectId))
	default:
		return data.ErrUndefinedLink
	}

	return nil
}
