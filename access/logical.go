package access

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

var ImmutableRecords = map[data.Kind]bool{
	models.ContextKind:    true,
	models.CredentialKind: true,
	models.GroupKind:      true,
	models.SessionKind:    true,
}

type Level int

const (
	None Level = iota
	Read
	Write
)

type Context interface {
	data.Record
	Contains(data.Record) bool
}

type Group interface {
	data.Record
	Contexts(data.DB) ([]Context, error)
	Access() int
}

type User interface {
	data.Record
	Groups(data.DB) ([]Group, error)
}

type ModelProperty interface {
	data.Record
	Owner(data.DB) (*models.User, error)
}

type Property interface {
	data.Record
	Owner(data.DB) (User, error)
}

// --- Context Implementation {{{
type context struct {
	*models.Context
}

func WrapContext(c *models.Context) Context {
	return &context{Context: c}
}

// --- }}}

// --- Group Implementation --- {{{

type group struct {
	*models.Group
}

func WrapGroup(g *models.Group) Group {
	return &group{Group: g}
}

func (g *group) Contexts(db data.DB) ([]Context, error) {
	contexts, err := g.Group.Contexts(db)

	if err != nil {
		return nil, err
	}

	c := make([]Context, len(contexts))

	for i := range contexts {
		c[i] = WrapContext(contexts[i])
	}

	return c, nil
}

func (g *group) Access() int {
	return g.Group.Access
}

// --- }}}

// --- User Implementation {{{

type user struct {
	*models.User
}

func WrapUser(u *models.User) User {
	return &user{User: u}
}

func (u *user) Groups(db data.DB) ([]Group, error) {
	mgs, err := u.User.Groups(db)

	if err != nil {
		return nil, err
	}

	g := make([]Group, len(mgs))
	for i := range mgs {
		g[i] = WrapGroup(mgs[i])
	}

	return g, nil
}

// --- }}}

// --- Property Implementation {{{

type property struct {
	ModelProperty
}

func WrapProperty(p ModelProperty) Property {
	return &property{ModelProperty: p}
}

func (p *property) Owner(db data.DB) (User, error) {
	u, err := p.ModelProperty.Owner(db)
	if err != nil {
		return nil, err
	}

	return WrapUser(u), nil
}

// --- }}}
