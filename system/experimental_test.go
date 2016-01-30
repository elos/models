package system_test

import (
	"testing"

	"github.com/elos/data/builtin/mem"
	"github.com/elos/models"
	"github.com/elos/models/system"
	"github.com/elos/models/task"
)

func TestExperimental(t *testing.T) {
	db := mem.NewDB()
	u := models.NewUser()
	u.SetID(db.NewID())
	if err := db.Save(u); err != nil {
		t.Fatal(err)
	}

	tsk := models.NewTask()

	system.DB(db).As(u).CompleteTask(tsk)

	if !task.IsComplete(tsk) {
		t.Fatal("task should be complete")
	}
}
