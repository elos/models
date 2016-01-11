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
type Note struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
	Id        string    `json:"id" bson:"_id,omitempty"`
	OwnerId   string    `json:"owner_id" bson:"owner_id"`
	Text      string    `json:"text" bson:"text"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewNote() *Note {
	return &Note{}
}

func FindNote(db data.DB, id data.ID) (*Note, error) {

	note := NewNote()
	note.SetID(id)

	return note, db.PopulateByID(note)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (note *Note) Kind() data.Kind {
	return NoteKind
}

// just returns itself for now
func (note *Note) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = note.ID()
	return foo
}

func (note *Note) SetID(id data.ID) {
	note.Id = id.String()
}

func (note *Note) ID() data.ID {
	return data.ID(note.Id)
}

func (note *Note) SetOwner(userArgument *User) error {
	note.OwnerId = userArgument.ID().String()
	return nil
}

func (note *Note) Owner(db data.DB) (*User, error) {
	if note.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	pid, _ := mongo.ParseObjectID(note.OwnerId)
	userArgument.SetID(data.ID(pid.Hex()))
	return userArgument, db.PopulateByID(userArgument)

}

func (note *Note) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := note.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := note.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(note); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (note *Note) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Text string `json:"text" bson:"text"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: note.CreatedAt,

		DeletedAt: note.DeletedAt,

		Text: note.Text,

		UpdatedAt: note.UpdatedAt,

		OwnerId: note.OwnerId,
	}, nil

}

func (note *Note) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Text string `json:"text" bson:"text"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	note.CreatedAt = tmp.CreatedAt

	note.DeletedAt = tmp.DeletedAt

	note.Id = tmp.Id.Hex()

	note.Text = tmp.Text

	note.UpdatedAt = tmp.UpdatedAt

	note.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (note *Note) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		note.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		note.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		note.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		note.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["text"]; ok {
		note.Text = val.(string)
	}

	if val, ok := structure["owner_id"]; ok {
		note.OwnerId = val.(string)
	}

}

var NoteStructure = map[string]metis.Primitive{

	"deleted_at": 4,

	"text": 3,

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"owner_id": 9,
}
