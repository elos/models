package models

import "time"

func NewSessionForUser(u *User) *Session {
	s := &Session{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Token:        NewToken(),
		ExpiresAfter: 3600,
	}

	s.SetUser(u)

	return s
}

func (s *Session) Valid() bool {
	return time.Now().Sub(s.CreatedAt.Add(time.Duration(s.ExpiresAfter)*time.Second)) < 0
}
