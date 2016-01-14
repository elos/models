package access

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/models"
)

func CanCreate(u *models.User, k data.Kind) bool {
	if k == models.UserKind {
		return false
	} else {
		return true
	}
}

func CanRead(db data.DB, u *models.User, record data.Record) (bool, error) {
	property, ok := record.(Property)

	if !ok {
		return data.Equivalent(u, record), nil
	}

	if owner, err := property.Owner(db); err != nil {
		if err == models.ErrEmptyLink {
			log.Print("Model without an owner!")
			return false, nil
		}

		return false, err
	} else {
		if data.Equivalent(u, owner) {
			return true, nil
		} else {
			log.Print("user is not the model's owner")
		}
	}

	groups, err := u.Groups(db)

	if err == models.ErrEmptyLink {
		return false, nil
	} else if err != nil {
		return false, err
	}

	for _, g := range groups {
		contexts, err := g.Contexts(db)

		if err != nil {
			return false, err
		}

		for _, c := range contexts {
			if c.Contains(property) {
				if Level(g.Access) > None {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func CanWrite(db data.DB, u *models.User, record data.Record) (bool, error) {
	property, ok := record.(Property)

	if !ok {
		return data.Equivalent(u, record), nil
	}

	if _, immutable := ImmutableRecords[record.Kind()]; immutable {
		return false, nil
	}

	if owner, err := property.Owner(db); err != nil {
		if err == models.ErrEmptyLink {
			log.Print("Model without an owner!")
			return false, nil
		}

		return false, err
	} else {
		if data.Equivalent(u, owner) {
			return true, nil
		} else {
			log.Print("user is not the model's owner")
		}
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
				if Level(g.Access) > Read {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func CanDelete(db data.DB, u *models.User, record data.Record) (bool, error) {
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
