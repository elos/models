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
type Schedule struct {
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	DeletedAt   time.Time `json:"deleted_at" bson:"deleted_at"`
	EndTime     time.Time `json:"end_time" bson:"end_time"`
	FixturesIds []string  `json:"fixtures_ids" bson:"fixtures_ids"`
	Id          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	OwnerId     string    `json:"owner_id" bson:"owner_id"`
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
	otherID := fixture.ID().String()
	for i := range schedule.FixturesIds {
		if schedule.FixturesIds[i] == otherID {
			return
		}
	}
	schedule.FixturesIds = append(schedule.FixturesIds, otherID)
}

func (schedule *Schedule) ExcludeFixture(fixture *Fixture) {
	tmp := make([]string, 0)
	id := fixture.ID().String()
	for _, s := range schedule.FixturesIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	schedule.FixturesIds = tmp
}

func (schedule *Schedule) FixturesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(schedule.FixturesIds), db), nil
}

func (schedule *Schedule) Fixtures(db data.DB) (fixtures []*Fixture, err error) {
	fixtures = make([]*Fixture, len(schedule.FixturesIds))
	fixture := NewFixture()
	for i, id := range schedule.FixturesIds {
		fixture.Id = id
		if err = db.PopulateByID(fixture); err != nil {
			return
		}

		fixtures[i] = fixture
		fixture = NewFixture()
	}

	return
}

func (schedule *Schedule) SetOwner(userArgument *User) error {
	schedule.OwnerId = userArgument.ID().String()
	return nil
}

func (schedule *Schedule) Owner(db data.DB) (*User, error) {
	if schedule.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(schedule.OwnerId)
	userArgument.SetID(id)
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

		FixturesIds []string `json:"fixtures_ids" bson:"fixtures_ids"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: schedule.CreatedAt,

		DeletedAt: schedule.DeletedAt,

		EndTime: schedule.EndTime,

		Name: schedule.Name,

		StartTime: schedule.StartTime,

		UpdatedAt: schedule.UpdatedAt,

		FixturesIds: schedule.FixturesIds,

		OwnerId: schedule.OwnerId,
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

		FixturesIds []string `json:"fixtures_ids" bson:"fixtures_ids"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
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

	schedule.FixturesIds = tmp.FixturesIds

	schedule.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (schedule *Schedule) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["id"]; ok {
		schedule.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		schedule.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		schedule.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		schedule.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["name"]; ok {
		schedule.Name = val.(string)
	}

	if val, ok := structure["start_time"]; ok {
		schedule.StartTime = val.(time.Time)
	}

	if val, ok := structure["end_time"]; ok {
		schedule.EndTime = val.(time.Time)
	}

	if val, ok := structure["fixtures_ids"]; ok {
		schedule.FixturesIds = val.([]string)
	}

	if val, ok := structure["owner_id"]; ok {
		schedule.OwnerId = val.(string)
	}

}

var ScheduleStructure = map[string]metis.Primitive{

	"start_time": 4,

	"end_time": 4,

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"name": 3,

	"owner_id": 9,

	"fixtures_ids": 10,
}
