package habit

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/tag"
)

func CheckinFor(db data.DB, h *models.Habit, note string, date time.Time) (*models.Event, error) {
	now := time.Now()
	user, err := h.Owner(db)
	if err != nil {
		return nil, err
	}

	e := models.NewEvent()
	e.SetID(db.NewID())
	e.CreatedAt = now
	e.SetOwner(user)
	e.Name = h.Name + " Checkin"
	e.Time = date

	n := models.NewNote()
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
	ht, err := tag.ForName(db, user, tag.Habit)
	if err != nil {
		return nil, err
	}
	e.IncludeTag(ht)
	ct, err := tag.ForName(db, user, tag.Checkin)
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

func sameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

func DidCheckinOn(db data.DB, h *models.Habit, t time.Time) (bool, error) {
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

func Create(db data.DB, user *models.User, name string) (*models.Habit, error) {
	h := models.NewHabit()
	h.SetID(db.NewID())
	h.CreatedAt = time.Now()
	h.SetOwner(user)
	h.Name = name

	t := models.NewTag()
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
