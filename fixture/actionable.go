package fixture

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

func NextAction(f models.Fixture, access data.Access) (models.Action, error) {
	// Allow an actionable to hijack
	if f.HasActionable() {
		if actionable, err := f.Actionable(access); err == nil {
			return actionable.NextAction(access)
		} else {
			return nil, err
		}
	}

	m, err := access.ModelFor(models.ActionKind)
	if err != nil {
		return nil, err
	}

	action, ok := m.(models.Action)
	if !ok {
		return nil, models.CastError(models.ActionKind)
	}

	user, err := f.User(access)
	if err != nil {
		return nil, err
	}

	action.SetName(f.Name())
	action.SetActionable(f)
	action.SetUser(user)
	f.IncludeAction(action)

	access.Save(user)
	access.Save(action)
	access.Save(f)

	return action, nil
}

func StartAction(f models.Fixture, access data.Access, action models.Action) error {
	return nil
}

func CompleteAction(f models.Fixture, access data.Access, action models.Action) error {
	return nil
}
