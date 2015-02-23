package models

import (
	"log"

	"github.com/elos/data"
)

const (
	DataVersion            = 1
	UserKind     data.Kind = "user"
	EventKind    data.Kind = "event"
	TaskKind     data.Kind = "task"
	ActionKind   data.Kind = "action"
	RoutineKind  data.Kind = "routine"
	SetKind      data.Kind = "set"
	ScheduleKind data.Kind = "schedule"
	FixtureKind  data.Kind = "fixture"
	CalendarKind data.Kind = "calendar"

	OntologyKind data.Kind = "ontology"
	ClassKind    data.Kind = "class"
	ObjectKind   data.Kind = "object"
)

const (
	UserEvents        data.LinkName = "events"
	UserTasks         data.LinkName = "tasks"
	UserActions       data.LinkName = "actions"
	UserCurrentAction data.LinkName = "current_action"
	UserRoutines      data.LinkName = "routines"
	UserCalendar      data.LinkName = "calendar"
	UserOntology      data.LinkName = "ontology"

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

	CalendarUser      data.LinkName = "user"
	CalendarBase      data.LinkName = "base"
	CalendarMon       data.LinkName = "monday"
	CalendarTue       data.LinkName = "tuesday"
	CalendarWed       data.LinkName = "wednesday"
	CalendarThu       data.LinkName = "thursday"
	CalendarFri       data.LinkName = "friday"
	CalendarSat       data.LinkName = "saturday"
	CalendarSun       data.LinkName = "sunday"
	CalendarSchedules data.LinkName = "schedules"

	ScheduleUser     data.LinkName = "user"
	ScheduleFixtures data.LinkName = "fixtures"

	FixtureUser     data.LinkName = "user"
	FixtureSchedule data.LinkName = "schedule"

	// Experimental
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
		UserActions: data.Link{
			Name:    UserActions,
			Kind:    data.MulLink,
			Other:   ActionKind,
			Inverse: ActionUser,
		},
		UserCurrentAction: data.Link{
			Name:  UserCurrentAction,
			Kind:  data.OneLink,
			Other: ActionKind,
		},
		UserRoutines: data.Link{
			Name:    UserRoutines,
			Kind:    data.MulLink,
			Other:   RoutineKind,
			Inverse: RoutineUser,
		},
		UserCalendar: data.Link{
			Name:    UserCalendar,
			Kind:    data.OneLink,
			Other:   CalendarKind,
			Inverse: CalendarUser,
		},
		UserOntology: data.Link{
			Name:    UserOntology,
			Kind:    data.OneLink,
			Other:   OntologyKind,
			Inverse: OntologyUser,
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
	},

	CalendarKind: {
		CalendarUser: data.Link{
			Name:    CalendarUser,
			Kind:    data.OneLink,
			Other:   UserKind,
			Inverse: UserCalendar,
		},
		CalendarBase: data.Link{
			Name:  CalendarBase,
			Kind:  data.OneLink,
			Other: ScheduleKind,
		},
		CalendarMon: data.Link{
			Name:  CalendarMon,
			Kind:  data.OneLink,
			Other: ScheduleKind,
		},
		CalendarTue: data.Link{
			Name:  CalendarTue,
			Kind:  data.OneLink,
			Other: ScheduleKind,
		},
		CalendarWed: data.Link{
			Name:  CalendarWed,
			Kind:  data.OneLink,
			Other: ScheduleKind,
		},
		CalendarThu: data.Link{
			Name:  CalendarThu,
			Kind:  data.OneLink,
			Other: ScheduleKind,
		},
		CalendarFri: data.Link{
			Name:  CalendarFri,
			Kind:  data.OneLink,
			Other: ScheduleKind,
		},
		CalendarSat: data.Link{
			Name:  CalendarSat,
			Kind:  data.OneLink,
			Other: ScheduleKind,
		},
		CalendarSun: data.Link{
			Name:  CalendarSun,
			Kind:  data.OneLink,
			Other: ScheduleKind,
		},
		CalendarSchedules: data.Link{
			Name:  CalendarSchedules,
			Kind:  data.MulLink,
			Other: ScheduleKind,
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
