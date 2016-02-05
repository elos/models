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
type User struct {
	AuthorizationsIds []string  `json:"authorizations_ids" bson:"authorizations_ids"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
	CredentialsIds    []string  `json:"credentials_ids" bson:"credentials_ids"`
	DeletedAt         time.Time `json:"deleted_at" bson:"deleted_at"`
	GroupsIds         []string  `json:"groups_ids" bson:"groups_ids"`
	Id                string    `json:"id" bson:"_id,omitempty"`
	Password          string    `json:"password" bson:"password"`
	SessionsIds       []string  `json:"sessions_ids" bson:"sessions_ids"`
	UpdatedAt         time.Time `json:"updated_at" bson:"updated_at"`
}

func NewUser() *User {
	return &User{}
}

func FindUser(db data.DB, id data.ID) (*User, error) {

	user := NewUser()
	user.SetID(id)

	return user, db.PopulateByID(user)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (user *User) Kind() data.Kind {
	return UserKind
}

// just returns itself for now
func (user *User) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = user.ID()
	return foo
}

func (user *User) SetID(id data.ID) {
	user.Id = id.String()
}

func (user *User) ID() data.ID {
	return data.ID(user.Id)
}

func (user *User) IncludeAuthorization(authorization *Group) {
	otherID := authorization.ID().String()
	for i := range user.AuthorizationsIds {
		if user.AuthorizationsIds[i] == otherID {
			return
		}
	}
	user.AuthorizationsIds = append(user.AuthorizationsIds, otherID)
}

func (user *User) ExcludeAuthorization(authorization *Group) {
	tmp := make([]string, 0)
	id := authorization.ID().String()
	for _, s := range user.AuthorizationsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	user.AuthorizationsIds = tmp
}

func (user *User) AuthorizationsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(user.AuthorizationsIds), db), nil
}

func (user *User) Authorizations(db data.DB) (authorizations []*Group, err error) {
	authorizations = make([]*Group, len(user.AuthorizationsIds))
	authorization := NewGroup()
	for i, id := range user.AuthorizationsIds {
		authorization.Id = id
		if err = db.PopulateByID(authorization); err != nil {
			return
		}

		authorizations[i] = authorization
		authorization = NewGroup()
	}

	return
}

func (user *User) IncludeCredential(credential *Credential) {
	otherID := credential.ID().String()
	for i := range user.CredentialsIds {
		if user.CredentialsIds[i] == otherID {
			return
		}
	}
	user.CredentialsIds = append(user.CredentialsIds, otherID)
}

func (user *User) ExcludeCredential(credential *Credential) {
	tmp := make([]string, 0)
	id := credential.ID().String()
	for _, s := range user.CredentialsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	user.CredentialsIds = tmp
}

func (user *User) CredentialsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(user.CredentialsIds), db), nil
}

func (user *User) Credentials(db data.DB) (credentials []*Credential, err error) {
	credentials = make([]*Credential, len(user.CredentialsIds))
	credential := NewCredential()
	for i, id := range user.CredentialsIds {
		credential.Id = id
		if err = db.PopulateByID(credential); err != nil {
			return
		}

		credentials[i] = credential
		credential = NewCredential()
	}

	return
}

func (user *User) IncludeGroup(group *Group) {
	otherID := group.ID().String()
	for i := range user.GroupsIds {
		if user.GroupsIds[i] == otherID {
			return
		}
	}
	user.GroupsIds = append(user.GroupsIds, otherID)
}

func (user *User) ExcludeGroup(group *Group) {
	tmp := make([]string, 0)
	id := group.ID().String()
	for _, s := range user.GroupsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	user.GroupsIds = tmp
}

func (user *User) GroupsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(user.GroupsIds), db), nil
}

func (user *User) Groups(db data.DB) (groups []*Group, err error) {
	groups = make([]*Group, len(user.GroupsIds))
	group := NewGroup()
	for i, id := range user.GroupsIds {
		group.Id = id
		if err = db.PopulateByID(group); err != nil {
			return
		}

		groups[i] = group
		group = NewGroup()
	}

	return
}

func (user *User) IncludeSession(session *Session) {
	otherID := session.ID().String()
	for i := range user.SessionsIds {
		if user.SessionsIds[i] == otherID {
			return
		}
	}
	user.SessionsIds = append(user.SessionsIds, otherID)
}

func (user *User) ExcludeSession(session *Session) {
	tmp := make([]string, 0)
	id := session.ID().String()
	for _, s := range user.SessionsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	user.SessionsIds = tmp
}

func (user *User) SessionsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(user.SessionsIds), db), nil
}

func (user *User) Sessions(db data.DB) (sessions []*Session, err error) {
	sessions = make([]*Session, len(user.SessionsIds))
	session := NewSession()
	for i, id := range user.SessionsIds {
		session.Id = id
		if err = db.PopulateByID(session); err != nil {
			return
		}

		sessions[i] = session
		session = NewSession()
	}

	return
}

// BSON {{{
func (user *User) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Password string `json:"password" bson:"password"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		AuthorizationsIds []string `json:"authorizations_ids" bson:"authorizations_ids"`

		CredentialsIds []string `json:"credentials_ids" bson:"credentials_ids"`

		GroupsIds []string `json:"groups_ids" bson:"groups_ids"`

		SessionsIds []string `json:"sessions_ids" bson:"sessions_ids"`
	}{

		CreatedAt: user.CreatedAt,

		DeletedAt: user.DeletedAt,

		Password: user.Password,

		UpdatedAt: user.UpdatedAt,

		AuthorizationsIds: user.AuthorizationsIds,

		CredentialsIds: user.CredentialsIds,

		GroupsIds: user.GroupsIds,

		SessionsIds: user.SessionsIds,
	}, nil

}

func (user *User) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Password string `json:"password" bson:"password"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		AuthorizationsIds []string `json:"authorizations_ids" bson:"authorizations_ids"`

		CredentialsIds []string `json:"credentials_ids" bson:"credentials_ids"`

		GroupsIds []string `json:"groups_ids" bson:"groups_ids"`

		SessionsIds []string `json:"sessions_ids" bson:"sessions_ids"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	user.CreatedAt = tmp.CreatedAt

	user.DeletedAt = tmp.DeletedAt

	user.Id = tmp.Id.Hex()

	user.Password = tmp.Password

	user.UpdatedAt = tmp.UpdatedAt

	user.AuthorizationsIds = tmp.AuthorizationsIds

	user.CredentialsIds = tmp.CredentialsIds

	user.GroupsIds = tmp.GroupsIds

	user.SessionsIds = tmp.SessionsIds

	return nil

}

// BSON }}}

func (user *User) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		user.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		user.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		user.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		user.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["password"]; ok {
		user.Password = val.(string)
	}

	if val, ok := structure["authorizations_ids"]; ok {
		user.AuthorizationsIds = val.([]string)
	}

	if val, ok := structure["sessions_ids"]; ok {
		user.SessionsIds = val.([]string)
	}

	if val, ok := structure["credentials_ids"]; ok {
		user.CredentialsIds = val.([]string)
	}

	if val, ok := structure["groups_ids"]; ok {
		user.GroupsIds = val.([]string)
	}

}

var UserStructure = map[string]metis.Primitive{

	"deleted_at": 4,

	"password": 3,

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"groups_ids": 10,

	"authorizations_ids": 10,

	"sessions_ids": 10,

	"credentials_ids": 10,
}
