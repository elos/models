package user

import (
	"fmt"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

/*
	Authenticates checks the id and key credentials for a user
	model.

	It returns a populated user model if the credentials are valid
	(meaning it could locate the user).

	It is not safe to ignore the error (this is intentional because
	we could possibly "locate" the user, without athenticating them).

		var store data.Store
		id := "23913123123"
		key := "asdfkljaskdfjasjdfjalsdfjadskfa"

		u, authed, err := Authenticate(s, id, key)

		if authed {
			// Only true of the authentication was successful, in this
			// scenario the user has been populated
		} else {
			// The User is not populated, and is nil! don't use it
			if err == data.ErrNotFound {
				// bad id
			} else if err.Error() == "Invalid key" { // we need a type for this (TODO)
				// bad key
			} else {
				// undefined
			}
		}
*/
func Authenticate(s data.Store, id string, key string) (models.User, bool, error) {
	user, err := Find(s, mongo.NewObjectIDFromHex(id))

	if err != nil {
		return nil, false, err
	}

	if user.Key() != key {
		return nil, false, fmt.Errorf("Invalid key")
	}

	return user, true, nil
}

/*
	Find locates a user by an id.

		var store data.Store
		id := "2341234"

		u, err := Find(s, id)

	The error could be data.ErrInvalidDBType, data.ErrInvalidID,
	or an error from store.PopulateByID
*/
func Find(s data.Store, id data.ID) (models.User, error) {
	user := New(s)

	id, ok := id.(bson.ObjectId)
	if !ok {
		return nil, data.ErrInvalidID
	}

	user.SetID(id)

	// Find a user that has specified id
	if err := s.PopulateByID(user); err != nil {
		return nil, err
	}

	return user, nil
}

/*
	FindBy locates a user by some field and it's value

		var store data.Store

		u, err := FindBy(s, "name", "Nick Landolfi")

	Error could be data.ErrInvalidDBType, or any PopulateByField
	error
*/
func FindBy(s data.Store, field string, value interface{}) (models.User, error) {
	user := New(s)
	return user, s.PopulateByField(field, value, user)
}

/*
	NewWithName instantiates and saves a new user model with
	the provided name.
*/
func NewWithName(s data.Store, n string) (models.User, error) {
	user, err := CreateAttrs(s, data.AttrMap{
		"name": n,
	})

	if err != nil {
		return user, err
	}

	return user, s.Save(user)
}
