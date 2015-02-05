package event

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

var (
	User data.LinkName
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

func New(s data.Store) (models.Event, error) {
	switch s.Type() {
	case mongo.DBType:
		return &mongoEvent{}, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}

func Create(s data.Store, a data.AttrMap) (models.Event, error) {
	event, err := New(s)
	if err != nil {
		return event, err
	}

	switch s.Type() {
	case mongo.DBType:
		if id, ok := a["id"].(bson.ObjectId); ok {
			event.SetID(id)
		} else {
			event.SetID(mongo.NewObjectID().(bson.ObjectId))
		}

	default:
		return event, data.ErrInvalidDBType
	}

	if ca, ok := a["created_at"].(time.Time); ok {
		event.SetCreatedAt(ca)
	} else {
		event.SetCreatedAt(time.Now())
	}

	if n, ok := a["name"].(string); ok {
		event.SetName(n)
	}

	// Try linking to user?

	if err := s.Save(event); err != nil {
		return nil, err
	} else {
		return event, nil
	}
}

func Find(s data.Store, id data.ID) (models.Event, error) {
	event, err := New(s)
	if err != nil {
		return event, err
	}

	id, ok := id.(bson.ObjectId)
	if !ok {
		return event, data.ErrInvalidID
	}

	event.SetID(id)

	if err := s.PopulateByID(event); err != nil {
		return event, err
	}

	return event, nil
}

func FindBy(s data.Store, field string, value interface{}) (models.Event, error) {
	event, err := New(s)
	if err != nil {
		return event, err
	}

	if err = s.PopulateByField(field, value, event); err != nil {
		return event, err
	}

	return event, nil
}

func Validate(e models.Event) (bool, error) {

	if e.Name() == "" {
		return false, data.NewAttrError("name", "be defined")
	}

	if e.StartTime().IsZero() {
		return false, data.NewAttrError("start_time", "be non-zero")
	}

	if e.EndTime().IsZero() {
		return false, data.NewAttrError("end_time", "be non-zero")
	}

	switch e.(type) {
	case *mongoEvent:
		if !e.(*mongoEvent).UserID.Valid() {
			return false, data.NewAttrError("user", "be set and valid")
		}
	}

	return true, nil
}
