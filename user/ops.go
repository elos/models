package user

import "github.com/elos/models"

func NewCalendar(s models.Store, u models.User) error {
	if _, err := u.Calendar(s); err == nil {
		return nil
	} else if err != models.ErrEmptyRelationship {
		return err
	}

	c := s.Calendar()

	if err := s.Save(c); err != nil {
		return err
	}

	if err := u.SetCalendar(c); err != nil {
		return err
	}

	return s.Save(u)
}
