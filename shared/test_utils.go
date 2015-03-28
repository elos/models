package shared

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/models"
	"gopkg.in/mgo.v2/bson"
)

/*
	ExpecteAccessDenial is a helper that ensures err
	is data.ErrAccessDenial and prints a failure message
	if not
*/
func ExpectAccessDenial(property string, err error, t *testing.T) {
	if err != data.ErrAccessDenial {
		t.Errorf("Expected access denial on %s, got %s", property, err)
	}
}

func ExpectEmptyLinkError(property string, err error, t *testing.T) {
	_, ok := err.(*data.EmptyLinkError)
	if !ok {
		t.Errorf("Expected empty link on %s, got %s", property, err)
	}
}

func ExpectEmptyRelationship(property string, err error, t *testing.T) {
	if err != models.ErrEmptyRelationship {
		t.Errorf("Expected empty relationship on %s, got %s", property, err)
	}
}

func ExpectNoError(op string, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("Error while %s: %s", op, err)
	}
}

/*
	TestAnonReadAccess ensures that an anonymous access can not
	read a model
*/
func TestAnonReadAccess(s data.Store, m data.Model, t *testing.T) {
	if err := s.Save(m); err != nil {
		t.Fatalf("Error while saving model: %s", err)
	}

	access := data.NewAnonAccess(s)

	m, err := access.Unmarshal(m.Kind(), data.AttrMap{
		"id": m.ID().(bson.ObjectId).Hex(),
	})

	if err != nil {
		t.Errorf("Error while unmarshalling user: %s", err)
	}

	ExpectAccessDenial("Reading User Anonymously", access.PopulateByID(m), t)
}
