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
type Fixture struct {
	ActionableId   string      `json:"actionable_id" bson:"actionable_id"`
	ActionableKind string      `json:"actionable_kind" bson:"actionable_kind"`
	ActionsIds     []string    `json:"actions_ids" bson:"actions_ids"`
	CreatedAt      time.Time   `json:"created_at" bson:"created_at"`
	DeletedAt      time.Time   `json:"deleted_at" bson:"deleted_at"`
	EndTime        time.Time   `json:"end_time" bson:"end_time"`
	EventableId    string      `json:"eventable_id" bson:"eventable_id"`
	EventableKind  string      `json:"eventable_kind" bson:"eventable_kind"`
	EventsIds      []string    `json:"events_ids" bson:"events_ids"`
	Exceptions     []time.Time `json:"exceptions" bson:"exceptions"`
	ExpiresAt      time.Time   `json:"expires_at" bson:"expires_at"`
	Id             string      `json:"id" bson:"_id,omitempty"`
	Label          bool        `json:"label" bson:"label"`
	Name           string      `json:"name" bson:"name"`
	OwnerId        string      `json:"owner_id" bson:"owner_id"`
	Rank           int         `json:"rank" bson:"rank"`
	StartTime      time.Time   `json:"start_time" bson:"start_time"`
	UpdatedAt      time.Time   `json:"updated_at" bson:"updated_at"`
}

func NewFixture() *Fixture {
	return &Fixture{}
}

func FindFixture(db data.DB, id data.ID) (*Fixture, error) {

	fixture := NewFixture()
	fixture.SetID(id)

	return fixture, db.PopulateByID(fixture)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (fixture *Fixture) Kind() data.Kind {
	return FixtureKind
}

// just returns itself for now
func (fixture *Fixture) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = fixture.ID()
	return foo
}

func (fixture *Fixture) SetID(id data.ID) {
	fixture.Id = id.String()
}

func (fixture *Fixture) ID() data.ID {
	return data.ID(fixture.Id)
}

func (fixture *Fixture) SetActionable(actionableArgument Actionable) error {
	fixture.ActionableId = actionableArgument.ID().String()
	return nil
}

func (fixture *Fixture) Actionable(db data.DB) (Actionable, error) {
	if fixture.ActionableId == "" {
		return nil, ErrEmptyLink
	}

	m := ModelFor(data.Kind(fixture.ActionableKind))
	actionable := m.(Actionable)

	pid, _ := mongo.ParseObjectID(fixture.ActionableId)

	actionable.SetID(data.ID(pid.Hex()))
	return actionable, db.PopulateByID(actionable)

}

func (fixture *Fixture) IncludeAction(action *Action) {
	otherID := action.ID().String()
	for i := range fixture.ActionsIds {
		if fixture.ActionsIds[i] == otherID {
			return
		}
	}
	fixture.ActionsIds = append(fixture.ActionsIds, otherID)
}

func (fixture *Fixture) ExcludeAction(action *Action) {
	tmp := make([]string, 0)
	id := action.ID().String()
	for _, s := range fixture.ActionsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	fixture.ActionsIds = tmp
}

func (fixture *Fixture) ActionsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(fixture.ActionsIds), db), nil
}

func (fixture *Fixture) Actions(db data.DB) (actions []*Action, err error) {
	actions = make([]*Action, len(fixture.ActionsIds))
	action := NewAction()
	for i, id := range fixture.ActionsIds {
		action.Id = id
		if err = db.PopulateByID(action); err != nil {
			return
		}

		actions[i] = action
		action = NewAction()
	}

	return
}

func (fixture *Fixture) SetEventable(eventableArgument Eventable) error {
	fixture.EventableId = eventableArgument.ID().String()
	return nil
}

func (fixture *Fixture) Eventable(db data.DB) (Eventable, error) {
	if fixture.EventableId == "" {
		return nil, ErrEmptyLink
	}

	m := ModelFor(data.Kind(fixture.EventableKind))
	eventable := m.(Eventable)

	pid, _ := mongo.ParseObjectID(fixture.EventableId)

	eventable.SetID(data.ID(pid.Hex()))
	return eventable, db.PopulateByID(eventable)

}

func (fixture *Fixture) IncludeEvent(event *Event) {
	otherID := event.ID().String()
	for i := range fixture.EventsIds {
		if fixture.EventsIds[i] == otherID {
			return
		}
	}
	fixture.EventsIds = append(fixture.EventsIds, otherID)
}

func (fixture *Fixture) ExcludeEvent(event *Event) {
	tmp := make([]string, 0)
	id := event.ID().String()
	for _, s := range fixture.EventsIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	fixture.EventsIds = tmp
}

func (fixture *Fixture) EventsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(fixture.EventsIds), db), nil
}

func (fixture *Fixture) Events(db data.DB) (events []*Event, err error) {
	events = make([]*Event, len(fixture.EventsIds))
	event := NewEvent()
	for i, id := range fixture.EventsIds {
		event.Id = id
		if err = db.PopulateByID(event); err != nil {
			return
		}

		events[i] = event
		event = NewEvent()
	}

	return
}

func (fixture *Fixture) SetOwner(userArgument *User) error {
	fixture.OwnerId = userArgument.ID().String()
	return nil
}

func (fixture *Fixture) Owner(db data.DB) (*User, error) {
	if fixture.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	pid, _ := mongo.ParseObjectID(fixture.OwnerId)
	userArgument.SetID(data.ID(pid.Hex()))
	return userArgument, db.PopulateByID(userArgument)

}

func (fixture *Fixture) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := fixture.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := fixture.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(fixture); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (fixture *Fixture) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Exceptions []time.Time `json:"exceptions" bson:"exceptions"`

		ExpiresAt time.Time `json:"expires_at" bson:"expires_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Label bool `json:"label" bson:"label"`

		Name string `json:"name" bson:"name"`

		Rank int `json:"rank" bson:"rank"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ActionableId string `json:"actionable_id" bson:"actionable_id"`

		ActionableKind string `json:"actionable_kind" bson:"actionable_kind"`

		ActionsIds []string `json:"actions_ids" bson:"actions_ids"`

		EventableId string `json:"eventable_id" bson:"eventable_id"`

		EventableKind string `json:"eventable_kind" bson:"eventable_kind"`

		EventsIds []string `json:"events_ids" bson:"events_ids"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: fixture.CreatedAt,

		DeletedAt: fixture.DeletedAt,

		EndTime: fixture.EndTime,

		Exceptions: fixture.Exceptions,

		ExpiresAt: fixture.ExpiresAt,

		Label: fixture.Label,

		Name: fixture.Name,

		Rank: fixture.Rank,

		StartTime: fixture.StartTime,

		UpdatedAt: fixture.UpdatedAt,

		ActionableId: fixture.ActionableId,

		ActionableKind: fixture.ActionableKind,

		ActionsIds: fixture.ActionsIds,

		EventableId: fixture.EventableId,

		EventableKind: fixture.EventableKind,

		EventsIds: fixture.EventsIds,

		OwnerId: fixture.OwnerId,
	}, nil

}

func (fixture *Fixture) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Exceptions []time.Time `json:"exceptions" bson:"exceptions"`

		ExpiresAt time.Time `json:"expires_at" bson:"expires_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Label bool `json:"label" bson:"label"`

		Name string `json:"name" bson:"name"`

		Rank int `json:"rank" bson:"rank"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		ActionableId string `json:"actionable_id" bson:"actionable_id"`

		ActionableKind string `json:"actionable_kind" bson:"actionable_kind"`

		ActionsIds []string `json:"actions_ids" bson:"actions_ids"`

		EventableId string `json:"eventable_id" bson:"eventable_id"`

		EventableKind string `json:"eventable_kind" bson:"eventable_kind"`

		EventsIds []string `json:"events_ids" bson:"events_ids"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	fixture.CreatedAt = tmp.CreatedAt

	fixture.DeletedAt = tmp.DeletedAt

	fixture.EndTime = tmp.EndTime

	fixture.Exceptions = tmp.Exceptions

	fixture.ExpiresAt = tmp.ExpiresAt

	fixture.Id = tmp.Id.Hex()

	fixture.Label = tmp.Label

	fixture.Name = tmp.Name

	fixture.Rank = tmp.Rank

	fixture.StartTime = tmp.StartTime

	fixture.UpdatedAt = tmp.UpdatedAt

	fixture.ActionableId = tmp.ActionableId

	fixture.ActionableKind = tmp.ActionableKind

	fixture.ActionsIds = tmp.ActionsIds

	fixture.EventableId = tmp.EventableId

	fixture.EventableKind = tmp.EventableKind

	fixture.EventsIds = tmp.EventsIds

	fixture.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (fixture *Fixture) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		fixture.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		fixture.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		fixture.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["end_time"]; ok {
		fixture.EndTime = val.(time.Time)
	}

	if val, ok := structure["exceptions"]; ok {
		fixture.Exceptions = val.([]time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		fixture.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["name"]; ok {
		fixture.Name = val.(string)
	}

	if val, ok := structure["start_time"]; ok {
		fixture.StartTime = val.(time.Time)
	}

	if val, ok := structure["rank"]; ok {
		fixture.Rank = val.(int)
	}

	if val, ok := structure["label"]; ok {
		fixture.Label = val.(bool)
	}

	if val, ok := structure["expires_at"]; ok {
		fixture.ExpiresAt = val.(time.Time)
	}

	if val, ok := structure["actionable_id"]; ok {
		fixture.ActionableId = val.(string)
	}

	if val, ok := structure["actionable_kind"]; ok {
		fixture.ActionableKind = val.(string)
	}

	if val, ok := structure["eventable_id"]; ok {
		fixture.EventableId = val.(string)
	}

	if val, ok := structure["eventable_kind"]; ok {
		fixture.EventableKind = val.(string)
	}

	if val, ok := structure["actions_ids"]; ok {
		fixture.ActionsIds = val.([]string)
	}

	if val, ok := structure["events_ids"]; ok {
		fixture.EventsIds = val.([]string)
	}

	if val, ok := structure["owner_id"]; ok {
		fixture.OwnerId = val.(string)
	}

}

var FixtureStructure = map[string]metis.Primitive{

	"name": 3,

	"start_time": 4,

	"rank": 1,

	"label": 0,

	"expires_at": 4,

	"deleted_at": 4,

	"created_at": 4,

	"updated_at": 4,

	"end_time": 4,

	"exceptions": 8,

	"id": 9,

	"events_ids": 10,

	"owner_id": 9,

	"actionable_id": 9,

	"actionable_kind": 3,

	"eventable_id": 9,

	"eventable_kind": 3,

	"actions_ids": 10,
}
