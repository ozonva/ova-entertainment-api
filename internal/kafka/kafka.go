package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type Producer interface {
	Send(message Message) error
	Close() error
}

type producer struct {
	syncProducer sarama.SyncProducer
	topic        string
}

type EventType int

const (
	Create EventType = iota
	MultiCreate
	Update
	Remove
)

type Message struct {
	EventType EventType
	Value     interface{}
}

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

func (p *producer) Close() error {
	return p.syncProducer.Close()
}
