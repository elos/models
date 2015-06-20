package models

import (
	"errors"

	"github.com/elos/data"
)

func Authenticate(db data.DB, public, private string) (*Credential, error) {
	credentialsIter, err := db.NewQuery(CredentialKind).Select(data.AttrMap{"public": public}).Execute()
	if err != nil {
		return nil, err
	}

	credential := NewCredential()
	credentialsIter.Next(credential)

	if credential.Challenge(private) {
		return credential, nil
	}

	return nil, errors.New("challenge failed")
}

type Property interface {
	data.Record
	Owner(data.DB) (*User, error)
}

type AccessLevel int

const (
	Ownership = iota
	Write
	Read
	Discovery
	None
)

var AccessLevels = map[string]AccessLevel{
	"write":     Write,
	"read":      Read,
	"discovery": Discovery,
}

func (u *User) Owns(db data.DB, property Property) (bool, error) {
	if owner, err := property.Owner(db); err == nil {
		if owner.ID().String() == u.ID().String() {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, err
	}
}

func (u *User) HasAccess(db data.DB, property Property) (bool, AccessLevel, error) {
	if owns, err := u.Owns(db, property); err == nil {
		if owns {
			return true, Ownership, nil
		}
	} else {
		// we need to be transparent about errors
		return false, None, err
	}

	// we don't automatically return false if u.Owns returns false
	// false cause the user may still have access through a group

	groups, err := u.Groups(db)
	if err != nil {
		return false, None, err
	}

	for _, group := range groups {
		if access, level, err := group.HasAccess(db, property); err == nil {
			if access {
				return true, level, nil
			}
		} else {
			return false, None, err
		}
	}

	return false, None, nil
}

func (u *User) CanRead(db data.DB, property Property) (bool, error) {
	access, level, err := u.HasAccess(db, property)

	if err != nil {
		return false, err
	}

	if !access {
		return false, nil
	}

	return (level <= Read), nil
}

func (u *User) CanWrite(db data.DB, property Property) (bool, error) {
	access, level, err := u.HasAccess(db, property)

	if err != nil {
		return false, err
	}

	if !access {
		return false, nil
	}

	return (level <= Write), nil
}
