package models

import "github.com/elos/data"

// See: https://github.com/elos/documentation/blob/master/data/models/ontology.md
type Ontology interface {
	data.Model
	Userable

	IncludeClass(Class) error
	ExcludeClass(Class) error
	ClassesIter(data.Access) (data.ModelIterator, error)
	Classes(data.Access) ([]Class, error)

	IncludeObject(Object) error
	ExcludeObject(Object) error
	ObjectsIter(data.Access) (data.ModelIterator, error)
	Objects(data.Access) ([]Object, error)
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
	Ontology(data.Access) (Ontology, error)

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
	ObjectsIter(data.Access) (data.ModelIterator, error)
	Objects(data.Access) ([]Object, error)

	NewObject(a data.Access) (Object, error)
}

// See: https://github.com/elos/documentation/blob/master/data/models/object.md
type Object interface {
	data.Model
	data.Nameable

	SetOntology(Ontology) error
	Ontology(data.Access) (Ontology, error)

	SetClass(Class) error
	Class(data.Access) (Class, error)

	SetTrait(data.Access, string, string) error
	AddRelationship(data.Access, string, Object) error
	DropRelationship(data.Access, string, Object) error
}
