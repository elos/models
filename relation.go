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
type Relation struct {
	Codomain     string    `json:"codomain" bson:"codomain"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	DeletedAt    time.Time `json:"deleted_at" bson:"deleted_at"`
	Id           string    `json:"id" bson:"_id,omitempty"`
	Inverse      string    `json:"inverse" bson:"inverse"`
	LinksIds     []string  `json:"links_ids" bson:"links_ids"`
	ModelId      string    `json:"model_id" bson:"model_id"`
	Multiplicity string    `json:"multiplicity" bson:"multiplicity"`
	Name         string    `json:"name" bson:"name"`
	OwnerId      string    `json:"owner_id" bson:"owner_id"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

func NewRelation() *Relation {
	return &Relation{}
}

func FindRelation(db data.DB, id data.ID) (*Relation, error) {

	relation := NewRelation()
	relation.SetID(id)

	return relation, db.PopulateByID(relation)

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

func (relation *Relation) IncludeLink(link *Link) {
	otherID := link.ID().String()
	for i := range relation.LinksIds {
		if relation.LinksIds[i] == otherID {
			return
		}
	}
	relation.LinksIds = append(relation.LinksIds, otherID)
}

func (relation *Relation) ExcludeLink(link *Link) {
	tmp := make([]string, 0)
	id := link.ID().String()
	for _, s := range relation.LinksIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	relation.LinksIds = tmp
}

func (relation *Relation) LinksIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(relation.LinksIds), db), nil
}

func (relation *Relation) Links(db data.DB) (links []*Link, err error) {
	links = make([]*Link, len(relation.LinksIds))
	link := NewLink()
	for i, id := range relation.LinksIds {
		link.Id = id
		if err = db.PopulateByID(link); err != nil {
			return
		}

		links[i] = link
		link = NewLink()
	}

	return
}

func (relation *Relation) SetModel(modelArgument *Model) error {
	relation.ModelId = modelArgument.ID().String()
	return nil
}

func (relation *Relation) Model(db data.DB) (*Model, error) {
	if relation.ModelId == "" {
		return nil, ErrEmptyLink
	}

	modelArgument := NewModel()
	pid, _ := mongo.ParseObjectID(relation.ModelId)
	modelArgument.SetID(data.ID(pid.Hex()))
	return modelArgument, db.PopulateByID(modelArgument)

}

func (relation *Relation) ModelOrCreate(db data.DB) (*Model, error) {
	model, err := relation.Model(db)

	if err == ErrEmptyLink {
		model := NewModel()
		model.SetID(db.NewID())
		if err := relation.SetModel(model); err != nil {
			return nil, err
		}

		if err := db.Save(model); err != nil {
			return nil, err
		}

		if err := db.Save(relation); err != nil {
			return nil, err
		}

		return model, nil
	} else {
		return model, err
	}
}

func (relation *Relation) SetOwner(userArgument *User) error {
	relation.OwnerId = userArgument.ID().String()
	return nil
}

func (relation *Relation) Owner(db data.DB) (*User, error) {
	if relation.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	pid, _ := mongo.ParseObjectID(relation.OwnerId)
	userArgument.SetID(data.ID(pid.Hex()))
	return userArgument, db.PopulateByID(userArgument)

}

func (relation *Relation) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := relation.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := relation.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(relation); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (relation *Relation) GetBSON() (interface{}, error) {

	return struct {
		Codomain string `json:"codomain" bson:"codomain"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Inverse string `json:"inverse" bson:"inverse"`

		Multiplicity string `json:"multiplicity" bson:"multiplicity"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		LinksIds []string `json:"links_ids" bson:"links_ids"`

		ModelId string `json:"model_id" bson:"model_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		Codomain: relation.Codomain,

		CreatedAt: relation.CreatedAt,

		DeletedAt: relation.DeletedAt,

		Inverse: relation.Inverse,

		Multiplicity: relation.Multiplicity,

		Name: relation.Name,

		UpdatedAt: relation.UpdatedAt,

		LinksIds: relation.LinksIds,

		ModelId: relation.ModelId,

		OwnerId: relation.OwnerId,
	}, nil

}

func (relation *Relation) SetBSON(raw bson.Raw) error {

	tmp := struct {
		Codomain string `json:"codomain" bson:"codomain"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Inverse string `json:"inverse" bson:"inverse"`

		Multiplicity string `json:"multiplicity" bson:"multiplicity"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		LinksIds []string `json:"links_ids" bson:"links_ids"`

		ModelId string `json:"model_id" bson:"model_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	relation.Codomain = tmp.Codomain

	relation.CreatedAt = tmp.CreatedAt

	relation.DeletedAt = tmp.DeletedAt

	relation.Id = tmp.Id.Hex()

	relation.Inverse = tmp.Inverse

	relation.Multiplicity = tmp.Multiplicity

	relation.Name = tmp.Name

	relation.UpdatedAt = tmp.UpdatedAt

	relation.LinksIds = tmp.LinksIds

	relation.ModelId = tmp.ModelId

	relation.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (relation *Relation) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["name"]; ok {
		relation.Name = val.(string)
	}

	if val, ok := structure["multiplicity"]; ok {
		relation.Multiplicity = val.(string)
	}

	if val, ok := structure["codomain"]; ok {
		relation.Codomain = val.(string)
	}

	if val, ok := structure["inverse"]; ok {
		relation.Inverse = val.(string)
	}

	if val, ok := structure["id"]; ok {
		relation.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		relation.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		relation.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		relation.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["owner_id"]; ok {
		relation.OwnerId = val.(string)
	}

	if val, ok := structure["model_id"]; ok {
		relation.ModelId = val.(string)
	}

	if val, ok := structure["links_ids"]; ok {
		relation.LinksIds = val.([]string)
	}

}

var RelationStructure = map[string]metis.Primitive{

	"inverse": 3,

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"name": 3,

	"multiplicity": 3,

	"codomain": 3,

	"model_id": 9,

	"links_ids": 10,

	"owner_id": 9,
}
