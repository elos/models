package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/metis"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Media struct {
	Codec     string    `json:"codec" bson:"codec"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
	Id        string    `json:"id" bson:"_id,omitempty"`
	OwnerId   string    `json:"owner_id" bson:"owner_id"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewMedia() *Media {
	return &Media{}
}

func FindMedia(db data.DB, id data.ID) (*Media, error) {

	media := NewMedia()
	media.SetID(id)

	return media, db.PopulateByID(media)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (media *Media) Kind() data.Kind {
	return MediaKind
}

// just returns itself for now
func (media *Media) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = media.ID()
	return foo
}

func (media *Media) SetID(id data.ID) {
	media.Id = id.String()
}

func (media *Media) ID() data.ID {
	return data.ID(media.Id)
}

func (media *Media) SetOwner(userArgument *User) error {
	media.OwnerId = userArgument.ID().String()
	return nil
}

func (media *Media) Owner(db data.DB) (*User, error) {
	if media.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(media.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (media *Media) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := media.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := media.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(media); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (media *Media) GetBSON() (interface{}, error) {

	return struct {
		Codec string `json:"codec" bson:"codec"`

		Content string `json:"content" bson:"content"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		Codec: media.Codec,

		Content: media.Content,

		CreatedAt: media.CreatedAt,

		DeletedAt: media.DeletedAt,

		UpdatedAt: media.UpdatedAt,

		OwnerId: media.OwnerId,
	}, nil

}

func (media *Media) SetBSON(raw bson.Raw) error {

	tmp := struct {
		Codec string `json:"codec" bson:"codec"`

		Content string `json:"content" bson:"content"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	media.Codec = tmp.Codec

	media.Content = tmp.Content

	media.CreatedAt = tmp.CreatedAt

	media.DeletedAt = tmp.DeletedAt

	media.Id = tmp.Id.Hex()

	media.UpdatedAt = tmp.UpdatedAt

	media.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (media *Media) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["created_at"]; ok {
		media.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		media.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		media.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["content"]; ok {
		media.Content = val.(string)
	}

	if val, ok := structure["codec"]; ok {
		media.Codec = val.(string)
	}

	if val, ok := structure["id"]; ok {
		media.Id = val.(string)
	}

	if val, ok := structure["owner_id"]; ok {
		media.OwnerId = val.(string)
	}

}

var MediaStructure = map[string]metis.Primitive{

	"codec": 3,

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"content": 3,

	"owner_id": 9,
}
