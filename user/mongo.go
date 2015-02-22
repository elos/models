package user

import (
	"errors"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoUser struct {
	models.MongoID     `bson:",inline"`
	models.Named       `bson:",inline"`
	models.Timestamped `bson:",inline"`
	EKey               string `json:"key" bson:"key"`

	EventIDs   mongo.IDSet `json:"event_ids" bson:"event_ids"`
	TaskIDs    mongo.IDSet `json:"task_ids" bson:"task_ids"`
	RoutineIDs mongo.IDSet `json:"routine_ids", bson:"routine_ids"`

	CalendarID      bson.ObjectId `json:"calendar_id" bson:"calendar_id,omitempty"`
	CurrentActionID bson.ObjectId `json:"current_action_id" bson:"current_action_id,omitempty"`
	ActionableKind  data.Kind     `json:"actionable_kind" bson:"actionable_kind"`
	ActionableID    bson.ObjectId `json:"actionable_id" bson:"actionable_id,omitempty"`
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
	case CurrentAction:
		u.CurrentActionID = m.ID().(bson.ObjectId)
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
	case CurrentAction:
		if u.CurrentActionID == m.ID().(bson.ObjectId) {
			u.CurrentActionID = *new(bson.ObjectId)
		}

		return nil
	default:
		return data.ErrUndefinedLink
	}
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

func (u *mongoUser) DropEvent(e models.Event) error {
	return u.Schema().Unlink(u, e, Events)
}

func (u *mongoUser) Events(a *data.Access) (data.RecordIterator, error) {
	if u.CanWrite(a.Client) {
		return mongo.NewIDIter(u.EventIDs, a.Store), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (u *mongoUser) AddTask(t models.Task) error {
	return u.Schema().Link(u, t, Tasks)
}

func (u *mongoUser) DropTask(t models.Task) error {
	return u.Schema().Unlink(u, t, Tasks)
}

func (u *mongoUser) Tasks(a *data.Access) (data.RecordIterator, error) {
	if u.CanWrite(a.Client) {
		return mongo.NewIDIter(u.TaskIDs, a.Store), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (u *mongoUser) SetCurrentAction(a models.Action) {
	u.Schema().Link(u, a, CurrentAction)
}

func (u *mongoUser) CurrentAction(a *data.Access, action models.Action) error {
	action.SetID(u.CurrentActionID)
	return a.PopulateByID(action)
}

func (u *mongoUser) SetCurrentActionable(a models.Actionable) {
	u.ActionableKind = a.Kind()
	u.ActionableID = a.ID().(bson.ObjectId)
}

func (u *mongoUser) CurrentActionable(a *data.Access) (models.Actionable, error) {
	m, err := a.ModelFor(u.ActionableKind)
	if err != nil {
		return nil, err
	}

	m.SetID(u.ActionableID)
	if err = a.PopulateByID(m); err != nil {
		return nil, err
	}

	actionable, ok := m.(models.Actionable)
	if !ok {
		return nil, errors.New("idk")
	} else {
		return actionable, nil
	}
}

func (u *mongoUser) CanRead(c data.Client) bool {
	if c.Kind() != models.UserKind {
		return false
	}

	if u.EID.Valid() && c.ID().(bson.ObjectId) != u.EID {
		return false
	}

	return true
}

func (u *mongoUser) CanWrite(c data.Client) bool {
	cid := c.ID()
	cid, ok := cid.(bson.ObjectId)
	if ok {
		if u.EID.Valid() && c.ID().(bson.ObjectId) != u.EID {
			return false
		}

		return true
	} else {
		if u.EID == "" && c.Kind() == data.Anonymous {
			return true
		}

		return false
	}
}
