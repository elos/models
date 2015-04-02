package action_test

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/action"
	"github.com/elos/models/persistence"
	"github.com/elos/models/shared"
)

func TestMongo(t *testing.T) {
	s := persistence.Store(persistence.MongoMemoryDB())
	a, err := action.New(s)
	if err != nil {
		t.Errorf("Error from calendar.New, expected no error but got %s", err)
	}

	testAction(s, a, t)

	if a.Version() != 1 {
		t.Errorf("Expected mongoAction version to be 1, got %d", a.Version())
	}

	if a.Kind() != models.ActionKind {
		t.Errorf("Expected mongoAction kind to equal models.ActionKind, got %s", a.Kind())
	}

	if a.Schema() != models.Schema {
		t.Errorf("Expected mongoAction schema to be models.Schema")
	}
}

func testAction(s data.Store, a models.Action, t *testing.T) {
	access := data.NewAnonAccess(s)

	testActioned(access, a, t)
	testCompleted(access, a, t)
	testTask(access, a, t)
	testComplete(access, a, t)
	testAccessProtection(s, a, t)

	shared.TestUserable(s, a, t)
	shared.TestUserOwnedAccessRights(s, a, t)
	shared.TestAnonReadAccess(s, a, t)
}

func testActioned(access data.Access, a models.Action, t *testing.T) {
}

func testCompleted(access data.Access, a models.Action, t *testing.T) {
}

func testTask(access data.Access, a models.Action, t *testing.T) {
}

func testComplete(access data.Access, a models.Action, t *testing.T) {
}

func testAccessProtection(s data.Store, a models.Action, t *testing.T) {
}
