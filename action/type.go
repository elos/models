package action

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User data.LinkName = models.ActionUser
	Task data.LinkName = models.ActionTask
)

var (
	kind    data.Kind   = models.ActionKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Action, error) {
	switch s.Type() {
	case mongo.DBType:
		a := &mongoAction{}
		a.SetID(s.NewID())
		return a, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}

func Create(s data.Store, a data.AttrMap) (models.Action, error) {
	action, err := New(s)
	if err != nil {
		return action, err
	}

	id, present := a["id"]
	id, valid := id.(data.ID)
	if present && valid {
		if err := action.SetID(id.(data.ID)); err != nil {
			return action, err
		}
	} else {
		if err := action.SetID(s.NewID()); err != nil {
			return action, err
		}
	}

	if err := s.Save(action); err != nil {
		return action, err
	} else {
		return action, nil
	}

}
