package fixture

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

func Action(a data.Access, f models.Fixture) (models.Action, error) {
	// Allow an actionable to hijack
	if f.HasActionable() {
		actionable, err := f.Actionable(a)
		if err != nil {
			return nil, err
		}

		return actionable.NextAction(a)
	}

	m, err := a.ModelFor(models.ActionKind)
	if err != nil {
		return nil, err
	}

	action, ok := m.(models.Action)
	if !ok {
		return nil, models.CastError(models.ActionKind)
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
	// Allow an eventable to hijack
	if f.HasEventable() {
		e, err := f.Eventable(a)
		if err != nil {
			return nil, err
		}

		return e.Event(a)
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
	event.SetUserID(f.UserID())
	f.IncludeEvent(event)

	a.Save(event)
	a.Save(f)

	return event, nil
}

func Expired(f models.Fixture) bool {
	return f.Expires().Before(time.Now())
}

func CanonicalYearlyDay(t time.Time) int {
	return 100*int(t.Month()) + t.Day()
}

var cDate = CanonicalYearlyDay

// only checks to month/day/year validity
func OmitOnDate(f models.Fixture, t time.Time) bool {
	exceptions := f.DateExceptions()
	for i := range exceptions {
		exception := exceptions[i]
		if cDate(t) == cDate(exception) && t.Year() == exception.Year() {
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
