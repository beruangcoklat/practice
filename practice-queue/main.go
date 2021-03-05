package main

import (
	"fmt"

	"github.com/practice-queue/consumer"
	"github.com/practice-queue/mq"
	"github.com/practice-queue/producer"
)

func main() {
	var queue mq.MQItf = mq.New()
	var c1 consumer.ConsumerItf = consumer.New(1)

	queue.AddConsumer("topic 1", 2, c1, c1.Consume)
	queue.AddConsumer("topic 2", 1, c1, c1.Consume2)

	var p1 producer.ProducerItf = producer.New(queue)
	p1.Publish("topic 1", "msg 1")
	p1.Publish("topic 1", "msg 2")

	p1.Publish("topic 2", "wow 1")
	p1.Publish("topic 2", "wow 2")

	wait()
}

func wait() {
	var str string
	fmt.Scan(&str)
}
