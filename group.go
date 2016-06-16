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
type Group struct {
	Access      int       `json:"access" bson:"access"`
	ContextsIds []string  `json:"contexts_ids" bson:"contexts_ids"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	DeletedAt   time.Time `json:"deleted_at" bson:"deleted_at"`
	GranteesIds []string  `json:"grantees_ids" bson:"grantees_ids"`
	Id          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	OwnerId     string    `json:"owner_id" bson:"owner_id"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func NewGroup() *Group {
	return &Group{}
}

func FindGroup(db data.DB, id data.ID) (*Group, error) {

	group := NewGroup()
	group.SetID(id)

	return group, db.PopulateByID(group)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (group *Group) Kind() data.Kind {
	return GroupKind
}

// just returns itself for now
func (group *Group) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = group.ID()
	return foo
}

func (group *Group) SetID(id data.ID) {
	group.Id = id.String()
}

func (group *Group) ID() data.ID {
	return data.ID(group.Id)
}

func (group *Group) IncludeContext(context *Context) {
	otherID := context.ID().String()
	for i := range group.ContextsIds {
		if group.ContextsIds[i] == otherID {
			return
		}
	}
	group.ContextsIds = append(group.ContextsIds, otherID)
}

func (group *Group) ExcludeContext(context *Context) {
	tmp := make([]string, 0)
	id := context.ID().String()
	for _, s := range group.ContextsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	group.ContextsIds = tmp
}

func (group *Group) ContextsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(group.ContextsIds), db), nil
}

func (group *Group) Contexts(db data.DB) (contexts []*Context, err error) {
	contexts = make([]*Context, len(group.ContextsIds))
	context := NewContext()
	for i, id := range group.ContextsIds {
		context.Id = id
		if err = db.PopulateByID(context); err != nil {
			return
		}

		contexts[i] = context
		context = NewContext()
	}

	return
}

func (group *Group) IncludeGrantee(grantee *User) {
	otherID := grantee.ID().String()
	for i := range group.GranteesIds {
		if group.GranteesIds[i] == otherID {
			return
		}
	}
	group.GranteesIds = append(group.GranteesIds, otherID)
}

func (group *Group) ExcludeGrantee(grantee *User) {
	tmp := make([]string, 0)
	id := grantee.ID().String()
	for _, s := range group.GranteesIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	group.GranteesIds = tmp
}

func (group *Group) GranteesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(group.GranteesIds), db), nil
}

func (group *Group) Grantees(db data.DB) (grantees []*User, err error) {
	grantees = make([]*User, len(group.GranteesIds))
	grantee := NewUser()
	for i, id := range group.GranteesIds {
		grantee.Id = id
		if err = db.PopulateByID(grantee); err != nil {
			return
		}

		grantees[i] = grantee
		grantee = NewUser()
	}

	return
}

func (group *Group) SetOwner(userArgument *User) error {
	group.OwnerId = userArgument.ID().String()
	return nil
}

func (group *Group) Owner(db data.DB) (*User, error) {
	if group.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(group.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (group *Group) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := group.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := group.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(group); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (group *Group) GetBSON() (interface{}, error) {

	return struct {
		Access int `json:"access" bson:"access"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ContextsIds []string `json:"contexts_ids" bson:"contexts_ids"`

		GranteesIds []string `json:"grantees_ids" bson:"grantees_ids"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		Access: group.Access,

		CreatedAt: group.CreatedAt,

		DeletedAt: group.DeletedAt,

		Name: group.Name,

		UpdatedAt: group.UpdatedAt,

		ContextsIds: group.ContextsIds,

		GranteesIds: group.GranteesIds,

		OwnerId: group.OwnerId,
	}, nil

}

func (group *Group) SetBSON(raw bson.Raw) error {

	tmp := struct {
		Access int `json:"access" bson:"access"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ContextsIds []string `json:"contexts_ids" bson:"contexts_ids"`

		GranteesIds []string `json:"grantees_ids" bson:"grantees_ids"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	group.Access = tmp.Access

	group.CreatedAt = tmp.CreatedAt

	group.DeletedAt = tmp.DeletedAt

	group.Id = tmp.Id.Hex()

	group.Name = tmp.Name

	group.UpdatedAt = tmp.UpdatedAt

	group.ContextsIds = tmp.ContextsIds

	group.GranteesIds = tmp.GranteesIds

	group.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (group *Group) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		group.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		group.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		group.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		group.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["name"]; ok {
		group.Name = val.(string)
	}

	if val, ok := structure["access"]; ok {
		group.Access = val.(int)
	}

	if val, ok := structure["owner_id"]; ok {
		group.OwnerId = val.(string)
	}

	if val, ok := structure["grantees_ids"]; ok {
		group.GranteesIds = val.([]string)
	}

	if val, ok := structure["contexts_ids"]; ok {
		group.ContextsIds = val.([]string)
	}

}

var GroupStructure = map[string]metis.Primitive{

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"name": 3,

	"access": 1,

	"id": 9,

	"contexts_ids": 10,

	"owner_id": 9,

	"grantees_ids": 10,
}
