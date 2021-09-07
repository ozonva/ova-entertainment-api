package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics интерфейс для отправки метрик
type Metrics interface {
	IncCounterSuccessResponseForCreate()
	IncCounterSuccessResponseForMultiCreate()
	IncCounterSuccessResponseForUpdate()
	IncCounterSuccessResponseForList()
	IncCounterSuccessResponseForRemove()
}

type metrics struct {
	createCounter      prometheus.Counter
	multiCreateCounter prometheus.Counter
	updateCounter      prometheus.Counter
	listCounter        prometheus.Counter
	removeCounter      prometheus.Counter
}

// IncCounterSuccessResponseForCreate Инкремент метрики об успешном создании
func (m *metrics) IncCounterSuccessResponseForCreate() {
	m.createCounter.Inc()
}

// IncCounterSuccessResponseForMultiCreate Инкремент метрики об успешном мульти создании
func (m *metrics) IncCounterSuccessResponseForMultiCreate() {
	m.multiCreateCounter.Inc()
}

// IncCounterSuccessResponseForUpdate Инкремент метрики об успешном обновлении
func (m *metrics) IncCounterSuccessResponseForUpdate() {
	m.updateCounter.Inc()
}

// IncCounterSuccessResponseForList Инкремент метрики об успешном получении списка
func (m *metrics) IncCounterSuccessResponseForList() {
	m.listCounter.Inc()
}

// IncCounterSuccessResponseForRemove Инкремент метрики об успешном удалении
func (m *metrics) IncCounterSuccessResponseForRemove() {
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
