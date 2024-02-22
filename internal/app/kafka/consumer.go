package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (mq *MQ) ConsumeMessage() {
	mq.CClient.SubscribeTopics([]string{"broker-repo"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		fmt.Println("Waiting for the next message")
		msg, err := mq.CClient.ReadMessage(-1)
		if err != nil && !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		} else {
			fmt.Println(string(msg.Value))
		}

	}
}
