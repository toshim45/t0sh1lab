package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

var (
	topics         []string = []string{"test"}
	brokers        []string = []string{"localhost:9092"}
	zookeeperNodes []string = []string{"localhost:2181"}
	consumerGrupID string   = "test-cg-1"
	msgText        string   = "test-produce-"
)

func main() {
	fmt.Println("starting")
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = 5 * time.Second

	consumer, err := consumergroup.JoinConsumerGroup(consumerGrupID, topics, zookeeperNodes, config)
	if err != nil {
		fmt.Printf("fail to join consumer group %v\n", err)
	}

	// set terminate service
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if err := consumer.Close(); err != nil {
			sarama.Logger.Println("Error closing the consumer", err)
		}
	}()

	go func() {
		for err := range consumer.Errors() {
			fmt.Println(err)
		}
	}()

	// get event message
	eventCount := 0
	offsets := make(map[string]map[int32]int64)
	for message := range consumer.Messages() {
		if offsets[message.Topic] == nil {
			offsets[message.Topic] = make(map[int32]int64)
		}

		eventCount += 1
		if offsets[message.Topic][message.Partition] != 0 && offsets[message.Topic][message.Partition] != message.Offset-1 {
			fmt.Printf("Unexpected offset on %s:%d. Expected %d, found %d, diff %d.\n", message.Topic, message.Partition, offsets[message.Topic][message.Partition]+1, message.Offset, message.Offset-offsets[message.Topic][message.Partition]+1)
		}

		// Simulate processing time
		fmt.Printf("Processing %v\n", message.Value)

		offsets[message.Topic][message.Partition] = message.Offset
		consumer.CommitUpto(message)
	}

	fmt.Printf("Processed %d events.", eventCount)
	fmt.Printf("%+v", offsets)
}
