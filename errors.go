package models

import (
	"errors"
	"fmt"

	"github.com/elos/data"
)

type ErrCast error

func CastError(toKind data.Kind) ErrCast {
	return ErrCast(errors.New(fmt.Sprintf("cast error: failed to cast model to %s kind", toKind)))
}
