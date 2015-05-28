package models

import (
	"time"

	"github.com/elos/data"
)

func WeekdayKey(t time.Time) int {
	return int(t.Weekday())
}

func YeardayKey(t time.Time) int {
	return int(t.Month())*100 + t.Day()
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

func (c *Calendar) YeardaySchedule(t time.Time, db data.DB) (*Schedule, error) {
	id, ok := c.YeardaySchedules[YeardayKey(t)]
	if !ok {
		return nil, ErrEmptyLink
	}

	schedule := NewSchedule()
	schedule.SetID(data.ID(id))
	return schedule, db.PopulateByID(schedule)
}

func (c *Calendar) SchedulesForDate(t time.Time, db data.DB) ([]*Schedule, error) {
	schedules := make([]*Schedule, 3)

	if base, err := c.BaseSchedule(db); err == nil {
		schedules[0] = base
	} else {
		return schedules, err
	}

	if weekday, err := c.WeekdaySchedule(t, db); err == nil {
		schedules[1] = weekday
	} else {
		return schedules, err
	}

	if yearday, err := c.YeardaySchedule(t, db); err == nil {
		schedules[2] = yearday
	} else {
		return schedules, err
	}

	return schedules, nil
}
