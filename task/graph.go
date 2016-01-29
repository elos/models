package task

import (
	"sort"

	"github.com/elos/models"
	"github.com/nlandolfi/graph"
)

type Graph struct {
	// map from the ID of the task to it's node index
	book  map[string]int
	tasks []*models.Task
	nodes []graph.Node
}

func NewGraph(tasks []*models.Task) *Graph {
	if len(tasks) == 0 {
		panic("shouldn't have 0 tasks")
	}

	sort.Sort(BySalience(tasks))

	g := &Graph{
		book:  make(map[string]int),
		tasks: tasks,
		nodes: make([]graph.Node, len(tasks)),
	}

	// Construct nodes for all the tasks
	for i := range tasks {
		g.nodes[i] = graph.NewNode(i)
		g.book[tasks[i].ID().String()] = i
	}

	// add prerequisite edges
	// So this task's node points to it's prereqs
	for i, task := range tasks {
		// the nodes which are this task's prereqs
		prereqs := make([]graph.Node, 0)

		for _, id := range task.PrerequisitesIds {
			prereqs = append(prereqs, g.nodes[g.book[id]])
		}

		g.nodes[i].SetEdges(prereqs)
	}

	return g
}

// searches for the most salient prereq of this task
// may be the task itself, walking up the graph
func (g *Graph) searchMax(t *models.Task) *models.Task {
	if len(t.PrerequisitesIds) == 0 {
		return t
	}

	edges := g.nodes[g.book[t.ID().String()]].Edges()

	// Get this tasks' prerequisites, and sort them by salience
	prerequisites := make([]*models.Task, len(edges))
	for i, edge := range edges {
		prerequisites[i] = g.tasks[edge.ID()]
	}
	sort.Sort(BySalience(prerequisites))

	if len(prerequisites) > 0 {
		return g.searchMax(prerequisites[0])
	} else {
		panic("did not find prerequisite tasks")
	}
}

// Implementation of elos' suggest algorithm, currently just
// finds the most salient task, and then works to it's most salient
// prerequisite
func (g *Graph) Suggest() *models.Task {
	return g.searchMax(g.tasks[0])
}
