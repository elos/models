package models

import (
	"time"

	"github.com/elos/data"
)

func (c *Credential) Challenge(public, private string) bool {
	// eventually need to use spec mechanism

	if public != c.Public {
		return false
	}

	// eventually need encryption here
	if private != c.Private {
		return false
	}

	return true
}

func (c *Credential) NewSession(db data.DB, expiresAfter time.Duration) (*Session, error) {
	session := NewSession()
	session.SetID(db.NewID())

	user, err := c.Owner(db)
	if err != nil {
		return nil, err
	}

	if err := session.SetOwner(user); err != nil {
		return nil, err
	}
	if err := session.SetCredential(c); err != nil {
		return nil, err
	}

	now := time.Now()
	session.CreatedAt = now
	session.UpdatedAt = now
	session.ExpiresAfter = int(expiresAfter)

	user.IncludeSession(session)
	c.IncludeSession(session)

	if err := db.Save(session); err != nil {
		return nil, err
	}

	if err := db.Save(user); err != nil {
		return nil, err
	}

	if err := db.Save(c); err != nil {
		return nil, err
	}

	return session, nil
}
