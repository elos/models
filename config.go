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
)

const (
	UserEvents       data.LinkName = "events"
	UserTasks        data.LinkName = "tasks"
	UserCurrentTask  data.LinkName = "current_task"
	EventUser        data.LinkName = "user"
	TaskUser         data.LinkName = "user"
	TaskDependencies data.LinkName = "dependencies"
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
		UserCurrentTask: data.Link{
			Name:  UserCurrentTask,
			Kind:  data.OneLink,
			Other: TaskKind,
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
}

func SetupSchema() data.Schema {
	sch, err := data.NewSchema(&RMap, DataVersion)

	if err != nil {
		log.Fatal(err)
	}

	return sch
}

var Schema data.Schema = SetupSchema()
