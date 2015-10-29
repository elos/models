package models

import "github.com/elos/data"

func (g *Group) Contains(db data.DB, record data.Record) (bool, error) {
	contexts, err := g.Contexts(db)

	if err != nil {
		goto Denied
	}

	for _, c := range contexts {
		if c.Contains(record) {
			return true, nil
		}
	}

Denied:
	return false, nil
}
