/*
	Package user contains implementations for elos' user model.

	The user model is the entry point for the elos application layer in
	the sense that you can get to all other models (theoretically) from
	a user's base model.

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

	Another valid method of instaniation is the Create function.

		var store data.Store
		u, err := user.Create(s, data.AttrMap{
			"name": "Nick Landolfi",
		}

	If you know what you are doing use New though. Create is for conveneince,
	mostly for the case you already have the data from another source, and you
	want it to all be good or not. Create isn't that great.

	The user package explicitly does NOT export the types of the underlying
	data structures. The user package should be used at the models.User level.
*/
package user
