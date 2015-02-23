package class

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoClass struct {
	models.MongoModel `bson:",inline"`
	models.UserOwned  `bson:",inline"`
	models.Named      `bson:",inline"`

	EOntologyID bson.ObjectId `json:"ontology_id" bson:"ontology_id,inline"`
	ObjectIDs   mongo.IDSet   `json:"object_ids" bson:"object_ids"`

	TraitIDs        mongo.IDSet `json:"trait_ids" bson:"trait_ids"`
	RelationshipIDs mongo.IDSet `json:"relationship_ids" bson:"relationship_ids"`
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

func (c *mongoClass) SetUser(u models.User) error {
	return c.Schema().Link(c, u, User)
}

func (c *mongoClass) SetOntology(o models.Ontology) error {
	return c.Schema().Link(c, o, Ontology)
}

func (c *mongoClass) Ontology(a data.Access, o models.Ontology) error {
	if !data.Compatible(c, o) {
		return data.ErrIncompatibleModels
	}

	if c.CanRead(a.Client()) {
		o.SetID(c.EOntologyID)
		return a.PopulateByID(o)
	} else {
		return data.ErrAccessDenial
	}
}

func (c *mongoClass) IncludeTrait(t models.Trait) error {
	return c.Schema().Link(c, t, Traits)
}

func (c *mongoClass) ExcludeTrait(t models.Trait) error {
	return c.Schema().Unlink(c, t, Traits)
}

func (c *mongoClass) Traits(a data.Access) (data.ModelIterator, error) {
	if c.CanRead(a.Client()) {
		return mongo.NewIDIter(c.TraitIDs, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (c *mongoClass) IncludeRelationship(r models.Relationship) error {
	return c.Schema().Link(c, r, Relationships)
}

func (c *mongoClass) ExcludeRelationship(r models.Relationship) error {
	return c.Schema().Unlink(c, r, Relationships)
}

func (c *mongoClass) Relationships(a data.Access) (data.ModelIterator, error) {
	if c.CanRead(a.Client()) {
		return mongo.NewIDIter(c.RelationshipIDs, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (c *mongoClass) IncludeObject(o models.Object) error {
	return c.Schema().Link(c, o, Objects)
}

func (c *mongoClass) ExcludeObject(o models.Object) error {
	return c.Schema().Unlink(c, o, Objects)
}

func (c *mongoClass) Objects(a data.Access) (data.ModelIterator, error) {
	if c.CanRead(a.Client()) {
		return mongo.NewIDIter(c.ObjectIDs, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (c *mongoClass) Link(m data.Model, l data.Link) error {
	if !data.Compatible(c, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		return c.SetUserID(m.ID())
	default:
		return data.NewLinkError(c, m, l)
	}
	return nil
}

func (c *mongoClass) Unlink(m data.Model, l data.Link) error {

	if !data.Compatible(c, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		c.DropUserID()
	default:
		return data.NewLinkError(c, m, l)
	}
	return nil
}

func (c *mongoClass) HasTrait(a data.Access, name string) bool {
	m, _ := a.ModelFor(models.TraitKind)
	t := m.(models.Trait)
	iter, _ := c.Traits(a)

	for iter.Next(t) {
		if t.Name() == name {
			iter.Close()
			return true
		}
	}

	iter.Close()
	return false
}

func (c *mongoClass) RelationshipWithName(a data.Access, name string) (models.Relationship, error) {
	m, _ := a.ModelFor(models.RelationshipKind)
	r := m.(models.Relationship)

	iter, err := c.Relationships(a)

	for iter.Next(r) {
		if r.Name() == name {
			return r, nil
		}
	}

	err = iter.Close()

	return nil, err
}
