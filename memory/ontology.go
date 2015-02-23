package memory

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

type Ontology struct {
	space *Space          `json:"-"`
	model models.Ontology `json:"-"`

	ID        string   `json:"id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	UserID    string   `json:"user_id"`
	ClassIDs  []string `json:"class_ids"`
	ObjectIDs []string `json:"object_ids"`
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

func NewOntology(s *Space) *Ontology {
	m, _ := s.Access.ModelFor(models.OntologyKind)
	return OntologyModel(s, m.(models.Ontology))
}

func (o *Ontology) Reload() error {
	o.space.Access.PopulateByID(o.model)
	return data.TransferAttrs(o.model, o)
}

func (s *Space) FindOntology(id string) *Ontology {
	m, _ := s.Access.Unmarshal(models.OntologyKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return OntologyModel(s, m.(models.Ontology))
}
