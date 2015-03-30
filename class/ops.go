package class

import (
	"github.com/elos/data"
	"github.com/elos/models"
)

func NewObject(c models.Class, a data.Access) (models.Object, error) {
	m, err := a.ModelFor(models.ObjectKind)
	if err != nil {
		return nil, err
	}
	obj, ok := m.(models.Object)
	if !ok {
		return nil, models.CastError(models.ObjectKind)
	}

	ont, err := c.Ontology(a)
	if err != nil {
		return nil, err
	}

	obj.SetOntology(ont)
	obj.SetClass(c)
	obj.SetName(c.Name())

	return obj, nil
}
