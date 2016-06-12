package access

import (
	"errors"
	"log"

	"github.com/elos/data"
	"github.com/elos/models"
)

const (
	publicField  = "public"
	passwordSpec = "password"
)

type (
	Error error

	ErrInternal   Error
	ErrBadPublic  Error
	ErrBadPrivate Error
)

var (
	errInternal   ErrInternal   = errors.New("access error: internal")
	errBadPublic  ErrBadPublic  = errors.New("access error: bad public credential")
	errBadPrivate ErrBadPrivate = errors.New("access error: bad private credential")
)

func Authenticate(db data.DB, public, private string) (*models.Credential, Error) {
	c := new(models.Credential)

	if err := db.PopulateByField(publicField, public, c); err != nil {
		switch err {
		case data.ErrNotFound:
			return nil, errBadPublic
		default:
			return nil, errInternal
		}
		return nil, errBadPublic
	}

	if challenge(c, private) {
		return c, nil
	}

	return nil, errBadPrivate
}

func challenge(c *models.Credential, private string) bool {
	switch c.Spec {
	default:
		log.Printf("Unrecognized credential spec: '%s'", c.Spec)
		fallthrough
	case passwordSpec:
		// eventually need encryption here
		return private == c.Private
	}
}
