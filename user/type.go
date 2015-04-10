package user

import (
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
	"github.com/elos/mongo"
)

var (
	kind    data.Kind   = models.UserKind
	schema  data.Schema = models.Schema
	version int         = models.DataVersion

	actions       data.LinkName = models.UserActions
	events        data.LinkName = models.UserEvents
	tasks         data.LinkName = models.UserTasks
	routines      data.LinkName = models.UserRoutines
	currentAction data.LinkName = models.UserCurrentAction
	calendar      data.LinkName = models.UserCalendar
	ontology      data.LinkName = models.UserOntology
)

/*
	NewM is like New except that is satisfies the data.ModelConstructor
	interface.

	When associating a user with a store, use the NewM function:

		var store data.Store
		store.Register("user", user.NewM)

	We need this NewM method because a store doesn't know anything
	about a particular model, and therefore store.ModelFor and
	store.Unmarshal only work at the level of the data.Model interface.
*/
func NewM(s data.Store) data.Model {
	return New(s)
}

/*
	New looks up the data.DBType of the store and returns
	the appropriate model implementation.

	The only existing implementation is for mongo.

	If the data.DBType is not implemented by the user package,
	New will return a data.ErrInvalidDBType error, in which case
	the first return value (the user) will be nil. Check errors,
	or you will get a nil pointer dereference error.

		var store data.Store
		u, err := user.New(store)
		if err != nil { // or if err == data.ErrInvalidDBType
			// do something
		}

	Note that New will set the ID, but not the key for a user (which is a required field).
*/
func New(s data.Store) models.User {
	switch s.Type() {
	case mongo.DBType:
		u := &mongoUser{}
		u.SetID(s.NewID())
		return u
	default:
		panic(data.ErrInvalidDBType)
	}
}

func Create(s data.Store) (models.User, error) {
	u := New(s)
	return u, s.Save(u)
}

/*
	CreateAttrs instantiates and *saves* user using the the provided
	data.AttrMap.

	Create will look for an id attribute, a created_at attribute, and
	a name attriute. Create also generates and sets the key for a user.

		import (
			"time"
			"github.com/elos/data"
		)

		var store data.Store
		u, err := user.Create(s, data.AttrMap{
			"id": store.NewID(),
			"name": "Nick Landolfi",
			"created_at": time.Now(),
		})

	Most errors for create don't matter. But if the id is bad, Create complains.
*/
func CreateAttrs(s data.Store, a data.AttrMap) (models.User, error) {
	user := New(s)

	if idInterface, present := a["id"]; present {
		idString, ok := idInterface.(string)
		if ok {
			id, err := s.ParseID(idString)
			if err == nil {
				user.SetID(id)
			}
		}
	}

	ca, present := a["created_at"]
	ca, valid := ca.(time.Time)
	if present && valid {
		user.SetCreatedAt(ca.(time.Time))
	} else {
		user.SetCreatedAt(time.Now())
	}

	n, present := a["name"]
	n, valid = n.(string)
	if present && valid {
		user.SetName(n.(string))
	}

	user.SetKey(NewKey())

	if err := s.Save(user); err != nil {
		return user, err
	} else {
		return user, nil
	}
}

/*
	Validate confirms the presence of required attributes
	for a user's data to be considered "complete." It works
	at the models.User level.

	If the model is invalid, the second return value can be
	inspected for why.

		var u models.User
		valid, err := Validate(u)
		if !valid {
			// ... http.Error( ...
		}
*/
func Validate(u models.User) (bool, error) {
	if u.Name() == "" {
		return false, data.NewAttrError("name", "be present")
	}

	if u.Key() == "" {
		return false, data.NewAttrError("key", "be present")
	}

	return true, nil
}
