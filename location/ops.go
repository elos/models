package location

import (
	"time"

	"github.com/elos/models"
)

func NewCoords(alt, lat, lon float64) *models.Location {
	l := models.NewLocation()
	l.Altitude = alt
	l.Latitude = lat
	l.Longitude = lon
	l.CreatedAt = time.Now()
	l.UpdatedAt = l.CreatedAt
	return l
}
