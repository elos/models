package task

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User         data.LinkName = models.TaskUser
	Dependencies data.LinkName = models.TaskDependencies
)

var (
	kind    data.Kind   = models.TaskKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func Setup(s data.Schema, k data.Kind, v int) {
	schema, kind, version = s, k, v
}

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Task, error) {
	switch s.Type() {
	case mongo.DBType:
		return &mongoTask{}, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}

func Create(s data.Store, a data.AttrMap) (models.Task, error) {
	task, err := New(s)
	if err != nil {
		return task, err
	}

	id, present := a["id"]
	id, valid := id.(data.ID)
	if present && valid {
		if err := task.SetID(id.(data.ID)); err != nil {
			return task, err
		}
	} else {
		if err := task.SetID(s.NewID()); err != nil {
			return task, err
		}
	}

	if err := s.Save(task); err != nil {
		return task, err
	} else {
		return task, nil
	}
}
