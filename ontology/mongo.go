package ontology

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/shared"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoOntology struct {
	mongo.Model           `bson:",inline"`
	shared.MongoUserOwned `bson:",inline"`

	ClassIDs  mongo.IDSet `json:"class_ids" bson:"class_ids"`
	ObjectIDs mongo.IDSet `json:"object_ids" bson:"object_ids"`
}

func (o *mongoOntology) Kind() data.Kind {
	return kind
}

func (o *mongoOntology) Version() int {
	return version
}

func (o *mongoOntology) Schema() data.Schema {
	return schema
}

func (o *mongoOntology) SetUser(u models.User) error {
	return o.Schema().Link(o, u, User)
}

func (o *mongoOntology) IncludeClass(c models.Class) error {
	return o.Schema().Link(o, c, Classes)
}

func (o *mongoOntology) ExcludeClass(c models.Class) error {
	return o.Schema().Unlink(o, c, Classes)
}

func (o *mongoOntology) IncludeObject(obj models.Object) error {
	return o.Schema().Link(o, obj, Objects)
}

func (o *mongoOntology) ExcludeObject(obj models.Object) error {
	return o.Schema().Unlink(o, obj, Objects)
}

func (o *mongoOntology) Classes(a data.Access) (data.ModelIterator, error) {
	if o.CanRead(a.Client()) {
		return mongo.NewIDIter(o.ClassIDs, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (o *mongoOntology) Objects(a data.Access) (data.ModelIterator, error) {
	if o.CanRead(a.Client()) {
		return mongo.NewIDIter(o.ObjectIDs, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (o *mongoOntology) Link(m data.Model, l data.Link) error {
	if !data.Compatible(o, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		return o.SetUserID(m.ID())
	case Classes:
		o.ClassIDs = mongo.AddID(o.ClassIDs, m.ID().(bson.ObjectId))
	case Objects:
		o.ObjectIDs = mongo.AddID(o.ObjectIDs, m.ID().(bson.ObjectId))
	default:
		return data.NewLinkError(o, m, l)
	}
	return nil
}

func (o *mongoOntology) Unlink(m data.Model, l data.Link) error {
	if !data.Compatible(o, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		o.DropUserID()
	case Classes:
		o.ClassIDs = mongo.DropID(o.ClassIDs, m.ID().(bson.ObjectId))
	case Objects:
		o.ObjectIDs = mongo.DropID(o.ObjectIDs, m.ID().(bson.ObjectId))
	default:
		return data.NewLinkError(o, m, l)
	}
	return nil
}
