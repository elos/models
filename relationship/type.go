package relationship

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User  data.LinkName = models.RelationshipUser
	Class data.LinkName = models.RelationshipClass
)

var (
	kind    data.Kind   = models.RelationshipKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Relationship, error) {
	switch s.Type() {
	case mongo.DBType:
		r := &mongoRelationship{}
		r.SetID(s.NewID())
		return r, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}
