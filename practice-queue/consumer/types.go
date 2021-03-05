package consumer

type ConsumerItf interface {
	Consume(msg string)
	Consume2(msg string)
	ConsumeWrapper(msgs chan string, cFunc ConsumeFunc)
}

type Consumer struct {
	id int
}

type ConsumeFunc func(msgs string)
