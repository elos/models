package models

import (
	"strings"

	"github.com/elos/data"
)

func (g *Group) HasAccess(db data.DB, property Property) (bool, AccessLevel, error) {
	if strings.ToLower(string(property.Kind())) != strings.ToLower(g.Domain) {
		return false, None, nil
	}

	if !includes(g.Ids, property.ID().String()) {
		return false, None, nil
	}

	return true, AccessLevels[g.Access], nil
}

func includes(ss []string, s string) bool {
	for i := range ss {
		if ss[i] == s {
			return true
		}
	}

	return false
}
