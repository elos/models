package user

import (
	"errors"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

/*
	mongoUser is a models.User implementation for the
	"github.com/elos/mongo" data store.
*/
type mongoUser struct {
	mongo.Model `bson:",inline"`
	mongo.Named `bson:",inline"`

	EKey string `json:"key" bson:"key"`

	ActionIDs  mongo.IDSet `json:"action_ids"  bson:"action_ids"`
	EventIDs   mongo.IDSet `json:"event_ids"   bson:"event_ids"`
	TaskIDs    mongo.IDSet `json:"task_ids"    bson:"task_ids"`
	RoutineIDs mongo.IDSet `json:"routine_ids" bson:"routine_ids"`

	EOntologyID     bson.ObjectId `json:"ontology_id"       bson:"ontology_id,omitempty"`
	ECalendarID     bson.ObjectId `json:"calendar_id"       bson:"calendar_id,omitempty"`
	CurrentActionID bson.ObjectId `json:"current_action_id" bson:"current_action_id,omitempty"`
	ActionableKind  data.Kind     `json:"actionable_kind"   bson:"actionable_kind"`
	ActionableID    bson.ObjectId `json:"actionable_id"     bson:"actionable_id,omitempty"`
}

// Kind is derived from the models package and is
// defined in type.go, shared among implementations
func (u *mongoUser) Kind() data.Kind {
	return kind
}

// Schema is derived from the models package and is
// defined in type.go, shared among implementations
func (u *mongoUser) Schema() data.Schema {
	return schema
}

// Version is derived from the models package and is
// defined in type.go, shared among implementations
func (u *mongoUser) Version() int {
	return version
}

// Valid refers to whether the user is valid
// and is an implementation of the data.Validateable
// interface, please see those docs for the
// precise specification
func (u *mongoUser) Valid() bool {
	valid, _ := Validate(u)
	return valid
}

/*
	Concerned is for the implementation of
	the data.Record interface, and is how the
	application level store notifies subsribers
	that the user models has changed

	A user's only concern is itself.
*/
func (u *mongoUser) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = u.EID
	return a
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

/*
	Link satifies the data.Linkable interface

	It can return a data.ErrIncompatibleModels or
	data.ErrInvalidId errors or an instance of a data.LinkError
*/
func (u *mongoUser) Link(m data.Model, l data.Link) error {
	if !data.Compatible(u, m) {
		return data.ErrIncompatibleModels
	}

	id, ok := m.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	switch l.Name {
	case events:
		u.EventIDs = mongo.AddID(u.EventIDs, id)
	case tasks:
		u.TaskIDs = mongo.AddID(u.TaskIDs, id)
	case routines:
		u.RoutineIDs = mongo.AddID(u.RoutineIDs, id)
	case ontology:
		u.EOntologyID = id
	case calendar:
		u.ECalendarID = id
	case currentAction:
		u.CurrentActionID = id
	default:
		return data.NewLinkError(u, m, l)
	}

	return nil
}

/*
	Unlink satisfies the data.Linkable interface.

	Unlink can return a data.ErrIncompatibleModels or
	data.ErrInvalidID or an instance of a data.LinkError

	You will notice that the user follows the linkable spec
	that a requested unlink should only occur when that model
	was actually linked previously. Meaning that if you
	unlink a model and a calendar, if the calendar wasn't
	the user's current calendar, it doesn't just wipe the calendar
	that is the current user's calendar. These are the  if u.XXID == id
	checks
*/
func (u *mongoUser) Unlink(m data.Model, l data.Link) error {
	if !data.Compatible(u, m) {
		return data.ErrIncompatibleModels
	}

	id, ok := m.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	switch l.Name {
	case events:
		u.EventIDs = mongo.DropID(u.EventIDs, id)
	case tasks:
		u.TaskIDs = mongo.DropID(u.TaskIDs, id)
	case routines:
		u.RoutineIDs = mongo.DropID(u.RoutineIDs, id)
	case calendar:
		if u.ECalendarID == id {
			u.ECalendarID = *new(bson.ObjectId)
		}
	case ontology:
		if u.EOntologyID == id {
			u.EOntologyID = *new(bson.ObjectId)
		}
	case currentAction:
		if u.CurrentActionID == id {
			u.CurrentActionID = *new(bson.ObjectId)
		}
	default:
		return data.NewLinkError(u, m, l)
	}

	return nil
}

// Sets the user's key
func (u *mongoUser) SetKey(s string) {
	u.EKey = s
}

// Gets the user's key
func (u *mongoUser) Key() string {
	return u.EKey
}

// Sets the user current action, uses u.Link
func (u *mongoUser) SetCurrentAction(a models.Action) error {
	return u.Schema().Link(u, a, currentAction)
}

// Retrieves the current action
func (u *mongoUser) CurrentAction(a data.Access) (models.Action, error) {
	if a.Type() != u.DBType() {
		return nil, data.ErrInvalidDBType
	}

	m, err := a.ModelFor(models.ActionKind)
	if err != nil {
		return nil, models.UndefinedKindError(models.ActionKind)
	}

	action, ok := m.(models.Action)
	if !ok {
		return nil, models.CastError(models.ActionKind)
	}

	if u.CanRead(a.Client()) {
		if mongo.EmptyID(u.CurrentActionID) {
			return nil, models.ErrEmptyRelationship
		}

		action.SetID(u.CurrentActionID)
		return action, a.PopulateByID(action)
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (u *mongoUser) SetCurrentActionable(a models.Actionable) error {
	if !data.Compatible(u, a) {
		return data.ErrIncompatibleModels
	}

	id, ok := a.ID().(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	u.ActionableKind = a.Kind()
	u.ActionableID = id

	return nil
}

func (u *mongoUser) CurrentActionable(a data.Access) (models.Actionable, error) {
	if u.DBType() != a.Type() {
		return nil, data.ErrInvalidDBType
	}

	if u.ActionableKind == "" || mongo.EmptyID(u.ActionableID) {
		return nil, models.ErrEmptyRelationship
	}

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

func (u *mongoUser) ClearCurrentActionable() {
	u.ActionableKind = ""
	u.ActionableID = *new(bson.ObjectId)
}

func (u *mongoUser) SetCalendar(c models.Calendar) error {
	return u.Schema().Link(u, c, calendar)
}

func (u *mongoUser) Calendar(a data.Access) (models.Calendar, error) {
	if u.DBType() != a.Type() {
		return nil, data.ErrInvalidDBType
	}

	if mongo.EmptyID(u.ECalendarID) {
		return nil, models.ErrEmptyRelationship
	}

	if !u.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	m, err := a.ModelFor(models.CalendarKind)
	if err != nil {
		return nil, models.UndefinedKindError(models.CalendarKind)
	}

	c, ok := m.(models.Calendar)
	if !ok {
		return nil, models.CastError(models.CalendarKind)
	}

	c.SetID(u.ECalendarID)
	return c, a.PopulateByID(c)
}

// Sets the user's ontology, see u.Link for possible errors
func (u *mongoUser) SetOntology(o models.Ontology) error {
	return u.Schema().Link(u, o, ontology)
}

/*
	Retrieves user's current Ontology

	Can result in data.ErrInvalidDBType,
	an instance of models.ErrCast or models.ErrUndefinedKind

	It can also result it any PopulateByID error or, finally
	a data.ErrAccessDenial error

	If the user does not have an ontology this returns
	models.ErrEmptyRelationship

	Ir an error has occured, the ontology returned is nil, do not
	inspect to be able even to inspect the id of the returned ontology
	if there is an error
*/
func (u *mongoUser) Ontology(a data.Access) (models.Ontology, error) {
	if a.Type() != u.DBType() {
		return nil, data.ErrInvalidDBType
	}

	m, err := a.ModelFor(models.OntologyKind)
	if err != nil { // data.ErrUndefinedKind
		return nil, models.UndefinedKindError(models.OntologyKind)
	}

	o, ok := m.(models.Ontology)
	if !ok {
		return nil, models.CastError(models.OntologyKind)
	}

	if !u.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	if mongo.EmptyID(u.EOntologyID) {
		return nil, models.ErrEmptyRelationship
	}

	o.SetID(u.EOntologyID)
	return o, a.PopulateByID(o)
}

func (u *mongoUser) IncludeEvent(e models.Event) error {
	return u.Schema().Link(u, e, events)
}

func (u *mongoUser) ExcludeEvent(e models.Event) error {
	return u.Schema().Unlink(u, e, events)
}

func (u *mongoUser) Events(a data.Access) (data.ModelIterator, error) {
	if u.CanRead(a.Client()) {
		return mongo.NewIDIter(u.EventIDs, a), nil
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (u *mongoUser) IncludeTask(t models.Task) error {
	return u.Schema().Link(u, t, tasks)
}

func (u *mongoUser) ExcludeTask(t models.Task) error {
	return u.Schema().Unlink(u, t, tasks)
}

func (u *mongoUser) Tasks(a data.Access) (data.ModelIterator, error) {
	if u.CanRead(a.Client()) {
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
	return u.Schema().Link(u, r, routines)
}

func (u *mongoUser) ExcludeRoutine(r models.Routine) error {
	return u.Schema().Unlink(u, r, routines)
}
