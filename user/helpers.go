package user

import (
	"fmt"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

/*
	Authenticates a user, returning a populated user model
	If the second return value is true, the user's credentials have been validated
	otherwise, the user's credentials were malformed.
*/
func Authenticate(s data.Store, id string, key string) (data.Client, bool, error) {
	user, err := Find(s, mongo.NewObjectIDFromHex(id))

	if err != nil {
		return user, false, err
	}

	if user.Key() != key {
		return user, false, fmt.Errorf("Invalid key")
	}

	return user, true, nil
}

// Finds a user model by an id
func Find(s data.Store, id data.ID) (models.User, error) {
	user, err := New(s)
	if err != nil {
		return user, err
	}

	id, ok := id.(bson.ObjectId)
	if !ok {
		return user, data.ErrInvalidID
	}

	user.SetID(id)

	// Find a user that has specified id
	if err := s.PopulateByID(user); err != nil {
		return user, err
	}

	return user, nil
}

// Finds a user by some field and its value
func FindBy(s data.Store, field string, value interface{}) (models.User, error) {
	user, err := New(s)
	if err != nil {
		return user, err
	}

	if err = s.PopulateByField(field, value, user); err != nil {
		return user, err
	}

	return user, nil
}

func NewWithName(s data.Store, n string) (models.User, error) {
	user, err := Create(s, data.AttrMap{
		"name": n,
	})

	if err != nil {
		return user, err
	}

	return user, s.Save(user)
}
