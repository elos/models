package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Link struct {
	CreatedAt  time.Time      `json:"created_at" bson:"created_at"`
	DeletedAt  time.Time      `json:"deleted_at" bson:"deleted_at"`
	Id         string         `json:"id" bson:"_id,omitempty"`
	Ids        map[int]string `json:"ids" bson:"ids"`
	ObjectId   string         `json:"object_id" bson:"object_id"`
	OwnerId    string         `json:"owner_id" bson:"owner_id"`
	RelationId string         `json:"relation_id" bson:"relation_id"`
	UpdatedAt  time.Time      `json:"updated_at" bson:"updated_at"`
}

func NewLink() *Link {
	return &Link{}
}

func FindLink(db data.DB, id data.ID) (*Link, error) {

	link := NewLink()
	link.SetID(id)

	return link, db.PopulateByID(link)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (link *Link) Kind() data.Kind {
	return LinkKind
}

// just returns itself for now
func (link *Link) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = link.ID()
	return foo
}

func (link *Link) SetID(id data.ID) {
	link.Id = id.String()
}

func (link *Link) ID() data.ID {
	return data.ID(link.Id)
}

func (link *Link) SetObject(objectArgument *Object) error {
	link.ObjectId = objectArgument.ID().String()
	return nil
}

func (link *Link) Object(db data.DB) (*Object, error) {
	if link.ObjectId == "" {
		return nil, ErrEmptyLink
	}

	objectArgument := NewObject()
	id, _ := db.ParseID(link.ObjectId)
	objectArgument.SetID(id)
	return objectArgument, db.PopulateByID(objectArgument)

}

func (link *Link) ObjectOrCreate(db data.DB) (*Object, error) {
	object, err := link.Object(db)

	if err == ErrEmptyLink {
		object := NewObject()
		object.SetID(db.NewID())
		if err := link.SetObject(object); err != nil {
			return nil, err
		}

		if err := db.Save(object); err != nil {
			return nil, err
		}

		if err := db.Save(link); err != nil {
			return nil, err
		}

		return object, nil
	} else {
		return object, err
	}
}

func (link *Link) SetOwner(userArgument *User) error {
	link.OwnerId = userArgument.ID().String()
	return nil
}

func (link *Link) Owner(db data.DB) (*User, error) {
	if link.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(link.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (link *Link) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := link.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := link.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(link); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

func (link *Link) SetRelation(relationArgument *Relation) error {
	link.RelationId = relationArgument.ID().String()
	return nil
}

func (link *Link) Relation(db data.DB) (*Relation, error) {
	if link.RelationId == "" {
		return nil, ErrEmptyLink
	}

	relationArgument := NewRelation()
	id, _ := db.ParseID(link.RelationId)
	relationArgument.SetID(id)
	return relationArgument, db.PopulateByID(relationArgument)

}

func (link *Link) RelationOrCreate(db data.DB) (*Relation, error) {
	relation, err := link.Relation(db)

	if err == ErrEmptyLink {
		relation := NewRelation()
		relation.SetID(db.NewID())
		if err := link.SetRelation(relation); err != nil {
			return nil, err
		}

		if err := db.Save(relation); err != nil {
			return nil, err
		}

		if err := db.Save(link); err != nil {
			return nil, err
		}

		return relation, nil
	} else {
		return relation, err
	}
}

// BSON {{{
func (link *Link) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Ids map[int]string `json:"ids" bson:"ids"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ObjectId string `json:"object_id" bson:"object_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		RelationId string `json:"relation_id" bson:"relation_id"`
	}{

		CreatedAt: link.CreatedAt,

		DeletedAt: link.DeletedAt,

		Ids: link.Ids,

		UpdatedAt: link.UpdatedAt,

		ObjectId: link.ObjectId,

		OwnerId: link.OwnerId,

		RelationId: link.RelationId,
	}, nil

}

func (link *Link) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Ids map[int]string `json:"ids" bson:"ids"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ObjectId string `json:"object_id" bson:"object_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		RelationId string `json:"relation_id" bson:"relation_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	link.CreatedAt = tmp.CreatedAt

	link.DeletedAt = tmp.DeletedAt

	link.Id = tmp.Id.Hex()

	link.Ids = tmp.Ids

	link.UpdatedAt = tmp.UpdatedAt

	link.ObjectId = tmp.ObjectId

	link.OwnerId = tmp.OwnerId

	link.RelationId = tmp.RelationId

	return nil

}

// BSON }}}

func (link *Link) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["deleted_at"]; ok {
		link.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["ids"]; ok {
		link.Ids = val.(map[int]string)
	}

	if val, ok := structure["id"]; ok {
		link.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		link.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		link.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["owner_id"]; ok {
		link.OwnerId = val.(string)
	}

	if val, ok := structure["object_id"]; ok {
		link.ObjectId = val.(string)
	}

	if val, ok := structure["relation_id"]; ok {
		link.RelationId = val.(string)
	}

}

var LinkStructure = map[string]metis.Primitive{

	"ids": 12,

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"owner_id": 9,

	"object_id": 9,

	"relation_id": 9,
}
