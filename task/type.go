package task

import (
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

var (
	User         data.LinkName
	Dependencies data.LinkName
)

var (
	kind    data.Kind
	schema  data.Schema
	version int
)

func Setup(s data.Schema, k data.Kind, v int) {
	schema, kind, version = s, k, v
}

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

func New(s data.Store) (models.Task, error) {
	switch s.Type() {
	case mongo.DBType:
		return &mongoTask{}, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}

func Create(s data.Store, a data.AttrMap) (models.Task, error) {
	task, err := New(s)
	if err != nil {
		return task, err
	}

	switch s.Type() {
	case mongo.DBType:
		if id, ok := a["id"].(bson.ObjectId); ok {
			task.SetID(id)
		} else {
			task.SetID(mongo.NewObjectID().(bson.ObjectId))
		}

	default:
		return task, data.ErrInvalidDBType
	}

	return task, nil
}
