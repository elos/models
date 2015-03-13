package action_test

import (
	"testing"
	"time"

	"github.com/elos/data"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"

	"github.com/elos/models"
	. "github.com/elos/models/action"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
)

func TestMongoAction(t *testing.T) {
	// a new mock store, TODO: should  move this to mongo?
	// like mongo.NewMockStore(s Schema)
	store := data.NewRecorderStore(data.NewRecorderDBWithType(mongo.DBType), models.Schema)
	store.Register(models.ActionKind, NewM)
	store.Register(models.UserKind, user.NewM)
	store.Register(models.TaskKind, task.NewM)

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

	u, err := user.New(store)
	if err != nil {
		t.Fatalf("user.New() returned %s", err)
	}

	u.SetName(testName)

	ucp, _ := user.New(store)

	access := data.NewAccess(u, store)

	err = a.User(access, ucp)
	if err != nil {
		t.Fatalf("Error populating user model: %s", err)
	}

	if ucp.ID() != bson.ObjectId("") { // this is kinda a hacky check
		t.Errorf("User should start as nil")
	}

	a.SetUser(u)

	err = a.User(access, ucp)
	if err != nil {
		t.Fatalf("Error populateing user model: %s", err)
	}

	if ucp.ID() != u.ID() {
		t.Fatalf("The user copy model should have populated to the id of the set user")
	}

	id := bson.NewObjectId()

	a.SetUserID(id)
	if a.UserID() != id {
		t.Errorf("SetUserID or UserID() doesn't work")
	}

	if a.Completed() != false {
		t.Errorf("Completed() should start as false for an action")
	}

	a.Complete()

	if a.Completed() != true {
		t.Errorf("Complete() should complete an action")
	}

	tt, err := a.Task(access)

	if err != nil {
		t.Errorf("Error getting task %s", err)
	}

	if tt.ID() != *new(bson.ObjectId) {
		t.Errorf("Current task id should be zero id")
	}

	ts, err := task.New(store)

	if err != nil {
		t.Fatalf("task.New() returned %s", err)
	}

	a.SetTask(ts)

	tt, err = a.Task(access)

	if err != nil {
		t.Errorf("Error getting task %s", err)
	}

	if tt.ID() != ts.ID() {
		t.Errorf("Did not appropiately set task")
	}

}
