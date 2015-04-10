package fixture

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

func NextEvent(f models.Fixture, store models.Store) (models.Event, error) {
	// Allow an eventable to hijack
	if f.HasEventable() {
		e, err := f.Eventable(store)
		if err != nil {
			return nil, err
		}

		return e.NextEvent(store)
	}

	event := store.Event()

	event.SetName(f.Name())
	u, err := store.Unmarshal(models.UserKind, data.AttrMap{
		"id": f.UserID().(bson.ObjectId).Hex(),
	})
	if err != nil {
		return nil, err
	}
	event.SetUser(u.(models.User))
	f.IncludeEvent(event)

	store.Save(u)
	store.Save(event)
	store.Save(f)

	return event, nil
}
