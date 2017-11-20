package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"fmt"
)

func main() {
	producer, err := nsq.NewProducer("10.5.1.167:4150", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	err = producer.Ping()
	if nil != err {
		producer.Stop()
		log.Fatal(err)
	}

	err = producer.Publish("go_learning_test_nsq", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("finish")

}