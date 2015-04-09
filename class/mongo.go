package class

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/shared"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoClass struct {
	mongo.Model           `bson:",inline"`
	mongo.Named           `bson:",inline"`
	shared.MongoUserOwned `bson:",inline"`

	EOntologyID bson.ObjectId `json:"ontology_id" bson:"ontology_id,omitempty"`

	ETraits        map[string]*models.Trait        `json:"traits" bson:"traits"`
	ERelationships map[string]*models.Relationship `json:"relationships" bson:"relationships"`

	shared.MongoObjects `bson:",inline"`
}

func (c *mongoClass) Kind() data.Kind {
	return kind
}

func (c *mongoClass) Version() int {
	return version
}

func (c *mongoClass) Schema() data.Schema {
	return schema
}

func (c *mongoClass) Link(m data.Model, l data.Link) error {
	if !data.Compatible(c, m) {
		return data.ErrIncompatibleModels
	}

	id := m.ID().(bson.ObjectId)

	switch l.Name {
	case user:
		return c.SetUserID(id)
	case objects:
		c.MongoObjects.IncludeObjectID(id)
	default:
		return data.NewLinkError(c, m, l)
	}

	return nil
}

func (c *mongoClass) Unlink(m data.Model, l data.Link) error {
	if !data.Compatible(c, m) {
		return data.ErrIncompatibleModels
	}

	id := m.ID().(bson.ObjectId)

	switch l.Name {
	case user:
		c.DropUserID()
	case objects:
		c.MongoObjects.ExcludeObjectID(id)
	default:
		return data.NewLinkError(c, m, l)
	}

	return nil
}

func (c *mongoClass) SetUser(u models.User) error {
	return c.Schema().Link(c, u, user)
}

func (c *mongoClass) SetOntology(o models.Ontology) error {
	return c.Schema().Link(c, o, ontology)
}

func (c *mongoClass) Ontology(a data.Access) (models.Ontology, error) {
	if c.DBType() != a.Type() {
		return nil, data.ErrInvalidDBType
	}

	if mongo.EmptyID(c.EOntologyID) {
		return nil, models.ErrEmptyRelationship
	}

	if !c.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	m, err := a.ModelFor(models.OntologyKind)
	if err != nil {
		return nil, err
	}

	o, ok := m.(models.Ontology)
	if !ok {
		return nil, models.CastError(models.OntologyKind)
	}

	o.SetID(c.EOntologyID)
	return o, a.PopulateByID(o)
}

func (c *mongoClass) IncludeTrait(t *models.Trait) {
	c.ETraits[t.Name] = t
}

func (c *mongoClass) ExcludeTrait(t *models.Trait) {
	delete(c.ETraits, t.Name)
}

func (c *mongoClass) Traits() []*models.Trait {
	ts := make([]*models.Trait, 0)
	for _, val := range c.ETraits {
		ts = append(ts, val)
	}
	return ts
}

func (c *mongoClass) Trait(name string) (*models.Trait, bool) {
	t, ok := c.ETraits[name]
	return t, ok
}

func (c *mongoClass) IncludeRelationship(r *models.Relationship) {
	c.ERelationships[r.Name] = r
}

func (c *mongoClass) ExcludeRelationship(r *models.Relationship) {
	delete(c.ERelationships, r.Name)
}

func (c *mongoClass) Relationships() []*models.Relationship {
	rs := make([]*models.Relationship, 0)
	for _, val := range c.ERelationships {
		rs = append(rs, val)
	}
	return rs
}

func (c *mongoClass) Relationship(name string) (*models.Relationship, bool) {
	r, ok := c.ERelationships[name]
	return r, ok
}

func (c *mongoClass) IncludeObject(obj models.Object) error {
	return c.Schema().Link(c, obj, objects)
}

func (c *mongoClass) ExcludeObject(obj models.Object) error {
	return c.Schema().Unlink(c, obj, objects)
}

func (c *mongoClass) ObjectsIter(a data.Access) (data.ModelIterator, error) {
	if !c.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	return c.MongoObjects.ObjectsIter(a), nil
}

func (c *mongoClass) Objects(a data.Access) ([]models.Object, error) {
	if !c.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	return c.MongoObjects.Objects(a)
}

func (c *mongoClass) NewObject(a data.Access) (models.Object, error) {
	return NewObject(c, a)
}
