package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Session struct {
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	CredentialId string    `json:"credential_id" bson:"credential_id"`
	DeletedAt    time.Time `json:"deleted_at" bson:"deleted_at"`
	ExpiresAfter int       `json:"expires_after" bson:"expires_after"`
	Id           string    `json:"id" bson:"_id,omitempty"`
	OwnerId      string    `json:"owner_id" bson:"owner_id"`
	Token        string    `json:"token" bson:"token"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

func NewSession() *Session {
	return &Session{}
}

func FindSession(db data.DB, id data.ID) (*Session, error) {

	session := NewSession()
	session.SetID(id)

	return session, db.PopulateByID(session)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (session *Session) Kind() data.Kind {
	return SessionKind
}

// just returns itself for now
func (session *Session) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = session.ID()
	return foo
}

func (session *Session) SetID(id data.ID) {
	session.Id = id.String()
}

func (session *Session) ID() data.ID {
	return data.ID(session.Id)
}

func (session *Session) SetCredential(credentialArgument *Credential) error {
	session.CredentialId = credentialArgument.ID().String()
	return nil
}

func (session *Session) Credential(db data.DB) (*Credential, error) {
	if session.CredentialId == "" {
		return nil, ErrEmptyLink
	}

	credentialArgument := NewCredential()
	id, _ := db.ParseID(session.CredentialId)
	credentialArgument.SetID(id)
	return credentialArgument, db.PopulateByID(credentialArgument)

}

func (session *Session) CredentialOrCreate(db data.DB) (*Credential, error) {
	credential, err := session.Credential(db)

	if err == ErrEmptyLink {
		credential := NewCredential()
		credential.SetID(db.NewID())
		if err := session.SetCredential(credential); err != nil {
			return nil, err
		}

		if err := db.Save(credential); err != nil {
			return nil, err
		}

		if err := db.Save(session); err != nil {
			return nil, err
		}

		return credential, nil
	} else {
		return credential, err
	}
}

func (session *Session) SetOwner(userArgument *User) error {
	session.OwnerId = userArgument.ID().String()
	return nil
}

func (session *Session) Owner(db data.DB) (*User, error) {
	if session.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(session.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (session *Session) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := session.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := session.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(session); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (session *Session) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		ExpiresAfter int `json:"expires_after" bson:"expires_after"`

		Id string `json:"id" bson:"_id,omitempty"`

		Token string `json:"token" bson:"token"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		CredentialId string `json:"credential_id" bson:"credential_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: session.CreatedAt,

		DeletedAt: session.DeletedAt,

		ExpiresAfter: session.ExpiresAfter,

		Token: session.Token,

		UpdatedAt: session.UpdatedAt,

		CredentialId: session.CredentialId,

		OwnerId: session.OwnerId,
	}, nil

}

func (session *Session) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		ExpiresAfter int `json:"expires_after" bson:"expires_after"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Token string `json:"token" bson:"token"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		CredentialId string `json:"credential_id" bson:"credential_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	session.CreatedAt = tmp.CreatedAt

	session.DeletedAt = tmp.DeletedAt

	session.ExpiresAfter = tmp.ExpiresAfter

	session.Id = tmp.Id.Hex()

	session.Token = tmp.Token

	session.UpdatedAt = tmp.UpdatedAt

	session.CredentialId = tmp.CredentialId

	session.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (session *Session) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		session.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		session.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		session.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		session.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["token"]; ok {
		session.Token = val.(string)
	}

	if val, ok := structure["expires_after"]; ok {
		session.ExpiresAfter = val.(int)
	}

	if val, ok := structure["owner_id"]; ok {
		session.OwnerId = val.(string)
	}

	if val, ok := structure["credential_id"]; ok {
		session.CredentialId = val.(string)
	}

}

var SessionStructure = map[string]metis.Primitive{

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"token": 3,

	"expires_after": 1,

	"credential_id": 9,

	"owner_id": 9,
}
