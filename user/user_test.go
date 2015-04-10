package user_test

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/action"
	"github.com/elos/models/calendar"
	"github.com/elos/models/event"
	"github.com/elos/models/ontology"
	"github.com/elos/models/persistence"
	"github.com/elos/models/routine"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
	"github.com/elos/testing/modeltest"
)

func TestMongo(t *testing.T) {
	s := persistence.Store(persistence.MongoMemoryDB())
	u := user.New(s)
	testUser(s, u, t)

	if u.Version() != 1 {
		t.Errorf("Expected mongoUser version to be 1, got %d", u.Version())
	}

	if u.Kind() != models.UserKind {
		t.Errorf("Expected mongoUser kind to equal models.UserKind, got %s", u.Kind())
	}

	if u.Schema() != models.Schema {
		t.Errorf("Expected mongoUser schema to be models.Schema")
	}
}

/*
	TestUser is the generic test suite for any
	implementations of models.User. Whenever a new
	implementation is written it should be tested
	at the least by using this function.

	TestUser does not test anything implementation
	specific, only at the heigh models.User interface level.
*/
func testUser(s data.Store, u models.User, t *testing.T) {
	store := persistence.ModelsStore(s)

	testName(store, u, t)
	testKey(store, u, t)
	testCurrentAction(store, u, t)
	testCurrentActionable(store, u, t)
	testCalendar(store, u, t)
	testOntology(store, u, t)
	testEvents(store, u, t)
	testTasks(store, u, t)
	testRoutines(store, u, t)

	modeltest.AnonReadAccess(s, u, t)
}

func testName(access models.Store, u models.User, t *testing.T) {
	testName := "Nick Landolfi Jr. III -hypens .periods" // all valid
	u.SetName(testName)
	if u.Name() != testName {
		t.Errorf("User should have name %s; got: %s", testName, u.Name())
	}
}

func testKey(access models.Store, u models.User, t *testing.T) {
	testKey := user.NewKey()
	u.SetKey(testKey)
	if u.Key() != testKey {
		t.Errorf("User should have key %s; got: %s", testKey, u.Key())
	}
}

func testCurrentAction(access models.Store, u models.User, t *testing.T) {
	act, err := action.Create(access)
	if err != nil {
		t.Fatalf("Error while creating action: %s", err)
	}

	if err = u.SetCurrentAction(act); err != nil {
		t.Errorf("Error while setting current action: %s", err)
	}

	act2, err := u.CurrentAction(access)
	if err != nil {
		t.Fatalf("Error while retrieving current action: %s", err)
	}

	if act.ID().String() != act2.ID().String() {
		t.Errorf("Retrieved current action id differs from set current action id expected %s, got: %s", act.ID(), act2.ID())
	}
}

func testCurrentActionable(access models.Store, u models.User, t *testing.T) {
	// We will use routine as the actionable
	r, err := routine.Create(access)
	if err != nil {
		t.Fatalf("Error while creating routine: %s", err)
	}

	if err = u.SetCurrentActionable(r); err != nil {
		t.Errorf("Error while setting current actionable: %s", err)
	}

	r2, err := u.CurrentActionable(access)
	if err != nil {
		t.Fatalf("Error while retrieving current actionable")
	}

	if r.ID().String() != r2.ID().String() {
		t.Errorf("Retrieved current actionable id differs from set current actionable id. Expected: %s, got: %s", r.ID(), r2.ID())
	}

	if r.Kind() != r2.Kind() {
		t.Errorf("Retrieved current actionable kind differs from set current actionable kind. Expected %s, got: %s", r.Kind(), r2.Kind())
	}

	u.ClearCurrentActionable()

	if _, err = u.CurrentActionable(access); err != models.ErrEmptyRelationship {
		t.Errorf("The user shouldn't have a current actionable after it has been cleared")
	}
}

func testCalendar(access models.Store, u models.User, t *testing.T) {
	c, err := calendar.Create(access)
	if err != nil {
		t.Fatalf("Error while creating calendar: %s", err)
	}

	if err = u.SetCalendar(c); err != nil {
		t.Errorf("Error while setting calendar: %s", err)
	}

	c2, err := u.Calendar(access)
	if err != nil {
		t.Fatalf("Error while retrieving calendar: %s", err)
	}

	if c2.ID().String() != c.ID().String() {
		t.Errorf("Retrieved calendar id doesn't equal set calendar id. Expected: %s, got: %s", c.ID(), c2.ID())
	}
}

func testOntology(access models.Store, u models.User, t *testing.T) {
	o, err := ontology.Create(access)
	if err != nil {
		t.Fatalf("Error while creating ontology: %s", err)
	}

	if err = u.SetOntology(o); err != nil {
		t.Errorf("Error while setting ontology: %s", err)
	}

	o2, err := u.Ontology(access)
	if err != nil {
		t.Fatalf("Error while retrieving ontology", err)
	}

	if o2.ID().String() != o.ID().String() {
		t.Errorf("Retrieved ontology id doesn't equal set ontology id. Expected %s, got: %s", o.ID(), o2.ID())
	}

}

func testEvents(access models.Store, u models.User, t *testing.T) {
	e1, err := event.Create(access)
	if err != nil {
		t.Fatalf("Error while creating event: %s", err)
	}
	e2, err := event.Create(access)
	if err != nil {
		t.Fatalf("Error while creating event: %s", err)
	}

	if err = u.IncludeEvent(e1); err != nil {
		t.Errorf("Error while including event: %s", err)
	}

	if err = u.IncludeEvent(e2); err != nil {
		t.Errorf("Error while includeing event: %s", err)
	}

	events, err := u.Events(access)
	if err != nil {
		t.Fatalf("Error retrieving user events: %s", err)
	}

	if len(events) != 2 {
		t.Errorf("Expected user to have 2 events, retrieved %d", len(events))
	}

	if err = u.ExcludeEvent(e1); err != nil {
		t.Errorf("Error while excluding event: %s", err)
	}

	events, err = u.Events(access)
	if err != nil {
		t.Fatalf("Error retrieving user evnets: %s", err)
	}

	if len(events) != 1 {
		t.Fatalf("Expected user to have 1 event, retrieved %d", len(events))
	}

	e2Copy := events[0]

	if e2Copy.ID().String() != e2.ID().String() {
		t.Errorf("Expected to find e2 as the only event")
	}
}

func testTasks(access models.Store, u models.User, t *testing.T) {
	t1, err := task.Create(access)
	if err != nil {
		t.Fatalf("Error while creating task: %s", err)
	}
	t2, err := task.Create(access)
	if err != nil {
		t.Fatalf("Error while creating task: %s", err)
	}

	if err = u.IncludeTask(t1); err != nil {
		t.Errorf("Error while including task: %s", err)
	}

	if err = u.IncludeTask(t2); err != nil {
		t.Errorf("Error while including task: %s", err)
	}

	tasks, err := u.Tasks(access)
	if err != nil {
		t.Fatalf("Error retrieving user tasks: %s", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected user to have 2 tasks, retrieved %d", len(tasks))
	}

	if err = u.ExcludeTask(t1); err != nil {
		t.Errorf("Error while excluding task: %s", err)
	}

	tasks, err = u.Tasks(access)
	if err != nil {
		t.Fatalf("Error retrieving user tasks: %s", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected user to have 1 task, retrieved %d", len(tasks))
	}

	t2Copy := tasks[0]

	if t2Copy.ID().String() != t2.ID().String() {
		t.Errorf("Expected to find t2 as the only event")
	}
}

func testRoutines(access models.Store, u models.User, t *testing.T) {
	r1, err := routine.Create(access)
	if err != nil {
		t.Fatalf("Error while creating routine: %s", err)
	}
	r2, err := routine.Create(access)
	if err != nil {
		t.Fatalf("Error while creating routine: %s", err)
	}

	if err = u.IncludeRoutine(r1); err != nil {
		t.Errorf("Error while including routine: %s", err)
	}

	if err = u.IncludeRoutine(r2); err != nil {
		t.Errorf("Error while including routine: %s", err)
	}

	routines, err := u.Routines(access)
	if err != nil {
		t.Fatalf("Error retrieving user routines: %s", err)
	}

	if len(routines) != 2 {
		t.Errorf("Expected user to have 2 routines, retrieved %d", len(routines))
	}

	if err = u.ExcludeRoutine(r1); err != nil {
		t.Errorf("Error while excluding routine: %s", err)
	}

	routines, err = u.Routines(access)
	if err != nil {
		t.Fatalf("Error retrieving user routine: %s", err)
	}

	if len(routines) != 1 {
		t.Fatalf("Expected user to have 1 routine, retrieved %d", len(routines))
	}

	r2Copy := routines[0]

	if r2Copy.ID().String() != r2.ID().String() {
		t.Errorf("Expected to find t2 as the only event")
	}
}
