package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Context struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
	Domain    string    `json:"domain" bson:"domain"`
	Id        string    `json:"id" bson:"_id,omitempty"`
	Ids       []string  `json:"ids" bson:"ids"`
	OwnerId   string    `json:"owner_id" bson:"owner_id"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewContext() *Context {
	return &Context{}
}

func FindContext(db data.DB, id data.ID) (*Context, error) {

	context := NewContext()
	context.SetID(id)

	return context, db.PopulateByID(context)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (context *Context) Kind() data.Kind {
	return ContextKind
}

// just returns itself for now
func (context *Context) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = context.ID()
	return foo
}

func (context *Context) SetID(id data.ID) {
	context.Id = id.String()
}

func (context *Context) ID() data.ID {
	return data.ID(context.Id)
}

func (context *Context) SetOwner(userArgument *User) error {
	context.OwnerId = userArgument.ID().String()
	return nil
}

func (context *Context) Owner(db data.DB) (*User, error) {
	if context.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(context.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (context *Context) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := context.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := context.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(context); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (context *Context) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Domain string `json:"domain" bson:"domain"`

		Id string `json:"id" bson:"_id,omitempty"`

		Ids []string `json:"ids" bson:"ids"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: context.CreatedAt,

		DeletedAt: context.DeletedAt,

		Domain: context.Domain,

		Ids: context.Ids,

		UpdatedAt: context.UpdatedAt,

		OwnerId: context.OwnerId,
	}, nil

}

func (context *Context) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Domain string `json:"domain" bson:"domain"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Ids []string `json:"ids" bson:"ids"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	context.CreatedAt = tmp.CreatedAt

	context.DeletedAt = tmp.DeletedAt

	context.Domain = tmp.Domain

	context.Id = tmp.Id.Hex()

	context.Ids = tmp.Ids

	context.UpdatedAt = tmp.UpdatedAt

	context.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (context *Context) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		context.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		context.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		context.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		context.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["domain"]; ok {
		context.Domain = val.(string)
	}

	if val, ok := structure["ids"]; ok {
		context.Ids = val.([]string)
	}

	if val, ok := structure["owner_id"]; ok {
		context.OwnerId = val.(string)
	}

}

var ContextStructure = map[string]metis.Primitive{

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"domain": 3,

	"ids": 10,

	"owner_id": 9,
}
