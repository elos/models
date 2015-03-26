package calendar_test

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/calendar"
	"github.com/elos/models/persistence"
)

func TestMongo(t *testing.T) {
	s := persistence.Store(persistence.MongoMemoryDB())
	c, err := calendar.New(s)
	if err != nil {
		t.Errorf("Error from calendar.New, expected no error but got %s", err)
	}

	testCalendar(s, c, t)

	if c.Version() != 1 {
		t.Errorf("Expected mongoCalendar version to be 1, got %d", c.Version())
	}

	if c.Kind() != models.CalendarKind {
		t.Errorf("Expected mongoCalendar kind to equal models.CalendarKind, got %s", c.Kind())
	}

	if c.Schema() != models.Schema {
		t.Errorf("Expected mongoCalendar schema to be models.Schema")
	}
}

func testCalendar(s data.Store, c models.Calendar, t *testing.T) {
	access := data.NewAnonAccess(s)

	testActionable(access, c, t)
	testBaseSchedule(access, c, t)
	testWeekdaySchedules(access, c, t)
	testYeardaySchedules(access, c, t)
	testCurrentFixture(access, c, t)
	testNextFixture(access, c, t)

	testAccessProtection(s, c, t)
	testAnonReadAccess(s, c, t)
}

func testActionable(access data.Access, c models.Calendar, t *testing.T) {
}

func testBaseSchedule(access data.Access, c models.Calendar, t *testing.T) {
}

func testWeekdaySchedules(access data.Access, c models.Calendar, t *testing.T) {
}

func testYeardaySchedules(access data.Access, c models.Calendar, t *testing.T) {
}

func testCurrentFixture(access data.Access, c models.Calendar, t *testing.T) {
}

func testNextFixture(access data.Access, c models.Calendar, t *testing.T) {
}

func testAccessProtection(s data.Store, c models.Calendar, t *testing.T) {
}

func testAnonReadAccess(s data.Store, c models.Calendar, t *testing.T) {
}
