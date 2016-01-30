package models

import (
	"log"
	"time"

	"github.com/elos/data"
)

func NewSessionForUser(u *User) *Session {
	log.Print("don't use") // <-- lol, what?
	s := &Session{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Token:        NewToken(),
		ExpiresAfter: 3600,
	}

	s.SetOwner(u)

	return s
}

func (s *Session) Valid() bool {
	return time.Now().Sub(s.CreatedAt.Add(time.Duration(s.ExpiresAfter)*time.Second)) < 0
}

func SessionForToken(db data.DB, token string) (*Session, error) {
	session := NewSession()
	if err := db.PopulateByField("token", token, session); err != nil {
		return nil, err
	}

	return session, nil
}
