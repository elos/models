package models

import (
	"fmt"
	"time"

	"github.com/elos/data"
)

func ValidWeekday(i int) bool {
	return i >= 0 && i <= 6
}

func WeekdayKey(t time.Time) string {
	return string(int(t.Weekday()))
}

func YeardayKey(t time.Time) string {
	return string(int(t.Month())*100 + t.Day())
}

func (c *Calendar) WeekdaySchedule(t time.Time, db data.DB) (*Schedule, error) {
	id, ok := c.WeekdaySchedules[WeekdayKey(t)]
	if !ok {
		return nil, ErrEmptyLink
	}

	schedule := NewSchedule()
	schedule.SetID(data.ID(id))
	return schedule, db.PopulateByID(schedule)
}

func (c *Calendar) WeekdayScheduleOrCreate(t time.Time, db data.DB) (*Schedule, error) {
	s, err := c.WeekdaySchedule(t, db)
	if err != nil {
		if err == ErrEmptyLink {
			schedule := NewSchedule()
			schedule.SetID(data.ID(db.NewID()))
			schedule.CreatedAt = time.Now()
			schedule.OwnerId = c.OwnerId
			schedule.Name = fmt.Sprintf("%s Schedule", t.Weekday())
			schedule.UpdatedAt = time.Now()

			err = db.Save(schedule)
			if err != nil {
				return nil, err
			}

			c.WeekdaySchedules[WeekdayKey(t)] = schedule.ID().String()

			err = db.Save(c)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	return s, nil
}

func (c *Calendar) YeardaySchedule(t time.Time, db data.DB) (*Schedule, error) {
	id, ok := c.YeardaySchedules[YeardayKey(t)]
	if !ok {
		return nil, ErrEmptyLink
	}

	schedule := NewSchedule()
	schedule.SetID(data.ID(id))
	return schedule, db.PopulateByID(schedule)
}

func (c *Calendar) YeardayScheduleOrCreate(t time.Time, db data.DB) (*Schedule, error) {
	s, err := c.YeardaySchedule(t, db)
	if err != nil {
		if err == ErrEmptyLink {
			schedule := NewSchedule()
			schedule.SetID(data.ID(db.NewID()))
			schedule.CreatedAt = time.Now()
			schedule.OwnerId = c.OwnerId
			schedule.Name = fmt.Sprintf("%d/%d Schedule", t.Year(), t.Day())
			schedule.UpdatedAt = time.Now()

			err = db.Save(schedule)
			if err != nil {
				return nil, err
			}

			c.YeardaySchedules[YeardayKey(t)] = schedule.ID().String()

			err = db.Save(c)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	return s, nil
}

func (c *Calendar) SchedulesForDate(t time.Time, db data.DB) []*Schedule {
	schedules := make([]*Schedule, 0)

	if base, err := c.BaseSchedule(db); err == nil {
		schedules = append(schedules, base)
	}

	if weekday, err := c.WeekdaySchedule(t, db); err == nil {
		schedules = append(schedules, weekday)
	}

	if yearday, err := c.YeardaySchedule(t, db); err == nil {
		schedules = append(schedules, yearday)
	}

	return schedules
}

func (c *Calendar) FixturesForDate(date time.Time, db data.DB) ([]*Fixture, error) {
	// TODO: perhaps this should return an error
	schedules := c.SchedulesForDate(date, db)

	fixtures, err := MergedFixtures(db, schedules...)
	if err != nil {
		return nil, err
	}

	return RelevantFixtures(date, fixtures), nil
}
