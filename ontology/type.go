package ontology

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	User    data.LinkName = models.OntologyUser
	Classes data.LinkName = models.OntologyClasses
	Objects data.LinkName = models.OntologyObjects
)

var (
	kind    data.Kind   = models.OntologyKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Ontology, error) {
	switch s.Type() {
	case mongo.DBType:
		o := &mongoOntology{}
		o.SetID(s.NewID())
		return o, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}

func Create(s data.Store) (models.Ontology, error) {
	o, err := New(s)
	if err != nil {
		return o, err
	}

	return o, s.Save(o)
}
