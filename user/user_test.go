package user_test

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/action"
	"github.com/elos/models/calendar"
	"github.com/elos/models/ontology"
	"github.com/elos/models/routine"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
	"github.com/elos/mongo"
)

func newMongoStore() data.Store {
	db := data.NewMemoryDBWithType(mongo.DBType)
	db.SetIDConstructor(func() data.ID {
		return mongo.NewObjectID()
	})
	store := data.NewStore(db, models.Schema)
	store.Register(models.UserKind, user.NewM)
	store.Register(models.ActionKind, action.NewM)
	store.Register(models.RoutineKind, routine.NewM)
	store.Register(models.TaskKind, task.NewM)
	store.Register(models.CalendarKind, calendar.NewM)
	store.Register(models.OntologyKind, ontology.NewM)

	return store
}

func TestMongo(t *testing.T) {
	s := newMongoStore()
	u, err := user.New(s)
	if err != nil {
		t.Errorf("Expected not error but got %s", err)
	}
	testUser(s, u, t)
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
	access := data.NewAccess(u, s)

	testName := "Nick Landolfi Jr. III -hypens .periods" // all valid
	u.SetName(testName)
	if u.Name() != testName {
		t.Errorf("User should have name %s; got: %s", testName, u.Name())
	}

	testKey := "5PCCM1qBxOx4q_vFmuLcTnxuOj4exGECiQlXWmx1xk9LeRuGPb6qDtomSUZhHTUM"
	u.SetKey(testKey)
	if u.Key() != testKey {
		t.Errorf("User should have key %s; got: %s", testKey, u.Key())
	}

	act, err := action.New(s)
	if err != nil {
		t.Fatalf("Error while creating action: %s", err)
	}

	if err = s.Save(act); err != nil {
		t.Fatalf("Error saving action: %s", err)
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

	// We will use routine as the actionable
	r, err := routine.New(s)
	if err != nil {
		t.Fatalf("Error while creating routine: %s", err)
	}

	if err = s.Save(r); err != nil {
		t.Fatalf("Error while saving routine: %s", err)
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

	if _, err = u.CurrentActionable(access); err != data.ErrNotFound {
		t.Errorf("The user shouldn't have a current actionable after it has been cleared")
	}

	c, err := calendar.New(s)
	if err != nil {
		t.Fatalf("Error while creating calendar: %s", err)
	}

	if err = s.Save(c); err != nil {
		t.Fatalf("Error while saving calendar: %s", err)
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

	o, err := ontology.New(s)
	if err != nil {
		t.Fatalf("Error while creating ontology: %s", err)
	}

	if err = s.Save(o); err != nil {
		t.Fatalf("Error while saving ontology: %s", err)
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
