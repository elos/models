package memory

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
	m, _ := s.Access.ModelFor(models.ObjectKind)
	return ObjectModel(s, m.(models.Object))
}

func (o *Object) Reload() error {
	o.space.Access.PopulateByID(o.model)
	return data.TransferAttrs(o.model, o)
}

func (s *Space) FindObject(id string) *Object {
	m, _ := s.Access.Unmarshal(models.ObjectKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return ObjectModel(s, m.(models.Object))
}
