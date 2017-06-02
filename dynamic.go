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

	case ContextKind:
		return NewContext()

	case CredentialKind:
		return NewCredential()

	case DatumKind:
		return NewDatum()

	case EventKind:
		return NewEvent()

	case FixtureKind:
		return NewFixture()

	case GroupKind:
		return NewGroup()

	case HabitKind:
		return NewHabit()

	case IntegrationKind:
		return NewIntegration()

	case LinkKind:
		return NewLink()

	case LocationKind:
		return NewLocation()

	case MediaKind:
		return NewMedia()

	case ModelKind:
		return NewModel()

	case NoteKind:
		return NewNote()

	case OauthKind:
		return NewOauth()

	case ObjectKind:
		return NewObject()

	case OntologyKind:
		return NewOntology()

	case PersonKind:
		return NewPerson()

	case ProfileKind:
		return NewProfile()

	case QuantityKind:
		return NewQuantity()

	case RecurrenceKind:
		return NewRecurrence()

	case RelationKind:
		return NewRelation()

	case RoutineKind:
		return NewRoutine()

	case ScheduleKind:
		return NewSchedule()

	case SessionKind:
		return NewSession()

	case TagKind:
		return NewTag()

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
