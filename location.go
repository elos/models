package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Location struct {
	Altitude  float64   `json:"altitude" bson:"altitude"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
	Id        string    `json:"id" bson:"_id,omitempty"`
	Latitude  float64   `json:"latitude" bson:"latitude"`
	Longitude float64   `json:"longitude" bson:"longitude"`
	OwnerId   string    `json:"owner_id" bson:"owner_id"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewLocation() *Location {
	return &Location{}
}

func FindLocation(db data.DB, id data.ID) (*Location, error) {

	location := NewLocation()
	location.SetID(id)

	return location, db.PopulateByID(location)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (location *Location) Kind() data.Kind {
	return LocationKind
}

// just returns itself for now
func (location *Location) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = location.ID()
	return foo
}

func (location *Location) SetID(id data.ID) {
	location.Id = id.String()
}

func (location *Location) ID() data.ID {
	return data.ID(location.Id)
}

func (location *Location) SetOwner(userArgument *User) error {
	location.OwnerId = userArgument.ID().String()
	return nil
}

func (location *Location) Owner(db data.DB) (*User, error) {
	if location.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(location.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (location *Location) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := location.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := location.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(location); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (location *Location) GetBSON() (interface{}, error) {

	return struct {
		Altitude float64 `json:"altitude" bson:"altitude"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Latitude float64 `json:"latitude" bson:"latitude"`

		Longitude float64 `json:"longitude" bson:"longitude"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		Altitude: location.Altitude,

		CreatedAt: location.CreatedAt,

		DeletedAt: location.DeletedAt,

		Latitude: location.Latitude,

		Longitude: location.Longitude,

		UpdatedAt: location.UpdatedAt,

		OwnerId: location.OwnerId,
	}, nil

}

func (location *Location) SetBSON(raw bson.Raw) error {

	tmp := struct {
		Altitude float64 `json:"altitude" bson:"altitude"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Latitude float64 `json:"latitude" bson:"latitude"`

		Longitude float64 `json:"longitude" bson:"longitude"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	location.Altitude = tmp.Altitude

	location.CreatedAt = tmp.CreatedAt

	location.DeletedAt = tmp.DeletedAt

	location.Id = tmp.Id.Hex()

	location.Latitude = tmp.Latitude

	location.Longitude = tmp.Longitude

	location.UpdatedAt = tmp.UpdatedAt

	location.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (location *Location) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["latitude"]; ok {
		location.Latitude = val.(float64)
	}

	if val, ok := structure["longitude"]; ok {
		location.Longitude = val.(float64)
	}

	if val, ok := structure["altitude"]; ok {
		location.Altitude = val.(float64)
	}

	if val, ok := structure["id"]; ok {
		location.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		location.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		location.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		location.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["owner_id"]; ok {
		location.OwnerId = val.(string)
	}

}

var LocationStructure = map[string]metis.Primitive{

	"altitude": 2,

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"latitude": 2,

	"longitude": 2,

	"owner_id": 9,
}
