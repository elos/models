package interactive

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

type Object struct {
	space *Space        `json:"-"`
	model models.Object `json:"-"`

	ID            string              `json:"id"`
	CreatedAt     string              `json:"created_at"`
	UpdatedAt     string              `json:"updated_at"`
	UserID        string              `json:"user_id"`
	Name          string              `json:"name"`
	OntologyID    string              `json:"ontology_id"`
	ClassID       string              `json:"class_id"`
	Traits        map[string]string   `json:"traits"`
	Relationships map[string][]string `json:"relationships"`
}

func ObjectModel(s *Space, m models.Object) *Object {
	c := &Object{
		space: s,
		model: m,
	}

	data.TransferAttrs(c.model, c)

	s.Register(c)

	return c
}

func NewObject(s *Space) *Object {
	m, _ := s.Store.ModelFor(models.ObjectKind)
	return ObjectModel(s, m.(models.Object))
}

func (o *Object) Reload() error {
	o.space.Store.PopulateByID(o.model)
	return data.TransferAttrs(o.model, o)
}

func (o *Object) Ontology() *Ontology {
	ontology, _ := o.model.Ontology(o.space.Store.(models.Store))
	return OntologyModel(o.space, ontology)
}

func (o *Object) Class() *Class {
	class, _ := o.model.Class(o.space.Store.(models.Store))
	return ClassModel(o.space, class)
}

func (s *Space) FindObject(id string) *Object {
	m, _ := s.Store.Unmarshal(models.ObjectKind, data.AttrMap{
		"id": id,
	})
	s.Store.PopulateByID(m)
	return ObjectModel(s, m.(models.Object))
}
