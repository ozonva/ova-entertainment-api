package healthcheck

import (
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-entertainment-api/internal/kafka"
	"sync"
	"time"
)

type Healthcheck interface {
	HealthDB() error
	HealthKafka() error
	Watch(errorCh chan<- struct{})
	Close()
}

type healthcheck struct {
	db       sqlx.DB
	producer kafka.Producer
	ticker   *time.Ticker
	ch       chan struct{}
}

func New(db sqlx.DB, producer kafka.Producer) Healthcheck {
	return &healthcheck{
		db:       db,
		producer: producer,
	}
}

func (h *healthcheck) HealthDB() error {
	_, err := h.db.Query("SELECT 1")

	return err
}

func (h *healthcheck) HealthKafka() error {
	err := h.producer.Send(kafka.Message{
		EventType: kafka.Ping,
		Value:     struct{}{},
	})

	return err
}

func (h *healthcheck) Watch(errorCh chan<- struct{}) {
	h.ticker = time.NewTicker(5 * time.Second)
	h.ch = make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)
	go func(ch <-chan struct{}) {
		wg.Done()
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					h.ticker.Stop()
					return
				}
			case <-h.ticker.C:
				err := h.HealthDB()
				if err != nil {
					errorCh <- struct{}{}
				}
				err = h.HealthKafka()
				if err != nil {
					errorCh <- struct{}{}
				}
			}
		}
	}(h.ch)

	wg.Wait()
}

func (h *healthcheck) Close() {
	close(h.ch)
}
