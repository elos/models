package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Relation struct {
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
	Id        string     `json:"id" bson:"_id,omitempty"`
	Ids       []string   `json:"ids" bson:"ids"`
	LinkID    string     `json:"link_id" bson:"link_id"`
	ObjectID  string     `json:"object_id" bson:"object_id"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}

func NewRelation() *Relation {
	return &Relation{}
}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (relation *Relation) Kind() data.Kind {
	return RelationKind
}

// just returns itself for now
func (relation *Relation) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = relation.ID()
	return foo
}

func (relation *Relation) SetID(id data.ID) {
	relation.Id = id.String()
}

func (relation *Relation) ID() data.ID {
	return data.ID(relation.Id)
}

func (relation *Relation) SetLink(link *Link) error {
	relation.LinkID = link.ID().String()
	return nil
}

func (relation *Relation) Link(store data.Store) (*Link, error) {
	if relation.LinkID == "" {
		return nil, ErrEmptyLink
	}

	link := NewLink()
	pid, _ := mongo.ParseObjectID(relation.LinkID)
	link.SetID(data.ID(pid.Hex()))
	return link, store.PopulateByID(link)

}

func (relation *Relation) SetObject(object *Object) error {
	relation.ObjectID = object.ID().String()
	return nil
}

func (relation *Relation) Object(store data.Store) (*Object, error) {
	if relation.ObjectID == "" {
		return nil, ErrEmptyLink
	}

	object := NewObject()
	pid, _ := mongo.ParseObjectID(relation.ObjectID)
	object.SetID(data.ID(pid.Hex()))
	return object, store.PopulateByID(object)

}

func (relation *Relation) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt *time.Time `json:"created_at" bson:"created_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Ids []string `json:"ids" bson:"ids"`

		UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`

		LinkID string `json:"link_id" bson:"link_id"`

		ObjectID string `json:"object_id" bson:"object_id"`
	}{

		CreatedAt: relation.CreatedAt,

		Ids: relation.Ids,

		UpdatedAt: relation.UpdatedAt,

		LinkID: relation.LinkID,

		ObjectID: relation.ObjectID,
	}, nil

}

func (relation *Relation) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt *time.Time `json:"created_at" bson:"created_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Ids []string `json:"ids" bson:"ids"`

		UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`

		LinkID string `json:"link_id" bson:"link_id"`

		ObjectID string `json:"object_id" bson:"object_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	relation.CreatedAt = tmp.CreatedAt

	relation.Id = tmp.Id.Hex()

	relation.Ids = tmp.Ids

	relation.UpdatedAt = tmp.UpdatedAt

	relation.LinkID = tmp.LinkID

	relation.ObjectID = tmp.ObjectID

	return nil

}
