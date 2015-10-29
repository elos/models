package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"gopkg.in/mgo.v2/bson"
)

// THIS FILE GENERATED BY METIS

// this type def generated by metis
type Schedule struct {
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	DeletedAt   time.Time `json:"deleted_at" bson:"deleted_at"`
	EndTime     time.Time `json:"end_time" bson:"end_time"`
	FixturesIDs []string  `json:"fixtures_ids" bson:"fixtures_ids"`
	Id          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	OwnerID     string    `json:"owner_id" bson:"owner_id"`
	StartTime   time.Time `json:"start_time" bson:"start_time"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func NewSchedule() *Schedule {
	return &Schedule{}
}

func FindSchedule(db data.DB, id data.ID) (*Schedule, error) {

	schedule := NewSchedule()
	schedule.SetID(id)

	return schedule, db.PopulateByID(schedule)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (schedule *Schedule) Kind() data.Kind {
	return ScheduleKind
}

// just returns itself for now
func (schedule *Schedule) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = schedule.ID()
	return foo
}

func (schedule *Schedule) SetID(id data.ID) {
	schedule.Id = id.String()
}

func (schedule *Schedule) ID() data.ID {
	return data.ID(schedule.Id)
}

func (schedule *Schedule) IncludeFixture(fixture *Fixture) {
	schedule.FixturesIDs = append(schedule.FixturesIDs, fixture.ID().String())
}

func (schedule *Schedule) ExcludeFixture(fixture *Fixture) {
	tmp := make([]string, 0)
	id := fixture.ID().String()
	for _, s := range schedule.FixturesIDs {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	schedule.FixturesIDs = tmp
}

func (schedule *Schedule) FixturesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(schedule.FixturesIDs), db), nil
}

func (schedule *Schedule) Fixtures(db data.DB) ([]*Fixture, error) {

	fixtures := make([]*Fixture, 0)
	iter := mongo.NewIDIter(mongo.NewIDSetFromStrings(schedule.FixturesIDs), db)
	fixture := NewFixture()
	for iter.Next(fixture) {
		fixtures = append(fixtures, fixture)
		fixture = NewFixture()
	}
	return fixtures, nil
}

func (schedule *Schedule) SetOwner(userArgument *User) error {
	schedule.OwnerID = userArgument.ID().String()
	return nil
}

func (schedule *Schedule) Owner(db data.DB) (*User, error) {
	if schedule.OwnerID == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	pid, _ := mongo.ParseObjectID(schedule.OwnerID)
	userArgument.SetID(data.ID(pid.Hex()))
	return userArgument, db.PopulateByID(userArgument)

}

func (schedule *Schedule) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := schedule.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := schedule.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(schedule); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (schedule *Schedule) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		FixturesIDs []string `json:"fixtures_ids" bson:"fixtures_ids"`

		OwnerID string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: schedule.CreatedAt,

		DeletedAt: schedule.DeletedAt,

		EndTime: schedule.EndTime,

		Name: schedule.Name,

		StartTime: schedule.StartTime,

		UpdatedAt: schedule.UpdatedAt,

		FixturesIDs: schedule.FixturesIDs,

		OwnerID: schedule.OwnerID,
	}, nil

}

func (schedule *Schedule) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		EndTime time.Time `json:"end_time" bson:"end_time"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		StartTime time.Time `json:"start_time" bson:"start_time"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		FixturesIDs []string `json:"fixtures_ids" bson:"fixtures_ids"`

		OwnerID string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	schedule.CreatedAt = tmp.CreatedAt

	schedule.DeletedAt = tmp.DeletedAt

	schedule.EndTime = tmp.EndTime

	schedule.Id = tmp.Id.Hex()

	schedule.Name = tmp.Name

	schedule.StartTime = tmp.StartTime

	schedule.UpdatedAt = tmp.UpdatedAt

	schedule.FixturesIDs = tmp.FixturesIDs

	schedule.OwnerID = tmp.OwnerID

	return nil

}

// BSON }}}
