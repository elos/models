package memory

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

type Class struct {
	space *Space       `json:"-"`
	model models.Class `json:"-"`

	ID            string                          `json:"id"`
	CreatedAt     string                          `json:"created_at"`
	UpdatedAt     string                          `json:"updated_at"`
	UserID        string                          `json:"user_id"`
	OntologyID    string                          `json:"ontology_id"`
	ObjectIDs     string                          `json:"object_ids"`
	Traits        map[string]*models.Trait        `json:"traits"`
	Relationships map[string]*models.Relationship `json:"relationships"`
}

func ClassModel(s *Space, m models.Class) *Class {
	c := &Class{
		space: s,
		model: m,
	}

	data.TransferAttrs(c.model, c)

	s.Register(c)

	return c
}

func NewClass(s *Space) *Class {
	m, _ := s.Access.ModelFor(models.ClassKind)
	return ClassModel(s, m.(models.Class))
}

func (c *Class) Reload() error {
	c.space.Access.PopulateByID(c.model)
	return data.TransferAttrs(c.model, c)
}

func (s *Space) FindClass(id string) *Class {
	m, _ := s.Access.Unmarshal(models.ClassKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return ClassModel(s, m.(models.Class))
}
