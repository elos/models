package fixture

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/shared"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoFixture struct {
	mongo.Model           `bson:",inline"`
	mongo.Named           `bson:",inline"`
	mongo.Timed           `bson:",inline"`
	shared.MongoUserOwned `bson:",inline"`

	EActionableKind data.Kind     `json:"actionable_kind" bson:"actionable_kind"`
	EActionableID   bson.ObjectId `json:"actionable_id" bson:"actionable_id,omitempty"`
	EEventableKind  data.Kind     `json:"eventable_kind" bson:"eventable_kind"`
	EEventableID    bson.ObjectId `json:"eventable_id" bson:"eventable_id,omitempty"`

	EDescription    string      `json:"decription" bson:"description"`
	ERank           int         `json:"rank" bson:"rank"`
	ELabel          bool        `json:"label" bson:"label"`
	EExpires        time.Time   `json:"expires" bson:"expires"`
	EDateExceptions []time.Time `json:"date_exceptions" bson:"date_exceptions"`

	EScheduleID bson.ObjectId `json:"schedule_id" bson:"schedule_id,omitempty"`
	EActionIDs  mongo.IDSet   `json:"action_ids" bson:"action_ids"`
	EEventIDs   mongo.IDSet   `json:"event_ids" bson:"event_ids"`
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

func (f *mongoFixture) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		return f.SetUserID(m.ID())
	case Schedule:
		f.EScheduleID = m.ID().(bson.ObjectId)
		return nil
	case Actions:
		f.EActionIDs = mongo.AddID(f.EActionIDs, m.ID().(bson.ObjectId))
	case Events:
		f.EEventIDs = mongo.AddID(f.EEventIDs, m.ID().(bson.ObjectId))
	default:
		return data.NewLinkError(f, m, l)
	}

	return nil
}

func (f *mongoFixture) Unlink(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		f.DropUserID()
	case Schedule:
		if f.EScheduleID == m.ID().(bson.ObjectId) {
			f.EScheduleID = *new(bson.ObjectId)
		}
	case Actions:
		f.EActionIDs = mongo.DropID(f.EActionIDs, m.ID().(bson.ObjectId))
	case Events:
		f.EEventIDs = mongo.DropID(f.EEventIDs, m.ID().(bson.ObjectId))
	default:
		return data.NewLinkError(f, m, l)
	}
	return nil
}

func (f *mongoFixture) SetUser(u models.User) error {
	return f.Schema().Link(f, u, User)
}

func (f *mongoFixture) NextAction(a data.Access) (models.Action, error) {
	return NextAction(f, a)
}

func (f *mongoFixture) StartAction(access data.Access, action models.Action) error {
	return nil
}

func (f *mongoFixture) CompleteAction(access data.Access, action models.Action) error {
	if _, present := f.EActionIDs.IndexID(action.ID().(bson.ObjectId)); !present {
		return data.ErrNotFound
	}

	// fixture can do any clean up, none right now

	action.Complete()
	return access.Save(action)
}

func (f *mongoFixture) NextEvent(a data.Access) (models.Event, error) {
	return NextEvent(a, f)
}

// Setting the fixture's actionable to itself is a no-op
func (f *mongoFixture) SetActionable(a models.Actionable) {
	if a.ID() == f.ID() {
		return
	}

	f.EActionableKind = a.Kind()
	f.EActionableID = a.ID().(bson.ObjectId)
}

func (f *mongoFixture) Actionable(a data.Access) (models.Actionable, error) {
	if !f.HasActionable() {
		return nil, models.ErrEmptyRelationship
	}

	m, err := a.ModelFor(f.EActionableKind)
	if err != nil {
		return nil, err
	}

	m.SetID(f.EActionableID)

	err = a.PopulateByID(m)

	return m.(models.Actionable), err
}

func (f *mongoFixture) DropActionable() {
	f.EActionableKind = data.Kind("")
	f.EActionableID = *new(bson.ObjectId)
}

func (f *mongoFixture) HasActionable() bool {
	return !mongo.EmptyID(f.EActionableID) && !data.EmptyKind(f.EActionableKind)
}

// setting to itself is a no-op
func (f *mongoFixture) SetEventable(e models.Eventable) {
	if f.ID() == e.ID() {
		return
	}

	f.EEventableKind = e.Kind()
	f.EEventableID = e.ID().(bson.ObjectId)
}

func (f *mongoFixture) Eventable(a data.Access) (models.Eventable, error) {
	m, err := a.ModelFor(f.EEventableKind)
	if err != nil {
		return nil, err
	}

	m.SetID(f.EEventableID)
	err = a.PopulateByID(m)
	return m.(models.Eventable), err
}

func (f *mongoFixture) DropEventable() {
	f.EEventableKind = data.Kind("")
	f.EEventableID = *new(bson.ObjectId)
}

func (f *mongoFixture) HasEventable() bool {
	return !mongo.EmptyID(f.EEventableID) && !data.EmptyKind(f.EEventableKind)
}

func (f *mongoFixture) SetDescription(s string) {
	f.EDescription = s
}

func (f *mongoFixture) Description() string {
	return f.EDescription
}

func (f *mongoFixture) SetRank(i int) {
	f.ERank = i
}

func (f *mongoFixture) Rank() int {
	return f.ERank
}

func (f *mongoFixture) SetLabel(b bool) {
	f.ELabel = b
}

func (f *mongoFixture) Label() bool {
	return f.ELabel
}

func (f *mongoFixture) AllDay() bool {
	return f.Label()
}

func (f *mongoFixture) SetExpires(t time.Time) {
	f.EExpires = t
}

func (f *mongoFixture) Expires() time.Time {
	return f.EExpires
}

func (f *mongoFixture) Expired() bool {
	return Expired(f)
}

func (f *mongoFixture) AddDateException(t time.Time) {
	f.EDateExceptions = append(f.EDateExceptions, t)
}

func (f *mongoFixture) DateExceptions() []time.Time {
	return f.EDateExceptions
}

func (f *mongoFixture) ShouldOmitOnDate(t time.Time) bool {
	return ShouldOmitOnDate(f, t)
}

func (f *mongoFixture) SetSchedule(s models.Schedule) error {
	return f.Schema().Link(f, s, Schedule)
}

func (f *mongoFixture) Schedule(a data.Access) (models.Schedule, error) {
	m, _ := a.ModelFor(models.ScheduleKind)
	s := m.(models.Schedule)
	s.SetID(f.EScheduleID)
	return s, a.PopulateByID(s)
}

func (f *mongoFixture) IncludeAction(a models.Action) error {
	return f.Schema().Link(f, a, Actions)
}

func (f *mongoFixture) ExcludeAction(a models.Action) error {
	return f.Schema().Unlink(f, a, Actions)
}

func (u *mongoFixture) ActionsIter(a data.Access) (data.ModelIterator, error) {
	if !u.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	return mongo.NewIDIter(u.EActionIDs, a), nil
}

func (u *mongoFixture) Actions(a data.Access) ([]models.Action, error) {
	if !u.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	actions := make([]models.Action, 0)
	iter, err := u.ActionsIter(a)
	if err != nil {
		return actions, err
	}

	m, err := a.ModelFor(models.ActionKind)
	if err != nil {
		return actions, err
	}

	for iter.Next(m) {
		action, ok := m.(models.Action)
		if !ok {
			return actions, models.CastError(models.ActionKind)
		}

		actions = append(actions, action)

		m, err = a.ModelFor(models.ActionKind)
		if err != nil {
			return actions, err
		}
	}

	return actions, nil
}

func (f *mongoFixture) IncludeEvent(e models.Event) error {
	return f.Schema().Link(f, e, Events)
}

func (f *mongoFixture) ExcludeEvent(e models.Event) error {
	return f.Schema().Unlink(f, e, Events)
}

func (u *mongoFixture) EventsIter(a data.Access) (data.ModelIterator, error) {
	if !u.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	return mongo.NewIDIter(u.EEventIDs, a), nil
}

func (u *mongoFixture) Events(a data.Access) ([]models.Event, error) {
	if !u.CanRead(a.Client()) {
		return nil, data.ErrAccessDenial
	}

	events := make([]models.Event, 0)

	iter, err := u.EventsIter(a)
	if err != nil {
		return events, err
	}

	m, err := a.ModelFor(models.EventKind)
	if err != nil {
		return events, err
	}

	for iter.Next(m) {
		e, ok := m.(models.Event)
		if !ok {
			return events, models.CastError(models.EventKind)
		}

		events = append(events, e)

		m, err = a.ModelFor(models.EventKind)
		if err != nil {
			return events, err
		}
	}

	return events, nil
}

func (f *mongoFixture) Conflicts(other models.Fixture) bool {
	return Conflicts(f, other)
}

func (f *mongoFixture) Order(other models.Fixture) (models.Fixture, models.Fixture) {
	return Order(f, other)
}

func (f *mongoFixture) Before(other models.Fixture) bool {
	return Before(f, other)
}
