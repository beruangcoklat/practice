package producer

import "github.com/practice-queue/mq"

type ProducerItf interface {
	Publish(topic, msg string)
}

type Producer struct {
	queue mq.MQItf
}
