package system

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/task"
)

type SystemEnvironment struct {
	As func(u *models.User) UserActions
}

func DB(db data.DB) SystemEnvironment {
	return SystemEnvironment{
		As: func(u *models.User) UserActions {
			return UserActions{
				CompleteTask: func(t *models.Task) error {
					task.StopAndComplete(t)
					// and this would include the evented logic
					return nil
				},
			}
		},
	}
}

type UserActions struct {
	CompleteTask func(t *models.Task) error
}
