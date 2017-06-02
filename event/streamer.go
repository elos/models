package event

import (
	"github.com/elos/data"
	"github.com/elos/models"
	"golang.org/x/net/context"
)

type Streamer interface {
	Stream(ctx context.Context, in <-chan *models.Event)
	Close() error
}

type streamer struct {
	db  data.DB
	err error
}

func NewStreamer(db data.DB) Streamer {
	return &streamer{
		db: db,
	}
}

func (s *streamer) Stream(ctx context.Context, in <-chan *models.Event) {
stream:
	for {
		select {
		case e := <-in:
			if err := s.db.Save(e); err != nil {
				s.err = err
				break stream
			}
		case <-ctx.Done():
			break stream
		}
	}
}

func (s *streamer) Close() error {
	return s.err
}
