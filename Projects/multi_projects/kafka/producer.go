package main

import (
	"github.com/Shopify/sarama"
)

// implementators: kafkaSyncProducer
type KafkaProducer interface {
	SendMessage(message string)
}
// implementators: status
type KafkaProducerStatus interface {
	GetPartition() int32
	GetOffset() int64
	GetError() error
	IsError() bool
}

type kafkaSyncProducer struct {
	brokerList []string
	topic string
	maxRetry int
	producer sarama.SyncProducer
}

/**
 * brokerList - lista klientow kafki
 * topic - topic do jakiego chcemy pisac
 * maxRetry - ile razy ma ponawiac probe dostarczenia komunikatu
 * callBack - funkcja jaka bedzie wywolana w razie problemow
 */
func NewKafkaSyncProducer(brokerList []string, topic string, maxRetry int, callBack func(error)) *kafkaSyncProducer {
	return &kafkaSyncProducer{
		brokerList: brokerList,
		topic: topic,
		maxRetry: maxRetry,
		producer: create(callBack),
	}
}

func create(callBack func(error)) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = maxRetry
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		callBack(err)
	}
	return producer
}

func (ksp *kafkaSyncProducer) SendMessage(message string) KafkaProducerStatus {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := ksp.producer.SendMessage(msg)
	if err != nil {
		return &status {
			error: err,
		}
	}
	return &status {
		partition: partition,
		offset: offset,
	}
}

type status struct {
	partition int32
	offset int64
	error error
}

func (s *status) GetPartition() int32 {
	return s.partition
}

func (s *status) GetOffset() int64 {
	return s.offset
}

func (s *status) GetError() error {
	return s.error
}

func (s *status) IsError() bool {
	return s.error != nil
}
