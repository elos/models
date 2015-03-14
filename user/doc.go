/*
	Package user contains implementations for elos' user model.

	The user model is the entry point for the elos application layer.

	When associating a user with a store, use the NewM function:

		var store data.Store
		store.Register("user", NewM)

	To initialize a user, you must provide a store. Remember that data.Access
	satisfies the data.Store interface, so if you want to restrict action pass
	that in. Remember that data.Access only restricts the actually database access
	so you may be allowed to work with the data structure in memory but unable
	to persist it. Anyway:

		var store data.Store
		u, err := user.New(s)

	New looks up the DBType of the store and returns the correct model implementation
	of a user. The only current implementation is for mongo. If the store is not supported
	New returns a data.ErrInvalidDBType error, and the user is nil. Always check errors!
	Note that New will set the ID, but not the key for a user (which is a required field).

	Another valid method of instaniation is the Create function. Create will look for an id
	attribute, a created_at attribute, and a name attriute. Create also generates and sets the
	key for a user.

		import (
			"time"
			"github.com/elos/data"
		)

		var store data.Store
		u, err := user.Create(s, data.AttrMap{
			"id": store.NewID(),
			"name": "Nick Landolfi",
			"created_at": time.Now(),
		})

	Most errors for create don't matter. But if the id is bad, Create complains.

	If you just need a new user model and you know what you are doing, then just use new.
	Create is a simple convenience mostly for the case you already have the data from another
	source, and you want it to all be good or not. Create isn't that great.
*/
package user
