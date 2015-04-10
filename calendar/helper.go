package calendar

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

func NewBaseSchedule(store models.Store, c models.Calendar) error {
	if _, err := c.BaseSchedule(store); err == nil {
		// meaning schedule already exists
		return nil
	} else if err != models.ErrEmptyRelationship {
		// meaning we can't help
		return err
	}

	s := store.Schedule()

	if err := store.Save(s); err != nil {
		return err
	}

	if err := c.SetBaseSchedule(s); err != nil {
		return err
	}

	return store.Save(c)
}

func Find(s data.Store, id data.ID) (models.Calendar, error) {
	m, err := s.Unmarshal(models.CalendarKind, data.AttrMap{
		"id": id.(bson.ObjectId).Hex(),
	})

	if err != nil {
		return nil, err
	}

	c, ok := m.(models.Calendar)
	if !ok {
		return nil, models.CastError(models.CalendarKind)
	}

	return c, s.PopulateByID(c)
}
