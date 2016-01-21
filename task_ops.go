package models

import (
	"sort"
	"time"

	"github.com/elos/data"
	"github.com/nlandolfi/graph"
)

func emptyTime(t time.Time) bool {
	return t == *new(time.Time)
}

func (t *Task) Salience() float64 {
	if t.Complete {
		return -1
	}

	if emptyTime(t.Deadline) {
		return 0
	}

	return 1 / (t.Deadline.Sub(time.Now()).Hours())
}

func (t *Task) InProgress() bool {
	return len(t.Stages)%2 == 1
}

func (t *Task) Start() {
	if !t.InProgress() {
		t.Stages = append(t.Stages, time.Now())
		t.UpdatedAt = time.Now()
	}
}

func (t *Task) Stop() {
	if t.InProgress() {
		t.Stages = append(t.Stages, time.Now())
		t.UpdatedAt = time.Now()
	}
}

func (t *Task) StopAndComplete() {
	t.Stop()
	t.Complete = true
	t.UpdatedAt = time.Now()
}

type bySalience []*Task

// Len is the number of elements in the collection.
func (b bySalience) Len() int {
	return len(b)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (b bySalience) Less(i, j int) bool {
	// highest salience in lowest index
	return b[i].Salience() > b[j].Salience()
}

// Swap swaps the elements with indexes i and j.
func (b bySalience) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type TaskGraph struct {
	book  map[string]int
	tasks []*Task
	nodes []graph.Node
}

func NewTaskGraph(tasks []*Task) *TaskGraph {
	if len(tasks) == 0 {
		panic("Shouldn't have 0 tasks")
	}

	sort.Sort(bySalience(tasks))

	g := &TaskGraph{
		book:  make(map[string]int),
		tasks: tasks,
		nodes: make([]graph.Node, len(tasks)),
	}

	for i := range tasks {
		g.nodes[i] = graph.NewNode(i)
		g.book[tasks[i].ID().String()] = i
	}

	for i, task := range tasks {
		prereqs := make([]graph.Node, 0)
		ids := make(map[string]bool)
		for _, id := range task.PrerequisitesIds {
			ids[id] = true
		}

		for i := range g.nodes {
			if _, ok := ids[tasks[i].ID().String()]; ok {
				prereqs = append(prereqs, g.nodes[i])
			}
		}

		g.nodes[i].SetEdges(prereqs)
	}

	return g
}

func (g *TaskGraph) searchMax(t *Task) *Task {
	if len(t.PrerequisitesIds) == 0 {
		return t
	}

	edges := g.nodes[g.book[t.ID().String()]].Edges()
	prereqs := make([]*Task, len(edges))
	for i, e := range edges {
		prereqs[i] = g.tasks[e.ID()]

	}
	sort.Sort(bySalience(prereqs))

	if len(prereqs) > 0 {
		return g.searchMax(prereqs[0])
	} else {
		// otherwise, for whatever reason those prerequisites were not
		// included in our list of tasks (perhaps they were completed)
		// in which case, ignore it
		return t
	}

}

func (g *TaskGraph) Suggest() *Task {
	return g.searchMax(g.tasks[0])
}

func (t *Task) Tag(db data.DB, name string) (*Tag, error) {
	u, err := t.Owner(db)
	if err != nil {
		return nil, err
	}

	tag, err := TagByName(db, name, u)
	if err != nil {
		return nil, err
	}

	t.IncludeTag(tag)

	if err := db.Save(t); err != nil {
		return nil, err
	}

	return tag, nil
}
