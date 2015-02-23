package relationship

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

type mongoRelationship struct {
	models.MongoModel `bson:,inline"`
	models.UserOwned  `bson:",inline"`
	models.Named      `bson:",inline"`

	EClassID bson.ObjectId `json:"class_id" bson:"class_id,omitempty"`
	EOther   string        `json:"other" bson:"other"`
	EInverse string        `json:"inverse" bson:"inverse"`
}

func (t *mongoRelationship) Kind() data.Kind {
	return kind
}

func (t *mongoRelationship) Version() int {
	return version
}

func (t *mongoRelationship) Schema() data.Schema {
	return schema
}

func (t *mongoRelationship) SetUser(u models.User) error {
	return t.Schema().Link(t, u, User)
}

func (t *mongoRelationship) SetOther(s string) {
	t.EOther = s
}

func (t *mongoRelationship) Other() string {
	return t.EOther
}

func (t *mongoRelationship) SetInverse(s string) {
	t.EInverse = s
}

func (t *mongoRelationship) Inverse() string {
	return t.EInverse
}

func (t *mongoRelationship) SetClass(c models.Class) error {
	return t.Schema().Link(t, c, Class)
}

func (t *mongoRelationship) Class(a data.Access) (models.Class, error) {
	m, _ := a.ModelFor(models.ClassKind)
	c := m.(models.Class)

	if t.CanRead(a.Client()) {
		c.SetID(t.EClassID)
		err := a.PopulateByID(c)
		return c, err
	} else {
		return nil, data.ErrAccessDenial
	}
}

func (t *mongoRelationship) Link(m data.Model, l data.Link) error {
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

func (t *mongoRelationship) Unlink(m data.Model, l data.Link) error {
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
