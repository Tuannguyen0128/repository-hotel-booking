package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type MQ struct {
	CClient *kafka.Consumer
}

func InitConnection() *MQ {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil
	}
	return &MQ{
		CClient: consumer,
	}
}
