package access

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

// The kind that is the user object in the ontology
const UserKind data.Kind = models.UserKind

type Level int

const (
	None Level = iota
	Read
	Write
)

var ImmutableRecords = map[data.Kind]bool{
	models.ContextKind:    true,
	models.CredentialKind: true,
	models.GroupKind:      true,
	models.SessionKind:    true,
}

type Property interface {
	data.Record
	Owner(data.DB) (*models.User, error)
}
