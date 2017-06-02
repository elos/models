package event_test

import (
	"testing"

	"github.com/elos/data/builtin/mem"
	"github.com/elos/models"
	"github.com/elos/models/event"
	"golang.org/x/net/context"
)

func TestStreamer(t *testing.T) {
	db := mem.NewDB()
	ctx, cancel := context.WithCancel(context.Background())
	s := event.NewStreamer(db)

	go s.Stream(ctx, make(chan *models.Event))
	cancel()

	if err := s.Close(); err != nil {
		t.Fatal("s.Close() error: %v", err)
	}
}
