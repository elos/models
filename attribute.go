package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Attribute struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Id        string    `json:"id" bson:"_id,omitempty"`
	ObjectID  string    `json:"object_id" bson:"object_id"`
	TraitID   string    `json:"trait_id" bson:"trait_id"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Value     string    `json:"value" bson:"value"`
}

func NewAttribute() *Attribute {
	return &Attribute{}
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

func (attribute *Attribute) SetObject(object *Object) error {
	attribute.ObjectID = object.ID().String()
	return nil
}

func (attribute *Attribute) Object(db data.DB) (*Object, error) {
	if attribute.ObjectID == "" {
		return nil, ErrEmptyLink
	}

	object := NewObject()
	pid, _ := mongo.ParseObjectID(attribute.ObjectID)
	object.SetID(data.ID(pid.Hex()))
	return object, db.PopulateByID(object)

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

func (attribute *Attribute) SetTrait(trait *Trait) error {
	attribute.TraitID = trait.ID().String()
	return nil
}

func (attribute *Attribute) Trait(db data.DB) (*Trait, error) {
	if attribute.TraitID == "" {
		return nil, ErrEmptyLink
	}

	trait := NewTrait()
	pid, _ := mongo.ParseObjectID(attribute.TraitID)
	trait.SetID(data.ID(pid.Hex()))
	return trait, db.PopulateByID(trait)

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

		Id string `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		Value string `json:"value" bson:"value"`

		ObjectID string `json:"object_id" bson:"object_id"`

		TraitID string `json:"trait_id" bson:"trait_id"`
	}{

		CreatedAt: attribute.CreatedAt,

		UpdatedAt: attribute.UpdatedAt,

		Value: attribute.Value,

		ObjectID: attribute.ObjectID,

		TraitID: attribute.TraitID,
	}, nil

}

func (attribute *Attribute) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		Value string `json:"value" bson:"value"`

		ObjectID string `json:"object_id" bson:"object_id"`

		TraitID string `json:"trait_id" bson:"trait_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	attribute.CreatedAt = tmp.CreatedAt

	attribute.Id = tmp.Id.Hex()

	attribute.UpdatedAt = tmp.UpdatedAt

	attribute.Value = tmp.Value

	attribute.ObjectID = tmp.ObjectID

	attribute.TraitID = tmp.TraitID

	return nil

}

// BSON }}}
