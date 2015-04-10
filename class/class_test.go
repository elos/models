package class_test

import (
	"testing"

	"github.com/elos/models/class"
	"github.com/elos/models/persistence"
	"github.com/elos/testing/expect"
)

func TestMongo(t *testing.T) {
	s := persistence.Store(persistence.MongoMemoryDB())
	_, err := class.New(s)

	expect.NoError("creating class", err, t)
}
