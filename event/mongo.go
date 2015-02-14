package event

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

type mongoEvent struct {
	models.MongoID     `bson:",inline"`
	models.Named       `bson:",inline"`
	models.Timestamped `bson:",inline"`
	models.Timed       `bson:",inline"`
	models.UserOwned   `bson:",inline"`
}

func (e *mongoEvent) Kind() data.Kind {
	return kind
}

func (e *mongoEvent) Schema() data.Schema {
	return schema
}

func (e *mongoEvent) Version() int {
	return version
}

func (e *mongoEvent) Valid() bool {
	valid, _ := Validate(e)
	return valid
}

func (u *mongoEvent) DBType() data.DBType {
	return mongo.DBType
}

func (e *mongoEvent) SetUser(u models.User) error {
	return e.Schema().Link(e, u, User)
}

func (e *mongoEvent) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		return e.SetUserID(m.ID())
	default:
		return data.ErrUndefinedLink
	}
}

func (e *mongoEvent) Unlink(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		e.DropUserID()
	default:
		return data.ErrUndefinedLink
	}

	return nil
}
