package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var (
	topic   string   = "test"
	brokers []string = []string{"localhost:9092"}
	msgText string   = "test-produce-3"
)

/*
func main() {
	fmt.Println("starting")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		fmt.Errorf("err creating sync producer %v", err)
		return
	}

	messageInput := sarama.StringEncoder(msgText)
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: messageInput,
	}

	partition, offset, err := producer.SendMessage(message)

	if err != nil {
		fmt.Errorf("error sending sync producer %v", err)
	}

	fmt.Printf("partition %v, offset %v", partition, offset)
}*/

func main() {
	fmt.Println("starting")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		fmt.Errorf("err creating async producer %v", err)
		return
	}
	defer producer.AsyncClose()

	messageInput := sarama.StringEncoder(msgText)
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: messageInput,
	}

	producer.Input() <- message

	select {
	case success := <-producer.Successes():
		fmt.Printf("success %v\n", success.Offset)
	case errors := <-producer.Errors():
		fmt.Printf("error %v\n", errors)
	}
}
