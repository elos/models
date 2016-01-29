package tag

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

// EventsFor retrieves all the events which are tagged by this tag
func EventsFor(db data.DB, t *models.Tag) ([]*models.Event, error) {
	iter, err := db.Query(models.EventKind).Select(data.AttrMap{
		"owner_id": t.OwnerId,
	}).Execute()

	if err != nil {
		return nil, err
	}

	e := models.NewEvent()
	events := make([]*models.Event, 0)

	for iter.Next(e) {
		if contains(e.TagsIds, t.Id) {
			events = append(events, e)
			e = models.NewEvent()
		}
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return events, nil
}

// TasksFor retrieves all tasks which are tagged by this tag
func TasksFor(db data.DB, t *models.Tag) ([]*models.Task, error) {
	iter, err := db.Query(models.TaskKind).Select(data.AttrMap{
		"owner_id": t.OwnerId,
	}).Execute()

	if err != nil {
		return nil, err
	}

	task := models.NewTask()
	tasks := make([]*models.Task, 0)

	for iter.Next(task) {
		if contains(task.TagsIds, t.Id) {
			tasks = append(tasks, task)
			task = models.NewTask()
		}
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func contains(ss []string, s string) bool {
	for _, t := range ss {
		if t == s {
			return true
		}
	}

	return false
}
