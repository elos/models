package user

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"github.com/elos/stack/util"
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

	id, present := a["id"]
	id, valid := id.(data.ID)
	if present && valid {
		if err := user.SetID(id.(data.ID)); err != nil {
			return user, err
		}
	} else {
		if err := user.SetID(s.NewID()); err != nil {
			return user, err
		}
	}

	ca, present := a["created_at"]
	ca, valid = ca.(time.Time)
	if present && valid {
		user.SetCreatedAt(ca.(time.Time))
	} else {
		user.SetCreatedAt(time.Now())
	}

	n, present := a["name"]
	n, valid = n.(string)
	if present && valid {
		user.SetName(n.(string))
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
