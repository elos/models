package calendar_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/calendar"
	"github.com/elos/models/fixture"
	"github.com/elos/models/persistence"
	"github.com/elos/models/schedule"
	"github.com/elos/testing/expect"
	"github.com/elos/testing/modeltest"
)

func TestMongo(t *testing.T) {
	s := persistence.Store(persistence.MongoMemoryDB())
	c := calendar.New(s)
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
	store := persistence.ModelsStore(s)

	testActionable(store, c, t)
	testBaseSchedule(store, c, t)
	testWeekdaySchedules(store, c, t)
	testYeardaySchedules(store, c, t)
	testCurrentFixture(store, c, t)
	testNextFixture(store, c, t)
	testIntegratedSchedule(store, c, t)

	modeltest.Userable(store, c, t)
	modeltest.UserOwnedAccessRights(store, c, t)
	modeltest.AnonReadAccess(s, c, t)
}

func testActionable(access models.Store, c models.Calendar, t *testing.T) {
}

func testBaseSchedule(access models.Store, c models.Calendar, t *testing.T) {
	_, err := c.BaseSchedule(access)
	expect.EmptyRelationship("BaseSchedule", err, t)

	s, err := schedule.Create(access)
	if err != nil {
		t.Errorf("Error creating new schedules: %s", err)
	}

	if err = c.SetBaseSchedule(s); err != nil {
		t.Errorf("Error while setting base schedule: %s", err)
	}

	if err = access.Save(c); err != nil {
		t.Fatalf("Error while saving calendar: %s", err)
	}

	c, err = calendar.Find(access, c.ID())
	if err != nil {
		t.Errorf("Error while finding calendar: %s", err)
	}

	sRetrieved, err := c.BaseSchedule(access)
	if err != nil {
		t.Errorf("Error while retrieving base schedule: %s", err)
	}

	if sRetrieved.ID().String() != s.ID().String() {
		t.Errorf("Retrieved schedules doesn't match set schedule")
	}
}

func testWeekdaySchedules(access models.Store, c models.Calendar, t *testing.T) {
	if err := access.Save(c); err != nil {
		t.Errorf("Error while saving calendar: %s", err)
	}

	for i := 0; i < 7; i++ {
		w := time.Weekday(i)

		_, err := c.WeekdaySchedule(access, w)
		expect.EmptyRelationship(fmt.Sprintf("Weekday: %s", w.String()), err, t)

		s, err := schedule.Create(access)
		if err != nil {
			t.Errorf("Error while creating schedule: %s", err)
		}

		if err := c.SetWeekdaySchedule(s, w); err != nil {
			t.Errorf("Error while setting weekday schedule for %s. err: %s", w, err)
		}

		if err = access.Save(c); err != nil {
			t.Error("Error while saving calendar")
		}

		c, err = calendar.Find(access, c.ID())

		sRetrieved, err := c.WeekdaySchedule(access, w)
		if err != nil {
			t.Errorf("Error while retrieving weekday schedules %s, err: %s", w, err)
		}

		if sRetrieved.ID().String() != s.ID().String() {
			t.Errorf("Retrieved schedules doesn't match set schedule")
		}
	}
}

var monthDays = map[time.Month]int{
	time.January:   31,
	time.February:  29,
	time.March:     31,
	time.April:     30,
	time.May:       31,
	time.June:      30,
	time.July:      31,
	time.August:    31,
	time.September: 30,
	time.October:   31,
	time.November:  30,
	time.December:  31,
}

func testYeardaySchedules(access models.Store, c models.Calendar, t *testing.T) {
	err := access.Save(c)
	expect.NoError("saving calendar", err, t)

	for month, days := range monthDays {
		for i := 1; i <= days; i++ {
			yearday := time.Date(2015, month, i, 0, 0, 0, 0, time.UTC)

			s, err := schedule.Create(access)
			expect.NoError("creating schedule", err, t)

			err = c.SetYeardaySchedule(s, yearday)
			expect.NoError("setting yearday schedules", err, t)

			err = access.Save(c)
			expect.NoError("saving calendar", err, t)

			c, err := calendar.Find(access, c.ID())
			expect.NoError("finding calendar", err, t)

			sRetrieved, err := c.YeardaySchedule(access, yearday)
			expect.NoError("retrieving yearday schedule", err, t)

			if sRetrieved.ID().String() != s.ID().String() {
				t.Errorf("Retrieved yearday schedule doesn't match set yearday schedule")
			}
		}
	}
}

func testCurrentFixture(access models.Store, c models.Calendar, t *testing.T) {
	f, err := fixture.Create(access)
	if err != nil {
		t.Fatalf("Error creating fixture: %s", err)
	}

	_, err = c.CurrentFixture(access)
	expect.EmptyRelationship("CurrentFixture", err, t)

	err = c.SetCurrentFixture(f)
	expect.NoError("setting current fixture", err, t)

	err = access.Save(c)
	expect.NoError("saving calendar", err, t)

	c, err = calendar.Find(access, c.ID())
	expect.NoError("finding calendar", err, t)

	fRetrieved, err := c.CurrentFixture(access)
	expect.NoError("retrieving current fixture", err, t)

	if !data.EqualModels(f, fRetrieved) {
		t.Errorf("Retrieved curret fixture doesn't match set current fixture")
	}
}

// TODO: fix this to be actually test hard scenarios
func testNextFixture(access models.Store, c models.Calendar, t *testing.T) {
	sBase, err := schedule.Create(access)
	expect.NoError("creating schedule", err, t)
	sWeek, err := schedule.Create(access)
	expect.NoError("creating schedule", err, t)
	sYear, err := schedule.Create(access)
	expect.NoError("creating schedule", err, t)

	testTime := time.Now()

	err = c.SetBaseSchedule(sBase)
	expect.NoError("setting base schedule", err, t)

	err = c.SetWeekdaySchedule(sWeek, testTime.Weekday())
	expect.NoError("setting weekday schedule", err, t)

	err = c.SetYeardaySchedule(sYear, testTime)
	expect.NoError("setting yearday schedule", err, t)

	f, err := fixture.Create(access)
	expect.NoError("creating fixture", err, t)

	f.SetStartTime(testTime.Add(1 * time.Hour))
	f.SetEndTime(testTime.Add(2 * time.Hour))

	err = sBase.IncludeFixture(f)
	expect.NoError("including fixture", err, t)

	err = access.Save(sBase)
	expect.NoError("saving schedule", err, t)
	err = access.Save(f)
	expect.NoError("saving fixture", err, t)

	nextF, err := c.NextFixture(access)
	expect.NoError("retrieving next fixture", err, t)

	if !data.EqualModels(f, nextF) {
		t.Errorf("Expected next fixture to be fixture on base schedule")
	}
}

func testIntegratedSchedule(access models.Store, c models.Calendar, t *testing.T) {
}
