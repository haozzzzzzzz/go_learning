package main

import (
	"github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
	"log"
	"fmt"
)

func main()  {
	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// init consumer
	brokers := []string{"10.5.1.82:9092"}
	topics := []string{"service-user-service"}
	consumer, err := cluster.NewConsumer(brokers, "test-consumer-group", topics, config)
	if err != nil {
		panic(err)
	    return
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}

			case <-signals:
				return
		}
	}
}