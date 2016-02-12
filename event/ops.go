package event

import "github.com/elos/models"

func ContainsTags(e *models.Event, tags ...*models.Tag) bool {

LookingThroughTags:
	for _, ta := range tags {
		for _, tid := range e.TagsIds {
			if ta.Id == tid {
				continue LookingThroughTags
			}
		}

		// Otherwise we didn't find it
		return false
	}

	return true
}
