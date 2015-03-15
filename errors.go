package models

import (
	"errors"
	"fmt"

	"github.com/elos/data"
)

type ErrCast error

func CastError(toKind data.Kind) ErrCast {
	return ErrCast(errors.New(fmt.Sprintf("models error: failed to cast model to %s kind", toKind)))
}

var ErrEmptyRelationship = errors.New("models error: empty relationship")
