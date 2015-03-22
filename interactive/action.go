package interactive

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

type Action struct {
	space *Space        `json:"-"`
	model models.Action `json:"-"`

	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Completed bool   `json:"completed"`
}

func (this *Action) Save() {
	data.TransferAttrs(this, this.model)
	this.space.Save(this.model)
	this.space.Reload()
}

func (this *Action) Delete() error {
	data.TransferAttrs(this, this.model)
	return this.space.Delete(this.model)
}

func NewAction(s *Space) *Action {
	a, _ := s.Access.ModelFor(models.ActionKind)
	a.SetID(s.NewID())
	return ActionModel(s, a.(models.Action))
}

func ActionModel(s *Space, m models.Action) *Action {
	a := &Action{
		space: s,
		model: m,
	}

	data.TransferAttrs(a.model, a)
	s.Register(a)
	return a
}

func Actions(s *Space, models []models.Action) []*Action {
	actions := make([]*Action, len(models))
	for i := range actions {
		actions[i] = ActionModel(s, models[i])
	}

	return actions
}

func (this *Action) Reload() error {
	this.space.Access.PopulateByID(this.model)
	data.TransferAttrs(this.model, this)
	return nil
}

func (this *Action) Complete() {
	this.model.Complete()
	this.space.Access.Save(this.model)
	this.space.Reload()
}

func (this *Action) Model() models.Action {
	return this.model
}
