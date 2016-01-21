package models

import (
	"time"

	"github.com/elos/data"
)

func CreateUser(db data.DB, username, password string) (*User, *Credential, error) {
	u := NewUser()
	u.SetID(db.NewID())
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	c := NewCredential()
	c.SetID(db.NewID())
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.Public = username
	c.Private = password
	c.Spec = "password"
	c.SetOwner(u)

	if err := db.Save(u); err != nil {
		return nil, nil, err
	}

	if err := db.Save(c); err != nil {
		return nil, nil, err
	}

	return u, c, nil
}

func (u *User) Tasks(db data.DB, completedOnly bool) ([]*Task, error) {
	taskQuery := db.Query(TaskKind)

	// only retrieve _incomplete_ tasks
	taskQuery.Select(data.AttrMap{
		"owner_id": u.Id,
		"complete": false,
	})

	iter, err := taskQuery.Execute()
	if err != nil {
		return nil, err
	}

	t := NewTask()
	tasks := make([]*Task, 0)
	for iter.Next(t) {
		tasks = append(tasks, t)
		t = NewTask()
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return tasks, nil
}
