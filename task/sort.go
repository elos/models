package task

import "github.com/elos/models"

// by salience implements the sort.Interface interface
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
