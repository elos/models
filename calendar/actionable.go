package calendar

import (
	"errors"

	"github.com/elos/data"
	"github.com/elos/models"
)

func NextAction(c models.Calendar, a data.Access) (action models.Action, err error) {
	current, err := c.CurrentFixture(a)
	if err != nil {
		return
	}

	action, err = current.NextAction(a)

	return
}

func StartAction(c models.Calendar, access data.Access, action models.Action) error {
	actionable, err := action.Actionable(access)
	if err != nil {
		return err
	}

	actionFixture, ok := actionable.(models.Fixture)
	if !ok {
		return errors.New("Actionable is not a fixture")
	}

	return c.SetCurrentFixture(actionFixture)
}

func CompleteAction(c models.Calendar, access data.Access, action models.Action) error {
	actionable, err := action.Actionable(access)
	if err != nil {
		return err
	}

	actionFixture, ok := actionable.(models.Fixture)
	if !ok {
		return errors.New("Actionable is not a fixture") // Invariant 1
	}

	cFixture, err := c.CurrentFixture(access)
	if err != nil {
		return err
	}

	if !data.EqualModels(actionFixture, cFixture) {
		return errors.New("Action's fixture does not match calendar's current fixture") // Invariant 2
	}

	if err = c.Schema().Unlink(c, cFixture, currentFixture); err != nil {
		return err
	}

	return cFixture.CompleteAction(access, action)
}
