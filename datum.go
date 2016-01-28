package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Datum struct {
	Context   string    `json:"context" bson:"context"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	EventId   string    `json:"event_id" bson:"event_id"`
	Id        string    `json:"id" bson:"_id,omitempty"`
	OwnerId   string    `json:"owner_id" bson:"owner_id"`
	PersonId  string    `json:"person_id" bson:"person_id"`
	Tags      []string  `json:"tags" bson:"tags"`
	Unit      string    `json:"unit" bson:"unit"`
	Value     float64   `json:"value" bson:"value"`
}

func NewDatum() *Datum {
	return &Datum{}
}

func FindDatum(db data.DB, id data.ID) (*Datum, error) {

	datum := NewDatum()
	datum.SetID(id)

	return datum, db.PopulateByID(datum)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (datum *Datum) Kind() data.Kind {
	return DatumKind
}

// just returns itself for now
func (datum *Datum) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = datum.ID()
	return foo
}

func (datum *Datum) SetID(id data.ID) {
	datum.Id = id.String()
}

func (datum *Datum) ID() data.ID {
	return data.ID(datum.Id)
}

func (datum *Datum) SetEvent(eventArgument *Event) error {
	datum.EventId = eventArgument.ID().String()
	return nil
}

func (datum *Datum) Event(db data.DB) (*Event, error) {
	if datum.EventId == "" {
		return nil, ErrEmptyLink
	}

	eventArgument := NewEvent()
	id, _ := db.ParseID(datum.EventId)
	eventArgument.SetID(id)
	return eventArgument, db.PopulateByID(eventArgument)

}

func (datum *Datum) EventOrCreate(db data.DB) (*Event, error) {
	event, err := datum.Event(db)

	if err == ErrEmptyLink {
		event := NewEvent()
		event.SetID(db.NewID())
		if err := datum.SetEvent(event); err != nil {
			return nil, err
		}

		if err := db.Save(event); err != nil {
			return nil, err
		}

		if err := db.Save(datum); err != nil {
			return nil, err
		}

		return event, nil
	} else {
		return event, err
	}
}

func (datum *Datum) SetOwner(userArgument *User) error {
	datum.OwnerId = userArgument.ID().String()
	return nil
}

func (datum *Datum) Owner(db data.DB) (*User, error) {
	if datum.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(datum.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (datum *Datum) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := datum.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := datum.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(datum); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

func (datum *Datum) SetPerson(personArgument *Person) error {
	datum.PersonId = personArgument.ID().String()
	return nil
}

func (datum *Datum) Person(db data.DB) (*Person, error) {
	if datum.PersonId == "" {
		return nil, ErrEmptyLink
	}

	personArgument := NewPerson()
	id, _ := db.ParseID(datum.PersonId)
	personArgument.SetID(id)
	return personArgument, db.PopulateByID(personArgument)

}

func (datum *Datum) PersonOrCreate(db data.DB) (*Person, error) {
	person, err := datum.Person(db)

	if err == ErrEmptyLink {
		person := NewPerson()
		person.SetID(db.NewID())
		if err := datum.SetPerson(person); err != nil {
			return nil, err
		}

		if err := db.Save(person); err != nil {
			return nil, err
		}

		if err := db.Save(datum); err != nil {
			return nil, err
		}

		return person, nil
	} else {
		return person, err
	}
}

// BSON {{{
func (datum *Datum) GetBSON() (interface{}, error) {

	return struct {
		Context string `json:"context" bson:"context"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Tags []string `json:"tags" bson:"tags"`

		Unit string `json:"unit" bson:"unit"`

		Value float64 `json:"value" bson:"value"`

		EventId string `json:"event_id" bson:"event_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PersonId string `json:"person_id" bson:"person_id"`
	}{

		Context: datum.Context,

		CreatedAt: datum.CreatedAt,

		Tags: datum.Tags,

		Unit: datum.Unit,

		Value: datum.Value,

		EventId: datum.EventId,

		OwnerId: datum.OwnerId,

		PersonId: datum.PersonId,
	}, nil

}

func (datum *Datum) SetBSON(raw bson.Raw) error {

	tmp := struct {
		Context string `json:"context" bson:"context"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Tags []string `json:"tags" bson:"tags"`

		Unit string `json:"unit" bson:"unit"`

		Value float64 `json:"value" bson:"value"`

		EventId string `json:"event_id" bson:"event_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PersonId string `json:"person_id" bson:"person_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	datum.Context = tmp.Context

	datum.CreatedAt = tmp.CreatedAt

	datum.Id = tmp.Id.Hex()

	datum.Tags = tmp.Tags

	datum.Unit = tmp.Unit

	datum.Value = tmp.Value

	datum.EventId = tmp.EventId

	datum.OwnerId = tmp.OwnerId

	datum.PersonId = tmp.PersonId

	return nil

}

// BSON }}}

func (datum *Datum) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		datum.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		datum.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["value"]; ok {
		datum.Value = val.(float64)
	}

	if val, ok := structure["unit"]; ok {
		datum.Unit = val.(string)
	}

	if val, ok := structure["tags"]; ok {
		datum.Tags = val.([]string)
	}

	if val, ok := structure["context"]; ok {
		datum.Context = val.(string)
	}

	if val, ok := structure["owner_id"]; ok {
		datum.OwnerId = val.(string)
	}

	if val, ok := structure["person_id"]; ok {
		datum.PersonId = val.(string)
	}

	if val, ok := structure["event_id"]; ok {
		datum.EventId = val.(string)
	}

}

var DatumStructure = map[string]metis.Primitive{

	"tags": 7,

	"context": 3,

	"id": 9,

	"created_at": 4,

	"value": 2,

	"unit": 3,

	"owner_id": 9,

	"person_id": 9,

	"event_id": 9,
}
