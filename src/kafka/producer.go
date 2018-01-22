package main

import (
	"github.com/Shopify/sarama"
	"fmt"
)

func main()  {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"10.5.1.82:9092"}, config)
	if err != nil {
		panic(err)
	    return
	}


	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "service-user-service",
	}

	for {
		msg.Value = sarama.ByteEncoder("nihao")
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("send message faild.")
		    return
		}

		fmt.Printf("Partition = %d, offset = %d \n", partition, offset)
	}
}