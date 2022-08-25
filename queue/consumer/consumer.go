package consumer

import (
	"fmt"
	"time"
)

type ConsumerItf interface {
	Consume(msgs chan string)
}

type Consumer struct {
	id int
}

func New(id int) *Consumer {
	return &Consumer{
		id: id,
	}
}

func (c *Consumer) consume(msg string) {
	fmt.Printf("id=%v msg=%v\n", c.id, msg)
	time.Sleep(time.Second * 2)
	fmt.Printf("id=%v msg=%v finished\n", c.id, msg)
}

func (c *Consumer) Consume(msgs chan string) {
	for msg := range msgs {
		c.consume(msg)
	}
}
