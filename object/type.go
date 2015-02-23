package object

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User     data.LinkName = models.ObjectUser
	Class    data.LinkName = models.ObjectClass
	Ontology data.LinkName = models.ObjectOntology
)

var (
	kind    data.Kind   = models.ObjectKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Object, error) {
	switch s.Type() {
	case mongo.DBType:
		o := &mongoObject{}
		o.SetID(s.NewID())
		o.Traits = make(map[string]string)
		o.Relationships = make(map[string]mongo.IDSet)
		return o, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}
