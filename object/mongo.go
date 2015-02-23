package object

import (
	"errors"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoObject struct {
	models.MongoModel `bson:,inline"`
	models.UserOwned  `bson:",inline"`
	models.Named      `bson:",inline"`

	EClassID      bson.ObjectId `json:"class_id" bson:"class_id,omitempty"`
	EOntologyID   bson.ObjectId `json:"ontology_id" bson:"ontology_id,omitempty"`
	Traits        map[string]string
	Relationships map[string]mongo.IDSet
}

func (t *mongoObject) Kind() data.Kind {
	return kind
}

func (t *mongoObject) Version() int {
	return version
}

func (t *mongoObject) Schema() data.Schema {
	return schema
}

func (t *mongoObject) SetUser(u models.User) error {
	return t.Schema().Link(t, u, User)
}

func (t *mongoObject) SetClass(c models.Class) error {
	return t.Schema().Link(t, c, Class)
}

func (t *mongoObject) Class(a data.Access) (models.Class, error) {
	m, _ := a.ModelFor(models.ClassKind)
	c := m.(models.Class)

	if t.CanRead(a.Client()) {
		c.SetID(t.EClassID)
		err := a.PopulateByID(c)
		return c, err
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (t *mongoObject) SetOntology(o models.Ontology) error {
	return t.Schema().Link(t, o, Ontology)
}

func (t *mongoObject) Ontology(a data.Access, o models.Ontology) error {
	if !data.Compatible(t, o) {
		return data.ErrIncompatibleModels
	}

	if t.CanRead(a.Client()) {
		o.SetID(t.EClassID)
		return a.PopulateByID(o)
	} else {
		return data.ErrAccessDenial
	}
}

func (t *mongoObject) Link(m data.Model, l data.Link) error {
	if !data.Compatible(t, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		t.DropUserID()
	case Class:
		t.EClassID = m.ID().(bson.ObjectId)
	case Ontology:
		t.EOntologyID = m.ID().(bson.ObjectId)
	default:
		return data.NewLinkError(t, m, l)
	}
	return nil
}

func (t *mongoObject) Unlink(m data.Model, l data.Link) error {
	if !data.Compatible(t, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		t.DropUserID()
	case Class:
		t.EClassID = *new(bson.ObjectId)
	case Ontology:
		t.EOntologyID = *new(bson.ObjectId)
	default:
		return data.NewLinkError(t, m, l)
	}
	return nil
}

func (o *mongoObject) SetTrait(a data.Access, name string, value string) error {
	c, _ := o.Class(a)

	if c.HasTrait(a, name) {
		o.Traits[name] = value
		return nil
	} else {
		return errors.New("Invalid trait name")
	}
}

func (o *mongoObject) AddRelationship(a data.Access, name string, other models.Object) error {
	c, _ := o.Class(a)

	r, err := c.RelationshipWithName(a, name)
	if err != nil {
		return err
	}

	if r == nil {
		return errors.New("Invalid relationship name")
	}

	if r.Other() != other.Name() {
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

	r, err := c.RelationshipWithName(a, name)

	if err != nil {
		return err
	}

	if r == nil {
		return errors.New("Invalid relationship name")
	}

	if r.Other() != other.Name() {
		return errors.New("Invalid other kind")
	}

	ids, ok := o.Relationships[name]

	if !ok {
		return nil // nothing to remove
	}

	o.Relationships[name] = mongo.DropID(ids, other.ID().(bson.ObjectId))
	return nil
}
