package models

import "github.com/elos/data"

func (c *Context) Contains(r data.Record) bool {
	if c.Domain != string(r.Kind()) {
		return false
	}

	return includes(c.Ids, r.ID().String())
}

func includes(ss []string, s string) bool {
	for i := range ss {
		if ss[i] == s {
			return true
		}
	}

	return false
}
