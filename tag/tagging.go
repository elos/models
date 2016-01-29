package tag

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

func Task(db data.DB, tsk *models.Task, name string) (*models.Tag, error) {
	u, err := tsk.Owner(db)
	if err != nil {
		return nil, err
	}

	tag, err := ByName(db, u, Name(name))
	if err != nil {
		return nil, err
	}

	tsk.IncludeTag(tag)

	if err := db.Save(tsk); err != nil {
		return nil, err
	}

	return tag, nil
}
