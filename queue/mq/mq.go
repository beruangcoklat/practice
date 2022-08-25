package mq

import (
	"github.com/beruangcoklat/practice/queue/consumer"
)

type MQItf interface {
	Publish(topic, msg string)
	AddConsumer(topic string, concurrency int, cons consumer.ConsumerItf)
}

type MQ struct {
	consumer map[string]consumerData
}

type consumerData struct {
	ch chan string
}

func New() *MQ {
	return &MQ{
		consumer: make(map[string]consumerData),
	}
}

func (m *MQ) Publish(topic, msg string) {
	data := m.consumer[topic]
	data.ch <- msg
}

func (m *MQ) AddConsumer(topic string, concurrency int, cons consumer.ConsumerItf) {
	ch := make(chan string, concurrency)

	data := consumerData{
		ch: ch,
	}
	m.consumer[topic] = data

	for i := 0; i < concurrency; i++ {
		go cons.Consume(ch)
	}
}
