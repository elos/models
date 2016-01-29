package tag_test

import (
	"log"
	"testing"

	"github.com/elos/data/builtin/mem"
	"github.com/elos/models"
	"github.com/elos/models/tag"
)

func TestTaggingTask(t *testing.T) {
	db := mem.NewDB()
	u := models.NewUser()
	u.SetID(db.NewID())
	if err := db.Save(u); err != nil {
		t.Fatal(err)
	}

	tsk := models.NewTask()
	tsk.SetID(db.NewID())
	tsk.SetOwner(u)
	tagName := "TAG"
	tag, err := tag.Task(db, tsk, tagName)
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("Tag:\n%+v", tag)
	t.Logf("Task:\n%+v", tsk)

	tags, err := tsk.Tags(db)
	if err != nil {
		t.Fatal(err)
	}

	if len(tags) != 1 {
		t.Fatal("The task should now have one tag")
	}

	if tags[0].Name != tagName {
		t.Fatal("The task's one tag should be the tag with the name tag.Task was called with")
	}
}
