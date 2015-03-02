package interactive

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/models/user"
	"github.com/robertkrimen/otto"
	"gopkg.in/mgo.v2/bson"
)

type ObjectKind string

type Credentials struct {
	ID  string
	Key string
}

type Space struct {
	*data.MemoryStore
	User *User
}

func Access(a data.Access) (space *Space, err error) {
	space = new(Space)

	space.MemoryStore = data.NewMemoryStore(a)
	space.User = space.FindUser(space.MemoryStore.Access.Client().ID().(bson.ObjectId).Hex())

	return
}

func NewSpace(c *Credentials, store data.Store) (*Space, error) {
	client, authed, err := user.Authenticate(store, c.ID, c.Key)
	if !authed {
		return nil, err
	}

	return Access(data.NewAccess(client, store))
}

func (s *Space) Expose(o *otto.Otto) {
	o.Set("FindUser", s.FindUser)
	o.Set("FindAction", s.FindAction)
	o.Set("FindRoutine", s.FindRoutine)
	o.Set("FindTask", s.FindTask)
	o.Set("FindFixture", s.FindFixture)
	o.Set("FindSchedule", s.FindSchedule)
	o.Set("FindOntology", s.FindOntology)
	o.Set("FindClass", s.FindClass)
	o.Set("FindObject", s.FindObject)

	o.Set("User", func() *User {
		return NewUser(s)
	})

	o.Set("Action", func() *Action {
		return NewAction(s)
	})

	o.Set("Routine", func() *Routine {
		return NewRoutine(s)
	})

	o.Set("Task", func() *Task {
		return NewTask(s)
	})

	o.Set("Fixture", func() *Fixture {
		return NewFixture(s)
	})

	o.Set("Schedule", func() *Schedule {
		return NewSchedule(s)
	})

	o.Set("Ontology", func() *Ontology {
		return NewOntology(s)
	})

	o.Set("Class", func() *Class {
		return NewClass(s)
	})

	o.Set("EObject", func() *Object {
		return NewObject(s)
	})

	o.Set("me", s.User)
	o.Set("StartAction", s.StartAction)
}

func (s *Space) FindUser(id string) *User {
	m, _ := s.Access.Unmarshal(models.UserKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return UserModel(s, m.(models.User))
}

func (s *Space) FindAction(id string) *Action {
	m, _ := s.Access.Unmarshal(models.ActionKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return ActionModel(s, m.(models.Action))
}

func (s *Space) FindRoutine(id string) *Routine {
	m, _ := s.Access.Unmarshal(models.RoutineKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return RoutineModel(s, m.(models.Routine))
}

func (s *Space) FindTask(id string) *Task {
	m, _ := s.Access.Unmarshal(models.TaskKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return TaskModel(s, m.(models.Task))
}

func (s *Space) FindFixture(id string) *Fixture {
	m, _ := s.Access.Unmarshal(models.FixtureKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return FixtureModel(s, m.(models.Fixture))
}

func (s *Space) FindSchedule(id string) *Schedule {
	m, _ := s.Access.Unmarshal(models.ScheduleKind, data.AttrMap{
		"id": id,
	})
	s.Access.PopulateByID(m)
	return ScheduleModel(s, m.(models.Schedule))
}

func (s *Space) StartAction(name string) *Action {
	a := NewAction(s)
	a.Name = name
	a.Save()
	s.User.SetCurrentAction(a)
	return a
}

func (s *Space) Reload() {
	s.MemoryStore.ReloadObjects()
}

func (s *Space) Register(o data.MemoryObject) {
	s.MemoryStore.RegisterObject(o)
}
