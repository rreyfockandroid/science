package main

import "fmt"

var (
	brokerList = []string{"node-1.kafka-2.dc-2.tools.dcwp.pl:9092"}
	topic = "adv.audit"
	maxRetry = 5
)

func main() {
	//consumer()
	producer()
}

func callBack(err error) {
	fmt.Printf("Error %v", err)
}

func producer() {
	producer := NewKafkaSyncProducer(brokerList, topic, maxRetry, callBack)
	producer.SendMessage(`{"mess":"test"}`)
}

