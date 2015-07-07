package models

import "time"

func includes(set []string, object string) bool {
	for _, o := range set {
		if o == object {
			return true
		}
	}

	return false
}

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
