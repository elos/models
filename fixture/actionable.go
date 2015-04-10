package fixture

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

func NextAction(f models.Fixture, store models.Store) (models.Action, error) {
	// Allow an actionable to hijack
	if f.HasActionable() {
		if actionable, err := f.Actionable(store); err == nil {
			return actionable.NextAction(store)
		} else {
			return nil, err
		}
	}

	action := store.Action()

	user, err := f.User(store)
	if err != nil {
		return nil, err
	}

	action.SetName(f.Name())
	action.SetActionable(f)
	action.SetUser(user)
	f.IncludeAction(action)

	store.Save(user)
	store.Save(action)
	store.Save(f)

	return action, nil
}

func StartAction(f models.Fixture, access data.Access, action models.Action) error {
	return nil
}

func CompleteAction(f models.Fixture, access data.Access, action models.Action) error {
	return nil
}
