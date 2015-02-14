package models

import (
	"log"

	"github.com/elos/data"
)

const (
	DataVersion           = 1
	UserKind    data.Kind = "user"
	EventKind   data.Kind = "event"
	TaskKind    data.Kind = "task"
	ActionKind  data.Kind = "action"
	RoutineKind data.Kind = "routine"
)

const (
	UserEvents        data.LinkName = "events"
	UserTasks         data.LinkName = "tasks"
	UserActions       data.LinkName = "actions"
	UserCurrentAction data.LinkName = "current_action"
	UserRoutines      data.LinkName = "routines"

	EventUser data.LinkName = "user"

	TaskUser         data.LinkName = "user"
	TaskDependencies data.LinkName = "dependencies"

	ActionUser data.LinkName = "user"
	ActionTask data.LinkName = "task"

	RoutineUser           data.LinkName = "user"
	RoutineTasks          data.LinkName = "tasks"
	RoutineCompletedTasks data.LinkName = "completed_tasks"
	RoutineActions        data.LinkName = "actions"
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
}

func SetupSchema() data.Schema {
	sch, err := data.NewSchema(&RMap, DataVersion)

	if err != nil {
		log.Fatal(err)
	}

	return sch
}

var Schema data.Schema = SetupSchema()
