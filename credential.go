package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Credential struct {
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	DeletedAt   time.Time `json:"deleted_at" bson:"deleted_at"`
	Id          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	OwnerID     string    `json:"owner_id" bson:"owner_id"`
	Private     string    `json:"private" bson:"private"`
	Public      string    `json:"public" bson:"public"`
	SessionsIDs []string  `json:"sessions_ids" bson:"sessions_ids"`
	Spec        string    `json:"spec" bson:"spec"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func NewCredential() *Credential {
	return &Credential{}
}

func FindCredential(db data.DB, id data.ID) (*Credential, error) {

	credential := NewCredential()
	credential.SetID(id)

	return credential, db.PopulateByID(credential)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (credential *Credential) Kind() data.Kind {
	return CredentialKind
}

// just returns itself for now
func (credential *Credential) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = credential.ID()
	return foo
}

func (credential *Credential) SetID(id data.ID) {
	credential.Id = id.String()
}

func (credential *Credential) ID() data.ID {
	return data.ID(credential.Id)
}

func (credential *Credential) SetOwner(user *User) error {
	credential.OwnerID = user.ID().String()
	return nil
}

func (credential *Credential) Owner(db data.DB) (*User, error) {
	if credential.OwnerID == "" {
		return nil, ErrEmptyLink
	}

	user := NewUser()
	pid, _ := mongo.ParseObjectID(credential.OwnerID)
	user.SetID(data.ID(pid.Hex()))
	return user, db.PopulateByID(user)

}

func (credential *Credential) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := credential.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := credential.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(credential); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

func (credential *Credential) IncludeSession(session *Session) {
	credential.SessionsIDs = append(credential.SessionsIDs, session.ID().String())
}

func (credential *Credential) ExcludeSession(session *Session) {
	tmp := make([]string, 0)
	id := session.ID().String()
	for _, s := range credential.SessionsIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	credential.SessionsIDs = tmp
}

func (credential *Credential) SessionsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(credential.SessionsIDs), db), nil
}

func (credential *Credential) Sessions(db data.DB) ([]*Session, error) {

	sessions := make([]*Session, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(credential.SessionsIDs), db)
	session := NewSession()
	for iter.Next(session) {
		sessions = append(sessions, session)
		session = NewSession()
	}
	return sessions, nil
}

// BSON {{{
func (credential *Credential) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		Private string `json:"private" bson:"private"`

		Public string `json:"public" bson:"public"`

		Spec string `json:"spec" bson:"spec"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerID string `json:"owner_id" bson:"owner_id"`

		SessionsIDs []string `json:"sessions_ids" bson:"sessions_ids"`
	}{

		CreatedAt: credential.CreatedAt,

		DeletedAt: credential.DeletedAt,

		Name: credential.Name,

		Private: credential.Private,

		Public: credential.Public,

		Spec: credential.Spec,

		UpdatedAt: credential.UpdatedAt,

		OwnerID: credential.OwnerID,

		SessionsIDs: credential.SessionsIDs,
	}, nil

}

func (credential *Credential) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		Private string `json:"private" bson:"private"`

		Public string `json:"public" bson:"public"`

		Spec string `json:"spec" bson:"spec"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerID string `json:"owner_id" bson:"owner_id"`

		SessionsIDs []string `json:"sessions_ids" bson:"sessions_ids"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	credential.CreatedAt = tmp.CreatedAt

	credential.DeletedAt = tmp.DeletedAt

	credential.Id = tmp.Id.Hex()

	credential.Name = tmp.Name

	credential.Private = tmp.Private

	credential.Public = tmp.Public

	credential.Spec = tmp.Spec

	credential.UpdatedAt = tmp.UpdatedAt

	credential.OwnerID = tmp.OwnerID

	credential.SessionsIDs = tmp.SessionsIDs

	return nil

}

// BSON }}}
