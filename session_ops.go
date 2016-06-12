package models

import (
	"time"

	"github.com/elos/data"
)

// TODO(nclandolfi): need to ensure the generated token is unique, can take the risk for now though
func NewSessionForUser(u *User) *Session {
	//log.Print("don't use") // <-- lol, what? (might be because session has a credential, so that if a credential is cancelled we can cancel the session. This is not currently implemented)

	s := &Session{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Token:        NewToken(),
		ExpiresAfter: 3600,
	}

	s.SetOwner(u)

	return s
}

func SessionForToken(db data.DB, token string) (*Session, error) {
	session := NewSession()
	if err := db.PopulateByField("token", token, session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Session) Valid() bool {
	return time.Now().Sub(s.CreatedAt.Add(time.Duration(s.ExpiresAfter)*time.Second)) < 0
}

func (s *Session) Expires() time.Time {
	return s.CreatedAt.Add(time.Duration(s.ExpiresAfter) * time.Second)
}
