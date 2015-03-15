package action_test

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/models"
	. "github.com/elos/models/action"
	"github.com/elos/mongo"
)

func Store() data.Store {
	s := data.NewRecorderStore(data.NewRecorderDBWithType(mongo.DBType), models.Schema)
	s.Register(models.ActionKind, NewM)
	return s
}

func BadStore() data.Store {
	return data.NewRecorderStore(data.NewRecorderDBWithType("alkdsfj"), models.Schema)
}

func TestNew(t *testing.T) {
	store := Store()

	m, err := store.ModelFor(models.ActionKind)

	if err != nil {
		t.Fatalf("ModelFor returned an error: %s", err)
	}

	_, ok := m.(models.Action)

	if !ok {
		t.Errorf("ModelFor didn't return an action, got %+v", m)
	}

	store = BadStore()
	m, err = New(store)

	if err != data.ErrInvalidDBType {
		t.Errorf("New(store) w/ faulty type should => ErrInvalidDBType")
	}

	if m != nil {
		t.Errorf("When New() errors it should return a nil model")
	}

	//TODO check if id gets set
}

func TestCreate(t *testing.T) {
	store := Store()

	id := mongo.NewObjectID()

	attrsID := data.AttrMap{
		"id":    id,
		"junk":  "garbage",
		"more":  123123123,
		"trash": []string{"1324", "1234", "1234"},
	}

	a, err := CreateAttrs(store, attrsID)

	if err != nil {
		t.Error(err)
	}

	if a.ID() != id {
		t.Errorf("Create failed to extract the id")
	}

	store = BadStore()

	a, err = CreateAttrs(store, attrsID)

	if err != data.ErrInvalidDBType || a != nil {
		t.Errorf("Create should propogate invalid db type errors")
	}

	attrsID["id"] = "trash"

	a, err = CreateAttrs(Store(), attrsID)
	if err != data.ErrInvalidID || a == nil {
		t.Errorf("ID was bad create shouldn't choke, go err %s, model %+v", err, a)
	}
}
