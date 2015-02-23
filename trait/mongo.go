package trait

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

type mongoTrait struct {
	models.MongoModel `bson:,inline"`
	models.UserOwned  `bson:",inline"`
	models.Named      `bson:",inline"`

	EClassID bson.ObjectId `json:"class_id" bson:"class_id,omitempty"`
	EType    string        `json:"type" bson:"type"`
}

func (t *mongoTrait) Kind() data.Kind {
	return kind
}

func (t *mongoTrait) Version() int {
	return version
}

func (t *mongoTrait) Schema() data.Schema {
	return schema
}

func (t *mongoTrait) SetUser(u models.User) error {
	return t.Schema().Link(t, u, User)
}

func (t *mongoTrait) SetType(s string) {
	t.EType = s
}

func (t *mongoTrait) Type() string {
	return t.EType
}

func (t *mongoTrait) SetClass(c models.Class) error {
	return t.Schema().Link(t, c, Class)
}

func (t *mongoTrait) Class(a data.Access, c models.Class) error {
	if !data.Compatible(t, c) {
		return data.ErrIncompatibleModels
	}

	if t.CanRead(a.Client()) {
		c.SetID(t.EClassID)
		return a.PopulateByID(c)
	} else {
		return data.ErrAccessDenial
	}
}

func (t *mongoTrait) Link(m data.Model, l data.Link) error {
	if !data.Compatible(t, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		t.DropUserID()
	case Class:
		t.EClassID = m.ID().(bson.ObjectId)
	default:
		return data.NewLinkError(t, m, l)
	}
	return nil
}

func (t *mongoTrait) Unlink(m data.Model, l data.Link) error {
	if !data.Compatible(t, m) {
		return data.ErrIncompatibleModels
	}

	switch l.Name {
	case User:
		t.DropUserID()
	case Class:
		t.EClassID = *new(bson.ObjectId)
	default:
		return data.NewLinkError(t, m, l)
	}
	return nil
}
