package user

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"github.com/elos/stack/util"
	"gopkg.in/mgo.v2/bson"
)

var (
	Events      data.LinkName = models.UserEvents
	Tasks       data.LinkName = models.UserTasks
	CurrentTask data.LinkName = models.UserCurrentTask
)

var (
	kind    data.Kind   = models.UserKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion
)

func NewM(s data.Store) (data.Model, error) {
	return New(s)
}

// Returns a new empty user struct.
// Note, if the DBType of the data.Store
// has not been implemented, it will return
// and data.ErrInvalidDBType
func New(s data.Store) (models.User, error) {
	switch s.Type() {
	case mongo.DBType:
		return &mongoUser{}, nil
	default:
		return nil, data.ErrInvalidDBType
	}
}

// Creates a new models.User with the attributes supplied in
// the second argument.
// Create will currently extrapolate "id", "created_at", and "name".
func Create(s data.Store, a data.AttrMap) (models.User, error) {
	user, err := New(s)
	if err != nil {
		return user, err
	}

	if id, ok := a["id"].(bson.ObjectId); ok {
		user.SetID(id)
	} else {
		user.SetID(mongo.NewObjectID().(bson.ObjectId))
	}

	if ca, ok := a["created_at"].(time.Time); ok {
		user.SetCreatedAt(ca)
	} else {
		user.SetCreatedAt(time.Now())
	}

	if n, ok := a["name"].(string); ok {
		user.SetName(n)
	}

	user.SetKey(util.RandomString(64))

	if err := s.Save(user); err != nil {
		return user, err
	} else {
		return user, nil
	}
}

// Validates user, the first return value determines
// overall validity. If the models is invalid the second
// return value can be insepcted for why
func Validate(u models.User) (bool, error) {
	if u.Name() == "" {
		return false, data.NewAttrError("name", "be present")
	}

	if u.Key() == "" {
		return false, data.NewAttrError("key", "be present")
	}

	return true, nil
}
