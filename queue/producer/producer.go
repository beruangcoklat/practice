package producer

import "github.com/beruangcoklat/practice/queue/mq"

type ProducerItf interface {
	Publish(topic, msg string)
}

type Producer struct {
	queue mq.MQItf
}

func New(queue mq.MQItf) *Producer {
	return &Producer{
		queue: queue,
	}
}

func (p *Producer) Publish(topic, msg string) {
	p.queue.Publish(topic, msg)
}
