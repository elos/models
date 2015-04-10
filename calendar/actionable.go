package calendar

import (
	"errors"

	"github.com/elos/data"
	"github.com/elos/models"
)

func NextAction(c models.Calendar, store models.Store) (action models.Action, err error) {
	current, err := c.CurrentFixture(store)
	if err != nil {
		return
	}

	action, err = current.NextAction(store)
	return
}

func StartAction(c models.Calendar, store models.Store, action models.Action) error {
	actionable, err := action.Actionable(store)
	if err != nil {
		return err
	}

	actionFixture, ok := actionable.(models.Fixture)
	if !ok {
		return errors.New("Actionable is not a fixture")
	}

	return c.SetCurrentFixture(actionFixture)
}

func CompleteAction(c models.Calendar, store models.Store, action models.Action) error {
	actionable, err := action.Actionable(store)
	if err != nil {
		return err
	}

	actionFixture, ok := actionable.(models.Fixture)
	if !ok {
		return errors.New("Actionable is not a fixture") // Invariant 1
	}

	cFixture, err := c.CurrentFixture(store)
	if err != nil {
		return err
	}

	if !data.EqualModels(actionFixture, cFixture) {
		return errors.New("Action's fixture does not match calendar's current fixture") // Invariant 2
	}

	if err = c.Schema().Unlink(c, cFixture, currentFixture); err != nil {
		return err
	}

	return cFixture.CompleteAction(store, action)
}
