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

func (u *User) Tasks(db data.DB, completedOnly bool) ([]*Task, error) {
	taskQuery := db.NewQuery(TaskKind)

	// only retrieve _incomplete_ tasks
	taskQuery.Select(data.AttrMap{
		"owner_id": u.Id,
		"complete": false,
	})

	iter, err := taskQuery.Execute()
	if err != nil {
		return nil, err
	}

	t := NewTask()
	tasks := make([]*Task, 0)
	for iter.Next(t) {
		tasks = append(tasks, t)
		t = NewTask()
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return tasks, nil
}
