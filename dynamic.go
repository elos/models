package models

import (
	"fmt"

	"github.com/elos/data"
)

func ModelFor(k data.Kind) data.Record {
	switch k {

	case ActionKind:
		return NewAction()

	case AttributeKind:
		return NewAttribute()

	case CalendarKind:
		return NewCalendar()

	case ClassKind:
		return NewClass()

	case EventKind:
		return NewEvent()

	case FixtureKind:
		return NewFixture()

	case LinkKind:
		return NewLink()

	case ObjectKind:
		return NewObject()

	case OntologyKind:
		return NewOntology()

	case RelationKind:
		return NewRelation()

	case RoutineKind:
		return NewRoutine()

	case ScheduleKind:
		return NewSchedule()

	case SessionKind:
		return NewSession()

	case TaskKind:
		return NewTask()

	case TraitKind:
		return NewTrait()

	case UserKind:
		return NewUser()

	default:
		panic(fmt.Sprintf("uknown kind: %s", k))
	}
}
