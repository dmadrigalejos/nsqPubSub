package main

import (
	"github.com/bitly/go-nsq"
	"log"
)

var (
	nsqAddr    = "127.0.0.1:4150"
	nsqTopic   = "TestTopic"
	nsqChannel = "TestChannel"
)

type handler struct {
	q *nsq.Consumer
}

// HandleMessage - message processor
func (h *handler) HandleMessage(message *nsq.Message) error {
	log.Println(string(message.Body))

	return nil
}

func main() {
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer(nsqTopic, nsqChannel, config)
	h := &handler{
		q: q,
	}
	q.SetHandler(h)

	err := q.ConnectToNSQD(nsqAddr)
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan bool)
}
