package fixture

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoFixture struct {
	mongo.Model      `bson:",inline"`
	mongo.Named      `bson:",inline"`
	mongo.Timed      `bson:",inline"`
	models.UserOwned `bson:",inline"`

	EScheduleID  bson.ObjectId `json:"schedule_id" bson:"schedule_id,omitempty"`
	EDescription string        `json:"decription" bson:"description"`
}

func (f *mongoFixture) Kind() data.Kind {
	return kind
}

func (f *mongoFixture) Version() int {
	return version
}

func (f *mongoFixture) Schema() data.Schema {
	return schema
}

func (f *mongoFixture) SetDescription(s string) {
	f.EDescription = s
}

func (f *mongoFixture) Description() string {
	return f.EDescription
}

func (f *mongoFixture) SetSchedule(s models.Schedule) error {
	return f.Schema().Link(f, s, Schedule)
}

func (f *mongoFixture) Schedule(a data.Access, s models.Schedule) error {
	s.SetID(f.EScheduleID)
	return a.PopulateByID(s)
}

func (f *mongoFixture) SetUser(u models.User) error {
	return f.Schema().Link(f, u, User)
}

func (f *mongoFixture) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		return f.SetUserID(m.ID())
	case Schedule:
		f.EScheduleID = m.ID().(bson.ObjectId)
		return nil
	default:
		return data.NewLinkError(f, m, l)
	}
}

func (f *mongoFixture) Unlink(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		f.DropUserID()
	case Schedule:
		if f.EScheduleID == m.ID().(bson.ObjectId) {
			f.EScheduleID = *new(bson.ObjectId)
		}
	default:
		return data.NewLinkError(f, m, l)
	}
	return nil
}
