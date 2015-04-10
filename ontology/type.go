package ontology

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	kind    data.Kind   = models.OntologyKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion

	user    data.LinkName = models.OntologyUser
	classes data.LinkName = models.OntologyClasses
	objects data.LinkName = models.OntologyObjects
)

func NewM(s data.Store) data.Model {
	return New(s)
}

func New(s data.Store) models.Ontology {
	switch s.Type() {
	case mongo.DBType:
		o := &mongoOntology{}
		o.SetID(s.NewID())
		return o
	default:
		panic(data.ErrInvalidDBType)
	}
}

func Create(s data.Store) (models.Ontology, error) {
	o := New(s)
	return o, s.Save(o)
}
