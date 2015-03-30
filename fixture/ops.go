package fixture

import (
	"time"

	"github.com/elos/models"
	"github.com/elos/models/calendar"
)

func Expired(f models.Fixture) bool {
	return f.Expires().Before(time.Now())
}

// only checks to month/day/year validity
func ShouldOmitOnDate(f models.Fixture, t time.Time) bool {
	exceptions := f.DateExceptions()
	for i := range exceptions {
		exception := exceptions[i]
		if calendar.Yearday(t) == calendar.Yearday(exception) && t.Year() == exception.Year() {
			return true
		}
	}

	return false
}

func Order(f1 models.Fixture, f2 models.Fixture) (first models.Fixture, second models.Fixture) {
	if f1.StartTime().Before(f2.StartTime()) {
		first = f1
		second = f2
	} else { // f2 starts before f1
		first = f2
		second = f1
	}
	return
}

func Conflicts(f1 models.Fixture, f2 models.Fixture) bool {
	first, second := Order(f1, f2)
	return second.StartTime().Before(first.EndTime())
}

func Before(f1 models.Fixture, f2 models.Fixture) bool {
	first, _ := Order(f1, f2)
	return first == f1
}
