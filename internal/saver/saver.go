package saver

import (
	"errors"
	"sync"
	"time"

	"github.com/ozonva/ova-entertainment-api/internal/flusher"
	"github.com/ozonva/ova-entertainment-api/internal/models"
)

const TickerInterval = 3

type Saver interface {
	Save(entity models.Entertainment) error
	Close()
}

type saver struct {
	sync.Mutex
	ticker   *time.Ticker
	ch       chan struct{}
	capacity uint
	flusher  flusher.Flusher
	models   []models.Entertainment
}

func NewSaver(capacity uint, flusher flusher.Flusher) Saver {
	s := &saver{
		capacity: capacity,
		flusher:  flusher,
		models:   make([]models.Entertainment, 0, capacity),
	}

	s.init()

	return s
}

func (s *saver) init() {
	s.ticker = time.NewTicker(TickerInterval * time.Second)
	s.ch = make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)
	go func(ch <-chan struct{}) {
		wg.Done()
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					s.ticker.Stop()
					return
				}
			case <-s.ticker.C:
				s.flush()
			}
		}
	}(s.ch)

	wg.Wait()
}

func (s *saver) Save(entity models.Entertainment) error {
	s.Lock()
	defer s.Unlock()

	if len(s.models) == cap(s.models) {
		return errors.New("Maximum slice length")
	}

	s.models = append(s.models, entity)
	return nil
}

func (s *saver) Close() {
	s.flush()
	close(s.ch)
}

func (s *saver) flush() {
	s.Lock()
	defer s.Unlock()

	if len(s.models) > 0 {
		s.flusher.Flush(s.models)
		s.models = make([]models.Entertainment, 0, s.capacity)
	}
}
