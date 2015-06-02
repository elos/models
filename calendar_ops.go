package models

import (
	"log"
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

func (c *Calendar) SchedulesForDate(t time.Time, db data.DB) []*Schedule {
	schedules := make([]*Schedule, 0)

	if base, err := c.BaseScheduleOrCreate(db); err == nil {
		log.Printf("Base, base: %+v, err: %s", base, err)
		schedules = append(schedules, base)
	}

	if weekday, err := c.WeekdaySchedule(t, db); err == nil {
		schedules = append(schedules, weekday)
	}

	if yearday, err := c.YeardaySchedule(t, db); err == nil {
		schedules = append(schedules, yearday)
	}

	log.Print(schedules)

	return schedules
}
