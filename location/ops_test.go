package location_test

import (
	"testing"

	"github.com/elos/models/location"
)

func TestNewCoords(t *testing.T) {
	al := 0.0
	la := 1.0
	lo := 2.0

	location.NewCoords(al, la, lo)
}
