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

	EOntologyID     bson.ObjectId `json:"ontology_id" bson:"ontology_id,omitempty"`
	ObjectIDs       mongo.IDSet   `json:"object_ids" bson:"object_ids"`
	TraitIDs        mongo.IDSet   `json:"trait_ids" bson:"trait_ids"`
	RelationshipIDs mongo.IDSet   `json:"relationship_ids" bson:"relationship_ids"`

	ETraits        map[string]*models.Trait        `json:"traits" bson:"traits"`
	ERelationships map[string]*models.Relationship `json:"relationships" bson:"relationships"`
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

func (c *mongoClass) Ontology(a data.Access) (models.Ontology, error) {
	m, _ := a.ModelFor(models.OntologyKind)
	o := m.(models.Ontology)

	if c.CanRead(a.Client()) {
		o.SetID(c.EOntologyID)
		err := a.PopulateByID(o)
		return o, err
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (c *mongoClass) IncludeTrait(t *models.Trait) error {
	c.ETraits[t.Name] = t
	return nil
}

func (c *mongoClass) ExcludeTrait(t *models.Trait) error {
	delete(c.ETraits, t.Name)
	return nil
}

func (c *mongoClass) Traits() []*models.Trait {
	ts := make([]*models.Trait, 0)
	for _, val := range c.ETraits {
		ts = append(ts, val)
	}
	return ts
}

func (c *mongoClass) IncludeRelationship(r *models.Relationship) error {
	c.ERelationships[r.Name] = r
	return nil
}

func (c *mongoClass) ExcludeRelationship(r *models.Relationship) error {
	delete(c.ERelationships, r.Name)
	return nil
}

func (c *mongoClass) Relationships() []*models.Relationship {
	rs := make([]*models.Relationship, 0)
	for _, val := range c.ERelationships {
		rs = append(rs, val)
	}
	return rs
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

	id := m.ID().(bson.ObjectId)

	switch l.Name {
	case User:
		return c.SetUserID(id)
	case Objects:
		c.ObjectIDs = mongo.AddID(c.ObjectIDs, id)
	case Traits:
		c.TraitIDs = mongo.AddID(c.TraitIDs, id)
	case Relationships:
		c.RelationshipIDs = mongo.AddID(c.RelationshipIDs, id)
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
	case User:
		c.DropUserID()
	case Objects:
		c.ObjectIDs = mongo.DropID(c.ObjectIDs, id)
	case Traits:
		c.TraitIDs = mongo.DropID(c.TraitIDs, id)
	case Relationships:
		c.RelationshipIDs = mongo.DropID(c.RelationshipIDs, id)
	default:
		return data.NewLinkError(c, m, l)
	}

	return nil
}

func (c *mongoClass) Relationship(name string) (*models.Relationship, bool) {
	r, ok := c.ERelationships[name]
	return r, ok
}

func (c *mongoClass) Trait(name string) (*models.Trait, bool) {
	t, ok := c.ETraits[name]
	return t, ok
}

func (c *mongoClass) NewObject(a data.Access) models.Object {
	m, _ := a.ModelFor(models.ObjectKind)
	obj := m.(models.Object)

	m, _ = a.Unmarshal(models.OntologyKind, data.AttrMap{
		"id": c.EOntologyID,
	})
	ont := m.(models.Ontology)

	obj.SetOntology(ont)
	obj.SetClass(c)
	obj.SetName(c.Name())

	return obj
}
