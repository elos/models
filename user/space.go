package user

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

type SpaceUser struct {
	memory *data.MemoryStore `json:"-"`
	model  models.User       `json:"-"`

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
}

func NewSpaceUser(s *data.MemoryStore) (u *SpaceUser, err error) {
	m, err := New(s.Access)
	if err != nil {
		return
	}

	u.model = m
	u.memory = s
	s.RegisterObject(u)

	err = data.TransferAttrs(u.model, u)

	return
}

func (u *SpaceUser) Reload() error {
	u.memory.Access.PopulateByID(u.model)
	return data.TransferAttrs(u.model, u)
}

func (u *SpaceUser) Save() {
	data.TransferAttrs(u, u.model)
	u.memory.Access.Save(u.model)
	u.memory.ReloadObjects()
}

func (u *SpaceUser) Delete() {
	data.TransferAttrs(u, u.model)
	u.memory.Access.Delete(u.model)
}
