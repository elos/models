package models

import "github.com/elos/data"

// See: https://github.com/elos/documentation/blob/master/data/models/ontology.md
type Ontology interface {
	data.Model
	Userable

	IncludeClass(Class) error
	ExcludeClass(Class) error
	ClassesIter(Store) (data.ModelIterator, error)
	Classes(Store) ([]Class, error)

	IncludeObject(Object) error
	ExcludeObject(Object) error
	ObjectsIter(Store) (data.ModelIterator, error)
	Objects(Store) ([]Object, error)
}

type Trait struct {
	Name string
	Type string
}

type Relationship struct {
	Name    string
	Other   string
	Inverse string
}

// See: https://github.com/elos/documentation/blob/master/data/models/class.md
type Class interface {
	data.Model
	data.Nameable
	Userable

	SetOntology(Ontology) error
	Ontology(Store) (Ontology, error)

	IncludeTrait(*Trait)
	ExcludeTrait(*Trait)
	Traits() []*Trait
	Trait(string) (*Trait, bool)

	IncludeRelationship(*Relationship)
	ExcludeRelationship(*Relationship)
	Relationships() []*Relationship
	Relationship(string) (*Relationship, bool)

	IncludeObject(Object) error
	ExcludeObject(Object) error
	ObjectsIter(Store) (data.ModelIterator, error)
	Objects(Store) ([]Object, error)

	NewObject(a Store) (Object, error)
}

// See: https://github.com/elos/documentation/blob/master/data/models/object.md
type Object interface {
	data.Model
	data.Nameable

	SetOntology(Ontology) error
	Ontology(Store) (Ontology, error)

	SetClass(Class) error
	Class(Store) (Class, error)

	SetTrait(Store, string, string) error
	AddRelationship(Store, string, Object) error
	DropRelationship(Store, string, Object) error
}
