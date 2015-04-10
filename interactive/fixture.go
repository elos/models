package interactive

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

type Fixture struct {
	space *Space         `json:"-"`
	model models.Fixture `json:"-"`

	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`

	UserID     string `json:"user_id"`
	ScheduleID string `json:"schedule_id"`
}

func (this *Fixture) Save() {
	data.TransferAttrs(this, this.model)
	this.space.Save(this.model)
	this.space.Reload()
}

func NewFixture(s *Space) *Fixture {
	f, _ := s.Store.ModelFor(models.FixtureKind)
	f.SetID(s.NewID())
	return FixtureModel(s, f.(models.Fixture))
}

func FixtureModel(s *Space, m models.Fixture) *Fixture {
	f := &Fixture{
		space: s,
		model: m,
	}

	data.TransferAttrs(f.model, f)
	s.Register(f)
	return f
}

func Fixtures(s *Space, models []models.Fixture) []*Fixture {
	fixtures := make([]*Fixture, len(models))
	for i := range fixtures {
		fixtures[i] = FixtureModel(s, models[i])
	}

	return fixtures
}

func (this *Fixture) Reload() error {
	this.space.Store.PopulateByID(this.model)
	data.TransferAttrs(this.model, this)
	return nil
}

func (this *Fixture) Model() models.Fixture {
	return this.model
}
