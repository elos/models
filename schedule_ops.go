package models

import (
	"time"

	"github.com/elos/data"
)

func (s *Schedule) FirstFixture(db data.DB) (*Fixture, error) {
	return s.EarliestSince(*new(time.Time), db)
}

func (s *Schedule) EarliestSince(start time.Time, db data.DB) (*Fixture, error) {
	fs, _ := s.Fixtures(db)
	return fs[0], nil
	//not implemented
}

func (s *Schedule) OrderedFixtures(db data.DB) ([]*Fixture, error) {
	return s.Fixtures(db)
}

func MergedFixtures(db data.DB, schedules ...*Schedule) ([]*Fixture, error) {
	fixtures := make([]*Fixture, 0)

	for _, s := range schedules {
		for _, id := range s.FixturesIds {
			fixture := NewFixture()
			fixture.SetID(data.ID(id))
			if err := db.PopulateByID(fixture); err != nil {
				return fixtures, err
			}

			fixtures = append(fixtures, fixture)
		}
	}

	return fixtures, nil
}

func DayEquivalent(t1 time.Time, t2 time.Time) bool {
	return (t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day())
}

func RelevantScheduleFixtures(t time.Time, fixtures []*Fixture) []*Fixture {
	filtered := make([]*Fixture, 0)
	now := time.Now()

Filtering:
	for _, fixture := range fixtures {
		if now.After(fixture.ExpiresAt) {
			continue Filtering
		}

		for _, exception := range fixture.Exceptions {
			if DayEquivalent(now, exception) {
				continue Filtering
			}
		}

		filtered = append(filtered, fixture)
	}

	return filtered
}
