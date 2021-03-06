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
type Model struct {
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	DeletedAt    time.Time `json:"deleted_at" bson:"deleted_at"`
	Id           string    `json:"id" bson:"_id,omitempty"`
	Name         string    `json:"name" bson:"name"`
	ObjectsIds   []string  `json:"objects_ids" bson:"objects_ids"`
	OntologyId   string    `json:"ontology_id" bson:"ontology_id"`
	OwnerId      string    `json:"owner_id" bson:"owner_id"`
	RelationsIds []string  `json:"relations_ids" bson:"relations_ids"`
	TraitsIds    []string  `json:"traits_ids" bson:"traits_ids"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

func NewModel() *Model {
	return &Model{}
}

func FindModel(db data.DB, id data.ID) (*Model, error) {

	model := NewModel()
	model.SetID(id)

	return model, db.PopulateByID(model)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (model *Model) Kind() data.Kind {
	return ModelKind
}

// just returns itself for now
func (model *Model) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = model.ID()
	return foo
}

func (model *Model) SetID(id data.ID) {
	model.Id = id.String()
}

func (model *Model) ID() data.ID {
	return data.ID(model.Id)
}

func (model *Model) IncludeObject(object *Object) {
	otherID := object.ID().String()
	for i := range model.ObjectsIds {
		if model.ObjectsIds[i] == otherID {
			return
		}
	}
	model.ObjectsIds = append(model.ObjectsIds, otherID)
}

func (model *Model) ExcludeObject(object *Object) {
	tmp := make([]string, 0)
	id := object.ID().String()
	for _, s := range model.ObjectsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	model.ObjectsIds = tmp
}

func (model *Model) ObjectsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(model.ObjectsIds), db), nil
}

func (model *Model) Objects(db data.DB) (objects []*Object, err error) {
	objects = make([]*Object, len(model.ObjectsIds))
	object := NewObject()
	for i, id := range model.ObjectsIds {
		object.Id = id
		if err = db.PopulateByID(object); err != nil {
			return
		}

		objects[i] = object
		object = NewObject()
	}

	return
}

func (model *Model) SetOntology(ontologyArgument *Ontology) error {
	model.OntologyId = ontologyArgument.ID().String()
	return nil
}

func (model *Model) Ontology(db data.DB) (*Ontology, error) {
	if model.OntologyId == "" {
		return nil, ErrEmptyLink
	}

	ontologyArgument := NewOntology()
	id, _ := db.ParseID(model.OntologyId)
	ontologyArgument.SetID(id)
	return ontologyArgument, db.PopulateByID(ontologyArgument)

}

func (model *Model) OntologyOrCreate(db data.DB) (*Ontology, error) {
	ontology, err := model.Ontology(db)

	if err == ErrEmptyLink {
		ontology := NewOntology()
		ontology.SetID(db.NewID())
		if err := model.SetOntology(ontology); err != nil {
			return nil, err
		}

		if err := db.Save(ontology); err != nil {
			return nil, err
		}

		if err := db.Save(model); err != nil {
			return nil, err
		}

		return ontology, nil
	} else {
		return ontology, err
	}
}

func (model *Model) SetOwner(userArgument *User) error {
	model.OwnerId = userArgument.ID().String()
	return nil
}

func (model *Model) Owner(db data.DB) (*User, error) {
	if model.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(model.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (model *Model) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := model.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := model.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(model); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

func (model *Model) IncludeRelation(relation *Relation) {
	otherID := relation.ID().String()
	for i := range model.RelationsIds {
		if model.RelationsIds[i] == otherID {
			return
		}
	}
	model.RelationsIds = append(model.RelationsIds, otherID)
}

func (model *Model) ExcludeRelation(relation *Relation) {
	tmp := make([]string, 0)
	id := relation.ID().String()
	for _, s := range model.RelationsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	model.RelationsIds = tmp
}

func (model *Model) RelationsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(model.RelationsIds), db), nil
}

func (model *Model) Relations(db data.DB) (relations []*Relation, err error) {
	relations = make([]*Relation, len(model.RelationsIds))
	relation := NewRelation()
	for i, id := range model.RelationsIds {
		relation.Id = id
		if err = db.PopulateByID(relation); err != nil {
			return
		}

		relations[i] = relation
		relation = NewRelation()
	}

	return
}

func (model *Model) IncludeTrait(trait *Trait) {
	otherID := trait.ID().String()
	for i := range model.TraitsIds {
		if model.TraitsIds[i] == otherID {
			return
		}
	}
	model.TraitsIds = append(model.TraitsIds, otherID)
}

func (model *Model) ExcludeTrait(trait *Trait) {
	tmp := make([]string, 0)
	id := trait.ID().String()
	for _, s := range model.TraitsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	model.TraitsIds = tmp
}

func (model *Model) TraitsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(model.TraitsIds), db), nil
}

func (model *Model) Traits(db data.DB) (traits []*Trait, err error) {
	traits = make([]*Trait, len(model.TraitsIds))
	trait := NewTrait()
	for i, id := range model.TraitsIds {
		trait.Id = id
		if err = db.PopulateByID(trait); err != nil {
			return
		}

		traits[i] = trait
		trait = NewTrait()
	}

	return
}

// BSON {{{
func (model *Model) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ObjectsIds []string `json:"objects_ids" bson:"objects_ids"`

		OntologyId string `json:"ontology_id" bson:"ontology_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		RelationsIds []string `json:"relations_ids" bson:"relations_ids"`

		TraitsIds []string `json:"traits_ids" bson:"traits_ids"`
	}{

		CreatedAt: model.CreatedAt,

		DeletedAt: model.DeletedAt,

		Name: model.Name,

		UpdatedAt: model.UpdatedAt,

		ObjectsIds: model.ObjectsIds,

		OntologyId: model.OntologyId,

		OwnerId: model.OwnerId,

		RelationsIds: model.RelationsIds,

		TraitsIds: model.TraitsIds,
	}, nil

}

func (model *Model) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ObjectsIds []string `json:"objects_ids" bson:"objects_ids"`

		OntologyId string `json:"ontology_id" bson:"ontology_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		RelationsIds []string `json:"relations_ids" bson:"relations_ids"`

		TraitsIds []string `json:"traits_ids" bson:"traits_ids"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	model.CreatedAt = tmp.CreatedAt

	model.DeletedAt = tmp.DeletedAt

	model.Id = tmp.Id.Hex()

	model.Name = tmp.Name

	model.UpdatedAt = tmp.UpdatedAt

	model.ObjectsIds = tmp.ObjectsIds

	model.OntologyId = tmp.OntologyId

	model.OwnerId = tmp.OwnerId

	model.RelationsIds = tmp.RelationsIds

	model.TraitsIds = tmp.TraitsIds

	return nil

}

// BSON }}}

func (model *Model) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		model.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		model.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		model.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		model.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["name"]; ok {
		model.Name = val.(string)
	}

	if val, ok := structure["owner_id"]; ok {
		model.OwnerId = val.(string)
	}

	if val, ok := structure["traits_ids"]; ok {
		model.TraitsIds = val.([]string)
	}

	if val, ok := structure["relations_ids"]; ok {
		model.RelationsIds = val.([]string)
	}

	if val, ok := structure["ontology_id"]; ok {
		model.OntologyId = val.(string)
	}

	if val, ok := structure["objects_ids"]; ok {
		model.ObjectsIds = val.([]string)
	}

}

var ModelStructure = map[string]metis.Primitive{

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"name": 3,

	"owner_id": 9,

	"traits_ids": 10,

	"relations_ids": 10,

	"ontology_id": 9,

	"objects_ids": 10,
}
