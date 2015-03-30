package shared

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoObjects struct {
	ObjectIDs mongo.IDSet `json:"object_ids" bson:"object_ids"`
}

func (c *MongoObjects) IncludeObjectID(id bson.ObjectId) {
	c.ObjectIDs = mongo.AddID(c.ObjectIDs, id)
}

func (c *MongoObjects) ExcludeObjectID(id bson.ObjectId) {
	c.ObjectIDs = mongo.DropID(c.ObjectIDs, id)
}

func (c *MongoObjects) IncludeObject(obj models.Object) error {
	id, ok := obj.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	c.IncludeObjectID(id)
	return nil
}

func (c *MongoObjects) ExcludeObject(obj models.Object) error {
	id, ok := obj.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	c.ExcludeObjectID(id)
	return nil
}

func (c *MongoObjects) ObjectsIter(a data.Access) data.ModelIterator {
	return mongo.NewIDIter(c.ObjectIDs, a)
}

func (c *MongoObjects) Objects(a data.Access) ([]models.Object, error) {
	objects := make([]models.Object, 0)
	iter := c.ObjectsIter(a)

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
