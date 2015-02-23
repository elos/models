package class

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User          data.LinkName = models.ClassUser
	Ontology      data.LinkName = models.ClassOntology
	Objects       data.LinkName = models.ClassObjects
	Traits        data.LinkName = models.ClassTraits
	Relationships data.LinkName = models.ClassRelationships
)

var (
	kind    data.Kind   = models.ClassKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Class, error) {
	switch s.Type() {
	case mongo.DBType:
		c := &mongoClass{}
		c.SetID(s.NewID())
		return c, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}
