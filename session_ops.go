package models

import (
	"log"
	"time"
)

func NewSessionForUser(u *User) *Session {
	log.Print("don't use")
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
