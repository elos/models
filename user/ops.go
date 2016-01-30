package user

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

// New performs a more proper instantiation of a
// user, rather than just structurally
func New(ider data.IDer) *models.User {
	return &models.User{
		Id:        ider.NewID().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Creates a user and a password credential
func Create(db data.DB, username, password string) (*models.User, *models.Credential, error) {
	u := New(db)
	c := &models.Credential{
		Id:        db.NewID().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Public:    username,
		Private:   password,
		Spec:      "password",
		OwnerId:   u.Id,
	}

	if err := db.Save(u); err != nil {
		return nil, nil, err
	}

	if err := db.Save(c); err != nil {
		return nil, nil, err
	}

	return u, c, nil
}

// ForPhone retrieves the user associated with a phone number
func ForPhone(db data.DB, phone string) (*models.User, error) {
	p := models.NewProfile()
	if err := db.PopulateByField("phone", phone, p); err != nil {
		return nil, err
	}
	return p.Owner(db)
}

// Tasks returns the tasks for this user
func Tasks(db data.DB, u *models.User, attrs data.AttrMap) ([]*models.Task, error) {
	attrs["owner_id"] = u.Id
	iter, err := db.Query(models.TaskKind).Select(attrs).Execute()

	if err != nil {
		return nil, err
	}

	t := models.NewTask()
	tasks := make([]*models.Task, 0)
	for iter.Next(t) {
		tasks = append(tasks, t)
		t = models.NewTask()
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return tasks, nil
}
