package user_test

import (
	"log"
	"testing"
	"time"

	"github.com/elos/data"
	"github.com/elos/data/builtin/mem"
	"github.com/elos/models"
	"github.com/elos/models/access"
	"github.com/elos/models/user"
)

func TestNew(t *testing.T) {
	db := mem.NewDB()
	u := user.New(db)

	if u.Id == "" {
		t.Fatal("Id should not be empty")
	}

	if u.CreatedAt.IsZero() || u.UpdatedAt.IsZero() {
		t.Fatal("Bookeeping traits should be set")
	}
}

func TestCreate(t *testing.T) {
	db := mem.NewDB()
	userCreated, credentialCreated, err := user.Create(db, "username", "password")
	if err != nil {
		t.Fatal(err)
	}

	credentialRetrieved, err := access.Authenticate(db, "username", "password")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Credential created:\n%+v", credentialCreated)
	t.Logf("Credential retrieved:\n%+v", credentialRetrieved)

	if !data.Equivalent(credentialCreated, credentialRetrieved) {
		t.Fatal("The credential should have been added to the database")
	}

	userRetrieved, err := credentialRetrieved.Owner(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("User created:\n%+v", userCreated)
	t.Logf("User retrieved:\n%+v", userRetrieved)

	if !data.Equivalent(userCreated, userRetrieved) {
		t.Fatal("The user should have been added to the database")
	}
}

func TestForPhone(t *testing.T) {
	db := mem.NewDB()
	u := user.New(db)
	p := models.NewProfile()
	p.SetID(db.NewID())
	phoneNumber := "123 456 7890"
	p.Phone = phoneNumber
	p.SetOwner(u)
	if err := db.Save(u); err != nil {
		t.Fatal(err)
	}
	if err := db.Save(p); err != nil {
		t.Fatal(err)
	}

	uRetrieved, err := user.ForPhone(db, phoneNumber)
	if err != nil {
		t.Fatal(err)
	}

	if !data.Equivalent(u, uRetrieved) {
		t.Fatal("Expected to retrieve the user with the profile")
	}
}

func TestTasks(t *testing.T) {
	db := mem.NewDB()
	u, _, err := user.Create(db, "u", "p")
	if err != nil {
		t.Fatal(err)
	}

	tsk := models.NewTask()
	tsk.SetID(db.NewID())
	tsk.SetOwner(u)

	if err := db.Save(tsk); err != nil {
		t.Fatal(err)
	}

	tasks, err := user.Tasks(db, u, data.AttrMap{})
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 1 {
		t.Fatal("Expected there to be one task for this user")
	}

	if !data.Equivalent(tsk, tasks[0]) {
		t.Fatal("Expected task to match the one we created")
	}
}

func TestProfile(t *testing.T) {
	db := mem.NewDB()
	u, _, err := user.Create(db, "username", "password")
	if err != nil {
		t.Fatal(err)
	}

	p := models.NewProfile()
	p.SetID(db.NewID())
	p.SetOwner(u)
	if err := db.Save(p); err != nil {
		log.Fatal(err)
	}

	pFound, err := user.Profile(db, u)
	if err != nil {
		t.Fatal(err)
	}

	if !data.Equivalent(pFound, p) {
		t.Fatal("Profile saved and retrieved should match")
	}
}

func TestMap(t *testing.T) {
	db := mem.NewDB()

	u1, _, err := user.Create(db, "username", "password")
	if err != nil {
		log.Fatal(err)
	}
	u2, _, err := user.Create(db, "newusername", "newpassword")
	if err != nil {
		log.Fatal(err)
	}

	if !u1.DeletedAt.IsZero() {
		t.Fatal("User DeletedAt should be zero")
	}

	if !u2.DeletedAt.IsZero() {
		t.Fatal("User DeletedAt should be zero")
	}

	user.Map(db, func(db data.DB, u *models.User) error {
		u.DeletedAt = time.Now()
		return db.Save(u)
	})

	if err := db.PopulateByID(u1); err != nil {
		t.Fatal(err)
	}

	if err := db.PopulateByID(u2); err != nil {
		t.Fatal(err)
	}

	if u1.DeletedAt.IsZero() {
		t.Fatal("User DeletedAt should not be zero")
	}

	if u2.DeletedAt.IsZero() {
		t.Fatal("User DeletedAt should not be zero")
	}
}
