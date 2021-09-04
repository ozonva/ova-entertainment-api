package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

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

func (m *metrics) CreateSuccessResponseIncCounter() {
	m.createCounter.Inc()
}

func (m *metrics) MultiCreateSuccessResponseIncCounter() {
	m.multiCreateCounter.Inc()
}

func (m *metrics) UpdateSuccessResponseIncCounter() {
	m.updateCounter.Inc()
}

func (m *metrics) ListSuccessResponseIncCounter() {
	m.listCounter.Inc()
}

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
