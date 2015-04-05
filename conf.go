package models

import (
	"log"

	"github.com/elos/data"
)

const (
	DataVersion = 1

	ActionKind   data.Kind = "action"
	CalendarKind data.Kind = "calendar"
	ClassKind    data.Kind = "class"
	EventKind    data.Kind = "event"
	FixtureKind  data.Kind = "fixture"
	ObjectKind   data.Kind = "object"
	OntologyKind data.Kind = "ontology"
	RoutineKind  data.Kind = "routine"
	ScheduleKind data.Kind = "schedule"
	SetKind      data.Kind = "set"
	TaskKind     data.Kind = "task"
	UserKind     data.Kind = "user"
)

const (
	UserActions       data.LinkName = "actions"
	UserEvents        data.LinkName = "events"
	UserTasks         data.LinkName = "tasks"
	UserRoutines      data.LinkName = "routines"
	UserOntology      data.LinkName = "ontology"
	UserCalendar      data.LinkName = "calendar"
	UserCurrentAction data.LinkName = "current_action"

	EventUser data.LinkName = "user"

	TaskUser         data.LinkName = "user"
	TaskDependencies data.LinkName = "dependencies"

	ActionUser data.LinkName = "user"
	ActionTask data.LinkName = "task"

	RoutineUser           data.LinkName = "user"
	RoutineTasks          data.LinkName = "tasks"
	RoutineCompletedTasks data.LinkName = "completed_tasks"
	RoutineActions        data.LinkName = "actions"
	RoutineCurrentAction  data.LinkName = "current_action"

	SetUser   data.LinkName = "user"
	SetModels data.LinkName = "models"

	CalendarUser             data.LinkName = "user"
	CalendarBaseSchedule     data.LinkName = "base_schedule"
	CalendarWeekdaySchedules data.LinkName = "weekday_schedules"
	CalendarSchedules        data.LinkName = "schedules"
	CalendarCurrentFixture   data.LinkName = "current_fixture"

	ScheduleUser     data.LinkName = "user"
	ScheduleFixtures data.LinkName = "fixtures"

	FixtureUser     data.LinkName = "user"
	FixtureSchedule data.LinkName = "schedule"
	FixtureActions  data.LinkName = "actions"
	FixtureEvents   data.LinkName = "events"

	OntologyUser    data.LinkName = "user"
	OntologyClasses data.LinkName = "classes"
	OntologyObjects data.LinkName = "objects"

	ClassUser          data.LinkName = "user"
	ClassOntology      data.LinkName = "ontology"
	ClassObjects       data.LinkName = "objects"
	ClassTraits        data.LinkName = "traits"
	ClassRelationships data.LinkName = "relationships"

	ObjectUser     data.LinkName = "user"
	ObjectClass    data.LinkName = "class"
	ObjectOntology data.LinkName = "ontology"
)

var RMap data.RelationshipMap = data.RelationshipMap{
	UserKind: {
		UserActions: data.Link{
			Name:    UserActions,
			Kind:    data.MulLink,
			Other:   ActionKind,
			Inverse: ActionUser,
		},
		UserEvents: data.Link{
			Name:    UserEvents,
			Kind:    data.MulLink,
			Other:   EventKind,
			Inverse: EventUser,
		},
		UserTasks: data.Link{
			Name:    UserTasks,
			Kind:    data.MulLink,
			Other:   TaskKind,
			Inverse: TaskUser,
		},
		UserRoutines: data.Link{
			Name:    UserRoutines,
			Kind:    data.MulLink,
			Other:   RoutineKind,
			Inverse: RoutineUser,
		},
		UserOntology: data.Link{
			Name:    UserOntology,
			Kind:    data.OneLink,
			Other:   OntologyKind,
			Inverse: OntologyUser,
		},
		UserCalendar: data.Link{
			Name:    UserCalendar,
			Kind:    data.OneLink,
			Other:   CalendarKind,
			Inverse: CalendarUser,
		},
		UserCurrentAction: data.Link{
			Name:  UserCurrentAction,
			Kind:  data.OneLink,
			Other: ActionKind,
		},
	},

	EventKind: {
		EventUser: data.Link{
			Name:    EventUser,
			Kind:    data.OneLink,
			Other:   UserKind,
			Inverse: UserEvents,
		},
	},

	TaskKind: {
		TaskUser: data.Link{
			Name:    TaskUser,
			Kind:    data.OneLink,
			Other:   UserKind,
			Inverse: UserTasks,
		},
		TaskDependencies: data.Link{
			Name:  TaskDependencies,
			Kind:  data.MulLink,
			Other: TaskKind,
		},
	},

	ActionKind: {
		ActionUser: data.Link{
			Name:    ActionUser,
			Kind:    data.OneLink,
			Other:   UserKind,
			Inverse: UserActions,
		},
		ActionTask: data.Link{
			Name:  ActionTask,
			Kind:  data.OneLink,
			Other: TaskKind,
		},
	},

	RoutineKind: {
		RoutineUser: data.Link{
			Name:    RoutineUser,
			Kind:    data.OneLink,
			Other:   UserKind,
			Inverse: UserActions,
		},
		RoutineTasks: data.Link{
			Name:  RoutineTasks,
			Kind:  data.MulLink,
			Other: TaskKind,
		},
		RoutineCompletedTasks: data.Link{
			Name:  RoutineCompletedTasks,
			Kind:  data.MulLink,
			Other: TaskKind,
		},
		RoutineActions: data.Link{
			Name:  RoutineActions,
			Kind:  data.MulLink,
			Other: ActionKind,
		},
	},

	SetKind: {
		SetUser: data.Link{
			Name:  SetUser,
			Kind:  data.OneLink,
			Other: UserKind,
		},
		SetModels: data.Link{
			Name: SetModels,
			Kind: data.MulLink,
		},
	},

	ScheduleKind: {
		ScheduleUser: data.Link{
			Name:  ScheduleUser,
			Kind:  data.OneLink,
			Other: UserKind,
		},
		ScheduleFixtures: data.Link{
			Name:    ScheduleFixtures,
			Kind:    data.MulLink,
			Other:   FixtureKind,
			Inverse: FixtureSchedule,
		},
	},

	FixtureKind: {
		FixtureUser: data.Link{
			Name:  FixtureUser,
			Kind:  data.OneLink,
			Other: UserKind,
		},

		FixtureSchedule: data.Link{
			Name:    FixtureSchedule,
			Kind:    data.OneLink,
			Other:   ScheduleKind,
			Inverse: ScheduleFixtures,
		},

		FixtureActions: data.Link{
			Name:  FixtureActions,
			Kind:  data.MulLink,
			Other: ActionKind,
		},

		FixtureEvents: data.Link{
			Name:  FixtureEvents,
			Kind:  data.MulLink,
			Other: EventKind,
		},
	},

	CalendarKind: {
		CalendarUser: data.Link{
			Name:    CalendarUser,
			Kind:    data.OneLink,
			Other:   UserKind,
			Inverse: UserCalendar,
		},
		CalendarBaseSchedule: data.Link{
			Name:  CalendarBaseSchedule,
			Kind:  data.OneLink,
			Other: ScheduleKind,
		},
		CalendarWeekdaySchedules: data.Link{
			Name:  CalendarWeekdaySchedules,
			Kind:  data.MulLink,
			Other: ScheduleKind,
		},
		CalendarSchedules: data.Link{
			Name:  CalendarSchedules,
			Kind:  data.MulLink,
			Other: ScheduleKind,
		},
		CalendarCurrentFixture: data.Link{
			Name:  CalendarCurrentFixture,
			Kind:  data.OneLink,
			Other: FixtureKind,
		},
	},

	OntologyKind: {
		OntologyUser: data.Link{
			Name:    OntologyUser,
			Kind:    data.OneLink,
			Other:   UserKind,
			Inverse: UserOntology,
		},
		OntologyClasses: data.Link{
			Name:    OntologyClasses,
			Kind:    data.MulLink,
			Other:   ClassKind,
			Inverse: ClassOntology,
		},
		OntologyObjects: data.Link{
			Name:    OntologyObjects,
			Kind:    data.MulLink,
			Other:   ObjectKind,
			Inverse: ObjectOntology,
		},
	},

	ClassKind: {
		ClassUser: data.Link{
			Name:  ClassUser,
			Kind:  data.OneLink,
			Other: UserKind,
		},
		ClassOntology: data.Link{
			Name:    ClassOntology,
			Kind:    data.OneLink,
			Other:   OntologyKind,
			Inverse: OntologyClasses,
		},
		ClassObjects: data.Link{
			Name:    ClassObjects,
			Kind:    data.MulLink,
			Other:   ObjectKind,
			Inverse: ObjectClass,
		},
	},

	ObjectKind: {
		ObjectUser: data.Link{
			Name:  ObjectUser,
			Kind:  data.OneLink,
			Other: UserKind,
		},
		ObjectClass: data.Link{
			Name:    ObjectClass,
			Kind:    data.OneLink,
			Other:   ClassKind,
			Inverse: ClassObjects,
		},
		ObjectOntology: data.Link{
			Name:    ObjectOntology,
			Kind:    data.OneLink,
			Other:   OntologyKind,
			Inverse: OntologyObjects,
		},
	},
}

func SetupSchema() data.Schema {
	sch, err := data.NewSchema(&RMap, DataVersion)

	if err != nil {
		log.Fatal(err)
	}

	return sch
}

var Schema data.Schema = SetupSchema()
