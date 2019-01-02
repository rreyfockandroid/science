package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
)

func consumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	master, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})
	messageCountStart := 0
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				messageCountStart++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	fmt.Println("Processed", messageCountStart, "messages")
}

