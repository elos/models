package models

import (
	"errors"

	"github.com/elos/data"
)

func Authenticate(db data.DB, public, private string) (*Credential, error) {
	credentialsIter, err := db.NewQuery(CredentialKind).Select(data.AttrMap{"public": public}).Execute()
	if err != nil {
		return nil, err
	}

	credential := NewCredential()
	credentialsIter.Next(credential)

	if credential.Challenge(private) {
		return credential, nil
	}

	return nil, errors.New("challenge failed")
}
