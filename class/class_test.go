package class_test

import (
	"testing"

	"github.com/elos/models/class"
	"github.com/elos/models/persistence"
)

func TestMongo(t *testing.T) {
	s := persistence.Store(persistence.MongoMemoryDB())
	class.New(s)
}
