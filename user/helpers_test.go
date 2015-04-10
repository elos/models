package user_test

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/models/persistence"
	"github.com/elos/models/user"
	"github.com/elos/mongo"
	"gopkg.in/mgo.v2/bson"
)

func AuthenticateTest(t *testing.T) {
	// Initialization
	s := persistence.Store(persistence.MongoMemoryDB())

	// Creating User
	u, err := user.Create(s)
	if err != nil {
		t.Errorf("Error while creating user: %s", err)
	}
	u.SetKey(user.NewKey())
	testName := "me"
	u.SetName(testName)
	if err = s.Save(u); err != nil {
		t.Errorf("Error while saving user: %s", err)
	}

	// Good Auth
	u2, ok, err := user.Authenticate(s, u.ID().(bson.ObjectId).Hex(), u.Key())

	if !ok || err != nil {
		t.Errorf("Failed to authenticate when auth should have succeeded, err: %s", err)
	}

	if u2.Name() != testName || u.ID().String() == u2.ID().String() {
		t.Errorf("Failed to retrieve correct user from authenticate")
	}

	// Bad ID

	// 1 User doesn't exist
	u2, ok, err = user.Authenticate(s, mongo.NewObjectID().Hex(), u.Key())
	if ok || err == nil {
		t.Errorf("auth should have failed")
	}

	if u2 != nil {
		t.Errorf("Don't return user when auth fails")
	}

	// 2 Invalid ID
	u2, ok, err = user.Authenticate(s, "invalid id", u.Key())
	if ok || err == nil {
		t.Errorf("auth should have failed")
	}

	if u2 != nil {
		t.Errorf("Don't return user when auth fails")
	}

	// Bad Key
	u2, ok, err = user.Authenticate(s, u.ID().(bson.ObjectId).Hex(), "not the key")

	if ok || err == nil {
		t.Errorf("auth should have failed")
	}

	if u2 != nil {
		t.Errorf("Don't return user when auth fails")
	}
}

func TestFind(t *testing.T) {
	// Initialization and creating user
	s := persistence.Store(persistence.MongoMemoryDB())
	u := user.New(s)

	testName := "Not a real name"
	testKey := user.NewKey()

	u.SetName(testName)
	u.SetKey(testKey)

	if err := s.Save(u); err != nil {
		t.Errorf("Error while saving user: %s", err)
	}

	// User exists and should be found
	u2, err := user.Find(s, u.ID())
	if err != nil {
		t.Errorf("Error while finding user: %s", err)
	}

	if u2.ID().String() != u.ID().String() ||
		u2.Name() != u.Name() ||
		u2.Key() != u.Key() {
		t.Errorf("User found doesn't match user searched for")
	}

	// User doesn't exist
	u2, err = user.Find(s, mongo.NewObjectID())
	if err != data.ErrNotFound {
		t.Errorf("User should not be found")
	}

	if u2 != nil {
		t.Errorf("Should not return user if err not found")
	}
}

/*
  MemoryDB PopulateByField is not implemented

func TestFindBy(t *testing.T) {
	// Initialization and creating user
	s := persistence.Store(persistence.MongoMemoryDB())
	u, err := user.New(s)
	if err != nil {
		t.Errorf("Error while creating user: %s", err)
	}

	testName := "Not a real name"
	testKey := user.NewKey()

	u.SetName(testName)
	u.SetKey(testKey)

	if err = s.Save(u); err != nil {
		t.Errorf("Error while saving user: %s", err)
	}

	// User should be found
	u2, err := user.FindBy(s, "key", testKey)
	if err != nil {
		t.Errorf("Error while finding user %s", err)
	}

	if u2.ID().String() != u.ID().String() {
		t.Errorf("User found doesn't match user saved")
	}

	// User shouldn't be found
	u2, err = user.FindBy(s, "key", "garbage")

	if err != data.ErrNotFound {
		t.Errorf("Should not have found any user")
	}

	if u2 != nil {
		t.Errorf("When err not found, user returned should be nil")
	}
}

*/

func TestNewWithName(t *testing.T) {
	s := persistence.Store(persistence.MongoMemoryDB())

	u, err := user.NewWithName(s, "Name")
	if err != nil {
		t.Errorf("Error while creating user: %s", err)
	}

	if u.Name() != "Name" {
		t.Errorf("NewWithName failed to assign name to user")
	}

	u2, err := user.Find(s, u.ID())
	if err != nil {
		t.Errorf("Should be able to find the user, aka it should have been saved")
	}

	if u2.Name() != u.Name() {
		t.Errorf("User names should match")
	}
}
