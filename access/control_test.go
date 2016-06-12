package access_test

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mem"
	"github.com/elos/models"
	"github.com/elos/models/access"
	"github.com/elos/models/user"
)

// --- record {{{

type record struct {
	id   data.ID
	kind data.Kind
}

func (r *record) ID() data.ID {
	return r.id
}

func (r *record) SetID(i data.ID) {
	r.id = i
}

func (r *record) Kind() data.Kind {
	return r.kind
}

// --- }}}

// --- property {{{

type property struct {
	record
	owner_id data.ID
}

func (p *property) Owner(db data.DB) (*models.User, error) {
	if p.owner_id == data.ID("") {
		return nil, models.ErrEmptyLink
	}

	u := &models.User{
		Id: p.owner_id.String(),
	}

	return u, db.PopulateByID(u)
}

// --- }}}

func TestCanCreate(t *testing.T) {
	db := mem.NewDB()

	u, _, err := user.Create(db, "username", "password")
	if err != nil {
		t.Fatalf("user.Create error: %s", err)
	}

	p := new(property)
	p.owner_id = u.ID()

	if ok, err := access.CanCreate(db, u, p); err != nil {
		t.Fatalf("access.CanCreate error: %s", err)
	} else if got, want := ok, true; got != want {
		t.Errorf("ok, _ := access.CanCreate: got %b, want %b", got, want)
	}

	o, _, err := user.Create(db, "emanresu", "drowssap")
	if err != nil {
		t.Fatalf("user.Create error: %s", err)
	}

	p.owner_id = o.ID()

	if ok, err := access.CanCreate(db, u, p); err != nil {
		t.Fatalf("access.CanCreate error: %s", err)
	} else if got, want := ok, false; got != want {
		t.Errorf("ok, _ := access.CanCreate: got %b, want %b", got, want)
	}

	p.owner_id = data.ID("")

	if ok, err := access.CanCreate(db, u, p); err != nil {
		t.Fatalf("access.CanCreate error: %s", err)
	} else if got, want := ok, false; got != want {
		t.Errorf("ok, _ := access.CanCreate: got %b, want %b", got, want)
	}
}
