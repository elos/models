package task

import "github.com/elos/models"

// BySalience implements the sort.Interface interface
// and orders the tasks by their salience
type BySalience []*models.Task

// Len is the number of elements in the collection
func (b BySalience) Len() int {
	return len(b)
}

// Less reports whether the element with
// index i should sort before the element with
// index j.
func (b BySalience) Less(i, j int) bool {
	// highest salience in lowest index
	return Salience(b[i]) > Salience(b[j])
}

// Swap swaps the elements with indices i and j.
func (b BySalience) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// ByCompletedAt implements the sort.Interface interface
// and orders the tasks by their completed at date
type ByCompletedAt []*models.Task

// Len is the number of elements in the collection
func (b ByCompletedAt) Len() int {
	return len(b)
}

// Less reports whetehr the element with index i
// should sort before the elemeent with index j.
func (b ByCompletedAt) Less(i, j int) bool {
	// earliest task in the lowest index
	return b[i].CompletedAt.Local().Before(b[j].CompletedAt.Local())
}

// Swap swaps the elements with indices i and j.
func (b ByCompletedAt) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
