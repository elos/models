package models

import (
	"time"

	"github.com/elos/data"
)

const (
	HabitTagName   = "HABIT"
	CheckinTagName = "CHECKIN"
	GoalTagName    = "GOAL"
)

// TagByName finds the tag indicated by the name parameter for the
// the given user, or creates it if it doesn't exist
func TagByName(db data.DB, name string, u *User) (*Tag, error) {
	q := db.Query(TagKind)
	q.Select(data.AttrMap{
		"owner_id": u.Id,
		"name":     name,
	})
	iter, err := q.Execute()
	if err != nil {
		return nil, err
	}
	t := NewTag()
	exists := iter.Next(t)
	if err := iter.Close(); err != nil {
		return nil, err
	}

	if !exists {
		t = NewTag()
		t.SetID(db.NewID())
		t.CreatedAt = time.Now()
		t.Name = name
		t.SetOwner(u)
		t.UpdatedAt = t.CreatedAt
		if err := db.Save(t); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func (t *Tag) Events(db data.DB) ([]*Event, error) {
	iter, err := db.Query(EventKind).Select(data.AttrMap{
		"owner_id": t.OwnerId,
	}).Execute()

	if err != nil {
		return nil, err
	}

	e := NewEvent()
	events := make([]*Event, 0)

	for iter.Next(e) {
		if contains(e.TagsIds, t.Id) {
			events = append(events, e)
			e = NewEvent()
		}
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return events, nil
}

func (t *Tag) Tasks(db data.DB) ([]*Task, error) {
	iter, err := db.Query(TaskKind).Select(data.AttrMap{
		"owner_id": t.OwnerId,
	}).Execute()

	if err != nil {
		return nil, err
	}

	task := NewTask()
	tasks := make([]*Task, 0)

	for iter.Next(task) {
		if contains(task.TagsIds, t.Id) {
			tasks = append(tasks, task)
			task = NewTask()
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
