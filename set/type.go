package set

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User   data.LinkName = models.SetUser
	Models data.LinkName = models.SetModels
)

var (
	kind    data.Kind   = models.ActionKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Set, error) {
	switch s.Type() {
	case mongo.DBType:
		set := &mongoSet{}
		set.SetID(s.NewID())
		return set, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}
