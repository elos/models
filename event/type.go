package event

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

var (
	User data.LinkName = models.EventUser
)

var (
	kind    data.Kind   = models.EventKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) data.Model {
	return New(s)
}

func New(s data.Store) models.Event {
	switch s.Type() {
	case mongo.DBType:
		e := &mongoEvent{}
		e.SetID(s.NewID())
		return e
	default:
		panic(data.ErrInvalidDBType)
	}
}

func Create(s data.Store) (models.Event, error) {
	e := New(s)
	return e, s.Save(e)
}

func CreateAttrs(s data.Store, a data.AttrMap) (models.Event, error) {
	event := New(s)

	id, present := a["id"]
	id, valid := id.(data.ID)
	if present && valid {
		if err := event.SetID(id.(data.ID)); err != nil {
			return event, err
		}
	} else {
		if err := event.SetID(s.NewID()); err != nil {
			return event, err
		}
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
	event := New(s)

	id, ok := id.(bson.ObjectId)
	if !ok {
		return event, data.ErrInvalidID
	}

	event.SetID(id)
	return event, s.PopulateByID(event)
}

func FindBy(s data.Store, field string, value interface{}) (models.Event, error) {
	event := New(s)
	return event, s.PopulateByField(field, value, event)
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
		if !e.(*mongoEvent).UserID().Valid() {
			return false, data.NewAttrError("user", "be set and valid")
		}
	}

	return true, nil
}
