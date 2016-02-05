package event

import "github.com/elos/models"

func NewComponents(u *models.User, loc *models.Location, n *models.Note, m *models.Media, p *models.Event) *models.Event {
	e := models.NewEvent()
	if u != nil {
		e.SetOwner(u)
	}
	if loc != nil {
		e.SetLocation(loc)
	}
	if n != nil {
		e.SetNote(n)
	}
	if m != nil {
		e.SetMedia(m)
	}
	if p != nil {
		e.SetPrior(p)
	}
	return e
}
