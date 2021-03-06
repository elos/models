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
type Trait struct {
	AttributesIds []string  `json:"attributes_ids" bson:"attributes_ids"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	DeletedAt     time.Time `json:"deleted_at" bson:"deleted_at"`
	Id            string    `json:"id" bson:"_id,omitempty"`
	ModelId       string    `json:"model_id" bson:"model_id"`
	Name          string    `json:"name" bson:"name"`
	OwnerId       string    `json:"owner_id" bson:"owner_id"`
	Primitive     string    `json:"primitive" bson:"primitive"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
}

func NewTrait() *Trait {
	return &Trait{}
}

func FindTrait(db data.DB, id data.ID) (*Trait, error) {

	trait := NewTrait()
	trait.SetID(id)

	return trait, db.PopulateByID(trait)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (trait *Trait) Kind() data.Kind {
	return TraitKind
}

// just returns itself for now
func (trait *Trait) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = trait.ID()
	return foo
}

func (trait *Trait) SetID(id data.ID) {
	trait.Id = id.String()
}

func (trait *Trait) ID() data.ID {
	return data.ID(trait.Id)
}

func (trait *Trait) IncludeAttribute(attribute *Attribute) {
	otherID := attribute.ID().String()
	for i := range trait.AttributesIds {
		if trait.AttributesIds[i] == otherID {
			return
		}
	}
	trait.AttributesIds = append(trait.AttributesIds, otherID)
}

func (trait *Trait) ExcludeAttribute(attribute *Attribute) {
	tmp := make([]string, 0)
	id := attribute.ID().String()
	for _, s := range trait.AttributesIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	trait.AttributesIds = tmp
}

func (trait *Trait) AttributesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(trait.AttributesIds), db), nil
}

func (trait *Trait) Attributes(db data.DB) (attributes []*Attribute, err error) {
	attributes = make([]*Attribute, len(trait.AttributesIds))
	attribute := NewAttribute()
	for i, id := range trait.AttributesIds {
		attribute.Id = id
		if err = db.PopulateByID(attribute); err != nil {
			return
		}

		attributes[i] = attribute
		attribute = NewAttribute()
	}

	return
}

func (trait *Trait) SetModel(modelArgument *Model) error {
	trait.ModelId = modelArgument.ID().String()
	return nil
}

func (trait *Trait) Model(db data.DB) (*Model, error) {
	if trait.ModelId == "" {
		return nil, ErrEmptyLink
	}

	modelArgument := NewModel()
	id, _ := db.ParseID(trait.ModelId)
	modelArgument.SetID(id)
	return modelArgument, db.PopulateByID(modelArgument)

}

func (trait *Trait) ModelOrCreate(db data.DB) (*Model, error) {
	model, err := trait.Model(db)

	if err == ErrEmptyLink {
		model := NewModel()
		model.SetID(db.NewID())
		if err := trait.SetModel(model); err != nil {
			return nil, err
		}

		if err := db.Save(model); err != nil {
			return nil, err
		}

		if err := db.Save(trait); err != nil {
			return nil, err
		}

		return model, nil
	} else {
		return model, err
	}
}

func (trait *Trait) SetOwner(userArgument *User) error {
	trait.OwnerId = userArgument.ID().String()
	return nil
}

func (trait *Trait) Owner(db data.DB) (*User, error) {
	if trait.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(trait.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (trait *Trait) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := trait.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := trait.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(trait); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (trait *Trait) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		Primitive string `json:"primitive" bson:"primitive"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		AttributesIds []string `json:"attributes_ids" bson:"attributes_ids"`

		ModelId string `json:"model_id" bson:"model_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: trait.CreatedAt,

		DeletedAt: trait.DeletedAt,

		Name: trait.Name,

		Primitive: trait.Primitive,

		UpdatedAt: trait.UpdatedAt,

		AttributesIds: trait.AttributesIds,

		ModelId: trait.ModelId,

		OwnerId: trait.OwnerId,
	}, nil

}

func (trait *Trait) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		Primitive string `json:"primitive" bson:"primitive"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		AttributesIds []string `json:"attributes_ids" bson:"attributes_ids"`

		ModelId string `json:"model_id" bson:"model_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	trait.CreatedAt = tmp.CreatedAt

	trait.DeletedAt = tmp.DeletedAt

	trait.Id = tmp.Id.Hex()

	trait.Name = tmp.Name

	trait.Primitive = tmp.Primitive

	trait.UpdatedAt = tmp.UpdatedAt

	trait.AttributesIds = tmp.AttributesIds

	trait.ModelId = tmp.ModelId

	trait.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (trait *Trait) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		trait.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		trait.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		trait.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		trait.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["name"]; ok {
		trait.Name = val.(string)
	}

	if val, ok := structure["primitive"]; ok {
		trait.Primitive = val.(string)
	}

	if val, ok := structure["owner_id"]; ok {
		trait.OwnerId = val.(string)
	}

	if val, ok := structure["model_id"]; ok {
		trait.ModelId = val.(string)
	}

	if val, ok := structure["attributes_ids"]; ok {
		trait.AttributesIds = val.([]string)
	}

}

var TraitStructure = map[string]metis.Primitive{

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"name": 3,

	"primitive": 3,

	"owner_id": 9,

	"model_id": 9,

	"attributes_ids": 10,
}
