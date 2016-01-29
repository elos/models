package task_test

import (
	"testing"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/task"
)

type graphTestCase struct {
	tasks     map[string]time.Time
	prereqs   map[string][]string
	suggested string
}

var graphTestCases = []*graphTestCase{
	// three <- two <- one   [one is the max]
	&graphTestCase{
		tasks: map[string]time.Time{
			"one":   time.Now().Add(1 * time.Minute),
			"two":   time.Now().Add(24 * time.Hour),
			"three": time.Now().Add(24 * time.Hour),
		},
		prereqs: map[string][]string{
			"one": []string{"two"},
			"two": []string{"three"},
		},
		suggested: "three",
	},
}

// TestGraph tests the task.Graph data structure, and it's Suggest algorithm
func TestGraph(t *testing.T) {
	t.Parallel()

	for _, testCase := range graphTestCases {
		indices := make(map[string]int)
		tasks := make([]*models.Task, len(testCase.tasks))
		i := 0
		for name, deadline := range testCase.tasks {
			t := models.NewTask()
			t.SetID(data.ID(string(i)))
			t.Name = name
			t.Deadline = deadline

			tasks[i] = t
			indices[name] = i
			i++
		}

		for name, reqs := range testCase.prereqs {
			for _, req := range reqs {
				tasks[indices[name]].IncludePrerequisite(tasks[indices[req]])
			}
		}

		g := task.NewGraph(tasks)
		suggested := g.Suggest()
		if suggested.Name != testCase.suggested {
			t.Fatalf("Expected graph to suggest '%s', instead got '%s'", testCase.suggested, suggested.Name)
		}
	}
}
