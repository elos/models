package task

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
var taskKey key = 0

// NewContext returns a new Context that carries value t.
func NewContext(ctx context.Context, t *models.Task) context.Context {
	return context.WithValue(ctx, taskKey, t)
}

// FromContext returns the Task value stored in ctx, if any.
func FromContext(ctx context.Context) (*models.Task, bool) {
	t, ok := ctx.Value(taskKey).(*models.Task)
	return t, ok
}
