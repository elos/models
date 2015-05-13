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
