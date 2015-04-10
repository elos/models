package interactive

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/models"
)

type MongoModel struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	space *Space      `json:"-"`
	model models.User `json:"-"`

	MongoModel

	Name string `json:"name"`
	Key  string `json:"key"`

	ActionIDs  []string `json:"action_ids"`
	EventIDs   []string `json:"event_ids"`
	TaskIDs    []string `json:"task_ids"`
	RoutineIDs []string `json:"routine_ids"`

	OntologyID      string `json:"ontology_id"`
	CalendarID      string `json:"calendar_id"`
	CurrentActionID string `json:"current_action_id"`
	ActionableKind  string `json:"actionable_kind"`
	ActionableID    string `json:"actionable_id"`
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
	u, _ := s.Store.ModelFor(models.UserKind)
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
	this.space.Store.PopulateByID(this.model)
	data.TransferAttrs(this.model, this)
	return nil
}

func (u *User) Add(v interface{}) {
	switch v.(type) {
	case *Action:
	case *Task:
	case *Routine:
	default:
		log.Printf("Can't add %+v", v)
	}
}

func (u *User) Actions() []*Action {
	actions, err := u.model.Actions(u.space.Store.(models.Store))
	if err != nil {
		log.Print(err)
		return nil
	}

	return Actions(u.space, actions)
}

func (u *User) CurrentAction() *Action {
	return u.space.FindAction(u.CurrentActionID)
}

func (u *User) SetCurrentAction(a *Action) {
	other := a.Model()
	u.model.SetCurrentAction(other)
	u.space.Store.Save(u.model)
	u.space.Store.Save(other)
	u.space.Reload()
}

func (u *User) SetCurrentRoutine(r *Routine) {
	other := r.Model()
	u.model.SetCurrentActionable(other)
	u.space.Store.Save(u.model)
	u.space.Store.Save(other)
	u.space.Reload()
}
