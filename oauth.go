package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Oauth struct {
	AccessToken  string    `json:"access_token" bson:"access_token"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	DeletedAt    time.Time `json:"deleted_at" bson:"deleted_at"`
	Expiry       time.Time `json:"expiry" bson:"expiry"`
	Id           string    `json:"id" bson:"_id,omitempty"`
	OwnerId      string    `json:"owner_id" bson:"owner_id"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token"`
	TokenType    string    `json:"token_type" bson:"token_type"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

func NewOauth() *Oauth {
	return &Oauth{}
}

func FindOauth(db data.DB, id data.ID) (*Oauth, error) {

	oauth := NewOauth()
	oauth.SetID(id)

	return oauth, db.PopulateByID(oauth)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (oauth *Oauth) Kind() data.Kind {
	return OauthKind
}

// just returns itself for now
func (oauth *Oauth) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = oauth.ID()
	return foo
}

func (oauth *Oauth) SetID(id data.ID) {
	oauth.Id = id.String()
}

func (oauth *Oauth) ID() data.ID {
	return data.ID(oauth.Id)
}

func (oauth *Oauth) SetOwner(userArgument *User) error {
	oauth.OwnerId = userArgument.ID().String()
	return nil
}

func (oauth *Oauth) Owner(db data.DB) (*User, error) {
	if oauth.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(oauth.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (oauth *Oauth) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := oauth.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := oauth.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(oauth); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (oauth *Oauth) GetBSON() (interface{}, error) {

	return struct {
		AccessToken string `json:"access_token" bson:"access_token"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Expiry time.Time `json:"expiry" bson:"expiry"`

		Id string `json:"id" bson:"_id,omitempty"`

		RefreshToken string `json:"refresh_token" bson:"refresh_token"`

		TokenType string `json:"token_type" bson:"token_type"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		AccessToken: oauth.AccessToken,

		CreatedAt: oauth.CreatedAt,

		DeletedAt: oauth.DeletedAt,

		Expiry: oauth.Expiry,

		RefreshToken: oauth.RefreshToken,

		TokenType: oauth.TokenType,

		UpdatedAt: oauth.UpdatedAt,

		OwnerId: oauth.OwnerId,
	}, nil

}

func (oauth *Oauth) SetBSON(raw bson.Raw) error {

	tmp := struct {
		AccessToken string `json:"access_token" bson:"access_token"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Expiry time.Time `json:"expiry" bson:"expiry"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		RefreshToken string `json:"refresh_token" bson:"refresh_token"`

		TokenType string `json:"token_type" bson:"token_type"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	oauth.AccessToken = tmp.AccessToken

	oauth.CreatedAt = tmp.CreatedAt

	oauth.DeletedAt = tmp.DeletedAt

	oauth.Expiry = tmp.Expiry

	oauth.Id = tmp.Id.Hex()

	oauth.RefreshToken = tmp.RefreshToken

	oauth.TokenType = tmp.TokenType

	oauth.UpdatedAt = tmp.UpdatedAt

	oauth.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (oauth *Oauth) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["updated_at"]; ok {
		oauth.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		oauth.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["access_token"]; ok {
		oauth.AccessToken = val.(string)
	}

	if val, ok := structure["token_type"]; ok {
		oauth.TokenType = val.(string)
	}

	if val, ok := structure["refresh_token"]; ok {
		oauth.RefreshToken = val.(string)
	}

	if val, ok := structure["expiry"]; ok {
		oauth.Expiry = val.(time.Time)
	}

	if val, ok := structure["id"]; ok {
		oauth.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		oauth.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["owner_id"]; ok {
		oauth.OwnerId = val.(string)
	}

}

var OauthStructure = map[string]metis.Primitive{

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"access_token": 3,

	"token_type": 3,

	"refresh_token": 3,

	"expiry": 4,

	"id": 9,

	"owner_id": 9,
}