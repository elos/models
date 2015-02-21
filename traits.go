package models

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoModel struct {
	MongoID     `bson:",inline"`
	Timestamped `bson:",inline"`
}

func (m *MongoModel) DBType() data.DBType {
	return mongo.DBType
}

func (m *MongoModel) Valid() bool {
	return true
}

type MongoID struct {
	EID bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

func (m *MongoID) SetID(id data.ID) error {
	vid, ok := id.(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}
	m.EID = vid
	return nil
}

func (m *MongoID) ID() data.ID {
	return m.EID
}

type Loaded struct {
	loadedAt time.Time
}

func (l *Loaded) SetLoadedAt(t time.Time) {
	l.loadedAt = t
}

func (l *Loaded) LoadedAt() time.Time {
	return l.loadedAt
}

type Timestamped struct {
	ECreatedAt time.Time `json:"created_at" bson:"created_at"`
	EUpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (t *Timestamped) SetCreatedAt(ca time.Time) {
	t.ECreatedAt = ca
}

func (t *Timestamped) CreatedAt() time.Time {
	return t.ECreatedAt
}

func (t *Timestamped) SetUpdatedAt(ua time.Time) {
	t.EUpdatedAt = ua
}

func (t *Timestamped) UpdatedAt() time.Time {
	return t.EUpdatedAt
}

type Named struct {
	EName string `json:"name" bson:"name"`
}

func (n *Named) SetName(name string) {
	n.EName = name
}

func (n *Named) Name() string {
	return n.EName
}

type Timed struct {
	EStartTime time.Time `json:"start_time" bson:"start_time"`
	EEndTime   time.Time `json:"end_time" bson:"end_time"`
}

func (t *Timed) StartTime() time.Time {
	return t.EStartTime
}

func (t *Timed) SetStartTime(st time.Time) {
	t.EStartTime = st
}

func (t *Timed) EndTime() time.Time {
	return t.EEndTime
}

func (t *Timed) SetEndTime(et time.Time) {
	t.EEndTime = et
}

type UserOwned struct {
	EUserID bson.ObjectId `json:"user_id" bson:"user_id,omitempty"`
}

func (o *UserOwned) SetUserID(id data.ID) error {
	id, ok := id.(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	o.EUserID = id.(bson.ObjectId)
	return nil
}

func (o *UserOwned) DropUserID() {
	o.EUserID = *new(bson.ObjectId)
}

func (o *UserOwned) UserID() data.ID {
	return o.EUserID
}

func (o *UserOwned) User(a *data.Access, u User) error {
	u.SetID(o.EUserID)
	return a.PopulateByID(u)
}

func (o *UserOwned) Concerned() []data.ID {
	concerns := make([]data.ID, 1)
	concerns[0] = o.UserID()
	return concerns
}

func (o *UserOwned) CanRead(c data.Client) bool {
	if c.Kind() != UserKind {
		return false
	}

	if o.UserID().Valid() && c.ID() != o.UserID() {
		return false
	}

	return true
}

func (o *UserOwned) CanWrite(c data.Client) bool {
	return o.CanRead(c)
}
