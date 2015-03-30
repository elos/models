package fixture

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

func NextEvent(a data.Access, f models.Fixture) (models.Event, error) {
	// Allow an eventable to hijack
	if f.HasEventable() {
		e, err := f.Eventable(a)
		if err != nil {
			return nil, err
		}

		return e.NextEvent(a)
	}

	m, err := a.ModelFor(models.EventKind)
	if err != nil {
		return nil, err
	}

	event, ok := m.(models.Event)
	if !ok {
		return nil, models.CastError(models.EventKind)
	}

	event.SetName(f.Name())
	u, err := a.Unmarshal(models.UserKind, data.AttrMap{
		"id": f.UserID().(bson.ObjectId).Hex(),
	})
	if err != nil {
		return nil, err
	}
	event.SetUser(u.(models.User))
	f.IncludeEvent(event)

	a.Save(u)
	a.Save(event)
	a.Save(f)

	return event, nil
}
