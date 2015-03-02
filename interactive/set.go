package interactive

import "github.com/elos/models"

type Set struct {
	space *Space     `json:"-"`
	model models.Set `json:"-"`

	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	StartTime string   `json:"start_time"`
	EndTime   string   `json:"end_time"`
	UserID    string   `json:"user_id"`
	ModelKind string   `json:"model_kind"`
	ModelIDs  []string `json:"model_ids"`
}

func (this *Set) Save() {
	transferAttrs(this, this.model)
	this.space.Save(this.model)
	this.space.Reload()
}

func NewSet(s *Space) *Set {
	f, _ := s.Access.ModelFor(models.SetKind)
	f.SetID(s.NewID())
	return SetModel(s, f.(models.Set))
}

func SetModel(s *Space, m models.Set) *Set {
	f := &Set{
		space: s,
		model: m,
	}

	transferAttrs(f.model, f)
	s.Register(f)
	return f
}

func (this *Set) Reload() error {
	this.space.Access.PopulateByID(this.model)
	transferAttrs(this.model, this)
	return nil
}

func (this *Set) Model() models.Set {
	return this.model
}
