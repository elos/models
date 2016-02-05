package calendar

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

func FixturesAppearingOn(db data.DB, c *models.Calendar, t time.Time) ([]*models.Fixture, error) {
	daily, _ := c.DailySchedule(db)
	weekly, _ := c.WeeklySchedule(db)
	monthly, _ := c.MonthlySchedule(db)
	yearly, _ := c.YearlySchedule(db)
	regular, _ := c.Schedule(db)

	dailys, _ := schedule.FixturesAppearingOn(db, daily, t)
	weeklys, _ := schedule.FixturesAppearingOn(db, weekly, t)
	dailys = append(dailys, weeklys...)
	monthlys, _ := schedule.FixturesAppearingOn(db, monthly, t)
	dailys = append(dailys, monthlys...)
	yearlys, _ := schedule.FixturesAppearingOn(db, yearly, t)
	dailys = append(dailys, yearlys...)
	regulars, _ := schedule.FixturesAppearingOn(db, regular, t)
	dailys = append(dailys, regulars...)

	return dailys, nil
}
