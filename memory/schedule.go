package memory

import "github.com/elos/models"

type Schedule struct {
	space *Space          `json:"-"`
	model models.Schedule `json:"-"`

	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`

	UserID     string `json:"user_id"`
	FixtureIDs string `json:"fixture_ids"`
}

func (this *Schedule) Save() {
	transferAttrs(this, this.model)
	this.space.Save(this.model)
	this.space.Reload()
}

func NewSchedule(s *Space) *Schedule {
	f, _ := s.Access.ModelFor(models.ScheduleKind)
	f.SetID(s.NewID())
	return ScheduleModel(s, f.(models.Schedule))
}

func ScheduleModel(s *Space, m models.Schedule) *Schedule {
	f := &Schedule{
		space: s,
		model: m,
	}

	transferAttrs(f.model, f)
	s.Register(f)
	return f
}

func (this *Schedule) Reload() error {
	this.space.Access.PopulateByID(this.model)
	transferAttrs(this.model, this)
	return nil
}

func (this *Schedule) Model() models.Schedule {
	return this.model
}
