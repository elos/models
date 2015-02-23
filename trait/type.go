package trait

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User  data.LinkName = models.TraitUser
	Class data.LinkName = models.TraitClass
)

var (
	kind    data.Kind   = models.TraitKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Trait, error) {
	switch s.Type() {
	case mongo.DBType:
		t := &mongoTrait{}
		t.SetID(s.NewID())
		return t, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}
