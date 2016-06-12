package access

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/models"
)

// --- Logging Apparatus {{{

type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})

	WithPrefix(p string) Logger
}

type nullLogger struct{}

func (l *nullLogger) Print(v ...interface{})            {}
func (l *nullLogger) Printf(f string, v ...interface{}) {}
func (l *nullLogger) WithPrefix(p string) Logger        { return l }

var Log Logger = new(nullLogger)

// --- }}}

// --- CanCreate {{{

// CanCreate determines whether a user can create a record. It should be called
// directly prior to saving the new record. The following conditions must hold
// in order for a user to create the record:
//  * The record owner must be the user
func CanCreate(db data.DB, u *models.User, p Property) (bool, error) {
	l := Log.WithPrefix("access.CanCreate: ")

	if owner, err := p.Owner(db); err == models.ErrEmptyLink {
		l.Printf("User(%s) tried to create ownerless property", u.Id)
		return false, nil
	} else if err != nil {
		return false, errInternal
	} else {
		return data.Equivalent(u, owner), nil
	}
}

// --- }}}

// --- CanRead {{{

// CanRead determines whether a user can read a record. It should be called
// directly prior to returning the record to the user. One of the following
// conditions must hold in order for a user to read the record:
//  * The record is the user record, for the user requesting read.
//  * The record is property, owned by the user.
//  * The record is in the context, of a permission group with access level at
//    least Read
func CanRead(db data.DB, u *models.User, record data.Record) (bool, error) {
	l := Log.WithPrefix("access.CanRead ")

	// Assert the model is of type Property, if it isn't, it can only be read
	// by a user if it is _that_ user's record.
	prop, ok := record.(Property)
	if !ok {
		return data.Equivalent(u, record), nil
	}

	l.Print("checking if user owns property")
	if owner, err := prop.Owner(db); err != nil {
		if err == models.ErrEmptyLink {
			l.Printf("(%s, %s) without an owner!", record.Kind(), record.ID())
			return false, nil
		}

		return false, errInternal
	} else {
		if data.Equivalent(u, owner) {
			return true, nil
		}
	}

	l.Print("cehcking if user in any groups granted permission to property")
	if ok, err := grantedPermission(db, u, prop, Read); err != nil {
		return false, errInternal
	} else {
		return ok, nil
	}
}

// --- }}}

// --- CanWrite {{{

// CanWrite determines whether a user can write a record. It should be called
// directly prior to saving the updated record . The record must be mutable.
// One of the following conditions must hold in order for a user to write the record:
//  * The record is the user record, for the user requesting read.
//  * The record is property, owned by the user.
//  * The record is in the context, of a permission group with access level at
//    least Write
func CanWrite(db data.DB, u *models.User, record data.Record) (bool, error) {
	l := Log.WithPrefix("access.CanWrite: ")

	// Assert the model is of type Property, if it isn't, it can only be writtenj
	// by a user if it is _that_ user's record.
	prop, ok := record.(Property)
	if !ok {
		return data.Equivalent(u, record), nil
	}

	// If the model is one of the immutable records, reject write
	if _, immutable := ImmutableRecords[record.Kind()]; immutable {
		return false, nil
	}

	if owner, err := prop.Owner(db); err != nil {
		if err == models.ErrEmptyLink {
			log.Printf("(%s, %s) without an owner!", record.Kind(), record.ID())
			return false, nil
		}

		return false, errInternal
	} else {
		if data.Equivalent(u, owner) {
			return true, nil
		}
	}

	l.Print("checking if user in any groups granted permission to property")
	if ok, err := grantedPermission(db, u, prop, Write); err != nil {
		return false, errInternal
	} else {
		return ok, nil
	}
}

// --- }}}

func grantedPermission(db data.DB, u *models.User, p Property, l Level) (bool, error) {
	groups, err := u.Groups(db)

	if err != nil {
		return false, err
	}

	for _, g := range groups {
		contexts, err := g.Contexts(db)

		if err != nil {
			return false, err
		}

		for _, c := range contexts {
			if c.Contains(p) {
				if Level(g.Access) >= Write {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func CanDelete(db data.DB, u *models.User, record data.Record) (bool, error) {
	property, ok := record.(Property)

	if !ok {
		return data.Equivalent(u, record), nil
	}

	owner, err := property.Owner(db)

	if err != nil {
		return false, err
	}

	return data.Equivalent(u, owner), nil
}
