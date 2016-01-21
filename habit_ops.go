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

func sameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

func (h *Habit) Checkin(db data.DB, note string, date time.Time) (*Event, error) {
	now := time.Now()
	user, err := h.Owner(db)
	if err != nil {
		return nil, err
	}

	e := NewEvent()
	e.SetID(db.NewID())
	e.CreatedAt = now
	e.SetOwner(user)
	e.Name = h.Name + " Checkin"
	e.Time = date

	n := NewNote()
	n.SetID(db.NewID())
	n.CreatedAt = now
	n.SetOwner(user)
	n.Text = note
	n.UpdatedAt = now

	e.SetNote(n)

	t, err := h.Tag(db)
	if err != nil {
		return nil, err
	}

	e.IncludeTag(t)
	ht, err := TagByName(db, HabitTagName, user)
	if err != nil {
		return nil, err
	}
	e.IncludeTag(ht)
	ct, err := TagByName(db, CheckinTagName, user)
	if err != nil {
		return nil, err
	}
	e.IncludeTag(ct)

	if err := db.Save(n); err != nil {
		return nil, err
	}

	if err := db.Save(e); err != nil {
		return nil, err
	}

	h.IncludeCheckin(e)
	if err := db.Save(h); err != nil {
		return nil, err
	}

	return e, nil
}

func (h *Habit) CheckedInOn(db data.DB, t time.Time) (bool, error) {
	checkins, err := h.Checkins(db)
	if err != nil {
		return false, err
	}

	for _, c := range checkins {
		if sameDay(t, c.CreatedAt) {
			return true, nil
		}
	}

	return false, nil
}
