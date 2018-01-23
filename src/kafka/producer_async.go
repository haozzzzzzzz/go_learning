package main

import (
	"github.com/Shopify/sarama"
)

func main()  {

	config := sarama.NewConfig()
	//config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewAsyncProducer([]string{"10.5.1.82:9092"}, config)
	if err != nil {
		panic(err)
		return
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "service-user-service",
	}

	for i := uint32(0); i < 10; i ++{
		msg.Value = sarama.ByteEncoder("nihao")
		producer.Input() <- msg
	}
}