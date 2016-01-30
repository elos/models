package task

import (
	"time"

	"github.com/elos/models"
)

// IsComplete determines whether a task has been
// completed
func IsComplete(t *models.Task) bool {
	return !t.CompletedAt.IsZero()
}

// InProgress determines whether a task has been
// "started," and ergo is currently in progress.
func InProgress(t *models.Task) bool {
	return len(t.Stages)%2 == 1
}

// Salience determines the salience of the task
func Salience(t *models.Task) float64 {
	if IsComplete(t) || t.Deadline.IsZero() {
		return 0
	}

	return 1 / (t.Deadline.Sub(time.Now()).Hours())
}

// Start starts the task. It appends the current time
// to the stages of the task.
func Start(t *models.Task) time.Time {
	now := time.Now()
	if !InProgress(t) {
		t.Stages = append(t.Stages, now)
		t.UpdatedAt = now
	}
	return now
}

// Stop stops the task. It appends the current time
// to the stages of the task.
func Stop(t *models.Task) time.Time {
	now := time.Now()
	if InProgress(t) {
		t.Stages = append(t.Stages, now)
		t.UpdatedAt = now
	}
	return now
}

// Stops and completes the current task
func StopAndComplete(t *models.Task) time.Time {
	stopTime := Stop(t)
	t.CompletedAt = stopTime
	t.UpdatedAt = stopTime
	return stopTime
}

func sumIntervals(intervals []time.Time) time.Duration {
	var duration time.Duration
	for i := 0; i < len(intervals); i += 2 {
		duration += intervals[i+1].Sub(intervals[i])
	}
	return duration
}

// TimeSpent calculuates the time spent on a task
func TimeSpent(t *models.Task) time.Duration {
	if !InProgress(t) {
		return sumIntervals(t.Stages)
	}

	// look at right now
	return sumIntervals(append(t.Stages, time.Now()))
}

// CollectiveTimeSpent
func CollectiveTimeSpent(tasks []*models.Task) time.Duration {
	var duration time.Duration
	for _, t := range tasks {
		duration += TimeSpent(t)
	}
	return duration
}
