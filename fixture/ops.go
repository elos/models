package fixture

import (
	"errors"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

type ActionFixture struct {
	models.Fixture
	data.Access
}

func NewActionFixture(a data.Access, f models.Fixture) *ActionFixture {
	return &ActionFixture{
		Access:  a,
		Fixture: f,
	}
}

// Actionable
func (f *ActionFixture) Next() (models.Action, bool) {
	return nil, false
	// hold that thought
}

func Action(a data.Access, f models.Fixture) (models.Action, error) {
	m, err := a.ModelFor(models.ActionKind)
	if err != nil {
		return nil, err
	}

	action, ok := m.(models.Action)
	if !ok {
		return nil, errors.New("TODO")
	}

	action.SetName(f.Name())
	action.SetActionable(f)
	action.SetUserID(f.UserID())
	f.IncludeAction(action)

	a.Save(action)
	a.Save(f)

	return action, nil
}

func Event(a data.Access, f models.Fixture) (models.Event, error) {
	m, err := a.ModelFor(models.EventKind)
	if err != nil {
		return nil, err
	}

	event, ok := m.(models.Event)
	if !ok {
		return nil, errors.New("TODO")
	}

	event.SetName(f.Name())
	event.SetUserID(f.UserID())
	f.IncludeEvent(event)

	a.Save(event)
	a.Save(f)

	return event, nil
}

func Expired(f models.Fixture) bool {
	return f.Expires().Before(time.Now())
}

func canonDate(t time.Time) int {
	return 100*int(t.Month()) + t.Day()
}

// only checks to month/day/year validity
func OmitOnDate(f models.Fixture, t time.Time) bool {
	exceptions := f.DateExceptions()
	for i := range exceptions {
		exception := exceptions[i]
		if canonDate(t) == canonDate(exception) && t.Year() == exception.Year() {
			return true
		}
	}

	return false
}

func Sort(f1 models.Fixture, f2 models.Fixture) (first models.Fixture, second models.Fixture) {
	if f1.StartTime().Before(f2.StartTime()) {
		first = f1
		second = f2
	} else { // f2 starts before f1
		first = f2
		second = f1
	}
	return
}

func Conflicting(f1 models.Fixture, f2 models.Fixture) bool {
	first, second := Sort(f1, f2)
	return second.StartTime().Before(first.EndTime())
}
