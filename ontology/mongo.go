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

func (o *mongoOntology) Link(m data.Model, l data.Link) error {
	if !data.Compatible(o, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case user:
		return o.SetUserID(m.ID())
	case classes:
		o.ClassIDs = mongo.AddID(o.ClassIDs, m.ID().(bson.ObjectId))
	case objects:
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
	case user:
		o.DropUserID()
	case classes:
		o.ClassIDs = mongo.DropID(o.ClassIDs, m.ID().(bson.ObjectId))
	case objects:
		o.ObjectIDs = mongo.DropID(o.ObjectIDs, m.ID().(bson.ObjectId))
	default:
		return data.NewLinkError(o, m, l)
	}

	return nil
}

func (o *mongoOntology) SetUser(u models.User) error {
	return o.Schema().Link(o, u, user)
}

func (o *mongoOntology) IncludeClass(c models.Class) error {
	return o.Schema().Link(o, c, classes)
}

func (o *mongoOntology) ExcludeClass(c models.Class) error {
	return o.Schema().Unlink(o, c, classes)
}

func (o *mongoOntology) ClassesIter(a data.Access) (data.ModelIterator, error) {
	if !o.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	return mongo.NewIDIter(o.ClassIDs, a), nil
}

func (o *mongoOntology) Classes(a data.Access) ([]models.Class, error) {
	if !o.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	classes := make([]models.Class, 0)
	iter, err := o.ClassesIter(a)
	if err != nil {
		return classes, err
	}

	m, err := a.ModelFor(models.ClassKind)
	if err != nil {
		return classes, err
	}

	for iter.Next(m) {
		class, ok := m.(models.Class)
		if !ok {
			return classes, models.CastError(models.ClassKind)
		}

		classes = append(classes, class)

		m, err = a.ModelFor(models.ClassKind)
		if err != nil {
			return classes, err
		}
	}

	return classes, err
}

func (o *mongoOntology) IncludeObject(obj models.Object) error {
	return o.Schema().Link(o, obj, objects)
}

func (o *mongoOntology) ExcludeObject(obj models.Object) error {
	return o.Schema().Unlink(o, obj, objects)
}

func (o *mongoOntology) ObjectsIter(a data.Access) (data.ModelIterator, error) {
	if !o.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	return mongo.NewIDIter(o.ObjectIDs, a), nil
}

func (o *mongoOntology) Objects(a data.Access) ([]models.Object, error) {
	if !o.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	objects := make([]models.Object, 0)
	iter, err := o.ObjectsIter(a)
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
