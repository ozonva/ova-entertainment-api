package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics интерфейс для отправки метрик
type Metrics interface {
	CreateSuccessResponseIncCounter()
	MultiCreateSuccessResponseIncCounter()
	UpdateSuccessResponseIncCounter()
	ListSuccessResponseIncCounter()
	RemoveSuccessResponseIncCounter()
}

type metrics struct {
	createCounter      prometheus.Counter
	multiCreateCounter prometheus.Counter
	updateCounter      prometheus.Counter
	listCounter        prometheus.Counter
	removeCounter      prometheus.Counter
}

// CreateSuccessResponseIncCounter Инкремент метрики об успешном создании
func (m *metrics) CreateSuccessResponseIncCounter() {
	m.createCounter.Inc()
}

// MultiCreateSuccessResponseIncCounter Инкремент метрики об успешном мульти создании
func (m *metrics) MultiCreateSuccessResponseIncCounter() {
	m.multiCreateCounter.Inc()
}

// UpdateSuccessResponseIncCounter Инкремент метрики об успешном обновлении
func (m *metrics) UpdateSuccessResponseIncCounter() {
	m.updateCounter.Inc()
}

// ListSuccessResponseIncCounter Инкремент метрики об успешном получении списка
func (m *metrics) ListSuccessResponseIncCounter() {
	m.listCounter.Inc()
}

// RemoveSuccessResponseIncCounter Инкремент метрики об успешном удалении
func (m *metrics) RemoveSuccessResponseIncCounter() {
	m.removeCounter.Inc()
}

func NewMetrics() Metrics {
	m := &metrics{
		createCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_success_create_response",
			Help: "Api Create success response",
		}),
		updateCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_success_describe_response",
			Help: "Api Describe success response",
		}),
		listCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_success_list_response",
			Help: "Api List success response",
		}),
		removeCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_success_remove_response",
			Help: "Api Remove success response",
		}),
		multiCreateCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_success_multi_create_response",
			Help: "Api Remove success response",
		}),
	}

	prometheus.MustRegister(m.createCounter)
	prometheus.MustRegister(m.updateCounter)
	prometheus.MustRegister(m.listCounter)
	prometheus.MustRegister(m.removeCounter)
	prometheus.MustRegister(m.multiCreateCounter)

	return m
}
