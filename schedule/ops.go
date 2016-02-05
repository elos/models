package schedule

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

type spanClass int

const (
	Hour spanClass = iota + 1
	Day
	Week
	Month
	Year
	None
)

func Span(s *models.Schedule) time.Duration {
	return s.EndTime.Sub(s.StartTime)
}

func span(start time.Time, end time.Time) time.Duration {
	return end.Sub(start)
}

func class(start time.Time, end time.Time) spanClass {
	d := span(start, end)

	if d == 0 {
		return None
	}

	if d <= 1*time.Hour {
		return Hour
	} else if d <= 24*time.Hour {
		return Day
	} else if d <= 7*24*time.Hour {
		return Week
	} else if d <= 31*7*24*time.Hour {
		return Month
	} else if d <= 52*7*24*time.Hour {
		return Year
	} else {
		return None
	}
}

func Class(s *models.Schedule) spanClass {
	d := Span(s)

	// StartTime == EndTime, the schedule is infinite
	if d == 0 {
		return None
	}

	if d <= 1*time.Hour {
		return Hour
	} else if d <= 24*time.Hour {
		return Day
	} else if d <= 7*24*time.Hour {
		return Week
	} else if d <= 31*7*24*time.Hour {
		return Month
	} else if d <= 52*7*24*time.Hour {
		return Year
	} else {
		return None
	}
}

func FixturesAppearingOn(db data.DB, s *models.Schedule, t time.Time) {
}
