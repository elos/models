package models

import (
	"github.com/elos/data"
	"gopkg.in/mgo.v2/bson"
)

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

func (o *UserOwned) User(a data.Access, u User) error {
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
