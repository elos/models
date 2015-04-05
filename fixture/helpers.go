package fixture

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

/*
	Find locates a fixture by an id.

		var store data.Store
		id := "2341234"

		f, err := Find(s, id)

	The error could be data.ErrInvalidDBType, data.ErrInvalidID,
	or an error from store.PopulateByID
*/
func Find(s data.Store, id data.ID) (models.Fixture, error) {
	fixture, err := New(s)
	if err != nil {
		return nil, err
	}

	fixture.SetID(id)

	// Find a user that has specified id
	return fixture, s.PopulateByID(fixture)
}
