package habit_test

import (
	"testing"
	"time"

	"github.com/elos/data/builtin/mem"
	"github.com/elos/models/habit"
	"github.com/elos/models/user"
)

// super basic
func TestCheckinBasic(t *testing.T) {
	db := mem.NewDB()
	u, _, err := user.Create(db, "public", "private")
	if err != nil {
		t.Fatal(err)
	}
	h, err := habit.Create(db, u, "Habitual")
	if err != nil {
		t.Fatal(err)
	}

	if ok, err := habit.DidCheckinOn(db, h, time.Now()); err != nil {
		t.Fatal(err)
	} else if ok {
		t.Fatal("Should not be checked in")
	}

	if _, err = habit.CheckinFor(db, h, "notes", time.Now()); err != nil {
		t.Fatal(err)
	}

	if ok, err := habit.DidCheckinOn(db, h, time.Now()); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fatal("Should now be checked in")
	}
}
