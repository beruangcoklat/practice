package mq

import (
	"github.com/practice-queue/consumer"
)

func New() *MQ {
	return &MQ{
		consumer: make(map[string][]consumerData),
	}
}

func (m *MQ) Publish(topic, msg string) {
	datas := m.consumer[topic]
	for _, data := range datas {
		data.ch <- msg
	}
}

func (m *MQ) AddConsumer(topic string, concurrency int, cons consumer.ConsumerItf, cFunc consumer.ConsumeFunc) {
	_, ok := m.consumer[topic]

	ch := make(chan string, concurrency)

	data := consumerData{
		ch: ch,
	}

	if ok {
		m.consumer[topic] = append(m.consumer[topic], data)
	} else {
		m.consumer[topic] = []consumerData{data}
	}

	for i := 0; i < concurrency; i++ {
		go cons.ConsumeWrapper(ch, cFunc)
	}
}
