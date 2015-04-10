package interactive

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/calendar"
)

type Calendar struct {
	space *Space          `json:"-"`
	model models.Calendar `json:"-"`

	MongoModel

	UserID string `json:"user_id"`

	BaseScheduleID   string            `json:"base_schedule_id"`
	WeekdaySchedules map[string]string `json:"weekday_schedules"`
	Schedules        map[string]string `json:"schedules"`

	CurrentFixtureID string `json:"current_fixture_id"`
}

func CalendarModel(s *Space, m models.Calendar) *Calendar {
	c := &Calendar{
		space: s,
		model: m,
	}
	data.TransferAttrs(c.model, c)
	s.Register(c)
	return c
}

func NewCalendar(s *Space) *Calendar {
	c := calendar.New(s.Store)
	return CalendarModel(s, c)
}

func (this *Calendar) Save() {
	data.TransferAttrs(this, this.model)
	this.space.Save(this.model)
	this.space.Reload()
}

func (this *Calendar) Delete() error {
	data.TransferAttrs(this, this.model)
	return this.space.Delete(this.model)
}

func (this *Calendar) Reload() error {
	this.space.Store.PopulateByID(this.model)
	data.TransferAttrs(this.model, this)
	return nil
}

func (c *Calendar) SetBaseSchedule(b *Schedule) {
	c.BaseScheduleID = b.ID
	c.Save()
}

func (c *Calendar) BaseSchedule() *Schedule {
	return c.space.FindSchedule(c.BaseScheduleID)
}

func (c *Calendar) SetWeekdaySchedule(s *Schedule, weekday int) {
	c.WeekdaySchedules[string(weekday)] = s.ID
}

func (c *Calendar) WeekdaySchedule(weekday int) *Schedule {
	return c.space.FindSchedule(c.WeekdaySchedules[string(weekday)])
}
