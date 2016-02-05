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
type Object struct {
	AttributesIds []string  `json:"attributes_ids" bson:"attributes_ids"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	DeletedAt     time.Time `json:"deleted_at" bson:"deleted_at"`
	Id            string    `json:"id" bson:"_id,omitempty"`
	LinksIds      []string  `json:"links_ids" bson:"links_ids"`
	ModelId       string    `json:"model_id" bson:"model_id"`
	OntologyId    string    `json:"ontology_id" bson:"ontology_id"`
	OwnerId       string    `json:"owner_id" bson:"owner_id"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
}

func NewObject() *Object {
	return &Object{}
}

func FindObject(db data.DB, id data.ID) (*Object, error) {

	object := NewObject()
	object.SetID(id)

	return object, db.PopulateByID(object)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (object *Object) Kind() data.Kind {
	return ObjectKind
}

// just returns itself for now
func (object *Object) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = object.ID()
	return foo
}

func (object *Object) SetID(id data.ID) {
	object.Id = id.String()
}

func (object *Object) ID() data.ID {
	return data.ID(object.Id)
}

func (object *Object) IncludeAttribute(attribute *Attribute) {
	otherID := attribute.ID().String()
	for i := range object.AttributesIds {
		if object.AttributesIds[i] == otherID {
			return
		}
	}
	object.AttributesIds = append(object.AttributesIds, otherID)
}

func (object *Object) ExcludeAttribute(attribute *Attribute) {
	tmp := make([]string, 0)
	id := attribute.ID().String()
	for _, s := range object.AttributesIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	object.AttributesIds = tmp
}

func (object *Object) AttributesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(object.AttributesIds), db), nil
}

func (object *Object) Attributes(db data.DB) (attributes []*Attribute, err error) {
	attributes = make([]*Attribute, len(object.AttributesIds))
	attribute := NewAttribute()
	for i, id := range object.AttributesIds {
		attribute.Id = id
		if err = db.PopulateByID(attribute); err != nil {
			return
		}

		attributes[i] = attribute
		attribute = NewAttribute()
	}

	return
}

func (object *Object) IncludeLink(link *Link) {
	otherID := link.ID().String()
	for i := range object.LinksIds {
		if object.LinksIds[i] == otherID {
			return
		}
	}
	object.LinksIds = append(object.LinksIds, otherID)
}

func (object *Object) ExcludeLink(link *Link) {
	tmp := make([]string, 0)
	id := link.ID().String()
	for _, s := range object.LinksIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	object.LinksIds = tmp
}

func (object *Object) LinksIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(object.LinksIds), db), nil
}

func (object *Object) Links(db data.DB) (links []*Link, err error) {
	links = make([]*Link, len(object.LinksIds))
	link := NewLink()
	for i, id := range object.LinksIds {
		link.Id = id
		if err = db.PopulateByID(link); err != nil {
			return
		}

		links[i] = link
		link = NewLink()
	}

	return
}

func (object *Object) SetModel(modelArgument *Model) error {
	object.ModelId = modelArgument.ID().String()
	return nil
}

func (object *Object) Model(db data.DB) (*Model, error) {
	if object.ModelId == "" {
		return nil, ErrEmptyLink
	}

	modelArgument := NewModel()
	id, _ := db.ParseID(object.ModelId)
	modelArgument.SetID(id)
	return modelArgument, db.PopulateByID(modelArgument)

}

func (object *Object) ModelOrCreate(db data.DB) (*Model, error) {
	model, err := object.Model(db)

	if err == ErrEmptyLink {
		model := NewModel()
		model.SetID(db.NewID())
		if err := object.SetModel(model); err != nil {
			return nil, err
		}

		if err := db.Save(model); err != nil {
			return nil, err
		}

		if err := db.Save(object); err != nil {
			return nil, err
		}

		return model, nil
	} else {
		return model, err
	}
}

func (object *Object) SetOntology(ontologyArgument *Ontology) error {
	object.OntologyId = ontologyArgument.ID().String()
	return nil
}

func (object *Object) Ontology(db data.DB) (*Ontology, error) {
	if object.OntologyId == "" {
		return nil, ErrEmptyLink
	}

	ontologyArgument := NewOntology()
	id, _ := db.ParseID(object.OntologyId)
	ontologyArgument.SetID(id)
	return ontologyArgument, db.PopulateByID(ontologyArgument)

}

func (object *Object) OntologyOrCreate(db data.DB) (*Ontology, error) {
	ontology, err := object.Ontology(db)

	if err == ErrEmptyLink {
		ontology := NewOntology()
		ontology.SetID(db.NewID())
		if err := object.SetOntology(ontology); err != nil {
			return nil, err
		}

		if err := db.Save(ontology); err != nil {
			return nil, err
		}

		if err := db.Save(object); err != nil {
			return nil, err
		}

		return ontology, nil
	} else {
		return ontology, err
	}
}

func (object *Object) SetOwner(userArgument *User) error {
	object.OwnerId = userArgument.ID().String()
	return nil
}

func (object *Object) Owner(db data.DB) (*User, error) {
	if object.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(object.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (object *Object) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := object.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := object.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(object); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (object *Object) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		AttributesIds []string `json:"attributes_ids" bson:"attributes_ids"`

		LinksIds []string `json:"links_ids" bson:"links_ids"`

		ModelId string `json:"model_id" bson:"model_id"`

		OntologyId string `json:"ontology_id" bson:"ontology_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: object.CreatedAt,

		DeletedAt: object.DeletedAt,

		UpdatedAt: object.UpdatedAt,

		AttributesIds: object.AttributesIds,

		LinksIds: object.LinksIds,

		ModelId: object.ModelId,

		OntologyId: object.OntologyId,

		OwnerId: object.OwnerId,
	}, nil

}

func (object *Object) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		AttributesIds []string `json:"attributes_ids" bson:"attributes_ids"`

		LinksIds []string `json:"links_ids" bson:"links_ids"`

		ModelId string `json:"model_id" bson:"model_id"`

		OntologyId string `json:"ontology_id" bson:"ontology_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	object.CreatedAt = tmp.CreatedAt

	object.DeletedAt = tmp.DeletedAt

	object.Id = tmp.Id.Hex()

	object.UpdatedAt = tmp.UpdatedAt

	object.AttributesIds = tmp.AttributesIds

	object.LinksIds = tmp.LinksIds

	object.ModelId = tmp.ModelId

	object.OntologyId = tmp.OntologyId

	object.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (object *Object) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		object.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		object.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		object.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		object.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["links_ids"]; ok {
		object.LinksIds = val.([]string)
	}

	if val, ok := structure["model_id"]; ok {
		object.ModelId = val.(string)
	}

	if val, ok := structure["ontology_id"]; ok {
		object.OntologyId = val.(string)
	}

	if val, ok := structure["owner_id"]; ok {
		object.OwnerId = val.(string)
	}

	if val, ok := structure["attributes_ids"]; ok {
		object.AttributesIds = val.([]string)
	}

}

var ObjectStructure = map[string]metis.Primitive{

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"owner_id": 9,

	"attributes_ids": 10,

	"links_ids": 10,

	"model_id": 9,

	"ontology_id": 9,
}
