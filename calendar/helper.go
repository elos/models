package calendar

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

func NewBaseSchedule(a data.Access, c models.Calendar) error {
	_, err := c.BaseSchedule(a)

	// meaning schedule already exists
	if err == nil {
		return nil
	}

	// meaning we can't help
	if err != models.ErrEmptyRelationship {
		return err
	}

	m, err := a.ModelFor(models.ScheduleKind)
	if err != nil {
		return err
	}

	s, ok := m.(models.Schedule)
	if !ok {
		return models.CastError(models.ScheduleKind)
	}

	if err = a.Save(s); err != nil {
		return err
	}

	if err = c.SetBaseSchedule(s); err != nil {
		return err
	}

	return a.Save(c)
}
