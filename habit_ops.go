package models

import (
	"time"

	"github.com/elos/data"
)

func CreateHabit(db data.DB, user *User, name string) (*Habit, error) {
	h := NewHabit()
	h.SetID(db.NewID())
	h.CreatedAt = time.Now()
	h.SetOwner(user)
	h.Name = name

	t := NewTag()
	t.SetID(db.NewID())
	t.CreatedAt = time.Now()
	t.SetOwner(user)
	t.Name = h.Name

	h.SetTag(t)

	t.UpdatedAt = time.Now()
	h.UpdatedAt = time.Now()

	if err := db.Save(t); err != nil {
		return nil, err
	}

	if err := db.Save(h); err != nil {
		return nil, err
	}

	return h, nil
}
