package tag

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

type Name string

func (n Name) String() string {
	return string(n)
}

const (
	Habit    Name = "HABIT"
	Checkin       = "CHECKIN"
	Goal          = "GOAL"
	Location      = "LOCATION"
	Update        = "UPDATE"
	Mobile        = "MOBILE"
)

// ForName finds the tag indicated by the name parameter for the
// given user, or creates it if it doesn't exist
func ForName(db data.DB, u *models.User, name Name) (*models.Tag, error) {
	iter, err := db.Query(models.TagKind).Select(data.AttrMap{
		"owner_id": u.Id,
		"name":     name.String(),
	}).Execute()
	if err != nil {
		return nil, err
	}

	t := models.NewTag()
	exists := iter.Next(t)
	if err := iter.Close(); err != nil {
		return nil, err
	}

	if !exists {
		t = models.NewTag()
		t.SetID(db.NewID())
		t.CreatedAt = time.Now()
		t.Name = name.String()
		t.SetOwner(u)
		t.UpdatedAt = t.CreatedAt
		if err := db.Save(t); err != nil {
			return nil, err
		}
	}

	return t, nil
}
