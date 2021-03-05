package mq

import "github.com/practice-queue/consumer"

type MQItf interface {
	Publish(topic, msg string)
	AddConsumer(topic string, concurrency int, cons consumer.ConsumerItf, cFunc consumer.ConsumeFunc)
}

type MQ struct {
	consumer map[string][]consumerData
}

type consumerData struct {
	ch chan string
}
