package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Quantity struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
	Id        string    `json:"id" bson:"_id,omitempty"`
	OwnerId   string    `json:"owner_id" bson:"owner_id"`
	Unit      string    `json:"unit" bson:"unit"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Value     float64   `json:"value" bson:"value"`
}

func NewQuantity() *Quantity {
	return &Quantity{}
}

func FindQuantity(db data.DB, id data.ID) (*Quantity, error) {

	quantity := NewQuantity()
	quantity.SetID(id)

	return quantity, db.PopulateByID(quantity)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (quantity *Quantity) Kind() data.Kind {
	return QuantityKind
}

// just returns itself for now
func (quantity *Quantity) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = quantity.ID()
	return foo
}

func (quantity *Quantity) SetID(id data.ID) {
	quantity.Id = id.String()
}

func (quantity *Quantity) ID() data.ID {
	return data.ID(quantity.Id)
}

func (quantity *Quantity) SetOwner(userArgument *User) error {
	quantity.OwnerId = userArgument.ID().String()
	return nil
}

func (quantity *Quantity) Owner(db data.DB) (*User, error) {
	if quantity.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	pid, _ := mongo.ParseObjectID(quantity.OwnerId)
	userArgument.SetID(data.ID(pid.Hex()))
	return userArgument, db.PopulateByID(userArgument)

}

func (quantity *Quantity) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := quantity.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := quantity.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(quantity); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (quantity *Quantity) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Unit string `json:"unit" bson:"unit"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		Value float64 `json:"value" bson:"value"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: quantity.CreatedAt,

		DeletedAt: quantity.DeletedAt,

		Unit: quantity.Unit,

		UpdatedAt: quantity.UpdatedAt,

		Value: quantity.Value,

		OwnerId: quantity.OwnerId,
	}, nil

}

func (quantity *Quantity) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Unit string `json:"unit" bson:"unit"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		Value float64 `json:"value" bson:"value"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	quantity.CreatedAt = tmp.CreatedAt

	quantity.DeletedAt = tmp.DeletedAt

	quantity.Id = tmp.Id.Hex()

	quantity.Unit = tmp.Unit

	quantity.UpdatedAt = tmp.UpdatedAt

	quantity.Value = tmp.Value

	quantity.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (quantity *Quantity) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["value"]; ok {
		quantity.Value = val.(float64)
	}

	if val, ok := structure["unit"]; ok {
		quantity.Unit = val.(string)
	}

	if val, ok := structure["id"]; ok {
		quantity.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		quantity.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		quantity.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		quantity.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["owner_id"]; ok {
		quantity.OwnerId = val.(string)
	}

}

var QuantityStructure = map[string]metis.Primitive{

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"value": 2,

	"unit": 3,

	"owner_id": 9,
}
