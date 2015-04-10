package action_test

import (
	"testing"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/action"
	"github.com/elos/models/persistence"
	"github.com/elos/models/task"
	"github.com/elos/testing/expect"
	"github.com/elos/testing/modeltest"
)

func TestMongo(t *testing.T) {
	s := persistence.Store(persistence.MongoMemoryDB())
	a := action.New(s)
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
	store := persistence.ModelsStore(s)

	testActioned(store, a, t)
	testCompleted(store, a, t)
	testTask(store, a, t)
	testComplete(store, a, t)
	testAccessProtection(s, a, t)

	modeltest.Userable(s, a, t)
	modeltest.UserOwnedAccessRights(s, a, t)
}

func testActioned(access models.Store, a models.Action, t *testing.T) {
}

func testCompleted(access models.Store, a models.Action, t *testing.T) {
	if a.Completed() {
		t.Errorf("Expect a new action not to be comleted")
	}

	a.SetCompleted(false)

	if a.Completed() {
		t.Errorf("Set completed to false, should be false")
	}

	a.SetCompleted(true)

	if !a.Completed() {
		t.Errorf("Set completed to true, should be true")
	}
}

func testTask(access models.Store, a models.Action, t *testing.T) {
	_, err := a.Task(access)
	expect.EmptyLinkError("Task", err, t)

	tsk, err := task.Create(access)
	expect.NoError("creating task", err, t)

	err = a.SetTask(tsk)
	expect.NoError("setting task", err, t)

}

func testComplete(access models.Store, a models.Action, t *testing.T) {
	a.SetCompleted(false)

	a.Complete()

	if !a.Completed() {
		t.Errorf("Completed should be true")
	}

	if time.Now().Sub(a.EndTime()) > 1*time.Second {
		t.Errorf("Should complete task and set end time")
	}
}
