package user

import (
	"github.com/elos/models"
	"golang.org/x/net/context"
)

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// userKey is the key for user.User values in Contexts.  It is
// unexported; clients use user.NewContext and user.FromContext
// instead of using this key directly.
var userKey key = 0

// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context, u *models.User) context.Context {
	return context.WithValue(ctx, userKey, u)
}

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) (*models.User, bool) {
	if ctx == nil {
		return nil, false
	}

	v := ctx.Value(userKey)
	if v == nil {
		return nil, false
	}

	u, ok := v.(*models.User)
	return u, ok
}
