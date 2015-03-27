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
	data.Access
}

func NewActionRoutine(a data.Access, r models.Routine) *ActionRoutine {
	return &ActionRoutine{
		Access:  a,
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

	model, _ := r.Access.ModelFor(models.TaskKind)
	task := model.(models.Task)
	task.SetID(id)
	if err := r.Access.PopulateByID(task); err != nil {
		return nil, err
	}

	model, _ = r.Access.ModelFor(models.ActionKind)
	action := model.(models.Action)
	action.SetID(r.Access.NewID())
	action.SetName(task.Name())
	action.SetTask(task)

	u, err := r.Access.Unmarshal(models.UserKind, data.AttrMap{
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
	iter, err := r.Routine.Tasks(r.Access)
	if err != nil {
		return err
	}

	model, err := r.Access.ModelFor(models.TaskKind)
	task := model.(models.Task)

	for iter.Next(task) {
		go f(task)
	}

	return iter.Close()
}
