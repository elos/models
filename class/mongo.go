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

func (c *mongoClass) Link(m data.Model, l data.Link) error {
	if !data.Compatible(c, m) {
		return data.ErrIncompatibleModels
	}

	id := m.ID().(bson.ObjectId)

	switch l.Name {
	case user:
		return c.SetUserID(id)
	case objects:
		c.ObjectIDs = mongo.AddID(c.ObjectIDs, id)
	case traits:
		c.TraitIDs = mongo.AddID(c.TraitIDs, id)
	case relationships:
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
	case user:
		c.DropUserID()
	case objects:
		c.ObjectIDs = mongo.DropID(c.ObjectIDs, id)
	case traits:
		c.TraitIDs = mongo.DropID(c.TraitIDs, id)
	case relationships:
		c.RelationshipIDs = mongo.DropID(c.RelationshipIDs, id)
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

func (c *mongoClass) Trait(name string) (*models.Trait, bool) {
	t, ok := c.ETraits[name]
	return t, ok
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

	return mongo.NewIDIter(c.ObjectIDs, a), nil
}

func (c *mongoClass) Objects(a data.Access) ([]models.Object, error) {
	if !c.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	objects := make([]models.Object, 0)
	iter, err := c.ObjectsIter(a)
	if err != nil {
		return objects, err
	}

	m, err := a.ModelFor(models.ObjectKind)
	if err != nil {
		return objects, err
	}

	for iter.Next(m) {
		object, ok := m.(models.Object)
		if !ok {
			return objects, models.CastError(models.ObjectKind)
		}

		objects = append(objects, object)

		m, err = a.ModelFor(models.ObjectKind)
		if err != nil {
			return objects, err
		}
	}

	return objects, err
}

func (c *mongoClass) NewObject(a data.Access) (models.Object, error) {
	m, err := a.ModelFor(models.ObjectKind)
	if err != nil {
		return nil, err
	}
	obj, ok := m.(models.Object)
	if !ok {
		return nil, models.CastError(models.ObjectKind)
	}

	ont, err := c.Ontology(a)
	if err != nil {
		return nil, err
	}

	obj.SetOntology(ont)
	obj.SetClass(c)
	obj.SetName(c.Name())

	return obj, nil
}
