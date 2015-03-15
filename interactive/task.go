package interactive

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

type Task struct {
	space *Space      `json:"-"`
	model models.Task `json:"-"`

	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	StartTime string   `json:"start_time"`
	EndTime   string   `json:"end_time"`
	UserID    string   `json:"user_id"`
	TaskIDs   []string `json:"task_dependencies" bson:"task_dependencies"`
}

func (this *Task) Save() {
	data.TransferAttrs(this, this.model)
	this.space.Save(this.model)
	this.space.Reload()
}

func (this *Task) Delete() error {
	data.TransferAttrs(this, this.model)
	return this.space.Delete(this.model)
}

func (this *Task) Reload() error {
	this.space.Access.PopulateByID(this.model)
	data.TransferAttrs(this.model, this)
	return nil
}

func (this *Task) Model() models.Task {
	return this.model
}

func NewTask(s *Space) *Task {
	r, _ := s.Access.ModelFor(models.TaskKind)
	r.SetID(s.NewID())
	return TaskModel(s, r.(models.Task))
}

func TaskModel(s *Space, m models.Task) *Task {
	r := &Task{
		space: s,
		model: m,
	}

	data.TransferAttrs(r.model, r)
	s.Register(r)
	return r
}
