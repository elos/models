package set

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

type mongoSet struct {
	mongo.Model      `bson:",inline"`
	mongo.Named      `bson:",inline"`
	models.UserOwned `bson:",inline"`

	ModelKind data.Kind   `json:"model_kind" bson:"model_kind"`
	ModelIDs  mongo.IDSet `json:"model_ids" bson:"model_ids"`
}

func (s *mongoSet) DBType() data.DBType {
	return mongo.DBType
}

func (s *mongoSet) Kind() data.Kind {
	return kind
}

func (s *mongoSet) Schema() data.Schema {
	return schema
}

func (s *mongoSet) Version() int {
	return version
}

func (s *mongoSet) Valid() bool {
	return true
}

func (s *mongoSet) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = s.UserID()
	return a
}

func (s *mongoSet) IncludeModel(m data.Model) error {
	return s.LinkModel(m)
}

func (s *mongoSet) ExcludeModel(m data.Model) error {
	return s.UnlinkModel(m)
}

func (s *mongoSet) ElementKind() data.Kind {
	return s.ModelKind
}

func (s *mongoSet) LinkModel(m data.Model) error {
	if s.ModelKind == "" {
		s.ModelKind = m.Kind()
	}

	if s.ModelKind != m.Kind() {
		return data.ErrUndefinedLink
	}

	s.ModelIDs = mongo.AddID(s.ModelIDs, m.ID().(bson.ObjectId))

	return nil
}

func (s *mongoSet) UnlinkModel(m data.Model) error {
	s.ModelIDs = mongo.DropID(s.ModelIDs, m.ID().(bson.ObjectId))
	return nil
}

func (s *mongoSet) Link(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		return s.SetUserID(s.ID())
	case Models:
		return s.LinkModel(m)
	default:
		return data.NewLinkError(s, m, l)
	}
}

func (s *mongoSet) Unlink(m data.Model, l data.Link) error {
	switch l.Name {
	case User:
		s.DropUserID()
	case Models:
		s.UnlinkModel(m)
	default:
		return data.ErrUndefinedLink
	}
	return nil
}
