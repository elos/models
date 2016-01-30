package tag_test

import (
	"testing"

	"github.com/elos/data/builtin/mem"
	"github.com/elos/models"
	"github.com/elos/models/tag"
)

func TestEventsFor(t *testing.T) {
	db := mem.NewDB()

	u := models.NewUser()
	u.SetID(db.NewID())
	if err := db.Save(u); err != nil {
		t.Fatal(err)
	}

	e1 := models.NewEvent()
	e2 := models.NewEvent()
	e1.SetID(db.NewID())
	e2.SetID(db.NewID())
	e1.SetOwner(u)
	e2.SetOwner(u)

	tg, err := tag.ForName(db, u, tag.Name("yo"))
	if err != nil {
		t.Fatal(err)
	}

	e1.IncludeTag(tg)
	e2.IncludeTag(tg)

	if err := db.Save(e1); err != nil {
		t.Fatal(err)
	}

	if err := db.Save(e2); err != nil {
		t.Fatal(err)
	}

	t.Logf("Event 1:\n%+v", e1)
	t.Logf("Event 2:\n%+v", e2)

	events, err := tag.EventsFor(db, tg)
	if err != nil {
		t.Fatal(err)
	}

	if len(events) != 2 {
		t.Fatal("Expected there to be two events")
	}
}

func TestTagsFor(t *testing.T) {
	db := mem.NewDB()

	u := models.NewUser()
	u.SetID(db.NewID())
	if err := db.Save(u); err != nil {
		t.Fatal(err)
	}

	t1 := models.NewTask()
	t2 := models.NewTask()
	t1.SetID(db.NewID())
	t2.SetID(db.NewID())
	t1.SetOwner(u)
	t2.SetOwner(u)

	tg, err := tag.ForName(db, u, tag.Name("yo"))
	if err != nil {
		t.Fatal(err)
	}

	t1.IncludeTag(tg)
	t2.IncludeTag(tg)

	if err := db.Save(t1); err != nil {
		t.Fatal(err)
	}

	if err := db.Save(t2); err != nil {
		t.Fatal(err)
	}

	t.Logf("Task 1:\n%+v", t1)
	t.Logf("Task 2:\n%+v", t2)

	tasks, err := tag.TasksFor(db, tg)
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 2 {
		t.Fatal("Expected there to be two tasks")
	}
}
