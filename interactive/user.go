package interactive

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

type User struct {
	space *Space      `json:"-"`
	model models.User `json:"-"`

	ID              string   `json:"id"`
	Name            string   `json:"name"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	Key             string   `json:"key"`
	EventIDs        []string `json:"event_ids"`
	TaskIDs         []string `json:"task_ids"`
	CurrentActionID string   `json:"current_action_id"`
	ActionableKind  string   `json:"actionable_kind"`
	ActionableID    string   `json:"actionable_id"`
	OntologyID      string   `json:"ontology_id"`
	CalendarID      string   `json:"calendar_id"`
}

func (this *User) Space() *Space {
	return this.space
}

func (this *User) Model() data.Model {
	return this.model
}

func UserModel(s *Space, m models.User) *User {
	u := &User{
		space: s,
		model: m,
	}
	data.TransferAttrs(u.model, u)
	s.Register(u)
	return u
}

func NewUser(s *Space) *User {
	u, _ := s.Access.ModelFor(models.UserKind)
	u.SetID(s.NewID())
	return UserModel(s, u.(models.User))
}

func (this *User) Save() {
	data.TransferAttrs(this, this.model)
	this.space.Save(this.model)
	this.space.Reload()
}

func (this *User) Delete() error {
	data.TransferAttrs(this, this.model)
	return this.space.Delete(this.model)
}

func (this *User) Reload() error {
	this.space.Access.PopulateByID(this.model)
	data.TransferAttrs(this.model, this)
	return nil
}

func (u *User) CurrentAction() *Action {
	return u.space.FindAction(u.CurrentActionID)
}

func (u *User) SetCurrentAction(a *Action) {
	other := a.Model()
	u.model.SetCurrentAction(other)
	u.space.Access.Save(u.model)
	u.space.Access.Save(other)
	u.space.Reload()
}

func (u *User) SetCurrentRoutine(r *Routine) {
	other := r.Model()
	u.model.SetCurrentActionable(other)
	u.space.Access.Save(u.model)
	u.space.Access.Save(other)
	u.space.Reload()
}
