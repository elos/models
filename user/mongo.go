package user

import (
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

	ECalendarID     bson.ObjectId `json:"calendar_id"       bson:"calendar_id,omitempty"`
	EOntologyID     bson.ObjectId `json:"ontology_id"       bson:"ontology_id,omitempty"`
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
	case actions:
		u.ActionIDs = mongo.AddID(u.ActionIDs, id)
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
	case actions:
		u.ActionIDs = mongo.DropID(u.ActionIDs, id)
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
func (u *mongoUser) CurrentAction(s models.Store) (models.Action, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	if mongo.EmptyID(u.CurrentActionID) {
		return nil, models.ErrEmptyRelationship
	}

	action := s.Action()
	action.SetID(u.CurrentActionID)
	return action, s.PopulateByID(action)
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

func (u *mongoUser) CurrentActionable(s models.Store) (models.Actionable, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	if u.ActionableKind == "" || mongo.EmptyID(u.ActionableID) {
		return nil, models.ErrEmptyRelationship
	}

	m, err := s.ModelFor(u.ActionableKind)
	if err != nil {
		return nil, err
	}

	actionable, ok := m.(models.Actionable)
	if !ok {
		return nil, models.CastError("actionable")
	}

	actionable.SetID(u.ActionableID)
	return actionable, s.PopulateByID(actionable)
}

func (u *mongoUser) ClearCurrentActionable() {
	u.ActionableKind = ""
	u.ActionableID = *new(bson.ObjectId)
}

func (u *mongoUser) SetCalendar(c models.Calendar) error {
	return u.Schema().Link(u, c, calendar)
}

func (u *mongoUser) Calendar(s models.Store) (models.Calendar, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	if mongo.EmptyID(u.ECalendarID) {
		return nil, models.ErrEmptyRelationship
	}

	c := s.Calendar()
	c.SetID(u.ECalendarID)
	return c, s.PopulateByID(c)
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
func (u *mongoUser) Ontology(s models.Store) (models.Ontology, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	if mongo.EmptyID(u.EOntologyID) {
		return nil, models.ErrEmptyRelationship
	}

	o := s.Ontology()
	o.SetID(u.EOntologyID)
	return o, s.PopulateByID(o)
}

func (u *mongoUser) IncludeAction(a models.Action) error {
	return u.Schema().Link(u, a, actions)
}

func (u *mongoUser) ExcludeAction(a models.Action) error {
	return u.Schema().Unlink(u, a, actions)
}

func (u *mongoUser) ActionsIter(s models.Store) (data.ModelIterator, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(u.ActionIDs, s), nil
}

func (u *mongoUser) Actions(s models.Store) ([]models.Action, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	actions := make([]models.Action, 0)
	iter := mongo.NewIDIter(u.ActionIDs, s)

	action := s.Action()
	for iter.Next(action) {
		actions = append(actions, action)
		action = s.Action()
	}

	return actions, iter.Close()
}

func (u *mongoUser) IncludeEvent(e models.Event) error {
	return u.Schema().Link(u, e, events)
}

func (u *mongoUser) ExcludeEvent(e models.Event) error {
	return u.Schema().Unlink(u, e, events)
}

func (u *mongoUser) EventsIter(s models.Store) (data.ModelIterator, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(u.EventIDs, s), nil
}

func (u *mongoUser) Events(s models.Store) ([]models.Event, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	events := make([]models.Event, 0)

	iter := mongo.NewIDIter(u.EventIDs, s)

	event := s.Event()
	for iter.Next(event) {
		events = append(events, event)
		event = s.Event()
	}

	return events, iter.Close()
}

func (u *mongoUser) IncludeTask(t models.Task) error {
	return u.Schema().Link(u, t, tasks)
}

func (u *mongoUser) ExcludeTask(t models.Task) error {
	return u.Schema().Unlink(u, t, tasks)
}

func (u *mongoUser) TasksIter(s models.Store) (data.ModelIterator, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(u.TaskIDs, s), nil
}

func (u *mongoUser) Tasks(s models.Store) ([]models.Task, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	tasks := make([]models.Task, 0)
	iter := mongo.NewIDIter(u.TaskIDs, s)
	task := s.Task()
	for iter.Next(task) {
		tasks = append(tasks, task)
		task = s.Task()
	}

	return tasks, nil
}

func (u *mongoUser) IncludeRoutine(r models.Routine) error {
	return u.Schema().Link(u, r, routines)
}

func (u *mongoUser) ExcludeRoutine(r models.Routine) error {
	return u.Schema().Unlink(u, r, routines)
}

func (u *mongoUser) RoutinesIter(s models.Store) (data.ModelIterator, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	return mongo.NewIDIter(u.RoutineIDs, s), nil
}

func (u *mongoUser) Routines(s models.Store) ([]models.Routine, error) {
	if !s.Compatible(u) {
		return nil, data.ErrInvalidDBType
	}

	routines := make([]models.Routine, 0)
	iter := mongo.NewIDIter(u.RoutineIDs, s)
	routine := s.Routine()
	for iter.Next(routine) {
		routines = append(routines, routine)
		routine = s.Routine()
	}

	return routines, nil
}
