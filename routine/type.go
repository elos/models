package routine

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User           data.LinkName = models.RoutineUser
	Tasks          data.LinkName = models.RoutineTasks
	CompletedTasks data.LinkName = models.RoutineCompletedTasks
	Actions        data.LinkName = models.RoutineActions
	CurrentAction  data.LinkName = models.RoutineCurrentAction
)

var (
	kind    data.Kind   = models.RoutineKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Routine, error) {
	switch s.Type() {
	case mongo.DBType:
		return &mongoRoutine{}, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}

func Create(s data.Store, a data.AttrMap) (models.Routine, error) {
	routine, err := New(s)
	if err != nil {
		return routine, err
	}

	id, present := a["id"]
	id, valid := id.(data.ID)
	if present && valid {
		if err := routine.SetID(id.(data.ID)); err != nil {
			return routine, err
		}
	} else {
		if err := routine.SetID(s.NewID()); err != nil {
			return routine, err
		}
	}

	if err := s.Save(routine); err != nil {
		return routine, err
	} else {
		return routine, nil
	}
}
