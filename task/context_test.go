package task_test

import (
	"testing"

	"github.com/elos/models"
	"github.com/elos/models/task"
	"golang.org/x/net/context"
)

// TestTaskContext tests the utility functions that allow for
// storing a task on a context and subsequently retrieving it
func TestTaskContext(t *testing.T) {
	tsk := models.NewTask()
	ctx := context.Background()

	ctx = task.NewContext(ctx, tsk)

	if tskRetrieved, ok := task.FromContext(ctx); !ok {
		t.Fatal("Should have succesfully retrieved the task")
	} else {
		if tskRetrieved != tsk {
			t.Fatal("Should have retrieved the same task")
		}
	}
}
