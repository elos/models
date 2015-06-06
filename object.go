package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Object struct {
	AttributesIDs []string  `json:"attributes_ids" bson:"attributes_ids"`
	ClassID       string    `json:"class_id" bson:"class_id"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	Id            string    `json:"id" bson:"_id,omitempty"`
	Name          string    `json:"name" bson:"name"`
	OntologyID    string    `json:"ontology_id" bson:"ontology_id"`
	OwnerID       string    `json:"owner_id" bson:"owner_id"`
	RelationsIDs  []string  `json:"relations_ids" bson:"relations_ids"`
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
	object.AttributesIDs = append(object.AttributesIDs, attribute.ID().String())
}

func (object *Object) ExcludeAttribute(attribute *Attribute) {
	tmp := make([]string, 0)
	id := attribute.ID().String()
	for _, s := range object.AttributesIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	object.AttributesIDs = tmp
}

func (object *Object) AttributesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(object.AttributesIDs), db), nil
}

func (object *Object) Attributes(db data.DB) ([]*Attribute, error) {

	attributes := make([]*Attribute, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(object.AttributesIDs), db)
	attribute := NewAttribute()
	for iter.Next(attribute) {
		attributes = append(attributes, attribute)
		attribute = NewAttribute()
	}
	return attributes, nil
}

func (object *Object) SetClass(class *Class) error {
	object.ClassID = class.ID().String()
	return nil
}

func (object *Object) Class(db data.DB) (*Class, error) {
	if object.ClassID == "" {
		return nil, ErrEmptyLink
	}

	class := NewClass()
	pid, _ := mongo.ParseObjectID(object.ClassID)
	class.SetID(data.ID(pid.Hex()))
	return class, db.PopulateByID(class)

}

func (object *Object) ClassOrCreate(db data.DB) (*Class, error) {
	class, err := object.Class(db)

	if err == ErrEmptyLink {
		class := NewClass()
		class.SetID(db.NewID())
		if err := object.SetClass(class); err != nil {
			return nil, err
		}

		if err := db.Save(class); err != nil {
			return nil, err
		}

		if err := db.Save(object); err != nil {
			return nil, err
		}

		return class, nil
	} else {
		return class, err
	}
}

func (object *Object) SetOntology(ontology *Ontology) error {
	object.OntologyID = ontology.ID().String()
	return nil
}

func (object *Object) Ontology(db data.DB) (*Ontology, error) {
	if object.OntologyID == "" {
		return nil, ErrEmptyLink
	}

	ontology := NewOntology()
	pid, _ := mongo.ParseObjectID(object.OntologyID)
	ontology.SetID(data.ID(pid.Hex()))
	return ontology, db.PopulateByID(ontology)

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

func (object *Object) SetOwner(user *User) error {
	object.OwnerID = user.ID().String()
	return nil
}

func (object *Object) Owner(db data.DB) (*User, error) {
	if object.OwnerID == "" {
		return nil, ErrEmptyLink
	}

	user := NewUser()
	pid, _ := mongo.ParseObjectID(object.OwnerID)
	user.SetID(data.ID(pid.Hex()))
	return user, db.PopulateByID(user)

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

func (object *Object) IncludeRelation(relation *Relation) {
	object.RelationsIDs = append(object.RelationsIDs, relation.ID().String())
}

func (object *Object) ExcludeRelation(relation *Relation) {
	tmp := make([]string, 0)
	id := relation.ID().String()
	for _, s := range object.RelationsIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	object.RelationsIDs = tmp
}

func (object *Object) RelationsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(object.RelationsIDs), db), nil
}

func (object *Object) Relations(db data.DB) ([]*Relation, error) {

	relations := make([]*Relation, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(object.RelationsIDs), db)
	relation := NewRelation()
	for iter.Next(relation) {
		relations = append(relations, relation)
		relation = NewRelation()
	}
	return relations, nil
}

// BSON {{{
func (object *Object) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		AttributesIDs []string `json:"attributes_ids" bson:"attributes_ids"`

		ClassID string `json:"class_id" bson:"class_id"`

		OntologyID string `json:"ontology_id" bson:"ontology_id"`

		OwnerID string `json:"owner_id" bson:"owner_id"`

		RelationsIDs []string `json:"relations_ids" bson:"relations_ids"`
	}{

		CreatedAt: object.CreatedAt,

		Name: object.Name,

		UpdatedAt: object.UpdatedAt,

		AttributesIDs: object.AttributesIDs,

		ClassID: object.ClassID,

		OntologyID: object.OntologyID,

		OwnerID: object.OwnerID,

		RelationsIDs: object.RelationsIDs,
	}, nil

}

func (object *Object) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		AttributesIDs []string `json:"attributes_ids" bson:"attributes_ids"`

		ClassID string `json:"class_id" bson:"class_id"`

		OntologyID string `json:"ontology_id" bson:"ontology_id"`

		OwnerID string `json:"owner_id" bson:"owner_id"`

		RelationsIDs []string `json:"relations_ids" bson:"relations_ids"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	object.CreatedAt = tmp.CreatedAt

	object.Id = tmp.Id.Hex()

	object.Name = tmp.Name

	object.UpdatedAt = tmp.UpdatedAt

	object.AttributesIDs = tmp.AttributesIDs

	object.ClassID = tmp.ClassID

	object.OntologyID = tmp.OntologyID

	object.OwnerID = tmp.OwnerID

	object.RelationsIDs = tmp.RelationsIDs

	return nil

}

// BSON }}}
