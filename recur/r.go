package recur

import (
	"time"

	"github.com/elos/models"
)

type Frequency string

const (
	Secondly Frequency = "SECONDLY"
	Minutely           = "MINUTELY"
	Hourly             = "HOURLY"
	Daily              = "DAILY"
	Weekly             = "WEEKLY"
	Monthly            = "MONTHLY"
	Yearly             = "YEARLY"
)

var Durations = map[Frequency]time.Duration{
	Secondly: time.Second,
	Minutely: time.Minute,
	Hourly:   time.Hour,
	Daily:    time.Hour * 24,
	Weekly:   time.Hour * 24 * 7,
	Monthly:  time.Hour * 24 * 7 * 31,
	Yearly:   time.Hour * 24 * 365,
}

func Expand(r *models.Recurrence, limit int) []time.Time {
	if r.Count != 0 && r.Count < limit {
		limit = r.Count
	}

	return make([]time.Time, 0)
}
