package interactive

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
	ObjectIDs     []string                        `json:"object_ids"`
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

func (c *Class) Reload() error {
	c.space.Access.PopulateByID(c.model)
	return data.TransferAttrs(c.model, c)
}

func (c *Class) NewObject() *Object {
	o, _ := c.model.NewObject(c.space.Access)
	return ObjectModel(c.space, o)
}

func (c *Class) AddTrait(name string, tipe string) {
	t := &models.Trait{
		Name: name,
		Type: tipe,
	}

	c.model.IncludeTrait(t)
	c.space.Access.Save(c.model)
	data.TransferAttrs(c.model, c)
}

func (c *Class) AddRelationship(name string, other string, inverse string) {
	r := &models.Relationship{
		Name:    name,
		Other:   other,
		Inverse: inverse,
	}

	c.model.IncludeRelationship(r)
	c.space.Access.Save(c.model)
	data.TransferAttrs(c.model, c)
}

func (c *Class) Ontology() *Ontology {
	o, _ := c.model.Ontology(c.space.Access)
	return OntologyModel(c.space, o)
}

func (c *Class) Objects() []*Object {
	data.TransferAttrs(c.model, c)
	objects := make([]*Object, 0)
	for _, id := range c.ObjectIDs {
		objects = append(objects, c.space.FindObject(id))
	}
	return objects
}

func NewClass(s *Space) *Class {
	m, _ := s.Access.ModelFor(models.ClassKind)
	return ClassModel(s, m.(models.Class))
}

func (s *Space) FindClass(id string) *Class {
	m, _ := s.Access.Unmarshal(models.ClassKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return ClassModel(s, m.(models.Class))
}
