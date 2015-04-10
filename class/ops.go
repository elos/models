package class

import "github.com/elos/models"

func NewObject(c models.Class, store models.Store) (models.Object, error) {
	obj := store.Object()
	ont, err := c.Ontology(store)
	if err != nil {
		return nil, err
	}

	obj.SetOntology(ont)
	obj.SetClass(c)
	obj.SetName(c.Name())
	return obj, nil
}
