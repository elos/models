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
type Calendar struct {
	BaseScheduleId    string            `json:"base_schedule_id" bson:"base_schedule_id"`
	CreatedAt         time.Time         `json:"created_at" bson:"created_at"`
	DeletedAt         time.Time         `json:"deleted_at" bson:"deleted_at"`
	FixturesIds       []string          `json:"fixtures_ids" bson:"fixtures_ids"`
	Id                string            `json:"id" bson:"_id,omitempty"`
	ManifestFixtureId string            `json:"manifest_fixture_id" bson:"manifest_fixture_id"`
	Name              string            `json:"name" bson:"name"`
	OwnerId           string            `json:"owner_id" bson:"owner_id"`
	UpdatedAt         time.Time         `json:"updated_at" bson:"updated_at"`
	WeekdaySchedules  map[string]string `json:"weekday_schedules" bson:"weekday_schedules"`
	YeardaySchedules  map[string]string `json:"yearday_schedules" bson:"yearday_schedules"`
}

func NewCalendar() *Calendar {
	return &Calendar{}
}

func FindCalendar(db data.DB, id data.ID) (*Calendar, error) {

	calendar := NewCalendar()
	calendar.SetID(id)

	return calendar, db.PopulateByID(calendar)

}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (calendar *Calendar) Kind() data.Kind {
	return CalendarKind
}

// just returns itself for now
func (calendar *Calendar) Concerned() []data.ID {
	foo := make([]data.ID, 1)
	foo[0] = calendar.ID()
	return foo
}

func (calendar *Calendar) SetID(id data.ID) {
	calendar.Id = id.String()
}

func (calendar *Calendar) ID() data.ID {
	return data.ID(calendar.Id)
}

func (calendar *Calendar) SetBaseSchedule(scheduleArgument *Schedule) error {
	calendar.BaseScheduleId = scheduleArgument.ID().String()
	return nil
}

func (calendar *Calendar) BaseSchedule(db data.DB) (*Schedule, error) {
	if calendar.BaseScheduleId == "" {
		return nil, ErrEmptyLink
	}

	scheduleArgument := NewSchedule()
	id, _ := db.ParseID(calendar.BaseScheduleId)
	scheduleArgument.SetID(id)
	return scheduleArgument, db.PopulateByID(scheduleArgument)

}

func (calendar *Calendar) BaseScheduleOrCreate(db data.DB) (*Schedule, error) {
	schedule, err := calendar.BaseSchedule(db)

	if err == ErrEmptyLink {
		schedule := NewSchedule()
		schedule.SetID(db.NewID())
		if err := calendar.SetBaseSchedule(schedule); err != nil {
			return nil, err
		}

		if err := db.Save(schedule); err != nil {
			return nil, err
		}

		if err := db.Save(calendar); err != nil {
			return nil, err
		}

		return schedule, nil
	} else {
		return schedule, err
	}
}

func (calendar *Calendar) IncludeFixture(fixture *Fixture) {
	otherID := fixture.ID().String()
	for i := range calendar.FixturesIds {
		if calendar.FixturesIds[i] == otherID {
			return
		}
	}
	calendar.FixturesIds = append(calendar.FixturesIds, otherID)
}

func (calendar *Calendar) ExcludeFixture(fixture *Fixture) {
	tmp := make([]string, 0)
	id := fixture.ID().String()
	for _, s := range calendar.FixturesIds {
		if s != id {
			tmp = append(tmp, s)
		}
	}
	calendar.FixturesIds = tmp
}

func (calendar *Calendar) FixturesIter(db data.DB) (data.Iterator, error) {
	// not yet completely general
	return mongo.NewIDIter(mongo.NewIDSetFromStrings(calendar.FixturesIds), db), nil
}

func (calendar *Calendar) Fixtures(db data.DB) (fixtures []*Fixture, err error) {
	fixtures = make([]*Fixture, len(calendar.FixturesIds))
	fixture := NewFixture()
	for i, id := range calendar.FixturesIds {
		fixture.Id = id
		if err = db.PopulateByID(fixture); err != nil {
			return
		}

		fixtures[i] = fixture
		fixture = NewFixture()
	}

	return
}

func (calendar *Calendar) SetManifestFixture(fixtureArgument *Fixture) error {
	calendar.ManifestFixtureId = fixtureArgument.ID().String()
	return nil
}

func (calendar *Calendar) ManifestFixture(db data.DB) (*Fixture, error) {
	if calendar.ManifestFixtureId == "" {
		return nil, ErrEmptyLink
	}

	fixtureArgument := NewFixture()
	id, _ := db.ParseID(calendar.ManifestFixtureId)
	fixtureArgument.SetID(id)
	return fixtureArgument, db.PopulateByID(fixtureArgument)

}

func (calendar *Calendar) ManifestFixtureOrCreate(db data.DB) (*Fixture, error) {
	fixture, err := calendar.ManifestFixture(db)

	if err == ErrEmptyLink {
		fixture := NewFixture()
		fixture.SetID(db.NewID())
		if err := calendar.SetManifestFixture(fixture); err != nil {
			return nil, err
		}

		if err := db.Save(fixture); err != nil {
			return nil, err
		}

		if err := db.Save(calendar); err != nil {
			return nil, err
		}

		return fixture, nil
	} else {
		return fixture, err
	}
}

func (calendar *Calendar) SetOwner(userArgument *User) error {
	calendar.OwnerId = userArgument.ID().String()
	return nil
}

func (calendar *Calendar) Owner(db data.DB) (*User, error) {
	if calendar.OwnerId == "" {
		return nil, ErrEmptyLink
	}

	userArgument := NewUser()
	id, _ := db.ParseID(calendar.OwnerId)
	userArgument.SetID(id)
	return userArgument, db.PopulateByID(userArgument)

}

func (calendar *Calendar) OwnerOrCreate(db data.DB) (*User, error) {
	user, err := calendar.Owner(db)

	if err == ErrEmptyLink {
		user := NewUser()
		user.SetID(db.NewID())
		if err := calendar.SetOwner(user); err != nil {
			return nil, err
		}

		if err := db.Save(user); err != nil {
			return nil, err
		}

		if err := db.Save(calendar); err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return user, err
	}
}

// BSON {{{
func (calendar *Calendar) GetBSON() (interface{}, error) {

	return struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id string `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		WeekdaySchedules map[string]string `json:"weekday_schedules" bson:"weekday_schedules"`

		YeardaySchedules map[string]string `json:"yearday_schedules" bson:"yearday_schedules"`

		BaseScheduleId string `json:"base_schedule_id" bson:"base_schedule_id"`

		FixturesIds []string `json:"fixtures_ids" bson:"fixtures_ids"`

		ManifestFixtureId string `json:"manifest_fixture_id" bson:"manifest_fixture_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{

		CreatedAt: calendar.CreatedAt,

		DeletedAt: calendar.DeletedAt,

		Name: calendar.Name,

		UpdatedAt: calendar.UpdatedAt,

		WeekdaySchedules: calendar.WeekdaySchedules,

		YeardaySchedules: calendar.YeardaySchedules,

		BaseScheduleId: calendar.BaseScheduleId,

		FixturesIds: calendar.FixturesIds,

		ManifestFixtureId: calendar.ManifestFixtureId,

		OwnerId: calendar.OwnerId,
	}, nil

}

func (calendar *Calendar) SetBSON(raw bson.Raw) error {

	tmp := struct {
		CreatedAt time.Time `json:"created_at" bson:"created_at"`

		DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`

		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

		Name string `json:"name" bson:"name"`

		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

		WeekdaySchedules map[string]string `json:"weekday_schedules" bson:"weekday_schedules"`

		YeardaySchedules map[string]string `json:"yearday_schedules" bson:"yearday_schedules"`

		BaseScheduleId string `json:"base_schedule_id" bson:"base_schedule_id"`

		FixturesIds []string `json:"fixtures_ids" bson:"fixtures_ids"`

		ManifestFixtureId string `json:"manifest_fixture_id" bson:"manifest_fixture_id"`

		OwnerId string `json:"owner_id" bson:"owner_id"`
	}{}

	err := raw.Unmarshal(&tmp)
	if err != nil {
		return err
	}

	calendar.CreatedAt = tmp.CreatedAt

	calendar.DeletedAt = tmp.DeletedAt

	calendar.Id = tmp.Id.Hex()

	calendar.Name = tmp.Name

	calendar.UpdatedAt = tmp.UpdatedAt

	calendar.WeekdaySchedules = tmp.WeekdaySchedules

	calendar.YeardaySchedules = tmp.YeardaySchedules

	calendar.BaseScheduleId = tmp.BaseScheduleId

	calendar.FixturesIds = tmp.FixturesIds

	calendar.ManifestFixtureId = tmp.ManifestFixtureId

	calendar.OwnerId = tmp.OwnerId

	return nil

}

// BSON }}}

func (calendar *Calendar) FromStructure(structure map[string]interface{}) {

	if val, ok := structure["name"]; ok {
		calendar.Name = val.(string)
	}

	if val, ok := structure["weekday_schedules"]; ok {
		calendar.WeekdaySchedules = val.(map[string]string)
	}

	if val, ok := structure["yearday_schedules"]; ok {
		calendar.YeardaySchedules = val.(map[string]string)
	}

	if val, ok := structure["id"]; ok {
		calendar.Id = val.(string)
	}

	if val, ok := structure["created_at"]; ok {
		calendar.CreatedAt = val.(time.Time)
	}

	if val, ok := structure["updated_at"]; ok {
		calendar.UpdatedAt = val.(time.Time)
	}

	if val, ok := structure["deleted_at"]; ok {
		calendar.DeletedAt = val.(time.Time)
	}

	if val, ok := structure["owner_id"]; ok {
		calendar.OwnerId = val.(string)
	}

	if val, ok := structure["base_schedule_id"]; ok {
		calendar.BaseScheduleId = val.(string)
	}

	if val, ok := structure["manifest_fixture_id"]; ok {
		calendar.ManifestFixtureId = val.(string)
	}

	if val, ok := structure["fixtures_ids"]; ok {
		calendar.FixturesIds = val.([]string)
	}

}

var CalendarStructure = map[string]metis.Primitive{

	"yearday_schedules": 11,

	"id": 9,

	"created_at": 4,

	"updated_at": 4,

	"deleted_at": 4,

	"name": 3,

	"weekday_schedules": 11,

	"owner_id": 9,

	"base_schedule_id": 9,

	"manifest_fixture_id": 9,

	"fixtures_ids": 10,
}
