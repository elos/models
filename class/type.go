package class

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	kind    data.Kind   = models.ClassKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion

	user          data.LinkName = models.ClassUser
	ontology      data.LinkName = models.ClassOntology
	objects       data.LinkName = models.ClassObjects
	traits        data.LinkName = models.ClassTraits
	relationships data.LinkName = models.ClassRelationships
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Class, error) {
	switch s.Type() {
	case mongo.DBType:
		c := &mongoClass{}
		c.SetID(s.NewID())
		c.ETraits = make(map[string]*models.Trait)
		c.ERelationships = make(map[string]*models.Relationship)
		return c, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}

func Create(s data.Store) (models.Class, error) {
	c, err := New(s)
	if err != nil {
		return nil, err
	}

	return c, s.Save(c)
}
