package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

// Producer отправляет в очередь сообщения
type Producer interface {
	Send(message Message) error
	Close() error
}

type producer struct {
	syncProducer sarama.SyncProducer
	topic        string
}

// EventType тип события которое отправили
type EventType int

const (
	Create EventType = iota
	MultiCreate
	Update
	Remove
)

// Message Сообщение которое отправляется в очередь
// для каждого типа действия Value может быть разный по типу данных
// Value будет преобразовано в json
type Message struct {
	EventType EventType
	Value     interface{}
}

// New инициализация и подключение к брокерам
func New(brokerList []string) (Producer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(brokerList, saramaConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Create producer failed")
		return nil, err
	}

	return &producer{
		topic:        "ova-entertainment-api",
		syncProducer: syncProducer,
	}, nil
}

// Send Отправка сообщения в топик
func (p *producer) Send(message Message) error {
	jsonMes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, _, err = p.syncProducer.SendMessage(
		&sarama.ProducerMessage{
			Topic:     p.topic,
			Partition: -1,
			Key:       sarama.StringEncoder(p.topic),
			Value:     sarama.StringEncoder(jsonMes),
		})
	return err
}

// Close Завершение работы продюсера
func (p *producer) Close() error {
	return p.syncProducer.Close()
}
