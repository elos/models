package user

import (
	"fmt"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	"github.com/elos/models"
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
func Authenticate(db data.DB, id string, key string) (*models.User, bool, error) {
	user, err := Find(db, mongo.NewObjectIDFromHex(id))

	if err != nil {
		return nil, false, err
	}

	if user.Key != key {
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
func Find(db data.DB, id data.ID) (*models.User, error) {
	user := models.NewUser()

	bid, err := mongo.ParseObjectID(id.String())
	if err != nil {
		return nil, err
	}

	user.SetID(data.ID(bid.Hex()))

	// Find a user that has specified id
	if err := db.PopulateByID(user); err != nil {
		return nil, err
	}

	return user, nil
}
