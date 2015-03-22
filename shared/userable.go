package shared

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

// MongoUserOwned implements the Userable interface for mongo models
type MongoUserOwned struct {
	EUserID bson.ObjectId `json:"user_id" bson:"user_id,omitempty"`
}

func (o *MongoUserOwned) SetUserID(id data.ID) error {
	id, ok := id.(bson.ObjectId)
	if !ok {
		return data.ErrInvalidID
	}

	o.EUserID = id.(bson.ObjectId)
	return nil
}

func (o *MongoUserOwned) DropUserID() {
	o.EUserID = *new(bson.ObjectId)
}

func (o *MongoUserOwned) UserID() data.ID {
	return o.EUserID
}

func (o *MongoUserOwned) User(a data.Access, u models.User) error {
	u.SetID(o.EUserID)
	return a.PopulateByID(u)
}

func (o *MongoUserOwned) Concerned() []data.ID {
	concerns := make([]data.ID, 1)
	concerns[0] = o.UserID()
	return concerns
}

func (o *MongoUserOwned) CanRead(c data.Client) bool {
	if c.Kind() != models.UserKind {
		return false
	}

	if o.UserID().Valid() && c.ID() != o.UserID() {
		return false
	}

	return true
}

func (o *MongoUserOwned) CanWrite(c data.Client) bool {
	return o.CanRead(c)
}