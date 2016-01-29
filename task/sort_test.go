package task_test

import (
	"sort"
	"testing"
	"time"

	"github.com/elos/models"
	"github.com/elos/models/task"
)

// a case declaration for salience calculation and sorting
type bySalienceTestCase struct {
	structure map[string]time.Time
	sorted    []string
}

// all the test cases we wish to issue
var bySalienceTestCases = []*bySalienceTestCase{
	&bySalienceTestCase{
		structure: map[string]time.Time{
			"two":   time.Now().Add(48 * time.Hour),
			"one":   time.Now().Add(24 * time.Hour),
			"three": time.Now().Add(72 * time.Hour),
		},
		sorted: []string{"one", "two", "three"},
	},
}

// TestBySalience runs the salience sorting test cases defined
// by "bySalienceTestCases.
func TestBySalience(t *testing.T) {
	t.Parallel()

	for _, testCase := range bySalienceTestCases {
		tasks := make([]*models.Task, len(testCase.structure))

		// construct tasks
		i := 0
		for name, deadline := range testCase.structure {
			t := models.NewTask()
			t.Name = name
			t.Deadline = deadline

			tasks[i] = t
			i++
		}

		sort.Sort(task.BySalience(tasks))

		t.Logf("Structure:\n%+v", testCase.structure)
		t.Logf("Correct Sort:\n%v", testCase.sorted)

		// extract the names for printing, and comparison
		names := make([]string, len(tasks))
		for i := range tasks {
			names[i] = tasks[i].Name
		}

		t.Logf("Actual Sort:\n%v", names)

		// Verify each name is in the correct position
		for i, name := range names {
			if testCase.sorted[i] != name {
				t.Fatalf("Expected '%s' in position %d, but got: '%s'", testCase.sorted[i], i, name)
			}
		}
	}
}
