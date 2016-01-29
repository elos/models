package tag_test

import (
	"testing"

	"github.com/elos/data/builtin/mem"
	"github.com/elos/models"
	"github.com/elos/models/tag"
)

func TestByName(t *testing.T) {
	n := tag.Name("NAME")

	db := mem.NewDB()
	u := models.NewUser()
	u.SetID(db.NewID())
	if err := db.Save(u); err != nil {
		t.Fatal(err)
	}

	tg, err := tag.ByName(db, u, n)

	if err != nil {
		t.Fatal(err)
	}

	if tg.Name != n.String() {
		t.Fatal("Created tag should have the same name as the name given to ByName")
	}
}
