package user

import (
	"errors"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoUser struct {
	mongo.Model `bson:",inline"`
	mongo.Named `bson:",inline"`

	EKey string `json:"key" bson:"key"`

	ActionIDs  mongo.IDSet `json:"action_ids" bson:"action_ids"`
	EventIDs   mongo.IDSet `json:"event_ids" bson:"event_ids"`
	TaskIDs    mongo.IDSet `json:"task_ids" bson:"task_ids"`
	RoutineIDs mongo.IDSet `json:"routine_ids", bson:"routine_ids"`

	EOntologyID     bson.ObjectId `json:"ontology_id" bson:"ontology_id,omitempty"`
	ECalendarID     bson.ObjectId `json:"calendar_id" bson:"calendar_id,omitempty"`
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

func (u *mongoUser) SetOntology(o models.Ontology) error {
	return u.Schema().Link(u, o, Ontology)
}

func (u *mongoUser) Ontology(a data.Access) (models.Ontology, error) {
	m, _ := a.ModelFor(models.OntologyKind)
	o := m.(models.Ontology)

	if u.CanRead(a.Client()) {
		o.SetID(u.EOntologyID)
		err := a.PopulateByID(o)
		return o, err
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (u *mongoUser) Link(m data.Model, l data.Link) error {
	if !data.Compatible(u, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case Events:
		return u.LinkEvent(m.ID().(bson.ObjectId))
	case Tasks:
		u.TaskIDs = mongo.AddID(u.TaskIDs, m.ID().(bson.ObjectId))
	case Routines:
		u.RoutineIDs = mongo.AddID(u.RoutineIDs, m.ID().(bson.ObjectId))
	case Ontology:
		u.EOntologyID = m.ID().(bson.ObjectId)
	case Calendar:
		u.ECalendarID = m.ID().(bson.ObjectId)
	case CurrentAction:
		u.CurrentActionID = m.ID().(bson.ObjectId)
	default:
		return data.NewLinkError(u, m, l)
	}
	return nil
}

func (u *mongoUser) Unlink(m data.Model, l data.Link) error {
	if !data.Compatible(u, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case Events:
		return u.UnlinkEvent(m.ID().(bson.ObjectId))
	case Tasks:
		u.TaskIDs = mongo.DropID(u.TaskIDs, m.ID().(bson.ObjectId))
	case Routines:
		u.RoutineIDs = mongo.DropID(u.RoutineIDs, m.ID().(bson.ObjectId))
	case Calendar:
		if u.ECalendarID == m.ID().(bson.ObjectId) {
			u.ECalendarID = *new(bson.ObjectId)
		}
	case Ontology:
		if u.EOntologyID == m.ID().(bson.ObjectId) {
			u.EOntologyID = *new(bson.ObjectId)
		}
	case CurrentAction:
		if u.CurrentActionID == m.ID().(bson.ObjectId) {
			u.CurrentActionID = *new(bson.ObjectId)
		}

	default:
		return data.ErrUndefinedLink
	}
	return nil
}

func (u *mongoUser) SetKey(s string) {
	u.EKey = s
}

func (u *mongoUser) Key() string {
	return u.EKey
}

func (u *mongoUser) IncludeEvent(e models.Event) error {
	return u.Schema().Link(u, e, Events)
}

func (u *mongoUser) ExcludeEvent(e models.Event) error {
	return u.Schema().Unlink(u, e, Events)
}

func (u *mongoUser) Events(a data.Access) (data.ModelIterator, error) {
	if u.CanWrite(a.Client()) {
		return mongo.NewIDIter(u.EventIDs, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (u *mongoUser) IncludeTask(t models.Task) error {
	return u.Schema().Link(u, t, Tasks)
}

func (u *mongoUser) ExcludeTask(t models.Task) error {
	return u.Schema().Unlink(u, t, Tasks)
}

func (u *mongoUser) Tasks(a data.Access) (data.ModelIterator, error) {
	if u.CanWrite(a.Client()) {
		return mongo.NewIDIter(u.TaskIDs, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (u *mongoUser) Routines(a data.Access) (data.ModelIterator, error) {
	if u.CanRead(a.Client()) {
		return mongo.NewIDIter(u.RoutineIDs, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (u *mongoUser) IncludeRoutine(r models.Routine) error {
	return u.Schema().Link(u, r, Routines)
}

func (u *mongoUser) ExcludeRoutine(r models.Routine) error {
	return u.Schema().Unlink(u, r, Routines)
}

func (u *mongoUser) SetCalendar(c models.Calendar) error {
	return u.Schema().Link(u, c, Calendar)
}

func (u *mongoUser) Calendar(a data.Access) (models.Calendar, error) {
	if !u.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	m, err := a.ModelFor(models.CalendarKind)
	if err != nil {
		return nil, err
	}

	c, _ := m.(models.Calendar)

	c.SetID(u.ECalendarID)
	err = a.PopulateByID(c)
	return c, err
}

func (u *mongoUser) SetCurrentAction(a models.Action) {
	u.Schema().Link(u, a, CurrentAction)
}

func (u *mongoUser) CurrentAction(a data.Access) (models.Action, error) {
	m, err := a.ModelFor(models.ActionKind)
	if err != nil {
		return nil, err
	}

	action, ok := m.(models.Action)
	if !ok {
		return nil, errors.New("TODO")
	}

	action.SetID(u.CurrentActionID)
	err = a.PopulateByID(action)

	return action, err
}

func (u *mongoUser) ClearCurrentActionable() {
	u.ActionableKind = ""
	u.ActionableID = *new(bson.ObjectId)
}

func (u *mongoUser) SetCurrentActionable(a models.Actionable) {
	u.ActionableKind = a.Kind()
	u.ActionableID = a.ID().(bson.ObjectId)
}

func (u *mongoUser) CurrentActionable(a data.Access) (models.Actionable, error) {
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
