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
type Event struct {
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	DeletedAt  time.Time `json:"deleted_at" bson:"deleted_at"`
	Id         string    `json:"id" bson:"_id,omitempty"`
	LocationId string    `json:"location_id" bson:"location_id"`
	MediaId    string    `json:"media_id" bson:"media_id"`
	Name       string    `json:"name" bson:"name"`
	NoteId     string    `json:"note_id" bson:"note_id"`
	OwnerId    string    `json:"owner_id" bson:"owner_id"`
	PriorId    string    `json:"prior_id" bson:"prior_id"`
	QuantityId string    `json:"quantity_id" bson:"quantity_id"`
	TagsIds    []string  `json:"tags_ids" bson:"tags_ids"`
	Time       time.Time `json:"time" bson:"time"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

func NewEvent() *Event {
	return &Event{}
}

func FindEvent(db data.DB, id data.ID) (*Event, error) {

	event := NewEvent()
	event.SetID(id)

	return event, db.PopulateByID(event)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (event *Event) Kind() data.Kind {
	return EventKind
}

// just returns itself for now
func (event *Event) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = event.ID()
	return foo
}

func (event *Event) SetID(id data.ID) {
	event.Id = id.String()
}

func (event *Event) ID() data.ID {
	return data.ID(event.Id)
}

func (event *Event) SetLocation(locationArgument *Location) error {
	event.LocationId = locationArgument.ID().String()
	return nil
}

func (event *Event) Location(db data.DB) (*Location, error) {
	if event.LocationId == "" {
		return nil, ErrEmptyLink
	}

	locationArgument := NewLocation()
	id, _ := db.ParseID(event.LocationId)
	locationArgument.SetID(id)
	return locationArgument, db.PopulateByID(locationArgument)

}

func (event *Event) LocationOrCreate(db data.DB) (*Location, error) {
	location, err := event.Location(db)

	if err == ErrEmptyLink {
		location := NewLocation()
		location.SetID(db.NewID())
		if err := event.SetLocation(location); err != nil {
			return nil, err
		}

		if err := db.Save(location); err != nil {
			return nil, err
		}

		if err := db.Save(event); err != nil {
			return nil, err
		}

		return location, nil
	} else {
		return location, err
	}
}

func (event *Event) SetMedia(mediaArgument *Media) error {
	event.MediaId = mediaArgument.ID().String()
	return nil
}

func (event *Event) Media(db data.DB) (*Media, error) {
	if event.MediaId == "" {
		return nil, ErrEmptyLink
	}

	mediaArgument := NewMedia()
	id, _ := db.ParseID(event.MediaId)
	mediaArgument.SetID(id)
	return mediaArgument, db.PopulateByID(mediaArgument)

}

func (event *Event) MediaOrCreate(db data.DB) (*Media, error) {
	media, err := event.Media(db)

	if err == ErrEmptyLink {
		media := NewMedia()
		media.SetID(db.NewID())
		if err := event.SetMedia(media); err != nil {
			return nil, err
		}

		if err := db.Save(media); err != nil {
			return nil, err
		}

		if err := db.Save(event); err != nil {
			return nil, err
		}

		return media, nil
	} else {
		return media, err
	}
}

func (event *Event) SetNote(noteArgument *Note) error {
	event.NoteId = noteArgument.ID().String()
	return nil
}

func (event *Event) Note(db data.DB) (*Note, error) {
	if event.NoteId == "" {
		return nil, ErrEmptyLink
	}

	noteArgument := NewNote()
	id, _ := db.ParseID(event.NoteId)
	noteArgument.SetID(id)
	return noteArgument, db.PopulateByID(noteArgument)

}

func (event *Event) NoteOrCreate(db data.DB) (*Note, error) {
	note, err := event.Note(db)

	if err == ErrEmptyLink {
		note := NewNote()
		note.SetID(db.NewID())
		if err := event.SetNote(note); err != nil {
			return nil, err
		}

		if err := db.Save(note); err != nil {
			return nil, err
		}

		if err := db.Save(event); err != nil {
			return nil, err
		}

		return note, nil
	} else {
		return note, err
	}
}

func (event *Event) SetOwner(userArgument *User) error {
	event.OwnerId = userArgument.ID().String()
	return nil
}

func (event *Event) Owner(db data.DB) (*User, error) {
	if event.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(event.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (event *Event) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := event.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := event.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(event); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

func (event *Event) SetPrior(eventArgument *Event) error {
	event.PriorId = eventArgument.ID().String()
	return nil
}

func (event *Event) Prior(db data.DB) (*Event, error) {
	if event.PriorId == "" {
		return nil, ErrEmptyLink
	}

	eventArgument := NewEvent()
	id, _ := db.ParseID(event.PriorId)
	eventArgument.SetID(id)
	return eventArgument, db.PopulateByID(eventArgument)

}

func (event *Event) PriorOrCreate(db data.DB) (*Event, error) {
	event, err := event.Prior(db)

	if err == ErrEmptyLink {
		event := NewEvent()
		event.SetID(db.NewID())
		if err := event.SetPrior(event); err != nil {
			return nil, err
		}

		if err := db.Save(event); err != nil {
			return nil, err
		}

		if err := db.Save(event); err != nil {
			return nil, err
		}

		return event, nil
	} else {
		return event, err
	}
}

func (event *Event) SetQuantity(quantityArgument *Quantity) error {
	event.QuantityId = quantityArgument.ID().String()
	return nil
}

func (event *Event) Quantity(db data.DB) (*Quantity, error) {
	if event.QuantityId == "" {
		return nil, ErrEmptyLink
	}

	quantityArgument := NewQuantity()
	id, _ := db.ParseID(event.QuantityId)
	quantityArgument.SetID(id)
	return quantityArgument, db.PopulateByID(quantityArgument)

}

func (event *Event) QuantityOrCreate(db data.DB) (*Quantity, error) {
	quantity, err := event.Quantity(db)

	if err == ErrEmptyLink {
		quantity := NewQuantity()
		quantity.SetID(db.NewID())
		if err := event.SetQuantity(quantity); err != nil {
			return nil, err
		}

		if err := db.Save(quantity); err != nil {
			return nil, err
		}

		if err := db.Save(event); err != nil {
			return nil, err
		}

		return quantity, nil
	} else {
		return quantity, err
	}
}

func (event *Event) IncludeTag(tag *Tag) {
	otherID := tag.ID().String()
	for i := range event.TagsIds {
		if event.TagsIds[i] == otherID {
			return
		}
	}
	event.TagsIds = append(event.TagsIds, otherID)
}

func (event *Event) ExcludeTag(tag *Tag) {
	tmp := make([]string, 0)
	id := tag.ID().String()
	for _, s := range event.TagsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	event.TagsIds = tmp
}

func (event *Event) TagsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(event.TagsIds), db), nil
}

func (event *Event) Tags(db data.DB) (tags []*Tag, err error) {
	tags = make([]*Tag, len(event.TagsIds))
	tag := NewTag()
	for i, id := range event.TagsIds {
		tag.Id = id
		if err = db.PopulateByID(tag); err != nil {
			return
		}

		tags[i] = tag
		tag = NewTag()
	}

	return
}

// BSON {{{
func (event *Event) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		Time time.Time `json:"time" bson:"time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		LocationId string `json:"location_id" bson:"location_id"`

		MediaId string `json:"media_id" bson:"media_id"`

		NoteId string `json:"note_id" bson:"note_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PriorId string `json:"prior_id" bson:"prior_id"`

		QuantityId string `json:"quantity_id" bson:"quantity_id"`

		TagsIds []string `json:"tags_ids" bson:"tags_ids"`
	}{

		CreatedAt: event.CreatedAt,

		DeletedAt: event.DeletedAt,

		Name: event.Name,

		Time: event.Time,

		UpdatedAt: event.UpdatedAt,

		LocationId: event.LocationId,

		MediaId: event.MediaId,

		NoteId: event.NoteId,

		OwnerId: event.OwnerId,

		PriorId: event.PriorId,

		QuantityId: event.QuantityId,

		TagsIds: event.TagsIds,
	}, nil

}

func (event *Event) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		Time time.Time `json:"time" bson:"time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		LocationId string `json:"location_id" bson:"location_id"`

		MediaId string `json:"media_id" bson:"media_id"`

		NoteId string `json:"note_id" bson:"note_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`

		PriorId string `json:"prior_id" bson:"prior_id"`

		QuantityId string `json:"quantity_id" bson:"quantity_id"`

		TagsIds []string `json:"tags_ids" bson:"tags_ids"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	event.CreatedAt = tmp.CreatedAt

	event.DeletedAt = tmp.DeletedAt

	event.Id = tmp.Id.Hex()

	event.Name = tmp.Name

	event.Time = tmp.Time

	event.UpdatedAt = tmp.UpdatedAt

	event.LocationId = tmp.LocationId

	event.MediaId = tmp.MediaId

	event.NoteId = tmp.NoteId

	event.OwnerId = tmp.OwnerId

	event.PriorId = tmp.PriorId

	event.QuantityId = tmp.QuantityId

	event.TagsIds = tmp.TagsIds

	return nil

}

// BSON }}}

func (event *Event) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["updated_at"]; ok {
		event.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		event.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["name"]; ok {
		event.Name = val.(string)
	}

	if val, ok := structure["time"]; ok {
		event.Time = val.(time.Time)
	}

	if val, ok := structure["id"]; ok {
		event.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		event.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["note_id"]; ok {
		event.NoteId = val.(string)
	}

	if val, ok := structure["location_id"]; ok {
		event.LocationId = val.(string)
	}

	if val, ok := structure["tags_ids"]; ok {
		event.TagsIds = val.([]string)
	}

	if val, ok := structure["media_id"]; ok {
		event.MediaId = val.(string)
	}

	if val, ok := structure["owner_id"]; ok {
		event.OwnerId = val.(string)
	}

	if val, ok := structure["prior_id"]; ok {
		event.PriorId = val.(string)
	}

	if val, ok := structure["quantity_id"]; ok {
		event.QuantityId = val.(string)
	}

}

var EventStructure = map[string]metis.Primitive{

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"name": 3,

	"time": 4,

	"id": 9,

	"tags_ids": 10,

	"media_id": 9,

	"owner_id": 9,

	"prior_id": 9,

	"quantity_id": 9,

	"note_id": 9,

	"location_id": 9,
}
