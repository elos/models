package routine

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/shared"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoRoutine struct {
	mongo.Model           `bson:",inline"`
	mongo.Named           `bson:",inline"`
	mongo.Timed           `bson:",inline"`
	shared.MongoUserOwned `bson:",inline"`

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

func (r *mongoRoutine) SetUser(u models.User) error {
	return r.Schema().Link(r, u, User)
}

func (r *mongoRoutine) IncludeTask(t models.Task) error {
	return r.Schema().Link(r, t, Tasks)
}

func (r *mongoRoutine) ExcludeTask(t models.Task) error {
	return r.Schema().Unlink(r, t, Tasks)
}

func (r *mongoRoutine) TasksIter(store models.Store) (data.ModelIterator, error) {
	if !store.Compatible(r) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(r.ETaskIDs, store), nil
}

func (r *mongoRoutine) Tasks(store models.Store) ([]models.Task, error) {
	if !store.Compatible(r) {
		return nil, data.ErrInvalidDBType
	}

	tasks := make([]models.Task, 0)
	iter := mongo.NewIDIter(r.ETaskIDs, store)
	task := store.Task()
	for iter.Next(task) {
		tasks = append(tasks, task)
		task = store.Task()
	}

	return tasks, iter.Close()
}

func (r *mongoRoutine) CompleteTask(t models.Task) error {
	return r.Schema().Link(r, t, CompletedTasks)
}

func (r *mongoRoutine) UncompleteTask(t models.Task) error {
	return r.Schema().Unlink(r, t, CompletedTasks)
}

func (r *mongoRoutine) CompletedTasksIter(store models.Store) (data.ModelIterator, error) {
	if !store.Compatible(r) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(r.ECompletedTaskIDs, store), nil
}

func (r *mongoRoutine) CompletedTasks(store models.Store) ([]models.Task, error) {
	if !store.Compatible(r) {
		return nil, data.ErrInvalidDBType
	}

	tasks := make([]models.Task, 0)
	iter := mongo.NewIDIter(r.ECompletedTaskIDs, store)
	task := store.Task()
	for iter.Next(task) {
		tasks = append(tasks, task)
		task = store.Task()
	}

	return tasks, iter.Close()
}

func (r *mongoRoutine) ActionCount() int {
	return len(r.ETaskIDs) - len(r.ECompletedTaskIDs)
}

func (r *mongoRoutine) NextAction(store models.Store) (models.Action, error) {
	return NewActionRoutine(r, store).Next()
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

func (r *mongoRoutine) CurrentAction(store models.Store) (models.Action, error) {
	if !store.Compatible(r) {
		return nil, data.ErrInvalidDBType
	}

	if mongo.EmptyID(r.ECurrentActionID) {
		return nil, models.ErrEmptyRelationship
	}

	action := store.Action()
	action.SetID(r.ECurrentActionID)
	return action, store.PopulateByID(action)
}

func (r *mongoRoutine) StartAction(store models.Store, a models.Action) error {
	panic("not implemented")
}

func (r *mongoRoutine) CompleteAction(store models.Store, action models.Action) error {
	if action.ID() == r.ECurrentActionID {
		r.ECurrentActionID = ""
	}

	action.Complete()

	task, err := action.Task(store)
	if err != nil {
		return err
	}

	r.CompleteTask(task)

	if err := store.Save(action); err != nil {
		return err
	}
	if err := store.Save(task); err != nil {
		return err
	}
	if err := store.Save(r); err != nil {
		return err
	}

	return nil
}
