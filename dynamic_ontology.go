package models

import "github.com/elos/data"

type Ontology interface {
	data.Model
	Userable

	IncludeClass(Class) error
	ExcludeClass(Class) error

	IncludeObject(Object) error
	ExcludeObject(Object) error

	Classes(data.Access) (data.ModelIterator, error)
	Objects(data.Access) (data.ModelIterator, error)
}

type Class interface {
	data.Model
	data.Nameable
	Userable

	SetOntology(Ontology) error
	Ontology(data.Access) (Ontology, error)

	IncludeTrait(*Trait) error
	ExcludeTrait(*Trait) error
	Traits() []*Trait

	IncludeRelationship(*Relationship) error
	ExcludeRelationship(*Relationship) error
	Relationships() []*Relationship

	IncludeObject(Object) error
	ExcludeObject(Object) error
	Objects(data.Access) (data.ModelIterator, error)

	Trait(string) (*Trait, bool)
	Relationship(string) (*Relationship, bool)

	NewObject(a data.Access) Object
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
