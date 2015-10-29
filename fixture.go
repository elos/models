package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Fixture struct {
	ActionableID   string      `json:"actionable_id" bson:"actionable_id"`
	ActionableKind string      `json:"actionable_kind" bson:"actionable_kind"`
	ActionsIDs     []string    `json:"actions_ids" bson:"actions_ids"`
	CreatedAt      time.Time   `json:"created_at" bson:"created_at"`
	DeletedAt      time.Time   `json:"deleted_at" bson:"deleted_at"`
	EndTime        time.Time   `json:"end_time" bson:"end_time"`
	EventableID    string      `json:"eventable_id" bson:"eventable_id"`
	EventableKind  string      `json:"eventable_kind" bson:"eventable_kind"`
	EventsIDs      []string    `json:"events_ids" bson:"events_ids"`
	Exceptions     []time.Time `json:"exceptions" bson:"exceptions"`
	ExpiresAt      time.Time   `json:"expires_at" bson:"expires_at"`
	Id             string      `json:"id" bson:"_id,omitempty"`
	Label          bool        `json:"label" bson:"label"`
	Name           string      `json:"name" bson:"name"`
	OwnerID        string      `json:"owner_id" bson:"owner_id"`
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
	fixture.ActionableID = actionableArgument.ID().String()
	return nil
}

func (fixture *Fixture) Actionable(db data.DB) (Actionable, error) {
	if fixture.ActionableID == "" {
		return nil, ErrEmptyLink
	}

	m := ModelFor(data.Kind(fixture.ActionableKind))
	actionable := m.(Actionable)

	pid, _ := mongo.ParseObjectID(fixture.ActionableID)

	actionable.SetID(data.ID(pid.Hex()))
	return actionable, db.PopulateByID(actionable)

}

func (fixture *Fixture) IncludeAction(action *Action) {
	fixture.ActionsIDs = append(fixture.ActionsIDs, action.ID().String())
}

func (fixture *Fixture) ExcludeAction(action *Action) {
	tmp := make([]string, 0)
	id := action.ID().String()
	for _, s := range fixture.ActionsIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	fixture.ActionsIDs = tmp
}

func (fixture *Fixture) ActionsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(fixture.ActionsIDs), db), nil
}

func (fixture *Fixture) Actions(db data.DB) ([]*Action, error) {

	actions := make([]*Action, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(fixture.ActionsIDs), db)
	action := NewAction()
	for iter.Next(action) {
		actions = append(actions, action)
		action = NewAction()
	}
	return actions, nil
}

func (fixture *Fixture) SetEventable(eventableArgument Eventable) error {
	fixture.EventableID = eventableArgument.ID().String()
	return nil
}

func (fixture *Fixture) Eventable(db data.DB) (Eventable, error) {
	if fixture.EventableID == "" {
		return nil, ErrEmptyLink
	}

	m := ModelFor(data.Kind(fixture.EventableKind))
	eventable := m.(Eventable)

	pid, _ := mongo.ParseObjectID(fixture.EventableID)

	eventable.SetID(data.ID(pid.Hex()))
	return eventable, db.PopulateByID(eventable)

}

func (fixture *Fixture) IncludeEvent(event *Event) {
	fixture.EventsIDs = append(fixture.EventsIDs, event.ID().String())
}

func (fixture *Fixture) ExcludeEvent(event *Event) {
	tmp := make([]string, 0)
	id := event.ID().String()
	for _, s := range fixture.EventsIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	fixture.EventsIDs = tmp
}

func (fixture *Fixture) EventsIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(fixture.EventsIDs), db), nil
}

func (fixture *Fixture) Events(db data.DB) ([]*Event, error) {

	events := make([]*Event, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(fixture.EventsIDs), db)
	event := NewEvent()
	for iter.Next(event) {
		events = append(events, event)
		event = NewEvent()
	}
	return events, nil
}

func (fixture *Fixture) SetOwner(userArgument *User) error {
	fixture.OwnerID = userArgument.ID().String()
	return nil
}

func (fixture *Fixture) Owner(db data.DB) (*User, error) {
	if fixture.OwnerID == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	pid, _ := mongo.ParseObjectID(fixture.OwnerID)
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

		ActionableID string `json:"actionable_id" bson:"actionable_id"`

		ActionableKind string `json:"actionable_kind" bson:"actionable_kind"`

		ActionsIDs []string `json:"actions_ids" bson:"actions_ids"`

		EventableID string `json:"eventable_id" bson:"eventable_id"`

		EventableKind string `json:"eventable_kind" bson:"eventable_kind"`

		EventsIDs []string `json:"events_ids" bson:"events_ids"`

		OwnerID string `json:"owner_id" bson:"owner_id"`
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

		ActionableID: fixture.ActionableID,

		ActionableKind: fixture.ActionableKind,

		ActionsIDs: fixture.ActionsIDs,

		EventableID: fixture.EventableID,

		EventableKind: fixture.EventableKind,

		EventsIDs: fixture.EventsIDs,

		OwnerID: fixture.OwnerID,
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

		ActionableID string `json:"actionable_id" bson:"actionable_id"`

		ActionableKind string `json:"actionable_kind" bson:"actionable_kind"`

		ActionsIDs []string `json:"actions_ids" bson:"actions_ids"`

		EventableID string `json:"eventable_id" bson:"eventable_id"`

		EventableKind string `json:"eventable_kind" bson:"eventable_kind"`

		EventsIDs []string `json:"events_ids" bson:"events_ids"`

		OwnerID string `json:"owner_id" bson:"owner_id"`
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

	fixture.ActionableID = tmp.ActionableID

	fixture.ActionableKind = tmp.ActionableKind

	fixture.ActionsIDs = tmp.ActionsIDs

	fixture.EventableID = tmp.EventableID

	fixture.EventableKind = tmp.EventableKind

	fixture.EventsIDs = tmp.EventsIDs

	fixture.OwnerID = tmp.OwnerID

	return nil

}

// BSON }}}
