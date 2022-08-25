package main

import (
	"os"
	"os/signal"

	"github.com/beruangcoklat/practice/queue/consumer"
	"github.com/beruangcoklat/practice/queue/mq"
	"github.com/beruangcoklat/practice/queue/producer"
)

func main() {
	var queue mq.MQItf = mq.New()
	var c1 consumer.ConsumerItf = consumer.New(1)
	var c2 consumer.ConsumerItf = consumer.New(2)

	queue.AddConsumer("topic 1", 2, c1)
	queue.AddConsumer("topic 2", 1, c2)

	var p1 producer.ProducerItf = producer.New(queue)
	p1.Publish("topic 1", "msg 1")
	p1.Publish("topic 1", "msg 2")

	p1.Publish("topic 2", "wow 1")
	p1.Publish("topic 2", "wow 2")

	wait()
}

func wait() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
