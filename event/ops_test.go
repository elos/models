package event_test

import (
	"testing"

	"github.com/elos/data/builtin/mem"
	"github.com/elos/models"
	"github.com/elos/models/event"
)

func TestContainsTags(t *testing.T) {
	db := mem.NewDB()

	e := models.NewEvent()

	tag1 := models.NewTag()
	tag2 := models.NewTag()
	tag1.SetID(db.NewID())
	tag2.SetID(db.NewID())

	e.IncludeTag(tag1)

	if event.ContainsTags(e, tag1, tag2) {
		t.Fatal("Shouldn't have included both tags")
	}

	e.IncludeTag(tag2)

	if !event.ContainsTags(e, tag1, tag2) {
		t.Fatal("Should have included both tags")
	}
}
