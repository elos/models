package memory

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

type Ontology struct {
	space *Space          `json:"-"`
	model models.Ontology `json:"-"`

	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

func OntologyModel(s *Space, m models.Ontology) *Ontology {
	o := &Ontology{
		space: s,
		model: m,
	}

	data.TransferAttrs(o.model, o)

	s.Register(o)

	return o
}

func (o *Ontology) Reload() error {
	o.space.Access.PopulateByID(o.model)
	return data.TransferAttrs(o.model, o)
}
