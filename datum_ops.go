package models

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
