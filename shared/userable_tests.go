package shared

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/user"
)

func TestUserable(store data.Store, userable models.Userable, t *testing.T) {
	u, err := user.Create(store)
	if err != nil {
		t.Fatalf("Error while creating user: %s", err)
	}

	if err = userable.SetUser(u); err != nil {
		t.Errorf("Error while setting user: %s", err)
	}

	access := data.NewAccess(u, store)

	uRetrieved, err := userable.User(access)
	if err != nil {
		t.Errorf("Error while looking up user: %s", err)
	}

	if uRetrieved.ID().String() != u.ID().String() {
		t.Errorf("User retrieved does not match the user set")
	}
}

type userableModel interface {
	data.Model
	models.Userable
}

func TestUserOwnedAccessRights(store data.Store, userable userableModel, t *testing.T) {
	u, err := user.Create(store)
	if err != nil {
		t.Fatalf("Error while creating user: %s", err)
	}

	if err = userable.SetUser(u); err != nil {
		t.Errorf("Error while setting user: %s", err)
	}

	u2, err := user.Create(store)
	if err != nil {
		t.Fatalf("Error while creating user: %s", err)
	}

	if !userable.CanRead(u) {
		t.Errorf("Owning user should be able to read the userable")
	}

	if !userable.CanWrite(u) {
		t.Errorf("Owning user should be able to write the userable")
	}

	if userable.CanRead(u2) {
		t.Errorf("Non-owner user should not be able to read the userable")
	}

	if userable.CanWrite(u2) {
		t.Errorf("Non-owner user should not be able to write the userable")
	}
}
