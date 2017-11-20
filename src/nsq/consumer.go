package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

type NSQHandler struct {

}

func (m *NSQHandler) HandleMessage(message * nsq.Message) error {
	log.Println("recv:", string(message.Body))
	return nil
}

func main() {
	waiter := sync.WaitGroup{}
	waiter.Add(1)
	go func() {
		defer waiter.Done()

		consumer, err := nsq.NewConsumer("go_learning_test_nsq", "common", nsq.NewConfig())
		if err != nil {
			log.Fatal(err)
		}

		consumer.AddHandler(&NSQHandler{})
		err = consumer.ConnectToNSQLookupd("10.5.1.167:4161")
		if err != nil {
			log.Fatal(err)
		}

		select {

		}

	}()

	waiter.Wait()
}
