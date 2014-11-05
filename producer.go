package main

import (
	"github.com/bitly/go-nsq"
	"log"
)

var (
	nsqAddr  = "127.0.0.1:4150"
	nsqTopic = "TestTopic"
)

func main() {
	testMsg := "Hello"

	log.Println("Connecting")
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(nsqAddr, config)
	defer w.Stop()

	log.Println("Sending")
	err := w.Publish(nsqTopic, []byte(testMsg))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}
