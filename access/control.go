package access

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

func CanCreate(u User, k data.Kind) bool {
	if k == models.UserKind {
		return false
	} else {
		return true
	}
}

func CanRead(db data.DB, u User, record data.Record) (bool, error) {
	property, ok := record.(Property)

	if !ok {
		return data.Equivalent(u, record), nil
	}

	groups, err := u.Groups(db)

	if err != nil {
		return false, err
	}

	for _, g := range groups {
		contexts, err := g.Contexts(db)

		if err != nil {
			return false, err
		}

		for _, c := range contexts {
			if c.Contains(property) {
				if Level(g.Access()) > None {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func CanWrite(db data.DB, u User, record data.Record) (bool, error) {
	property, ok := record.(Property)

	if !ok {
		return data.Equivalent(u, record), nil
	}

	if _, immutable := ImmutableRecords[record.Kind()]; immutable {
		return false, nil
	}

	groups, err := u.Groups(db)

	if err != nil {
		return false, err
	}

	for _, g := range groups {
		contexts, err := g.Contexts(db)

		if err != nil {
			return false, err
		}

		for _, c := range contexts {
			if c.Contains(property) {
				if Level(g.Access()) > Read {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func CanDelete(db data.DB, u User, record data.Record) (bool, error) {
	property, ok := record.(Property)

	if !ok {
		return data.Equivalent(u, record), nil
	}

	owner, err := property.Owner(db)

	if err != nil {
		return false, err
	}

	return data.Equivalent(u, owner), nil
}
