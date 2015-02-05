package user

import (
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type mongoUser struct {
	EID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ECreatedAt time.Time     `json:"created_at" bson:"created_at"`
	EUpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	EName         string        `json:"name" bson:"name"`
	EKey          string        `json:"key" bson:"key"`
	EventIDs      mongo.IDSet   `json:"event_ids" bson:"event_ids"`
	TaskIDs       mongo.IDSet   `json:"task_ids" bson:"task_ids"`
	CurrentTaskID bson.ObjectId `json:"current_task" bson:"current_task,omitempty"`
}

func (u *mongoUser) DBType() data.DBType {
	return mongo.DBType
}

func (u *mongoUser) Kind() data.Kind {
	return kind
}

func (u *mongoUser) Schema() data.Schema {
	return schema
}

func (u *mongoUser) Version() int {
	return version
}

func (u *mongoUser) Valid() bool {
	valid, _ := Validate(u)
	return valid
}

func (u *mongoUser) Save(s data.Store) error {
	valid, err := Validate(u)
	if valid {
		return s.Save(u)
	} else {
		return err
	}
}

func (u *mongoUser) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = u.EID
	return a
}

func (u *mongoUser) LinkEvent(eventID bson.ObjectId) error {
	u.EventIDs = mongo.AddID(u.EventIDs, eventID)
	return nil
}

func (u *mongoUser) UnlinkEvent(eventID bson.ObjectId) error {
	u.EventIDs = mongo.DropID(u.EventIDs, eventID)
	return nil
}

func (u *mongoUser) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case Events:
		return u.LinkEvent(m.ID().(bson.ObjectId))
	case Tasks:
		u.TaskIDs = mongo.AddID(u.TaskIDs, m.ID().(bson.ObjectId))
		return nil
	case CurrentTask:
		u.CurrentTaskID = m.ID().(bson.ObjectId)
		return nil
	default:
		return data.NewLinkError(u, m, l)
	}
}

func (u *mongoUser) Unlink(m data.Model, l data.Link) error {
	switch l.Name {
	case Events:
		return u.UnlinkEvent(m.ID().(bson.ObjectId))
	case Tasks:
		u.TaskIDs = mongo.DropID(u.TaskIDs, m.ID().(bson.ObjectId))
		return nil
	case CurrentTask:
		if u.CurrentTaskID == m.ID().(bson.ObjectId) {
			u.CurrentTaskID = *new(bson.ObjectId)
		}

		return nil
	default:
		return data.ErrUndefinedLink
	}
}

func (u *mongoUser) SetID(id data.ID) {
	vid, ok := id.(bson.ObjectId)
	if ok {
		u.EID = vid
	}
}

func (u *mongoUser) ID() data.ID {
	return u.EID
}

func (u *mongoUser) SetName(name string) {
	u.EName = name
}

func (u *mongoUser) Name() string {
	return u.EName
}

func (u *mongoUser) SetCreatedAt(t time.Time) {
	u.ECreatedAt = t
}

func (u *mongoUser) CreatedAt() time.Time {
	return u.ECreatedAt
}

func (u *mongoUser) SetUpdatedAt(t time.Time) {
	u.EUpdatedAt = t
}

func (u *mongoUser) UpdatedAt() time.Time {
	return u.EUpdatedAt
}

func (u *mongoUser) SetKey(s string) {
	u.EKey = s
}

func (u *mongoUser) Key() string {
	return u.EKey
}

func (u *mongoUser) AddEvent(e models.Event) error {
	return u.Schema().Link(u, e, Events)
}

func (u *mongoUser) RemoveEvent(e models.Event) error {
	return u.Schema().Unlink(u, e, Events)
}

func (u *mongoUser) AddTask(t models.Task) error {
	return u.Schema().Link(u, t, Tasks)
}

func (u *mongoUser) RemoveTask(t models.Task) error {
	return u.Schema().Unlink(u, t, Tasks)
}
