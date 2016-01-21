package access

import (
	"errors"
	"log"

	"github.com/elos/data"
	"github.com/elos/models"
)

var (
	ErrBadPublic  = errors.New("bad public credential")
	ErrBadPrivate = errors.New("bad private credential")
)

func Authenticate(db data.DB, public, private string) (*models.Credential, error) {
	c := models.NewCredential()

	if err := db.PopulateByField("public", public, c); err != nil {
		return nil, ErrBadPublic
	}

	if challenge(c, private) {
		return c, nil
	}

	return nil, ErrBadPrivate
}

func challenge(c *models.Credential, private string) bool {
	switch c.Spec {
	default:
		log.Printf("Unrecognized credential spec: '%s'", c.Spec)
		fallthrough
	case "password":
		// eventually need encryption here
		return private == c.Private
	}
}
