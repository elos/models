package routine

import (
	"errors"
	"math/rand"

	"github.com/elos/data"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

type ActionRoutine struct {
	models.Routine
	models.Store
}

func NewActionRoutine(r models.Routine, store models.Store) *ActionRoutine {
	return &ActionRoutine{
		Store:   store,
		Routine: r,
	}
}

var genericError error = errors.New("TODO")

// Actionable
func (r *ActionRoutine) Next() (models.Action, error) {
	ids := r.Routine.IncompleteTaskIDs()
	if len(ids) < 1 {
		return nil, genericError
	}

	i := rand.Intn(len(ids))
	id := ids[i]

	task := r.Store.Task()
	task.SetID(id)
	if err := r.Store.PopulateByID(task); err != nil {
		return nil, err
	}

	action := r.Store.Action()
	action.SetID(r.Store.NewID())
	action.SetName(task.Name())
	action.SetTask(task)

	u, err := r.Store.Unmarshal(models.UserKind, data.AttrMap{
		"id": r.UserID().(bson.ObjectId).Hex(),
	})
	if err != nil {
		return nil, err
	}
	action.SetUser(u.(models.User))

	r.Routine.AddAction(action)
	r.Routine.SetCurrentAction(action)

	r.Save(u)
	r.Save(r.Routine)
	r.Save(action)
	r.Save(task)
	return action, nil
}

func (r *ActionRoutine) ForEachTask(f func(models.Task)) error {
	iter, err := r.Routine.TasksIter(r.Store)
	if err != nil {
		return err
	}

	task := r.Store.Task()

	for iter.Next(task) {
		go f(task)
	}

	return iter.Close()
}
