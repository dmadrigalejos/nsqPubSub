nsqPubSub
=========

A simple example of using nsq

### Requirement

1. golang
2. nsqd - download and run nsqd

    $ ./nsqd > /dev/null 2>&1 &

### Usage

Run the consumer

    $ go run consumer.go

Run the producer

    $ go run producer.go
