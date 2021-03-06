package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Attribute struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
	Id        string    `json:"id" bson:"_id,omitempty"`
	ObjectId  string    `json:"object_id" bson:"object_id"`
	OwnerId   string    `json:"owner_id" bson:"owner_id"`
	TraitId   string    `json:"trait_id" bson:"trait_id"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Value     string    `json:"value" bson:"value"`
}

func NewAttribute() *Attribute {
	return &Attribute{}
}

func FindAttribute(db data.DB, id data.ID) (*Attribute, error) {

	attribute := NewAttribute()
	attribute.SetID(id)

	return attribute, db.PopulateByID(attribute)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (attribute *Attribute) Kind() data.Kind {
	return AttributeKind
}

// just returns itself for now
func (attribute *Attribute) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = attribute.ID()
	return foo
}

func (attribute *Attribute) SetID(id data.ID) {
	attribute.Id = id.String()
}

func (attribute *Attribute) ID() data.ID {
	return data.ID(attribute.Id)
}

func (attribute *Attribute) SetObject(objectArgument *Object) error {
	attribute.ObjectId = objectArgument.ID().String()
	return nil
}

func (attribute *Attribute) Object(db data.DB) (*Object, error) {
	if attribute.ObjectId == "" {
		return nil, ErrEmptyLink
	}

	objectArgument := NewObject()
	id, _ := db.ParseID(attribute.ObjectId)
	objectArgument.SetID(id)
	return objectArgument, db.PopulateByID(objectArgument)

}

func (attribute *Attribute) ObjectOrCreate(db data.DB) (*Object, error) {
	object, err := attribute.Object(db)

	if err == ErrEmptyLink {
		object := NewObject()
		object.SetID(db.NewID())
		if err := attribute.SetObject(object); err != nil {
			return nil, err
		}

		if err := db.Save(object); err != nil {
			return nil, err
		}

		if err := db.Save(attribute); err != nil {
			return nil, err
		}

		return object, nil
	} else {
		return object, err
	}
}

func (attribute *Attribute) SetOwner(userArgument *User) error {
	attribute.OwnerId = userArgument.ID().String()
	return nil
}

func (attribute *Attribute) Owner(db data.DB) (*User, error) {
	if attribute.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(attribute.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (attribute *Attribute) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := attribute.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := attribute.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(attribute); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

func (attribute *Attribute) SetTrait(traitArgument *Trait) error {
	attribute.TraitId = traitArgument.ID().String()
	return nil
}

func (attribute *Attribute) Trait(db data.DB) (*Trait, error) {
	if attribute.TraitId == "" {
		return nil, ErrEmptyLink
	}

	traitArgument := NewTrait()
	id, _ := db.ParseID(attribute.TraitId)
	traitArgument.SetID(id)
	return traitArgument, db.PopulateByID(traitArgument)

}

func (attribute *Attribute) TraitOrCreate(db data.DB) (*Trait, error) {
	trait, err := attribute.Trait(db)

	if err == ErrEmptyLink {
		trait := NewTrait()
		trait.SetID(db.NewID())
		if err := attribute.SetTrait(trait); err != nil {
			return nil, err
		}

		if err := db.Save(trait); err != nil {
			return nil, err
		}

		if err := db.Save(attribute); err != nil {
			return nil, err
		}

		return trait, nil
	} else {
		return trait, err
	}
}

// BSON {{{
func (attribute *Attribute) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		Value string `json:"value" bson:"value"`

		ObjectId string `json:"object_id" bson:"object_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		TraitId string `json:"trait_id" bson:"trait_id"`
	}{

		CreatedAt: attribute.CreatedAt,

		DeletedAt: attribute.DeletedAt,

		UpdatedAt: attribute.UpdatedAt,

		Value: attribute.Value,

		ObjectId: attribute.ObjectId,

		OwnerId: attribute.OwnerId,

		TraitId: attribute.TraitId,
	}, nil

}

func (attribute *Attribute) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		Value string `json:"value" bson:"value"`

		ObjectId string `json:"object_id" bson:"object_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		TraitId string `json:"trait_id" bson:"trait_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	attribute.CreatedAt = tmp.CreatedAt

	attribute.DeletedAt = tmp.DeletedAt

	attribute.Id = tmp.Id.Hex()

	attribute.UpdatedAt = tmp.UpdatedAt

	attribute.Value = tmp.Value

	attribute.ObjectId = tmp.ObjectId

	attribute.OwnerId = tmp.OwnerId

	attribute.TraitId = tmp.TraitId

	return nil

}

// BSON }}}

func (attribute *Attribute) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["updated_at"]; ok {
		attribute.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		attribute.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["value"]; ok {
		attribute.Value = val.(string)
	}

	if val, ok := structure["id"]; ok {
		attribute.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		attribute.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["owner_id"]; ok {
		attribute.OwnerId = val.(string)
	}

	if val, ok := structure["object_id"]; ok {
		attribute.ObjectId = val.(string)
	}

	if val, ok := structure["trait_id"]; ok {
		attribute.TraitId = val.(string)
	}

}

var AttributeStructure = map[string]metis.Primitive{

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"value": 3,

	"id": 9,

	"owner_id": 9,

	"object_id": 9,

	"trait_id": 9,
}
