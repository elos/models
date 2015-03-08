package structures

import "github.com/elos/data"

// A timeable tree is a binary tree of timeable
// it keeps them orderd by start time. This data
// structure is useful for elos fixtures
type TimeableTree struct {
	Left  *TimeableTree
	Right *TimeableTree
	data.Timeable
}

var emptyTimeableTree = new(TimeableTree)

func (t *TimeableTree) Empty() bool {
	return *t == *emptyTimeableTree
}

var i int = 0

func (t *TimeableTree) Add(other data.Timeable) {
	if t.Empty() {
		t.Timeable = other
		return
	}

	if other.StartTime().Before(t.Timeable.StartTime()) {
		if t.Left == nil {
			t.Left = new(TimeableTree)
		}

		t.Left.Add(other)
		return
	} else {
		if t.Right == nil {
			t.Right = new(TimeableTree)
		}

		t.Right.Add(other)
		return
	}
}

func (t *TimeableTree) Stream() <-chan data.Timeable {
	c := make(chan data.Timeable, 1000)
	go t.walk(&c)
	return c
}

func (t *TimeableTree) walk(c *chan data.Timeable) {
	if t == nil || t.Empty() {
		return
	}

	if t.Left != nil {
		t.Left.walk(c)
	}

	*c <- t.Timeable

	if t.Right != nil {
		t.Right.walk(c)
	}
}
