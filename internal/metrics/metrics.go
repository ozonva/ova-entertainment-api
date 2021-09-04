package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	CreateSuccessResponseIncCounter()
	MultiCreateSuccessResponseIncCounter()
	DescribeSuccessResponseIncCounter()
	ListSuccessResponseIncCounter()
	RemoveSuccessResponseIncCounter()
}

type metrics struct {
	createCounter      prometheus.Counter
	multiCreateCounter prometheus.Counter
	describeCounter    prometheus.Counter
	listCounter        prometheus.Counter
	removeCounter      prometheus.Counter
}

func (m *metrics) CreateSuccessResponseIncCounter() {
	m.createCounter.Inc()
}

func (m *metrics) MultiCreateSuccessResponseIncCounter() {
	m.multiCreateCounter.Inc()
}

func (m *metrics) DescribeSuccessResponseIncCounter() {
	m.describeCounter.Inc()
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
		describeCounter: prometheus.NewCounter(prometheus.CounterOpts{
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
	// @todo что то не так
	prometheus.MustRegister(m.createCounter)
	prometheus.MustRegister(m.describeCounter)
	prometheus.MustRegister(m.listCounter)
	prometheus.MustRegister(m.removeCounter)
	prometheus.MustRegister(m.multiCreateCounter)

	return m
}
