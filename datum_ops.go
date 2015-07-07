package models

import "time"

func (d *Datum) Match(tags []string) bool {
	for _, tag := range tags {
		if !includes(d.Tags, tag) {
			return false
		}
	}

	return true
}

func (d *Datum) Between(start time.Time, end time.Time) bool {
	return d.CreatedAt.After(start) && d.CreatedAt.Before(end)
}
