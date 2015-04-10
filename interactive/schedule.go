package interactive

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/models"
)

type Schedule struct {
	space *Space          `json:"-"`
	model models.Schedule `json:"-"`

	MongoModel

	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	UserID    string `json:"user_id"`

	FixtureIDs []string `json:"fixture_ids"`
}

func (this *Schedule) Save() {
	data.TransferAttrs(this, this.model)
	this.space.Save(this.model)
	this.space.Reload()
}

func NewSchedule(s *Space) *Schedule {
	f, _ := s.Store.ModelFor(models.ScheduleKind)
	f.SetID(s.NewID())
	return ScheduleModel(s, f.(models.Schedule))
}

func ScheduleModel(s *Space, m models.Schedule) *Schedule {
	f := &Schedule{
		space: s,
		model: m,
	}

	data.TransferAttrs(f.model, f)
	s.Register(f)
	return f
}

func (this *Schedule) Reload() error {
	this.space.Store.PopulateByID(this.model)
	data.TransferAttrs(this.model, this)
	return nil
}

func (this *Schedule) Model() models.Schedule {
	return this.model
}

func (s *Schedule) Fixtures() []*Fixture {
	fixtures, err := s.model.Fixtures(s.space.Store.(models.Store))

	if err != nil {
		log.Print(err)
		return nil
	}

	return Fixtures(s.space, fixtures)
}

func (s *Schedule) IncludeFixture(f *Fixture) {
	f.Save()
	s.FixtureIDs = append(s.FixtureIDs, f.ID)
	s.Save()
}

func (s *Schedule) ExcludeFixture(f *Fixture) {
	fids := make([]string, 0)
	for _, id := range s.FixtureIDs {
		if id != f.ID {
			fids = append(fids, id)
		}
	}

	s.FixtureIDs = fids
	s.Save()
}
