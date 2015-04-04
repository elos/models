package action_test

import (
	"testing"
	"time"

	"github.com/elos/data"
	"github.com/elos/mongo"
	"github.com/elos/testing/expect"

	. "github.com/elos/models/action"
	"github.com/elos/models/persistence"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
)

func TestMongoAction(t *testing.T) {
	store := persistence.Store(persistence.MongoMemoryDB())

	a, err := New(store)

	if err != nil {
		t.Errorf("Expected a successful action New, but go error: %s", err)
	}

	if a.DBType() != mongo.DBType {
		t.Errorf("An action created with a mongo store should have the mongo db type")
	}

	if a.Name() != "" {
		t.Errorf("Action's name expected to be \"\", got %s", a.Name())
	}

	testName := "as21394183402!@34$@$kjdfadf"

	a.SetName(testName)

	if a.Name() != testName {
		t.Errorf("Action's name expected to be %s, got %s", testName, a.Name())
	}

	testTime := time.Now()

	if a.StartTime() != *new(time.Time) {
		t.Errorf("StartTime should be zero time")
	}

	a.SetStartTime(testTime)

	if a.StartTime() != testTime {
		t.Errorf("StartTime should be %s, got %s", testTime, a.StartTime())
	}

	if a.EndTime() != *new(time.Time) {
		t.Errorf("EndTime should be zero time")
	}

	a.SetEndTime(testTime)

	if a.EndTime() != testTime {
		t.Errorf("EndTime should be %s, got %s", testTime, a.EndTime())
	}

	u, err := user.Create(store)
	if err != nil {
		t.Fatalf("user.Create() returned %s", err)
	}

	u.SetName(testName)

	access := data.NewAccess(u, store)

	a.SetUser(u)

	ucp, err := a.User(access)
	if err != nil {
		t.Fatalf("Error populateing user model: %s", err)
	}

	if ucp.ID() != u.ID() {
		t.Fatalf("The user copy model should have populated to the id of the set user")
	}

	if a.Completed() != false {
		t.Errorf("Completed() should start as false for an action")
	}

	a.Complete()

	if a.Completed() != true {
		t.Errorf("Complete() should complete an action")
	}

	_, err = a.Task(access)
	expect.EmptyLinkError("Task", err, t)

	ts, err := task.Create(store)

	if err != nil {
		t.Fatalf("task.Create() returned %s", err)
	}

	a.SetTask(ts)

	tt, err := a.Task(access)

	if err != nil {
		t.Errorf("Error getting task %s", err)
	}

	if tt.ID() != ts.ID() {
		t.Errorf("Did not appropiately set task")
	}

}
