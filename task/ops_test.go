package task_test

import (
	"testing"
	"time"

	"github.com/elos/models"
	"github.com/elos/models/task"
)

// TestTaskCompletionBasic tests instantiating a task,
// it's completion state then, and then the completion
// of the task, and it's completeion state afterward.
func TestTaskCompletionBasic(t *testing.T) {
	t.Parallel()

	tsk := models.NewTask()

	t.Logf("Created task\n\t%+v", tsk)

	if task.IsComplete(tsk) {
		t.Fatal("Task should begin incomplete")
	}

	t.Log("Stopping and completing task")
	stopTime := task.StopAndComplete(tsk)
	t.Log("\tstopped and completed")

	t.Logf("Completed task:\n\t%+v", tsk)

	if !task.IsComplete(tsk) {
		t.Fatal("Task should now be incomplete")
	}

	if !tsk.UpdatedAt.Equal(stopTime) || !tsk.UpdatedAt.Equal(tsk.CompletedAt) {
		t.Fatal("Task should have an UpdatedAt time equal to it's stop time")
	}
}

// TestTaskCompletionInProgress test instantiting a task,
// starting it, and then completing it while it is in
// progress
func TestTaskCompletionInProgress(t *testing.T) {
	t.Parallel()

	tsk := models.NewTask()

	t.Logf("Created task\n\t%+v", tsk)

	if task.IsComplete(tsk) {
		t.Fatal("Task should begin incomplete")
	}

	startTime := task.Start(tsk)

	if !task.InProgress(tsk) {
		t.Fatal("Task should now be in progress")
	}

	t.Log("Stopping and completing task")
	stopTime := task.StopAndComplete(tsk)
	t.Log("\tstopped and completed")
	t.Logf("Completed task:\n\t%+v", tsk)

	if task.InProgress(tsk) {
		t.Fatal("Task should have been stopped")
	}

	if !task.IsComplete(tsk) {
		t.Fatal("Task should now be complete")
	}

	// sanity check
	if !stopTime.After(startTime) {
		t.Fatal("The start time should be before the stop time")
	}
}

// TestSalienceBasic tests that a task with a future deadline
// has a positive salience
func TestSalienceBasic(t *testing.T) {
	tsk := models.NewTask()
	tsk.Deadline = time.Now().Add(25 * time.Hour)

	if s := task.Salience(tsk); s <= 0 {
		t.Fatalf("Task should have positive salience, but had salience of: %f", s)
	}
}
