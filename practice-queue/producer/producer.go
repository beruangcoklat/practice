package producer

import "github.com/practice-queue/mq"

func New(queue mq.MQItf) *Producer {
	return &Producer{
		queue: queue,
	}
}

func (p *Producer) Publish(topic, msg string) {
	p.queue.Publish(topic, msg)
}
