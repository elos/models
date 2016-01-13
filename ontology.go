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
type Ontology struct {
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	DeletedAt  time.Time `json:"deleted_at" bson:"deleted_at"`
	Id         string    `json:"id" bson:"_id,omitempty"`
	ModelsIds  []string  `json:"models_ids" bson:"models_ids"`
	ObjectsIds []string  `json:"objects_ids" bson:"objects_ids"`
	OwnerId    string    `json:"owner_id" bson:"owner_id"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

func NewOntology() *Ontology {
	return &Ontology{}
}

func FindOntology(db data.DB, id data.ID) (*Ontology, error) {

	ontology := NewOntology()
	ontology.SetID(id)

	return ontology, db.PopulateByID(ontology)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (ontology *Ontology) Kind() data.Kind {
	return OntologyKind
}

// just returns itself for now
func (ontology *Ontology) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = ontology.ID()
	return foo
}

func (ontology *Ontology) SetID(id data.ID) {
	ontology.Id = id.String()
}

func (ontology *Ontology) ID() data.ID {
	return data.ID(ontology.Id)
}

func (ontology *Ontology) IncludeModel(model *Model) {
	otherID := model.ID().String()
	for i := range ontology.ModelsIds {
		if ontology.ModelsIds[i] == otherID {
			return
		}
	}
	ontology.ModelsIds = append(ontology.ModelsIds, otherID)
}

func (ontology *Ontology) ExcludeModel(model *Model) {
	tmp := make([]string, 0)
	id := model.ID().String()
	for _, s := range ontology.ModelsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	ontology.ModelsIds = tmp
}

func (ontology *Ontology) ModelsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(ontology.ModelsIds), db), nil
}

func (ontology *Ontology) Models(db data.DB) (models []*Model, err error) {
	models = make([]*Model, len(ontology.ModelsIds))
	model := NewModel()
	for i, id := range ontology.ModelsIds {
		model.Id = id
		if err = db.PopulateByID(model); err != nil {
			return
		}

		models[i] = model
		model = NewModel()
	}

	return
}

func (ontology *Ontology) IncludeObject(object *Object) {
	otherID := object.ID().String()
	for i := range ontology.ObjectsIds {
		if ontology.ObjectsIds[i] == otherID {
			return
		}
	}
	ontology.ObjectsIds = append(ontology.ObjectsIds, otherID)
}

func (ontology *Ontology) ExcludeObject(object *Object) {
	tmp := make([]string, 0)
	id := object.ID().String()
	for _, s := range ontology.ObjectsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	ontology.ObjectsIds = tmp
}

func (ontology *Ontology) ObjectsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(ontology.ObjectsIds), db), nil
}

func (ontology *Ontology) Objects(db data.DB) (objects []*Object, err error) {
	objects = make([]*Object, len(ontology.ObjectsIds))
	object := NewObject()
	for i, id := range ontology.ObjectsIds {
		object.Id = id
		if err = db.PopulateByID(object); err != nil {
			return
		}

		objects[i] = object
		object = NewObject()
	}

	return
}

func (ontology *Ontology) SetOwner(userArgument *User) error {
	ontology.OwnerId = userArgument.ID().String()
	return nil
}

func (ontology *Ontology) Owner(db data.DB) (*User, error) {
	if ontology.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	pid, _ := mongo.ParseObjectID(ontology.OwnerId)
	userArgument.SetID(data.ID(pid.Hex()))
	return userArgument, db.PopulateByID(userArgument)

}

func (ontology *Ontology) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := ontology.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := ontology.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(ontology); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (ontology *Ontology) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ModelsIds []string `json:"models_ids" bson:"models_ids"`

		ObjectsIds []string `json:"objects_ids" bson:"objects_ids"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: ontology.CreatedAt,

		DeletedAt: ontology.DeletedAt,

		UpdatedAt: ontology.UpdatedAt,

		ModelsIds: ontology.ModelsIds,

		ObjectsIds: ontology.ObjectsIds,

		OwnerId: ontology.OwnerId,
	}, nil

}

func (ontology *Ontology) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ModelsIds []string `json:"models_ids" bson:"models_ids"`

		ObjectsIds []string `json:"objects_ids" bson:"objects_ids"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	ontology.CreatedAt = tmp.CreatedAt

	ontology.DeletedAt = tmp.DeletedAt

	ontology.Id = tmp.Id.Hex()

	ontology.UpdatedAt = tmp.UpdatedAt

	ontology.ModelsIds = tmp.ModelsIds

	ontology.ObjectsIds = tmp.ObjectsIds

	ontology.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (ontology *Ontology) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["deleted_at"]; ok {
		ontology.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["id"]; ok {
		ontology.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		ontology.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		ontology.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["models_ids"]; ok {
		ontology.ModelsIds = val.([]string)
	}

	if val, ok := structure["objects_ids"]; ok {
		ontology.ObjectsIds = val.([]string)
	}

	if val, ok := structure["owner_id"]; ok {
		ontology.OwnerId = val.(string)
	}

}

var OntologyStructure = map[string]metis.Primitive{

	"updated_at": 4,

	"deleted_at": 4,

	"id": 9,

	"created_at": 4,

	"owner_id": 9,

	"models_ids": 10,

	"objects_ids": 10,
}
