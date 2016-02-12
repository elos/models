package event

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/location"
	"github.com/elos/models/tag"
)

func LocationUpdate(db data.DB, u *models.User, altitude, latitude, longitude float64, tags ...*models.Tag) (*models.Event, *models.Location, error) {
	loc := location.NewCoords(altitude, latitude, longitude)
	loc.SetID(db.NewID())
	loc.SetOwner(u)
	now := loc.CreatedAt

	e := models.NewEvent()
	e.CreatedAt = now
	e.SetID(db.NewID())
	e.SetOwner(u)
	e.Name = "Location Update"
	e.SetLocation(loc)
	e.Time = now
	e.UpdatedAt = now

	locationTag, err := tag.ForName(db, u, tag.Location)
	if err != nil {
		return nil, nil, err
	}
	updateTag, err := tag.ForName(db, u, tag.Update)
	if err != nil {
		return nil, nil, err
	}

	e.IncludeTag(locationTag)
	e.IncludeTag(updateTag)

	for _, t := range tags {
		e.IncludeTag(t)
	}

	if err := db.Save(loc); err != nil {
		return nil, nil, err
	}

	if err := db.Save(e); err != nil {
		return nil, nil, err
	}

	return e, loc, nil
}
