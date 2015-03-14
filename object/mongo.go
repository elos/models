package object

import (
	"errors"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/shared"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoObject struct {
	mongo.Model           `bson:",inline"`
	mongo.Named           `bson:",inline"`
	shared.MongoUserOwned `bson:",inline"`

	EClassID      bson.ObjectId          `json:"class_id" bson:"class_id,omitempty"`
	EOntologyID   bson.ObjectId          `json:"ontology_id" bson:"ontology_id,omitempty"`
	Traits        map[string]string      `json:"traits", bson:"traits"`
	Relationships map[string]mongo.IDSet `json:"relationships" bson:"relationships"`
}

func (o *mongoObject) Kind() data.Kind {
	return kind
}

func (o *mongoObject) Version() int {
	return version
}

func (o *mongoObject) Schema() data.Schema {
	return schema
}

func (o *mongoObject) SetUser(u models.User) error {
	return o.Schema().Link(o, u, User)
}

func (o *mongoObject) SetClass(c models.Class) error {
	return o.Schema().Link(o, c, Class)
}

func (o *mongoObject) Class(a data.Access) (models.Class, error) {
	m, _ := a.ModelFor(models.ClassKind)
	c := m.(models.Class)

	if o.CanRead(a.Client()) {
		c.SetID(o.EClassID)
		err := a.PopulateByID(c)
		return c, err
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (o *mongoObject) SetOntology(ont models.Ontology) error {
	return o.Schema().Link(o, ont, Ontology)
}

func (o *mongoObject) Link(m data.Model, l data.Link) error {
	if !data.Compatible(o, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		o.DropUserID()
	case Class:
		o.EClassID = m.ID().(bson.ObjectId)
	case Ontology:
		o.EOntologyID = m.ID().(bson.ObjectId)
	default:
		return data.NewLinkError(o, m, l)
	}
	return nil
}

func (o *mongoObject) Unlink(m data.Model, l data.Link) error {
	if !data.Compatible(o, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		o.DropUserID()
	case Class:
		o.EClassID = *new(bson.ObjectId)
	case Ontology:
		o.EOntologyID = *new(bson.ObjectId)
	default:
		return data.NewLinkError(o, m, l)
	}
	return nil
}

func (o *mongoObject) SetTrait(a data.Access, name string, value string) error {
	c, _ := o.Class(a)

	_, ok := c.Trait(name)
	if ok {
		o.Traits[name] = value
		return nil
	} else {
		return errors.New("Invalid trait name")
	}
}

func (o *mongoObject) AddRelationship(a data.Access, name string, other models.Object) error {
	c, _ := o.Class(a)

	r, ok := c.Relationship(name)
	if !ok || r == nil {
		return errors.New("Invalid relationship name")
	}

	if r.Other != other.Name() {
		return errors.New("Invalid other kind")
	}

	ids, ok := o.Relationships[name]
	if !ok {
		ids = make(mongo.IDSet, 0)
	}

	o.Relationships[name] = mongo.AddID(ids, other.ID().(bson.ObjectId))

	return nil
}

func (o *mongoObject) DropRelationship(a data.Access, name string, other models.Object) error {
	c, _ := o.Class(a)

	r, ok := c.Relationship(name)

	if !ok || r == nil {
		return errors.New("Invalid relationship name")
	}

	if r.Other != other.Name() {
		return errors.New("Invalid other kind")
	}

	ids, ok := o.Relationships[name]

	if !ok {
		return nil // nothing to remove
	}

	o.Relationships[name] = mongo.DropID(ids, other.ID().(bson.ObjectId))
	return nil
}

func (o *mongoObject) Ontology(a data.Access) (models.Ontology, error) {
	m, _ := a.ModelFor(models.OntologyKind)
	ontology := m.(models.Ontology)

	if o.CanRead(a.Client()) {
		ontology.SetID(o.EOntologyID)
		err := a.PopulateByID(ontology)
		return ontology, err
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (o *mongoObject) SetOntologyID(id data.ID) error {
	bid, ok := id.(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	o.EOntologyID = bid
	return nil
}
