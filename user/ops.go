package user

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

func NewCalendar(a data.Access, u models.User) error {
	_, err := u.Calendar(a)

	if err == nil {
		return nil
	}

	if err != models.ErrEmptyRelationship {
		return err
	}

	m, err := a.ModelFor(models.CalendarKind)
	if err != nil {
		return err
	}

	c, ok := m.(models.Calendar)
	if !ok {
		return models.CastError(models.CalendarKind)
	}

	if err = a.Save(c); err != nil {
		return err
	}

	if err = u.SetCalendar(c); err != nil {
		return err
	}

	return a.Save(u)
}
