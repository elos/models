package user_test

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mem"
	"github.com/elos/models/user"
	"golang.org/x/net/context"
)

func TestContextBasic(t *testing.T) {
	db := mem.NewDB()
	u, _, err := user.Create(db, "p", "p")
	if err != nil {
		t.Fatal(err)
	}

	ctx := user.NewContext(context.Background(), u)

	uRetrieved, ok := user.FromContext(ctx)

	if !ok {
		t.Fatal("Could not retrieve user from context")
	}

	if !data.Equivalent(u, uRetrieved) {
		t.Fatal("Retrieved a different user from the context")
	}
}
