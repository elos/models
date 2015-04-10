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

func (o *mongoOntology) ClassesIter(store models.Store) (data.ModelIterator, error) {
	if !store.Compatible(o) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(o.ClassIDs, store), nil
}

func (o *mongoOntology) Classes(store models.Store) ([]models.Class, error) {
	if !store.Compatible(o) {
		return nil, data.ErrInvalidDBType
	}

	classes := make([]models.Class, 0)
	iter := mongo.NewIDIter(o.ClassIDs, store)
	class := store.Class()
	for iter.Next(class) {
		classes = append(classes, class)
		class = store.Class()
	}

	return classes, iter.Close()
}

func (o *mongoOntology) IncludeObject(obj models.Object) error {
	return o.Schema().Link(o, obj, objects)
}

func (o *mongoOntology) ExcludeObject(obj models.Object) error {
	return o.Schema().Unlink(o, obj, objects)
}

func (o *mongoOntology) ObjectsIter(store models.Store) (data.ModelIterator, error) {
	if !store.Compatible(o) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(o.ObjectIDs, store), nil
}

func (o *mongoOntology) Objects(store models.Store) ([]models.Object, error) {
	if !store.Compatible(o) {
		return nil, data.ErrInvalidDBType
	}

	objects := make([]models.Object, 0)
	iter := mongo.NewIDIter(o.ObjectIDs, store)
	object := store.Object()
	for iter.Next(object) {
		objects = append(objects, object)
		object = store.Object()
	}

	return objects, iter.Close()
}
