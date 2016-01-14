package access

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

var ImmutableRecords = map[data.Kind]bool{
	models.ContextKind:    true,
	models.CredentialKind: true,
	models.GroupKind:      true,
	models.SessionKind:    true,
}

type Level int

const (
	None Level = iota
	Read
	Write
)

type Property interface {
	data.Record
	Owner(data.DB) (*models.User, error)
}
