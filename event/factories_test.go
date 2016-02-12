package event_test

import (
	"log"
	"testing"

	"github.com/elos/data/builtin/mem"
	"github.com/elos/models/event"
	"github.com/elos/models/user"
)

func TestLocationUpdateFactory(t *testing.T) {
	db := mem.NewDB()
	u, _, err := user.Create(db, "username", "password")
	if err != nil {
		t.Fatal(err)
	}

	e, l, err := event.LocationUpdate(db, u, 50, 50, 50)

	if err != nil {
		log.Fatal(err)
	}

	if e == nil {
		log.Fatal("Event should be non-nil")
	}

	if l == nil {
		log.Fatal("Location should be non-nil")
	}
}
