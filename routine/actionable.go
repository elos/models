package routine

import (
	"math/rand"

	"github.com/elos/data"
	"github.com/elos/models"
)

type ActionRoutine struct {
	models.Routine
	*data.Access
}

func NewActionRoutine(a *data.Access, r models.Routine) *ActionRoutine {
	return &ActionRoutine{
		Access:  a,
		Routine: r,
	}
}

// Actionable
func (r *ActionRoutine) Next() (models.Action, bool) {
	ids := r.Routine.IncompleteTaskIDs()
	i := rand.Intn(len(ids))
	id := ids[i]

	model, _ := r.Access.ModelFor(models.TaskKind)
	task := model.(models.Task)
	task.SetID(id)
	if err := r.Access.PopulateByID(task); err != nil {
		return nil, false
	}

	model, _ = r.Access.ModelFor(models.ActionKind)
	action := model.(models.Action)
	action.SetID(r.Access.NewID())
	action.SetName(task.Name())
	action.SetTask(task)
	action.SetUserID(r.UserID())

	r.Routine.AddAction(action)

	r.Save(r.Routine)
	r.Save(action)
	r.Save(task)
	return action, true
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
