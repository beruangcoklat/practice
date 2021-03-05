package consumer

import (
	"fmt"
	"time"
)

func New(id int) *Consumer {
	return &Consumer{
		id: id,
	}
}

func (c *Consumer) Consume(msg string) {
	fmt.Printf("1. id=%v msg=%v\n", c.id, msg)
	time.Sleep(time.Second * 2)
	fmt.Printf("1. id=%v msg=%v finished\n", c.id, msg)
}

func (c *Consumer) Consume2(msg string) {
	fmt.Printf("2. id=%v msg=%v\n", c.id, msg)
	time.Sleep(time.Second * 2)
	fmt.Printf("2. id=%v msg=%v finished\n", c.id, msg)
}

func (c *Consumer) ConsumeWrapper(msgs chan string, cFunc ConsumeFunc) {
	for msg := range msgs {
		cFunc(msg)
	}
}
